<h1 align="center">Capitolo 3: Programmazione Concorrente</h1>

<details open="open">
  <summary><h2 style="display: inline-block">Indice</h2></summary>
  <a href="#03---programmazione-concorrente">Programmazione Concorrente</a>
  <ul>
    <li><a href="#cenni-storici">Cenni Storici</a></li>
		<li><a href="#tipi-di-architettura">Tipi di architettura</a></li>
		<li><a href="#classificazione-delle-architetture">Classificazione delle Architetture</a>
			<ul>
				<li><a href="#single-processor">Single Processor</a></li>
				<li><a href="#shared-memory-multiprocessors">Shared-Memory Multiprocessors</a></li>
				<li><a href="#distributed-memory">Distributed-Memory</a>
					<ul>
						<li><a href="#multicomputers">Multicomputers</a></li>
						<li><a href="#network-systems">Network Systems</a></li>
					</ul>
				</li>
			</ul>
		</li>
		<li><a href="#tipi-di-applicazioni">Tipi di Applicazioni</a></li>
		<li><a href="#processi-non-sequenziali-e-tipi-di-iterazione">Processi Non Sequenziali e Tipi di Iterazione</a>
			<ul>
				<li><a href="#processo-sequenziale">Processo Sequenziale</a></li>
				<li><a href="#processo-non-sequenziale">Processo Non Sequenziale</a>
					<ul>
						<li><a href="#elaboratore-non-sequenziale">Elaboratore Non Sequenziale</a></li>
						<li><a href="#linguaggi-concorrenti">Linguaggi Concorrenti</a></li>
					</ul>
				</li>
				<li><a href="#scomposizione-di-un-processo-non-sequenziale">Scomposizione di un Processo Non Sequenziale</a>
					<ul>
						<li><a href="#interazione-tra-processi">Interazione tra Processi</a>
							<ul>
								<li><a href="#cooperazione">Cooperazione</a></li>
								<li><a href="#competizione">Competizione</a></li>
								<li><a href="#interferenza">Interferenza</a></li>
							</ul>
						</li>
					</ul>
				</li>
			</ul>
		</li>
		<li><a href="#architetture-e-linguaggi-per-la-programmazione-concorrente">Architetture e Linguaggi per la Programmazione Concorrente</a></li>
		<li><a href="#architettura-di-una-macchina-concorrente">Architettura di una Macchina Concorrente</a>
			<ul>
				<li><a href="#architettura-della-macchina-m">Architettura della Macchina M</a></li>
			</ul>
		</li>
		<li><a href="#costrutti-linguistici-per-la-specifica-della-concorrenza">Costrutti Linguistici per la Specifica della Concorrenza</a>
			<ul>
				<li><a href="#forkjoin">Fork/Join</a></li>
				<li><a href="#cobegincoend">Cobeign/Coend</a></li>
			</ul>
		</li>
		<li><a href="#proprietà-dei-programmi">Proprietà dei Programmi</a>
			<ul>
				<li><a href="#verifica-della-correttezza-di-un-programma">Verifica della Correttezza di un Programma</a></li>
				<li><a href="#proprietà-di-safety-e-liveness">Proprietà di Safety e Liveness</a>
					<ul>
						<li><a href="#proprietà-dei-programmi-sequenziali">Proprietà dei Programmi Sequenziali</a></li>
						<li><a href="#proprietà-dei-programmi-concorrenti">Proprietà dei Programmi Concorrenti</a></li>
					</ul>
				</li>
			</ul>
		</li>
  </ul>
</details>

<!-- lezione 2021/10/05 -->
## 03 - Programmazione Concorrente
La *programmazione concorrente* è l'insieme delle tecniche, metodologie e strumenti per il support all'esecuzione di sistemi software composti da *insiemi di attività svolte simultaneamente*.

### Cenni Storici
La programmazione concorrente nasce negli anni '60, proprio nell'ambito dei Sistemi Operativi, quando ci fu l'introduzione dei canali o controllori di dispositivi (hardware): questi consentono l'esecuzione concorrente di operazioni nei dispositivi ed istruzioni nei programmi eseguiti dall'unità di elaborazione centrale.

