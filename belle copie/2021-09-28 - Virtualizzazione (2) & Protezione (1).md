<!-- Lezione 2021-09-28 -->
<!--
RIPASSINO
Avevamo introdotto il tema della migrazione live, meccanismo che consente lo spostamento della VM da un nodo fisico ad un altro. Questo spostamento, se fatto in modo "live", significa che può essere eseguito senza neanche spegnere la macchina in questione.
In questo modo si ha un downtime davvero trascurabile (poco più di 100ms)

La capacità/caratteristica di indipendenza e isolamento dall'ambiente fisico è il vantaggio principale che ci consente di realizzare in modo molto semplice dal punto di vista tecnico, la migrazione live, ovvero suspend su un nodo e resume su un altro nodo.

Fra gli obbiettivi, sicuramente uno prioritario è minimizzare il downtime, in quanto se la macchina fornisce un servizio, vogliamo che questo rimanga disponibile per il maggior tempo possibile e dunque, appunto, che il downtime venga minimizzato.
Altri aspetti da tenere in considerazione:
- ridurre al minimo il tempo di migrazione, come tempo complessivo che richiede la migrazione;
- occupare meno banda possibile.

Naturalmente ci sono dei vantaggi nella realizzazione della migrazione se le architetture alle quali stiamo facendo riferimento hanno file system in comune (es. cluster), ovvero condividono gli stessi dischi.
-->

#### Soluzione: precopy
La soluzione più diffusa al giorno d'oggi si basa su un meccanismo di precopia, che viene attuata in una serie di passi:
1. **Pre-migrazione**: fase iniziale in cui si capisce quali sono i nodi interessati, ovvero si individua la VM da migrare (nodo A) e l'host di destinazione (nodo B);
2. **Reservation**: viene riservato un contenitore vuoto nel server di destinazione (reservation del posto per la macchina da migrare);
3. **Pre-copia iterativa delle pagine**: la VM da migrare chiaramente avrà un file immagine (tipicamente un file di stato mappato sui registri CPU). In questa fase viene eseguita una copia nell'host B di tutte le pagine allocate in memoria sull'host A. Poiché le pagine in memoria, soprattutto se la VM è in esecuzione, possono variare, ovviamente non è detto che vengano copiate una volta sola. Alla successiva iterazione, vengono copiate solo le pagine modificate (*dirty pages*), fino a quando il numero di queste è inferiore ad una certa soglia data;
4. **Sospensione della VM**: raggiunta la soglia (quando rimangono poche pagine), si applica la suspend sulla macchina d'origine (in seguito avverrà una resume sulla macchina di destinazione);
5. **Commit**: la copia della VM sul nodo di destinazione è completa, dunque si può procedere con una commit (ovvero ci si affranca completamente dal nodo di origine, dal quale la VM viene eliminata);
6. **Resume**: viene eseguita la resume sul nodo B, in cui si trova una macchina pronta a ripartire, completa sia come immagine sul file system, sia come stato presente nei registri.

Con questa modalità, si ha downtime solo durante la copia delle ultime dirty pages, ovvero quando si è raggiunta la soglia preimpostata.
```
NB: chiaramente la prima iterazione della precopia è quella che richiede più tempo,
quelle successive ne richiedono meno perché salvano solo le pagine modificate.
```
Sebbene la precopia sia la modalità oggi più diffusa, ne esistono anche altre, ad esempio *post-copy*, in cui la macchina viene sospesa e vengono copiate (non iterativamente) pagine e stato. Così facendo si ottiene un tempo totale di migrazione più basso, ma un downtime dei servizi forniti dalla VM molto più elevato.

### XEN
XEN è un progetto che nasce in ambito accademico a Cambridge. Nasce come hypervisor (VMM paravirtualizzato), richiede che le VM che girano sopra xen abbiano un kernel adattato all'interfaccia che xen offre ai propri utilizzatorii. Per quanto riguarda il porting di Linux ha coinvolto circa 3000 linee di codice del kernel, per adattarlo in modo che potesse dialogare con le API di XEN.
Dal punto di vista commerciale ha limitato la gamma di kernel installabili, per quanto riguarda i SO proprietari, nonostante un tentativo di porting dei Sistemi Operativi (ad esempio Windows, che non è stato portato a termine).

