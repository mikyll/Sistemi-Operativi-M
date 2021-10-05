<!-- lezione 2021/10/05 -->

## 03 - Programmazione Concorrente
La *programmazione concorrente* è l'insieme delle tecniche, metodologie e strumenti per il support all'esecuzione di sistemi software composti da *insiemi di attività svolte simultaneamente*.

### Cenni Storici
La programmazione concorrente nasce negli anni '60, proprio nell'ambito dei Sistemi Operativi, quando ci fu l'introduzione dei canali o controllori di dispositivi (hardware): questi consentono l'esecuzione concorrente di operazioni nei dispositivi ed istruzioni nei programmi eseguiti dall'unità di elaborazione centrale.

L'interazione tra dispositivi ed unità centrale di elaborazione (processore) è basata fortemente sul meccanismo delle interruzioni (segnali di interrupt).
Quando la CPU riceve un segnale di interrupt dalla periferica, può tempestivamente gestiree quel particolare evento, che potrebbe essere ad esempio il trasferimento di dati.
Questo meccanismo di interruzioni è stato poi importato ed utilizzato ampiamente in sistemi multiprogrammati time-sharing, in cui è impiegato il concetto di **quanto  di tempo** che consente di dividere equamente il tempo di CPU tra tutte le applicazioni in esecuzione su quel sistema/ambiente di esecuzione. Il modo per sancire il termine di un quanto di tempo assegnato ad un certo processo, che esegue un'applicazione, è ancora rappresentato dall'interruzione. Si ha lo scatto all'interruzione quando il quanto di tempo è esaurito, e dunque tempestivamente il Sistema Operativo si occupa di gestire il *cambio di contesto* tra un'applicazione e la successiva, secondo le politiche di scheduling che possiede.
Le interruzioni possono accadere ad istanti impredicibili, dunque in un sistema time-sharing parti di programmi possono essere eseguite in modo non predicibile. Infatti, una delle principali caratteristiche delle applicazioni concorrenti è il *non determinismo*: lo stesso programma eseguito in tempi diversi può comportare risultati diversi anche se il codice non cambia. Questo, ad esempio, si può rilevare quando cci sono parti di programmi che condividono le stesse variabili comuni: in questi casi, se non viene sincronizzato l'accesso a tali variabili, si possono creare delle interferenze.

Successivamente sono stati introdotti i sistemi multiprocessore, ovvero con più unità di elaborazione (parallelismo supportato a livello hardware). Se prima il parallelismo era puramente virtuale, con tali architetture il parallelismo era diventato effettivamente "reale", in quanto si potevano avere fisicamente diversi microprocessori che lavoravano in modo concorrente.
Ciò ha comportato diversi vantaggi, soprattutto in termini di prestazioni: in particolare, vengono abbattuti i tempi di esecuzione.