L'interazione tra dispositivi ed unità centrale di elaborazione (processore) è basata fortemente sul meccanismo delle interruzioni (segnali di interrupt).
Quando la CPU riceve un segnale di interrupt dalla periferica, può tempestivamente gestiree quel particolare evento, che potrebbe essere ad esempio il trasferimento di dati.\
Questo meccanismo di interruzioni è stato poi importato ed utilizzato ampiamente in sistemi multiprogrammati time-sharing, in cui è impiegato il concetto di **quanto  di tempo** che consente di dividere equamente il tempo di CPU tra tutte le applicazioni in esecuzione su quel sistema/ambiente di esecuzione. Il modo per sancire il termine di un quanto di tempo assegnato ad un certo processo, che esegue un'applicazione, è ancora rappresentato dall'interruzione. Si ha lo scatto all'interruzione quando il quanto di tempo è esaurito, e dunque tempestivamente il Sistema Operativo si occupa di gestire il *cambio di contesto* tra un'applicazione e la successiva, secondo le politiche di scheduling che possiede.\
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

### Tipi di Architettura

#### Single Processor
Si ha un solo processore che possiede delle memorie ad accesso rapido (tipicamente 2 cache) ed una memoria primaria. Non sono necessari ulteriori layer di comunicazione con altre unità di calcolo, in quanto ne è presente solo una.

<img width="20%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Single%20Processor.png" alt="Single Processor"/>

#### Shared-Memory Multiprocessors
Si tratta di un'architettura costituita da diversi nodi, ciascuno dei quali ha una propria unità di calcolo (microprocessore) e delle memorie ad accesso rapido (cache). Ogni nodo ha la possibilità di accedere a qualunque parte della memoria, grazie alla **rete di interconnessione**. È il più comune al giorno d'oggi.

<img width="45%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Shared-Memory%20Multiprocessors.png" alt="Shared-Memory Multiprocessors"/>

Possiamo distinguere due modelli di sistemi multiprocessore:
**UMA (Uniform Memory Access)**: sistemi a multiprocessore con un numero ridotto di processori (da 2 a circa 30). Sono caratterizzati da un'interconnessione realizzata tipicamente da memory bus o crossbar switch; *tempo di accesso alla memoria uniforme* (indipendentemente dal processore e dalla cella di memoria da accedere, il tempo di accesso rimane costante); sono chiamati anche SMP (Symmetric MultiProcessors).\
**NUMA (Non Uniform Memory Access)**: sistemi con un numero elevato di processori (decine o centinaia). Sono caratterizzati da: memoria organizzata gerarchicamente, per evitare la congestione del bus; rete di interconnessione strutturata anch'essa in modo gerarchico (insieme di switch e memorie strutturato ad albero) ed ogniprocessore ha memorie più vicine ed altre più lontane; tempo di accesso dipendente dalla distanza tra processore e memoria (NUMA).

#### Distributed-Memory
Nelle architetture con memoria distribuita ogni processore accede alla propria memoria che non è condivisa tra i nodi di elaborazione. La memoria è quindi specifica del processore a cui è associata ed un'unità di elaborazione non può fare riferimento alla memoria di un altro nodo. In questo tipo di architettura i nodi possono essere singoli processori o multiprocessori a memoria condivisa.\
Rientrano in questa categoria i *Multicomputers* ed i *Network Systems*.

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Distributed-Memory.png" alt="Distributed-Memory"/>

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
La classificazione dei sistemi di calcolo più utilizzata è la *Tassonomia di Flynn (1972)*, in cui vengono inquadrate architetture e sistemi di elaborazione secondo due parametri:
1. **parallelismo a livello di istruzioni**
	- **Single Instruction Stream**, può essere eseguito un solo singolo flusso di istruzioni;
	- **Multiple Instruction Stream**, possono essere eseguiti più flussi di istruzioni in parallelo.
2. **parallelismo a livello di dati**
	- **Single Data Stream**, l'architettura è in grado di elaborare un singolo flusso sequenziale di dati;
	- **Multiple Data Streams**, l'architettura è in grado di processare più flussi di dati paralleli.

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Tassonomia%20di%20Flynn%20(1972)%20(1).png" alt="Tassonomia di Flynn (1972) (1)"/><img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Tassonomia%20di%20Flynn%20(1972)%20(2).png" alt="Tassonomia di Flynn (1972) (2)"/>

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
**Algoritmo**: procedimento logico che deve essere eseguito per risolvere un determinato problema. È ciò che succede quando mettiamo in esecuzione un programma

**Programma**: descrizione di un algoritmo mediante un opportuno formalismo (linguaggio di programmazione), che rende possibile l'esecuzione dell'algoritmo da parte di un particolare elaboratore.

**Processo**: insieme ordinato degli eventi cui dà luogo un elaboratore quando opera sotto il controllo di un programma.

