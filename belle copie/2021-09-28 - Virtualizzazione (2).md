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

----
Come viene mappata ogni singola VM? Per motivi di efficienza si adotta il memory split
Memory split
Come viene mappata ogni singola VM? Si adotta un principio di memory split. Facciamo sempre il parallelo con sistema non virtualizzato/sistema virtualizzato:
in un sistema virtualizzato, ogni utilizzatore è una VM, quindi ogni entità che si interfaccia col VMM (equivalente del kernel) è una VM.
Così come accade nei sistemi non virtualizzati, in cui ogni processo ha un utilizzatore e un suo spazio di indirizzamento, anchenei sistemi virtualizzati ogni VM ha un suo spazio di indirizzamento virtuale (perché siamo in presenza di memoria virtuale).
Per motivi di efficienza, perché chiaramente nella commutazione tra una VM e l'altra c'è problema di reperire codice di XEN, nello spazio di indirizz di ogni VM viene riservata una parte in cui viene allocato XEN (in cui ogni volta che si commuta tra una VM e l'altra, non c'è bisogno di flushare la tabella delle pagine).

I guest si occupano della paginazione (gestiscono i page fault - quando c'è bisogno di stabilire cosa fare in seguito a un page fault, quello lo decide il pager del guest (politiche), quando c'è bisogno di scrivere nella PT di un certo processo interessato dal page fault, poiché il kernel non può scrivere nella memoria del VMM all'interno del quale sono allocate le PT, viene delegato l'hypervisor a farlo per la macchina virtuale).

Cosa succede quando viene creato un nuovo processo nello spazio del guest? Fra le altre cose dev'essere creata una PT associata a tale processo. ovviamente la creazione di una nuova PT non può essere fatta dal kernel che ospita quel processo (ring 1), quindi il guest richiede una nuova PT a hypervisor, il quale aggiunge alla PT anche lo spazio riservato a XEN; XEN quindi registra la tabella e naturalmente acquisisce il diritto di scrittura. Ad ogni successiva modifica da parte del guest, che tenterà di accedere alla struttura dati, ma questo darà origine ad una trap (protection fault), evento che verrà catturato e gestito dal VMM.


Politiche sopra (gestione della paginazione), meccanismi sotto (effettiva paginazione con allocazione delle pagine - ovvero della memoria).
Per com'è gestita la protezione in XEN, l'unica componente capace di allocare memoria è il VMM, ma può farlo solo in seguito a richeiste delle VM guest, in quanto come detto, le politiche si trovano in alto livello (ring 3), mentre i meccanismi a basso livello (ring 0). Poiché però può essere necessario al VMM, in alcuni casi, dover ottenere nuove pagine (es: attivazione nuova VM, dunque acquisire memoria necessaria per allocare lo spazio di indirizzamento di quella macchina virtuale). Questa possibilità, di richiedere pagine il VMM non ce l'ha. Può farlo solo in seguito a richieste da parte dei guest. Per farlo su XEN è stato adottata una soluzione, peculiare per la paravirtualizzazione, chiamata *balloon process*:
---


##### Paginazione: Balloon Process
La paginazione è a carico





La componente che si occupa della gestione della memoria virtuale è il pager (quando manca una pagina fa ciò che è necessario affinché quella pagina venga caricata in memoria e resa disponibile). All'interno del pager ci sono delle politiche: se la memoria è piena, il pager deve decidere con politiche sue proprie, di designare una vittima per far posto alla pagina da allocare in memoria centrale.
Problema: se avviene un page fault, viene notificato a livello hardware. Da un lato a livello alto ci sono le politiche che ne dovrebbero determinare il comportamento, ma il kernel della VM non si trova più a un livello privilegiato, quindi non ha la possibilità di cambiare la memoria e scrivere fisicamente il contenuto delle magine, in quanto questa prerogativa è compito esclusivamente del ring 0. Si ha dunque una tabella delle pagine per ogni processo, ma c'è il problema che il kernel della VM non può aggiornare direttamente il contenuto della tabella delle pagine, dunque in caso della necessità di aggiornamento della tabella delle pagine, bisogna delegare al VMM affinché vada ad aggiornarla.
Soluzione: si sfrutta il page fault