In un sistema concorrente i principali problemi sono:
- con quale criterio modellare l'applicazione concorrente;
- come suddividerla in attività concorrenti (quanti processi utilizzare);
- come garantire la corretta sincronizzazione delle loro operazioni (in generale le attività nelle quali si scompone l'applicaczione possono aver bisogno di interagire fra di loro, dunque è necessario imporre dei vincoli di precedenza).
Queste decisioni dipendono da:
- tipo di architettura hardware;
- tipo di applicazione.

### Tipi di architettura

#### Single Processor
Si ha un solo processore che possiede delle memorie ad accesso rapido (tipicamente 2 cache) ed una memoria primaria. Non sono necessari ulteriori layer di comunicazione con altre unità di calcolo, in quanto ne è presente solo una.

*foto slide 6*

#### Shared-Memory Multiprocessors
Si tratta di un'architettura costituita da diversi nodi, ciascuno dei quali ha una propria unità di calcolo (microprocessore) e delle memorie ad accesso rapido (cache). Ogni nodo ha la possibilità di accedere a qualunque parte della memoria, grazie alla **rete di interconnessione**. È il più comune al giorno d'oggi.

*foto*

Possiamo distinguere due modelli di sistemi multiprocessore:
**UMA (Uniform Memory Access)**: sistemi a multiprocessore con un numero ridotto di processori (da 2 a circa 30). Sono caratterizzati da un'interconnessione realizzata tipicamente da memory bus o crossbar switch; *tempo di accesso alla memoria uniforme* (indipendentemente dal processore e dalla cella di memoria da accedere, il tempo di accesso rimane costante); sono chiamati anche SMP (Symmetric MultiProcessors).\
**NUMA (Non Uniform Memory Access)**: sistemi con un numero elevato di processori (decine o centinaia). Sono caratterizzati da: memoria organizzata gerarchicamente, per evitare la congestione del bus; rete di interconnessione strutturata anch'essa in modo gerarchico (insieme di switch e memorie strutturato ad albero) ed ogniprocessore ha memorie più vicine ed altre più lontane; tempo di accesso dipendente dalla distanza tra processore e memoria (NUMA).

#### Distributed-Memory
Nelle architetture con memoria distribuita ogni processore accede alla propria memoria che non è condivisa tra i nodi di elaborazione. La memoria è quindi specifica del processore a cui è associata ed un'unità di elaborazione non può fare riferimento alla memoria di un altro nodo. In questo tipo di architettura i nodi possono essere singoli processori o multiprocessori a memoria condivisa.\
Rientrano in questa categoria i *Multicomputers* ed i *Network Systems*.

*foto*

##### Multicomputers
Modello in cui i nodi e la rete sono *fisicamente vicini*, ovvero nella stessa struttura fisica. La rete di interconnessione offre un cammino di comunicazione tra i processi ad alta velocità e larghezza di banda. Ad esempio i Cluster ed i sistemi ad alto parallelismo (HPC). I multicomputer sono fatti per essere aggregati in una stessa struttura fisica.
```
NB: un Cluster of Computers (CoW), un insieme di nodi, tipicamente chiamati server, fi-
sicamente vicini, in cui ogni nodo è una scheda inserita in una struttura fisica, detta
"rack", dove solitamente la rete di interconnessione è una linea ad alta velocità e 
con larghezza di banda sufficientemente ampia. 
```

##### Network Systems
Sistemi in cui i nodi sono collegati da una rete locale (es: Ethernet) o geografica (es: Internet).

### Classificazione delle Architetture
La classificazione dei sistemi di calcolo più utilizzata è la *tassonomia di Flynn (1972)*, in cui vengono inquadrate architetture e sistemi di elaborazione secondo due parametri:
1. **parallelismo a livello di istruzioni**
	- **Single Instruction Stream**, può essere eseguito un solo singolo flusso di istruzioni;
	- **Multiple Instruction Stream**, possono essere eseguiti più flussi di istruzioni in parallelo.
2. **parallelismo a livello di dati**
	- **Single Data Stream**, l'architettura è in grado di elaborare un singolo flusso sequenziale di dati;
	- **Multiple Data Streams**, l'architettura è in grado di processare più flussi di dati paralleli.

*schemino* + *foto slide 14?*

**SISD - Single Instruction (stream), Single Data (stream)**: sistemi monoprocessore che fanno riferimento all'architettura classica della macchina di Von Newman. Come dice il nome è in grado di gestire un singolo flusso di istruzioni (un programma) alla volta, su un singolo flusso di dati.

**SIMD - Single Instruction, Multiple Data**: architetture tipicamente parallele in cui vi sono diversi processori che, ad ogni istante, possono eseguire la stessa singola istruzione ma su dati diversi. Ad esempio rientrano in questa categoria gli array processors, di cui fanno parte anche le GPU.
```
NB: le GPU sono costituite da un insieme di nodi di elaborazione, a cui è assegnato una
singola control unit. Poiché elaborano dati che sono rappresentati da grandi matrici 
di informazioni (elaborazione di immagini), il modello SIMD risulta particolarmente ef-
ficace.
```
Rientrano in questa categoria anche i vector processors (migliaia di unità di elaborazione, non troppo potenti, ma che messe insieme e se controllate opportunamente, possono risolvere particolari classi di problemi in modo piuttosto efficiente e veloce).

**MIND - Multiple Instruction, Multiple Data**: insieme di nodi di elaborazione ognuno dei quali può eseguire flussi di istruzioni diverse su dati diversi. Ogni nodo può essere utilizzato da un processo che svolge operazioni diverse su dati differenti. Rientrano in questa categoria i sistemi multiprocessore (quelli che probabilmente conosciamo meglio), ma anche i MultiComputers.

**MISD - Multiple Instruction, Single Data**: il sistema è in grado di gestire un unico flusso di dati che ad ogni istante può essere elaborato con molteplici flussi di istruzioni. Non ci sono esempi particolarmente significativi da portare, ma è il caso dei "pipelined computer", dove lee diverse unità di elaborazione sono messe in cascata (pipeline), che lavora su quel flusso di dati, ognuna facendo qualcosa di differente.

### Tipi di Applicazioni
Ricapitolando, il progetto di applicazioni concorrenti dev'essere sviluppato in base al tipo di architettura, ma anche in base ai vincoli dati dal Sistema Operativo.

1. **multithreaded**:
	- si ha un'applicazione strutturata come un insieme di processi (thread) che:
		- permette di dominare la complessità del problema da risolvere;
		- aumentare l'efficienza, in quanto il carico di lavoro viene "scaricato" in parallelo;
		- semplificare la programmazione (secondo un modello di scomposizione dell'algoritmo in più parti che possono procedere contemporaneamente).
	- i processi possono condividere variabili;
	- sono caratterizzati dal fatto che generalmente esistono più processi che processori;
	- i processi sono schedulati ed eseguiti indipendentemente.
2. **sistemi multitasking/sistemi distribuiti**:
	- le componenti dell'applicazione (task) vengono eseguite su nodi (eventualmente virtuali) collegati tramite opportuni mezzi di interconnessione (es: canali);
	- i processi non possono condividere variabili, infatti comunicano scambiandosi messaggi;
	- questa organizzazione è tipica del modello client/server.
	I componenti in un sistema distribuito sono spesso multithreaded.
```
NB: in certi ambiti (sistemi distribuiti) esistono anche sistemi ibridi di applicazioni
in cui alcune parti sono multithreaded, mentre altre interagiscono a scambio di messag-
gio.
```
3. **applicazioni parallele**:
	- possiamo avere sia un modello in cui i processi condividono memoria, sia un modello a scambio di emssaggi;
	- hanno l'obbiettivo di risolvere il problema dato nel modo più veloce possibile, oppure un problema di dimensioni più grandi nello stesso tempo, sfruttando efficacemente il parallelismo disponibile a livello hardware;
	- sono eseguite su sistemi paralleli (es: HPC, array processors), facendo uso di algoritmi paralleli;
	- a seconda del modello architetturale, l'esecuzione è portata avanti da istruzioni/thread/processi paralleli che interagiscono utilizzando librerie specifiche.

### Processi Non Sequenziali e Tipi di Iterazione



ricontrollare i NB