**Elaboratore**: entità astratta realizzata in hardware e parzialmente in software, in grado di eseguire programmi (descritti in un dato linguaggio).

**Evento**: esecuzione di un'operazione tra quelle appartenenti all'insieme che l'elaboratore sa riconoscere ed eseguire. Ogni evento determina una transizione di stato dell'elaboratore.
```
NB: un programma descrive non un processo, ma un insieme di processi, ognuno dei quali
è relativo all'esecuzione del programma da parte dell'elaboratore per un determinato 
insieme di dati in ingresso.
```

#### Processo Sequenziale
Con *processo sequenziale* si intende il caso in cui l'insieme degli eventi che avvengono all'interno dell'elaboratore quando esegue un dato programma (l'insieme degli eventi che fanno parte dell'esecuzione prende il nome di "traccia del programma"), sia una vera e propria sequenza. Ovvero che gli eventi siano ordinati in modo sequenziale: per ogni evento, tranne il primo e l'ultimo, c'è sempre un solo evento che lo precede ed un solo evento che lo segue.

**Grafo di Precedenza**: è uno schema che permette di rappresentare, tramite un formalismo, la traccia del programma. Ogni nodo rappresenta un singolo evento durante l'esecuzione del programma, ogni arco rappresenta la *precedenza temporale* tra un nodo ed il successivo. Nel caso di un algoritmo strettamente sequenziale, il grafo di precedenza che lo rappresenta si dice ad **ordinamento totale** (qualunque coppia di nodi venga presa nel grafo, questa coppia è sempre ordinata).

<img width="60%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Esempio%20MCD%20(algoritmo).png" alt="Algoritmo MCD"/> <img width="11%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Esempio%20MCD%20(grafo).png" alt="Grafo MCD"/>

#### Processo Non Sequenziale
Con *processo non sequenziale* si intende il caso in cui l'insieme degli eventi che lo descrive è ordinato secondo una relazione d'ordine parziale. In altre parole, un processo si dice non sequenziale se il grafo di precedenza che lo descrive non è ordinato in modo totale, ma è caratterizzato da un **ordinamento parziale**.

<img width="40%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Esempio%20Elaborazione%20File%20(algoritmo).png" alt="Algoritmo Elaborazione File"/> <img width="9%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Esempio%20Elaborazione%20File%20(grafo).png" alt="Grafo Elaborazione File"/>

L'esecuzione di un processo non sequenziale richiede:
- innanzitutto che o a livello software o hardware l'*elaboratore* sia *non sequenziale*, ovvero ci dia la possibilità di eseguire operazioni simultanee;
- un *linguaggio di programmazione non sequenziale*.

###### Elaboratore Non Sequenziale
È in grado di eseguire più operazioni contemporaneamente e si hanno due possibilità:
- sistemei multielaboratori (a)
- sistemi monoelaboratori (b)

<img width="60%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Elaboratori%20Non%20Sequenziali.png" alt="Elaboratori Non Sequenziali"/>

###### Linguaggi Concorrenti
I linguaggi concorrenti (o non sequenziali) hanno la caratteristica comune di consentire, a livello di programma, la descrizione di un insieme di attività concorrenti, tramite moduli che possono essere eseguiti in parallelo (es: processi sequenziali).\
In generale, un linguaggio concorrente permette di esprimere il (potenziale) parallelismo nell'esecuzione di moduli differenti.

Tipicamente ci sono due modi in cui viene realizzato il modulo concorrente di un linguaggio:
- parallelismo espresso a livello di **singola istruzione**, oggi poco usato (es: CSP, Occam);
- parallelismo a livello di **sequenza di istruzioni**, molto più frequente (es: Java, Ada, Go, ...).

#### Scomposizione di un Processo Non Sequenziale
Se il linguaggio concorrente permette di esprimere il parallelismo a livello di sequenza di istruzioni, allora si può scomporre un processo non sequenziale in un insieme di processi sequenziali eseguiti contemporaneamente, e far fronte alla complessità di un algoritmo non sequenziale.\
Una volta noto l'algoritmo non sequenziale si tratta di ricavare dal suo grafo di precedenza una collezione di grafi di processi sequenziali, che chiaramente saranno legati fra di loro da vincoli di precedenza.\
Le attività rappresentate dai processi possono essere:
- **completamente indipententi**, se l'evoluzione del processo non influenza quella degli altri. Di fatto nel grafo abbiamo un unico punto di partenza ed un unico punto di arrivo, ma i nodi potrebbero esprimersi, ad esempio, come una serie di 3 sequenze di nodi, che non sono però legate fra loro da vincoli di precedenza (gli eventi che appartengono ad un processo non sono legati ad altri eventi appartenenti ad altri processi);
- **interagenti**, se sono assoggettati a vincoli di precedenza tra stati che appartengono a processi diversi (vincoli di precedenza fra le operazioni e vincoli di sincronizzazione).

