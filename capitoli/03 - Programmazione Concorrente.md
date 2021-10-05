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
			<li><a href="#scomposizione-di-un-processo-non-sequenziale">Scomposizione di un Processo Non Sequenziale</a></li>
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