#### Architettura
XEN è costituito da un VMM *hypervisor*, che si appoggia direttamente sull'hardware (virtualizzazione di sistema - quindi è necessario avere spazio e in caso togliere il SO preesistente) e si occupa della virtualizzazione della CPU, della memoria e dei dispositivi di ogni VM. In XEN le macchine virtuali vengono chiamate *domain* e su ogni sistema XEN c'è una VM speciale chiamata *domain 0* che è privilegiata: a livello architetturale è come tutte le altre ma, tramite un'interfaccia di controllo fornita da XEN, può amministrare tutto il sistema. Questa interfaccia è accessibile solo dal domain 0, ed è separata dall'hypervisor stesso, scelta che permette di ottenere una separazione dei meccanismi dalle politiche: all'interno delle applicazioni che consento la configurazione ed il controllo del sistema abbiamo le politiche (espresse dall'utente), che vengono poi implementate e messe in pratica dall'hypervisor. Infatti, tipicamente nel domain 0 girano applicazioni che consentono all'amministratore di configurare il sistema virtualizzato e operando sulla console di questa VM è possibile creare una VM guest (di domain U - utente), eliminarla, migrarla, ecc.

#### Realizzazione
Un VMM assomiglia per certi versi al kernel di un SO: deve gestire in modo appropriato l'hardware e fornire un accesso particolare agli utilizzatori (che nel caso di un sistema virtualizzato non sono gli utenti ma le VM.
Ogni VM vede una *CPU* come se fosse a lei esclusivamente dedicata, quando in realtà non è così: le risorse vengono condivise grazie all'attività dell'hypervisor tra tutti gli utilizzatori secondo politiche particolari (ad esempio per quanto riguarda la CPU l'hypervisor dovrà mettere in atto politiche di scheduling particolari).
Stessa cosa vale per la *memoria*, anch'essa dev'essere in qualche modo messa a disposizione per gli utilizzatori dal VMM, che deve garantire i criteri di sicurezza opportuni.
Altro compito importantissimo del VMM è quello della gestione dei *dispositivi* (quindi I/O).

qualche cenno sulle caratteristiche di XEN: noi facciamo riferimento a XEN "paravirtualizzato": in questi sistemi necessario separare il kernel dalla macchina virtuale e dalle applicazioni, in quanto XEN adotta una configurazione dei ring 0/1/3 (VMM esegue a ring 0, Sistemi Operativi a ring 1, le applicazioni a ring 3, così non si ha ring compression).
Le app possono utilizzare le system call per comunicare col sistema operativo, i sistemi operativi possono comunicare col VMM tramite delle hyper calls (sono come system call fornite dal nucleo dell'hypervisor per permettere agli SO di eseguire istruzioni particolari).

#### Gestione della Memoria e Paginazione
I SO guest gestiscono la memoria virtuale mediante la paginazione tradizionale: le page table delle VM vengono mappate in memoria fisica da XEN, il quale è l'unico a potervi accedere in *scrittura*, su richiesta delle VM. L'accesso in *lettura*, invece, è permesso anche ai sistemi operativi ospitati.

##### Protezione: Memory Split
Com'è strutturato lo spazio di indirizzamento delle singole VM guest? Si adotta un principio di **memory split**.
```
NB: Consideriamo sempre il parallelo con sistema non virtualizzato/sistema virtualizzato:
in un sistema virtualizzato, ogni utilizzatore è una VM, quindi ogni entità che si inter-
faccia col VMM (equivalente del kernel) è una VM. Così come accade nei sistemi non virtua-
lizzati, in cui ogni processo ha un utilizzatore e un suo spazio di indirizzamento, anche
nei sistemi virtualizzati ogni VM ha un suo spazio di indirizzamento virtuale (perché sia-
mo in presenza di memoria virtuale).
```
Per motivi di efficienza, poiché chiaramente nella commutazione tra una VM e l'altra c'è problema di reperire il codice di XEN, lo spazio di indirizzamento di ogni VM è strutturato a "segmenti": nei primi 64MiB viene allocato XEN (ring 0), poi c'è una parte relativa al Kernel del SO guest (ring 1), poi c'è lo spazio utente, che verrà utilizzato dalle applicazioni (ring 3). I VM guest si occupano delle politiche di gestione della paginazione, mentre i meccanismi, ovvero l'effettiva implementazione della paginazione, sono compito del VMM, in quanto il kernel del SO guest, non può occuparsene, non essendo nel ring privilegiato 0. Ciò garantisce maggiore protezione in quanto si ha separazione tra politiche (a carico dei guest - alto livello) e meccanismi (a carico del VMM - basso livello).
Con questa soluzione, quando viene creato un nuovo processo nello spazio del guest, fra le altre cose dev'essere creata una Tabella delle Pagine (PT) associata a tale processo. Ovviamente, poiché come detto tale operazione non può essere fatta dal kernel del sistema operativo che ospita quel processo (in quanto si trova a ring 1), dev'essere fatta da qualcun'altro. Quindi ciò che succede è che il guest richiede una nuova PT all'hypervisor, il quale la crea e vi aggiunge anche lo spazio riservato a XEN; così facendo XEN registra la tabella e acquisisce il diritto di scrittura esclusivo (i guesto potranno solo leggerle), e ogni volta che il guest di tale TP dovrà aggiornarla, proverà a scriverci generando un trap *protection fault*, che verrà catturata e gestita da XEN, permettendogli di verifcare la correttezza della richiesta ed aggiornare effettivamente, in seguito, la Tabella delle Pagine.

##### Protezione: Balloon Process
Per com'è gestita la protezione in XEN, l'unica componente capace di allocare memoria è il VMM (ring 0), ma può farlo solo in seguito a richeiste delle VM guest, in quanto come detto, le politiche si trovano in alto livello (ring 3), mentre i meccanismi a basso livello (ring 0). Però, in alcuni casi (es: attivazione nuova VM, operazione per la quale serve acquisire memoria necessaria per allocare lo spazio di indirizzamento di quella macchina virtuale), può essere necessario al VMM dover ottenere nuove pagine. Questa possibilità, ovvero di richiedere pagine, il VMM non ce l'ha. Può farlo solo in seguito a richieste da parte dei guest. Per risolvere questo problema, su XEN è stato adottata una soluzione (peculiare per la paravirtualizzazione) chiamata **balloon process**: in ogni guest c'è un processo in costante esecuzione, che è in grado di dialogare direttamente con l'hypervisor. In caso di necessità di pagine, il VMM può chiedere a tali processi di "gonfiarsi", ovvero richiedere al proprio SO ulteriori pagine. Tale richiesta provoca l'allocazione di nuove pagine al balloon process che, una volta ottenute, le cede al VMM.

#### Cenni su Virtualizzazione della CPU
Il VMM definisce un'architettura virtuale simile a quella del processore, nella quale però, le istruzioni privilegiate sono sostituite da opportune hypercalls:

Il VMM si occupa dello scheduling delle VM, seguendo un algoritmo molto generale (in grado di soddisfare dei vincoli temporali molto stringenti) chiamato *Borrowed Virtual Time*, che si basa sulla nozione di virtual-time: è un tempo che va avanti solo fintanto che la VM è attiva, ovvero se si trova in uno stato di sospensione il tempo si ferma e riprende quando viene attivato. Xen adotta due clock, uno relativo al real-time, l'altro al virtual-time.

#### Virtualizzazione dei dispositivi (I/O)
Le VM devono poter accedere ai dispositivi che sono disponibili a livello hardware. La scelta di XEN è quella, ovviamente di virtualizzare l'interfaccia di ogni dispositivo, ma farlo tramite due tipi di driver: *back-end driver* e *front-end driver*.<br/>
**Back-end driver** è il driver vero e proprio, che permette, tramite un'interfaccia del VMM chiamata *Safe Hardware Interface*, di comunicare ed utilizzare il dispostivo collegato a livello hardware. Tipicamente viene installato all'interno di una VM particolare che è sempre ancorata al nodo fisico (dominio 0 - solitamente qui vengono installati tutti i driver di ogni dispositivo presente connesso a livello fisico in quel nodo).<br/>
**Front-end driver** è un driver "astratto", generico, non riferito adun dispositivo particolare, che viene installato tipicamente nel kernel del SO di una VM guest. Questo driver, all'occorrenza si collega al back-end driver specifico.
```
NB: non c'è niente che vieti di installare un back-end direttametne su una VM di domain U,
ma può convenire concentrarli tutti nel domain 0, sia perché siamo certi che quella macchi-
na non si sposterà mai da lì, essendo ancorata all'hardware, sia per motivi di portabilità.
```

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/01%20-%20Virtualizzazione/Driver.png" alt="Front-end and Back-end Drivers"/>

Ovviamente, per consentire la comunicazione tra back-end driver e front-end driver, serve un meccanismo che gestica le richieste. Questo viene realizzato tramite delle strutture chiamate asynchronous I/O rings (buffer FIFO circolari) in cui ogni elemento è una specie di descrittore che rappresenta una particolare richiesta. Le richieste di accesso ad un particolare device vengono fatte dal guest tramite il front-end che deposita la richiesta nel ring relativo, mentre dall'altra parte c'è il back-end che le preleva e le gestisce.

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/01%20-%20Virtualizzazione/Asynchronous%20IO%20Ring.png" alt="Structure of Asynchronous I/O Rings"/>

**Vantaggi**: il driver viene scorporato in due parti, svingolando la VM dal particolare server fisico in cui risiede (il front-end driver della VM rimane lo stesso anche se questa viene spostata su un altro nodo), garantendo *portabilità*; inoltre, mantenendo i driver fuori dall'hypervisor, si ha che esso è più semplificato e leggero.<br/>
**Svantaggi**: il meccanismo di comunicazionee fra i due tipi di driver appesantisce l'accesso ai dispositivi.

#### Gestione delle Interruzioni
La gestione delle interruzioni viene virtualizzata in modo molto semplice: ogni interruzione viene gestita direttamente dal SO guest, eccezione fatta per la *page fault*, che richiede accesso al registro CR2, il quale contiene l'indirizzo che ha provocato il page fault. Poiché tale registro è accessibile solo a ring 0, la gestione del page fault deve coinvolgere il VMM: la routine di gesstione eseguita da XEN legge CR2, lo copia in una variabile nello spazio del SO ospitato, al quale viene restituito il controllo per poter gestire il page fault.

#### Migrazione Live
Il comando di migrazione viene eseguito da un demone di migrazione che si trova nel domain 0 del server di origine della macchina da migrare. La soluzione è basata sulla precopy e le pagine da migrare vengono compresse per ridurre l'occupazione di banda.


## 02 - Protezione
**Sicurezza**: riguarda l'insieme delle *tecniche per regolamentare l'accesso* degli utenti al sistema di elaborazione. La sicurezza impedisce accessi non autorizzati al sistema e i conseguenti tentativi dolosi di alterazione e distruzione dei dati. La sicurezza riguarda l'interfaccia del sistema verso il mondo esterno. Le tecnologie di sicurezza di un sistema informatico realizzano meccanismi per l'identificazione, l'autenticazione e l'autorizzazione di utenti "fidati".<br/>
**Protezione**: *insieme di attività volte a garantire il controllo dell'accesso* alle risorse logiche e fisiche da parte degli utenti autorizzati all'uso di un sistema di calcolo. Rispetto alla sicurezza ha un campo d'azione più interno al sistema. Per rendere un sistema "sicuro" è necessario stabilire per ogni utente autorizzato quali siano le risorse a cui può accedere e con quali operazioni può farlo. Ciò viene stabilito dal sistema di protezione tramite le tecniche di controllo degli accessi.

### Protezione: Modelli, Politiche e Meccanismi
In un sistema il controllo degli accessi si esprime tramite la definizione di tre livelli concettuali:
modelli, politiche e meccanismi.

#### Modelli
Un modello di protezione definisce i *soggetti*, gli *oggetti* ai quali i soggetti hanno accesso ed i *diritti* di accesso:
- **soggetti** - rappresentano la parte attiva di un sistema, ovvero le entità che possono richiedere l'accesso alle risorse. Ad esempio: gli utenti o i processi che eseguono per conto degli utenti;
- **oggetti** - costituiscono la parte passiva di un sistema, ovvero le risorse fisiche e logiche alle quali si può accedere e su cui si può operare. Ad esempio i file, o i processi intesi come risorsa (solitamente in un sistema di protezione i soggetti sono gli utenti, e i processo sono oggetti);
- **diritti di accesso** - sono le operazioni con le quali è possibile operare sugli oggetti. Ad esempio, in Linux i diritti di accesso sono lettura, scrittura, esecuzione.
```
NB: un soggetto può avere diritti di accesso anche per altri soggetti (es: processo che
controlla un altro processo.
```
Ad ogni soggetto è associato un dominio, che rappresenta l'*ambiente di protezione* nel quale il soggetto esegue e specifica i diritti di accesso posseduti da tale soggetto nei confronti di ogni risorsa.
Un dominio di protezione è unico per ogni soggetto, mentre un soggetto (ad esempio un processo) può cambiere dominio durante la sua esecuzione.

#### Politiche
Le politiche di protezione definiscono le regole con le quali i soggetti possono accedere agli oggetti. Mentre il modello è qualcosa di insito nel sistema, le politiche generalmente vengono scelte da chi opera su quel sistema. Si classificano in 3 diverse tipologie:
- **DAC** (Discretional Access Control) - il creatore di un oggetto controolla i diritti di accesso per quell'oggetto (tipologia adottata da UNIX, che fornisce un meccanismo per definire e interpretare per ciascun file i 3 bit di read, write ed execute, per il proprietario, il gruppo e gli altri). La definizione delle politiche è decentralizzata.
- **MAC** (Mandatory Access Control) - i diritti d'accesso vengono definiti in modo centralizzato. Questa soluzione viene utilizzata in sistemi di alta sicurezza per garantire assoluta confidenzialità e i diritti vengono gestiti da un'entità centrale.
- **RBAC** (Role Based Access Control) - ad un ruolo vengono assegnati specifici diritti di accesso sulle risorse. Gli utenti possono appartenere a diversi ruoli. I diritti attribuiti ad ogni ruolo vengono assegnati in modo centralizzato.
**Principio del Privilegio Minimo** (o POLA - Principle Of Least Authority): ad ogni soggetto sono garantiti i diritti di accesso solo agli oggetti strettamente necessari per la sua esecuzione. Questa è una caratteristica desiderabile per tutte le politiche di protezione.

#### Meccanismi
I meccanismi di protezione sono gli strumenti messi a disposizione dal sistema di protezione per imporre una determinata politica.
Principi di realizzazione:
- **flessibilità del sistema di protezione**: i meccanismi di protezione devono essere sufficientemente generali per consentire per consentire l'applicazione di diverse politiche di protezione;
- **separazione tra meccanismi e politiche**: la politica definisce *cosa va fatto* ed il meccanismo *come va fatto*. È desiderabile la massima indipendenza tra le due componenti.

### Dominio di Protezione
Un dominio definisce un insieme di coppie, ognuna contenente l'identificatore di un oggetto e l'insieme delle operazioni che il soggetto associato a tale dominio può eseguire su ciascun oggetto (diritti di accesso).

Es: D(S) = {\<o, diritti\> | o è un oggetto, diritti è un insieme di operazioni}

Ogni dominio è associato univocamente ad un soggetto, mentre un soggetto (ad esempio un processo) può eventualmente cambiare dominio durante la sua esecuzione; il soggetto può accedere solo agli oggetti definiti nel suo dominio, utilizzando i diritti specificati dal dominio.

Domini disgiunti o con diritti di accesso in comune: esiste la possibilità per due o più soggetti di effettuare alcune operazioni comuni su un oggetto condiviso. Le operazioni vengono svolte dai processi che operano per conto di soggetti (a cui sono associati i domini).
```
NB: in ogni istante della sua esecuzione, il processo esegue in uno ed un solo dominio.
```

#### Associazione tra Processo e Dominio
L'associazione tra processo e dominio può essere statica o dinamica.<br/>
**Statica**: l'insieme delle risorse disponibili ad un processo rimane fisso durante il suo tempo di vita. Osservazioni: questo tipo di associazione non è adatta al Principio del Privilegio Minimo, in quanto l'insieme globale delle risorse che un processo potrà usare può non essere un'informazione disponibile prima della sua esecuzione; inoltre, l'insieme minimo di risorse necessarie ad un processo per garantire tale Principio, può cambiare in modo dinamico durante l'esecuzione.<br/>
**Dinamica**: l'associazione tra processo e dominio varia durante l'esecuzione del processo. In questo modo si può mettere in pratica il *Principio del Privilegio Minimo*, in quanto in ogni sua fase di esecuzione il processo può acquisire diritti diversi (ovvero solo quelli strettamente necessari). Tuttavia in questo caso *occorre un meccanismo per consentire il passaggio da un dominio all'altro* del processo.

##### Esempio di cambio di dominio
**Standard dual mode** (kernel/user mode): ci sono due domini (ring) di protezione, quello dell'utente (user mode) e quello del kernel (monitor o kernel mode). Quando un processo deve eseguire un'istruzione privilegiata, chiama una system call ed avviene il cambio di dominio. Questo non realizza la protezione tra utenti, ma solo tra kernel e utente.

**UNIX**: il dominio è associato all'utente. Ogni *processo* è caratterizzato dall'attributo UserID (UID). Il cambio di dominio corrisponde al cambio temporaneo di identità (UID) del processo.
Ad ogni *file* invece sono associati il proprietario (user-id) ed un bit di dominio (set-uid). Se il un file ha il bit set-uid settato, quando un utente B, diverso dal proprietario A di tale file, ne lancia l'esecuzione, al processo che esegue viene assegnato lo user-id dell'utente B. Così facendo il file entra nel dominio di B.

### Matrice degli Accessi
Un sistema di protezione può essere rappresentato a livello astratto utilizzando il modello della matrice degli accessi: tale matrice consente di rappresentare il modello e le politiche valide nel sistema considerato.
- ogni riga rappresenta un soggetto (utente);
- ogni colonna rappresenta un oggetto (risorsa, file)
- ogni elemento rappresenta i diritti accordati ai soggetti sugli oggetti.
Le informazioni contenute nella matrice possono variare nel tempo: le informazioni contenute nella matrice all'istante t rappresentano lo stato di protezione del sistema in t. La matrice degli accessi offre ai meccanismi di protezione le informazioni che consentono di verificare il rispetto dei vincoli di accesso. Il meccanismo associato al modello:
- ha il compito di verificare se ogni richiesta di accesso che proviene da un processo (che opera in un determinato dominio) è consentita oppure no;
- autorizza l'esecuzione delle richieste consentite e impedisce quelle vietate;
- consente di modificare dinamicamente e in modo controllato il cambiamento dello stato di protezione.

#### Verifica del Rispetto dei Vincoli di Accesso
Il meccanismo consente di assicurare che un processo che opera nel dominio *Di* possa accedere solo agli oggetti specificati nella riga *i* e solo con i diritti di accesso indicati. Quando un'operazione M dev'essere eseguita nel dominio *Di* sull'oggetto *Oj*, il meccanismo consente di controllare che M sia contenuta nella casella *access(i,j)*. In caso affermativo l'operazione può essere eseguita, altrimenti viene restituito un errore.

#### Modifica dello Stato di Protezione
In base alla politica di protezione adottata, lo stato di protezione può essere modificato da entità differenti:
- nella politica DAC (es: sistemi UNIX-based) può essere modificato dai soggetti, ovvero gli utenti;
- nella politica MAC può essere fatto solo dall'entità centrale.

##### Modello Graham-Denning
Questo modello stabilisce quali sono i comandi che consentono una modifica dello stato di protezione, identificando 8 primitive:
- create object, aggiunge una colonna;
- delete object, rimuove una colonna;
- create subject, aggiunge una riga;
- delete subject, rimuove una riga;
- read access right, legge il diritto d'accesso;
- grant access right, assegna il diritto d'accesso;
- delete access right, rimuove il diritto d'accesso;
- transfer access right, propaga il diritto d'accesso.

##### Propagazione dei Diritti di Accesso (Copy Flag)
La possibilità di copiare un diritto di accesso per un oggetto da un dominio ad un altro nella matrice di accesso è indicata con un asterisco * (*copy flag*).
Un soggetto *Sa* può trasferire un diritto di accesso *a* (ad esempio 'read') per un oggetto *Ox* ad un altro soggetto *Sb* solo se *Sa* ha accesso a *Ox* con il diritto *a* e tale diritto ha il copy flag (ovvero solo se nella tabella delle matrici, l'elemento *Sa*\\*Ox* contiene *a*\*, ad esempio read*)

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/02%20-%20Protezione/Matrice%20degli%20Accessi%20Copy%20Flag.png" alt="Matrice degli Accessi e Copy Flag"/>

L'operazione di propagazione può essere realizzata in due modi:
- **trasferimento** del diritto, il soggetto iniziale perde il diritto di accesso, che viene spostato al nuovo soggetto;
- **copia** del diritto, il soggetto iniziale mantiene il diritto di accesso, "duplicandolo" al nuovo soggetto.