<!-- lezione 2021/10/06 -->
##### Interazione tra Processi
Esistono tre possibili tipi di interazione tra processi: *cooperazione*, *competizione*, *interferenza*.

###### Cooperazione
Comprende tutte le intearazioni *prevedibili* e *desiderate*, che sono in qualche modo dettate dall'algoritmo (date cioè dagli archi del grafo di precedenza ad ordinamento parziale). È insita nella logica che vogliamo rappresentare. Si può esprimere in 2 modi: **segnali temporali**, ovvero sincronizzazione pura, che esprime solo ed unicamente un vincolo di precedenza; **scambio di dati**, ovvero comunicazione vera e propria. In entrambi i casi esiste comunque un vincolo di precedenza tra gli eventi di processi diversi.\
C'è una relazione di causa ed effetto tra l'esecuzione dell'operazione di invio da parte del processo mittente e l'operazione di ricezione da parte del processo ricevente, con un vincolo di precedenza tra questi eventi (*sincronizzazione* di due processi). Il linguaggio di programmazione deve fornire i costrutti linguistici necessari a specificare la sincronizzazione e la eventuale comunicazione tra i processi.\
Esempio di cooperazione: interazione data da vincoli temporali (es: un processo esegue delle operazioni ogni 2 secondi, un altro ogni 3 ed un terzo li coordina attivando periodicamente tali processi).

###### Competizione
Consiste in un'interazione *prevedibile* e *non desiderata* (in quanto non fa parte dell'algoritmo che si vuole implementare, ma è solitamente dato da un limite della risorsa fisica o logica), ma *necessaria*. Infatti, la macchina concorrente, su cui i processi sono eseguiti, mette a disposizione un numero limitato di risorse condivise, disponibili nell'ambiente di esecuzione. Poiché alcune di queste non possono essere accedute o utilizzate contemporaneamente da più processi (o lo sono solo per un numero limitato), è necessario prevedere meccanismi che regolino la competizione, coordinando l'accesso alla risorsa da parte dei vari processi, in modo **mutuamente esclusivo**. Questo può determinare l'imposizione di vincoli di sincronizzazione (se una risorsa può essere usata da un solo processo alla volta, nella fase in cui sta venendo usata da un certo processo, nessun altro deve poterla utilizzare): un processo che tenta di accedere una risorsa già occupata (se non rispetta certi vincoli) dev'essere bloccato.\
**Sezione critica**: indica una sequenza di istruzioni con cui un processo accede ad una risorsa condivisa mutuamente esclusiva. Ad una risorsa possono essere associate, in casi particolari, anche più di una sezione critica. Se su una risorsa vale la mutua esclusione, sezioni critiche appartenenti alla stessa classe non possono eseguire contemporaneamente.\
Esempio di competizione: processi che devono accedere ad una stampante (risorsa mutuamente esclusiva).

###### Interferenza
È un tipo di interazione *non prevista* e *non desiderata*. Solitamente è provocata da errori del programmatore (infatti solitamente si cerca di eliminarle o escluderle), il quale non ha modellato correttamente l'interazione dei propri processi non sequenziali interagenti.\
Può non manifestarsi, in quanto a volte dipende dalla velocità relativa dei processi; gli errori possono manifestarsi nel corso dell'esecuzione del programma, a seconda delle diverse condizioni di velocità di esecuzione dei processi. In questi casi si parla di errori dipendenti dal tempo.\
Esempio tipico: deadlock.

### Architetture e Linguaggi per la Programmazione Concorrente
Avendo a disposizione una *macchina concorrente* **M** (in grado di eseguire più processi sequenziali contemporaneamente) e di un *linguaggio di programmazione* con il quale descrivere algoritmi non sequenziali, è possibile scrivere e far eseguire programmi concorrenti. L'elaborazione complessiva può essere descritta come un insieme di *processi sequenziali interagenti*.\
Le **proprietà di un linguaggio di programmazione concorrente** sono:
- fornire appositi costrutti con i quali sia possibile dichiarare moduli di programma destinati ad essere eseguiti come processi sequenziali distinti;
- non tutti i processi vengono eseguiti contemporaneamente. Alcuni processi vengono svolti se, dinamicamente, si verificano particolari condizioni. È quindi necessario poter specificare quando un processo deve essere attivato e termianto;
- devono essere presenti strumenti linguistici per specificare le interazioni che dinamicamente possono verificarsi tra i vari processi.

