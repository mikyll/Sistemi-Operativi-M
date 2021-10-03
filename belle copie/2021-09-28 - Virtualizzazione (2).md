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

Domande:
-Migrazione live tramite pre-copy
-Descrivere architettura, paginazione, gestione delle interruzioni e dei driver di XEN. Cos'è un balloon process?

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

### XEN (Approfondimento)
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
Le VM devono poter accedere ai dispositivi che sono disponibili a livello hardware. La scelta di XEN è quella, ovviamente di virtualizzare l'interfaccia di ogni dispositivo, ma farlo tramite due tipi di driver: *back-end driver* e *front-end driver*.
**Back-end driver** è il driver vero e proprio, che permette, tramite un'interfaccia del VMM chiamata *Safe Hardware Interface*, di comunicare ed utilizzare il dispostivo collegato a livello hardware. Tipicamente viene installato all'interno di una VM particolare che è sempre ancorata al nodo fisico (dominio 0 - solitamente qui vengono installati tutti i driver di ogni dispositivo presente connesso a livello fisico in quel nodo).
**Front-end driver** è un driver "astratto", generico, non riferito adun dispositivo particolare, che viene installato tipicamente nel kernel del SO di una VM guest. Questo driver, all'occorrenza si collega al back-end driver specifico.
```
NB: non c'è niente che vieti di installare un back-end direttametne su una VM di domain U,
ma può convenire concentrarli tutti nel domain 0, sia perché siamo certi che quella macchi-
na non si sposterà mai da lì, essendo ancorata all'hardware, sia per motivi di portabilità.
```
*foto*

Ovviamente, per consentire la comunicazione tra back-end driver e front-end driver, serve un meccanismo che gestica le richieste. Questo viene realizzato tramite delle strutture chiamate asynchronous I/O rings (buffer FIFO circolari) in cui ogni elemento è una specie di descrittore che rappresenta una particolare richiesta. Le richieste di accesso ad un particolare device vengono fatte dal guest tramite il front-end che deposita la richiesta nel ring relativo, mentre dall'altra parte c'è il back-end che le preleva e le gestisce.

*foto*

**Vantaggi**: il driver viene scorporato in due parti, svingolando la VM dal particolare server fisico in cui risiede (il front-end driver della VM rimane lo stesso anche se questa viene spostata su un altro nodo), garantendo *portabilità*; inoltre, mantenendo i driver fuori dall'hypervisor, si ha che esso è più semplificato e leggero.
**Svantaggi**: il meccanismo di comunicazionee fra i due tipi di driver appesantisce l'accesso ai dispositivi.


#### Gestione Interruzioni



xen - gestione interruzioni





NUOVO PDF - Protezione nei SO

Sicurezza: riguarda l'interfaccia del Sistema verso il mondo esterno.

Protezione: campo d'azione più interno al sistema (insieme di tutte quelle attività volte al controllo dell'accesso alle risorse -sia fisiche che soprattutto logiche-) Una volta superato il controllo di sicurezza sono autorizzato al controllo del sistema all'interno del quale ci sono
Opportuno controllo su ciò che gli utenti possono o non possono fare all'interno del sistema


Protezione
necessario esercitare un opportuno controllo su quali operazioni l'utente può eseguire e a quali risorse esso possa accedere

modello di protezione stabilisce soggetti, oggetti, diritti

un soggetto può essere ad esempio un utente, un processo

oggetto: risorse fisiche e logiche a cui si può accedere (es: file, processo -il processo può essere sia oggetto che soggetto, in quanto può anche essere qualcosa su cui si possono avere dei diritti di accesso-)

es diritti di accesso linux: lettura, scrittura, esecuzione.

ad ogni soggetto è associato un dominio di protezione, unico per ogni soggetto.

in un sistema di protezione spesso i soggetti sono associati agli utenti, quindi di solito i processi sono oggetti.

su unix i processi hanno anche uno UID, ovvero l'id dell'utente che l'ha generato




politiche si classificano in diverse tipologie (3):
- DAC
- MAC
- RBAC

indipendentemente dal modello considerato e politiche adottate, in tutti i sistemi di protezione di solito si assume il cosiddetto principio del privilegio minimo (POLA - Principle Of Least Authority).
Principio secondo cui ad ogni soggetto devono essere garantiti i diritti d'accesso strettamente necessari per la sua esecuzione.


Meccanismi: strumenti messi a disposizione dal sistema di protezione per imporre una determinata politica.

dominio di un certo sogetto S, è formato a coppie: ogni sominio stabilisce per ogni oggetto l'insieme di diritti che il soggetto S può esercitare su un oggetto


[...]

Con associazione dinamica, un processo può passare da un dominio ad un'altro in base ai diritti che necessita per svolgere certe attività

In questo modo si riesce meglio a soddisfare il principio del privilegio minimo.







Esempio: UNIX





un sistema di protezione, a livello astratto, può essere rappresentato tramite una struttura detta matrice degli accessi, in cui ogni colonna è associata ad una diversa risorsa, ogni riga è associata ad un oggetto.

            O1              O2              O3
    +---------------+---------------+---------------+
S1  | read, write   | execute       | write         |
    +---------------+---------------+---------------+
S2  |               | execute       | read, write   |
    +---------------+---------------+---------------+



visto che

la tabella è un'immagine del sistema di protezione in un certo momento, in quanto lo stato può variare a runtime.