in un sistema virtualizzato ogni utilizzatore è una VM, ognuna ha un suo spazio di indirizzamento virtuale.

Per motivi di efficienza, nello spazio di indirizz della VM viene riservato uno spazio in cui viene allocato xen (ogni volta che si commuta una macchina virtuale dall'altra non è necessario fare un TBN flush.

la situazione è quella nella slide seguente (Protezione)

si ha XEN, poi il kernel della VM, poi lo spazio User dedicato ai processo


xen: gestione della memoria

quando c'è bisogno di scrivere nella page table di un processo interessato, poiché il kernel del SO non può scrivere nella memoria del VMM (nelle cosiddette shadow table), viene delegato appunto il VMM per farlo al posto della VM.


xen - creazione processo
Quando nello spazio del guest viene creato un nuovo processo, dev'essere anche creata una nuova tabella delle pagine. La creazione come detto, non può essere fatta dal kernel che ospita quel processo, sempre per motivi di protezione.
I guest richiedono una nuova TdP (tabella delle pagine) al VMM di xen, che registra la tabella, e acquisisce diritto di scrittura esclusiva su quella tabella. Ogni sucessiva modifica del guest, tenterà di accedere alla tabella delle pagine, generando un page fault (tipo trap), il VMM lo rileva/cattura l'evento e gestisce l'eccezione, eseguendo l'operazione richiesta, se corretta.


NB: politiche nei VM, meccanismi nel VMM

Poiché la protezione fa sì che l'unica entità capace di fare aggiornamenti in area di memoria è il VMM (incorpa una serie di meccanismi che vengono sempre guidati dalle politiche dei guest, che stanno sopra), ma la

La possibilità di richiedere pagine il VMM non ce l'ha, perché è qualcosa che è compito del VM (il VMM le crea e basta, ma non le crea da solo, lo fa dopo delle richieste da parte delle VM).
Per questo motivo esiste un processo chiamato "balloon process", che è sempre attivo, in ciascuna VM, che è in grado di dialogare direttamente con l'hypervisor. Questa è chiaramente un po' una violazione dei principi di virtualizzazione (idea di dare ad ogni VM un ambiente di esecuzione completamente identico, dando l'illusione di trovarsi su una macchina fisica, ma col balloon process, si ha un processo conscio di trovarsi su una VM). Cosa fa?
Il sistema di paginazione del SO di ogni guest è in grado di reperire pagine su richiesta.
Problema da risolvere: come dare la possibilità al VMM di acquisire nuove pagine in memoria

Il VMM chiede al balloon process di gonfiarsi (e di richiedere nuove pagine). Ogni volta che si interpella il balloon process acquisisce le pagine ed accede al VMM.

E' purtroppo l'unico modo, perché l'hypervisor non può gestire la paginazione, questa viene gestita dalle varie VM, da come sono strutturate le politiche di quel particolare SO.


se la VM si trova in uno stato di sospensione, esiste un tempo "Borrowed Virtual Time" che non va avanti

xen ha il compito di virtualizzare

front-end driver: quello che tipicamente viene installato


front-end/back-end driver asynchronous I/O rings

le richieste di accesso ad un particolare device vengono fatte dal guest attraverso il front end, il quale andrà a depositare la richiesta in una struttura dati a "ring", dall'altra parte c'è il back-end che va a prendere le richieste in modo FIFO dal ring.


scelta di xen di scorporare l'utilizzo dei driver tramite una struttura del genere (front-end/back-end) è per motivi di portabilità, per disaccoppiare la VM dall'architettura fisica su cui esegue.


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
