### Architettura di una Macchina Concorrente
<img width="70%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Architettura%20Macchina%20Concorrente%20(1).png" alt="Architettura Macchina Concorrente (1)"/>
M offre un certo numero di unità di elaborazione virtuali, che però non sempre sono in numero sufficiente per supportare l'esecuzione contemporanea dei processi di un programma concorrente.\
M è una macchina astratta ottenuta tramite tecniche software (o hardware) basandosi su una macchina fisica M' generalmente più semplice (con un numero di unità di elaborazione solitamente minore del numero dei processi).\

<img width="60%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Architettura%20Macchina%20Concorrente%20(2).png" alt="Architettura Macchina Concorrente (2)"/>

Al proprio interno M contiene ciò che dev'essere messo in atto quando viene richiesta l'esecuzione di processi concorrenti e tutto ciò che riguarda l'interazione (sincronizzazione con scambio di informazioni).\
Il nucleo corrisponde al supporto a tempo di esecuzione del compilatore di un linguaggio concorrente e comprende sempre due funzionalità base:
- meccanismo di **multiprogrammazione**, preposto alla gestione delle unità di elaborazione della macchina M', ovvero le unità reali. Questo meccanismo è realizzato dal kernel del SO, il quale dà la possibilità ad ogni processo creato all'intero dell'ambiente, di avere una visione diversa, come se avesse una CPU completamente dedicata. Ciò permette ai vari processi eseguiti sulla macchina astratta M di condividere l'uso delle unità reali di elaborazione (tale virtualizzazione si basa sulle politiche di *scheduling*) tramite l'allocazione in modo esclusivo ad ogni processo di un'unità virtuale di elaborazione. Di fatto la macchina astratta M offre l'illusione che il sistema sia composto da tante unità di elaborazione, quanti siano i processi in esecuzione;
- meccanismo di **sincronizzazione** e **comunicazione**, estende le potenzialità delle unità reali di elaborazione, rendendo disponibile alle unità virtuali strumenti mediante i quali sincronizzarsi e comunicare.
Oltre ai meccanismi di multiprogrammazione e interazione, è presente anche il meccanismo di **protezione** (controllo degli accessi alle risorse): importante per rilevare eventuali interferenze tra i processi; può essere realizzato in hardware o software nel supporto a tempo di esecuzione; comprende capabilities e ACL.

#### Architettura della Macchina M
In base all'organizzazione logica di M vengono definiti due modelli di interazione tra i processi:
1. Modello a **memoria comune**, ovvero le macchine astratte M sono collegate ad un'unica memoria principale. La visione proposta è aderente al modello del *multiprocessore*. Se queste sono le caratteristiche della macchina astratta, le unità di elaborazione astratte/virtuali prevedono l'interazione dei processi tramite oggetti contenuti in memoria comune (modello ad ambiente globale).
2. Modello a **scambio di messaggi**, ovvero gli elaboratori astratti realizzati dalla macchina M non condividono memoria. Sono posti in collegamento da una rete di comunicazione, ma non hanno possibilità di accedere alle stesse aree di memoria (tipico dei sistemi *multicomputer*). Ciascuna di queste aree virtuali viene fornita ad un certo processo, e sarà compito della macchina M fornire dei meccanismi opportuni che consentano la comunicazione fra i processi che eseguono (modello ad ambiente locale).

### Costrutti Linguistici per la Specifica della Concorrenza
Qualunque siano le caratteristiche della macchina astratta, il linguaggio di programmazione (concorrente) deve fornire costrutti che consentano di gestire i processi.\
Esistono due modelli diversi:

#### Fork/Join
Questo modello comprende appunto due primitive fondamentali: *fork* e *join*.

**Fork**: permette di creare e attivare un processo che inizia la propria esecuzione in *parallelo* con quella del processo chiamante.
```
NB: non va confusa con la system call di UNIX: in questo caso riguarda un modello più
generale e, a differenza della primitiva UNIX, si passa una funzione, col codice da e-
seguire, alla fork.
```

La fork ha un comportamento simile ad una exec: mentre quest'ultima implica l'attivazione di un processo che esegue il programma chiamato e la sospensione del programma chiamante, la fork prevede che il programma chiamante prosegua contemporaneamente con l'esecuzione della funzione chiamata. Coincide infatti con una biforcazione del grafo.

**Join**: consente di sincronizzare un processo con la terminazione di un altro processo, precedentemente creato tramite una fork.

In un grafo di precedenza, il nodo che rappresenta l'evento join ha due predecessori.

```
NB: a differenza della wait UNIX, nella join è necessario specificare il processo da
attendere, mentre nella wait no, di conseguenza quest'ultima si mette in attesa della 
terminazione di uno qualunque dei processi figli.
```

#### Cobeign/Coend
Questo modello trae ispirazione dalla programmazione strutturata, permettendo di esprimere la concorrenza tramite opportuni blocchi da inserire nel codice di opportuni programmi concorrenti. Si basa su due primitive fondamentali: *cobegin* e *coend*.

**Cobegin**: specifica l'inizio di un blocco di codice che deve essere eseguito in parallelo. All'interno di questo blocco si possono specificare una serie di operazioni o processi: la caratteristica degli statement in questo blocco è che ognuno di essi verrà eseguito concorrentemente rispetto agli altri di tale blocco. Inoltre, è possibile innestare un blocco dentro l'altro. 

**Coend**: indica la fine di un blocco di istruzioni parallele.

*immagine dimostrativa con grafo di precedenza*

```
NB: fork/join è un formalismo più generale di cobegin/coend: tutti i grafi di preceden-
za possono essere espressi tramite fork/join ma non tutti possono essere espressi con 
cobegin/coend.
```

### Proprietà dei Programmi
I seguenti concetti permettono di specificare cosa succede quando il programma viene eseguito, di conseguenza sono utili per verificare la correttezza dei programmi realizzati.

**Traccia dell'esecuzione**: sequenza degli stati attraversati dal sistema di elaborazione durante l'esecuzione del programma. L'esecuzione di un programma è descritta dalla sua traccia.

**Stato**: insieme dei valori delle variabili definite nel programma più le variabili "implicite" (ad esempio il valore del program counter, o di altri registri).

#### Verifica della Correttezza di un Programma
**Programma sequenziale**: nei programmi sequenziali ogni esecuzione di un certo programma P su un particolare insieme di dati D genera sempre la stessa traccia (la verifica può essere svolta facilmente tramite debugging).\
**Programma concorrente**: nei programmi concorrenti l'esito dell'esecuzione dipende da quale sia l'effettiva sequenza cronologica di esecuzione delle istruzioni contenute, dunque ogni esecuzione di un certo programma P su un particolare insieme di dati D può dare origine a una traccia diversa, in quanto lo scheduling dei processi non è deterministico (la verifica è molto più difficile).

#### Proprietà di Safety e Liveness
**Proprietà di un programma**: attributo che è sempre vero, in ogni possibile traccia generata dalla sua esecuzione. Oltre alle proprietà di correttezza di un programma definite in precedenza, esistono anche altre proprietà, che solitamente si classificano in due categorie: *safety properties* e *liveness properties*.

**Safety**: garantisce che durante l'esecuzione di un programma *non si entrerà mai in uno stato "errato"*, ovvero in cui le variabili assumono valori non desiderati.

**Liveness**: garantisce che durante l'esecuzione del programma, *prima o poi si entrerà in uno stato "corretto"*, ovvero in cui le variabili assumono valori desiderati.

##### Proprietà dei Programmi Sequenziali
Le proprietà fondamentali che ogni programma sequenziale deve avere sono:
- *la correttezza del risultato finale*, ovvero che per ogni esecuzione, al termine del programma, il risultato ottenuto sia giusto -> **Safety**;
- *la terminazione*, ovvero prima o poi l'esecuzione del programma deve terminare -> **Liveness**.

<!-- lezione 2021/10/12 -->
prossima volta(?)

##### Proprietà dei Programmi Concorrenti
Le proprietà fondamentali che ogni programma concorrente deve avere sono:
- *correttezza del risultato finale* -> **Safety**;
- *terminazione*, -> **Liveness**;
- *mutua esclusione nell'accesso a risorse condivise*, ovvero per ogni esecuzione non accadrà mai che più di un processo acceda contemporaneamente alla stessa risorsa -> **Safety**;
- *assenza di deadlock*, ovvero per ogni esecuzione non si verificheranno mai situazioni di blocco critico -> **Safety**;
- *asseenza di starvation*, ovvero prima o poi ogni processo potrà accedere alle risorse richieste -> **Liveness**.
