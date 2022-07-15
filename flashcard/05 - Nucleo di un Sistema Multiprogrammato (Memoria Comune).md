[[Index](https://github.com/mikyll/Sistemi-Operativi-M/tree/main/flashcard)]&nbsp;&nbsp;
[[<<](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/01%20-%20Virtualizzazione.md)]
[[&nbsp;<&nbsp;](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/04%20-%20Modello%20a%20Memoria%20Comune.md)]
[[&nbsp;>&nbsp;](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/06%20-%20Modello%20a%20Scambio%20di%20Messaggi.md)]
[[>>](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/11%20-%20HPC.md)]

<h1 align="center">NUCLEO DI UN SISTEMA MULTIPROGRAMMATO (MEMORIA COMUNE)</h1>

### 1. Spiegare Cos'è il Nucleo e quali sono le sue Funzioni Fondamentali

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Il nucleo (o kernel) è il modulo realizzato in SW, HW o FW che supporta il concetto di processo e realizza gli strumenti necessari per la loro gestione e per la loro sincronizzazione. È il livello più interno di qualunque sistema basato su processi ed è l'unico conscio dell'esistenza delle interruzioni (sono invisibili ai processi).<br/>
  Caratteristiche fondamentali del nucleo:
  - **efficienza**, in quanto condiziona l'intera struttura a processi (alcuni sistemi prevedono l'esecuzione di operazioni del nucleo su hardware o tramite microprogrammi);
  - **dimensioni ridotte**, in quanto le funzioni richieste al nucleo sono estremamente semplici;
  - **separazione meccanismi e politiche**, il nucleo deve il più possibile contenere solo *meccanismi*, consentendo ai processi (in base ai meccanismi offerti dal nucleo) di scegliere ed applicare diverse politiche di gestione a seconda del tipo di applicazione.
  
  Stati di un processo (in un sistema in cui il numero di processi è maggiore del numero delle unità di elaborazione):
  - **esecuzione** (running), quando al processo è assegnata l'unità di elaborazione;
  - **pronto** (ready), quando al processo è revocata l'unità di elaborazione;
  - **bloccato** (waiting), quando il processo non è attivo (P sospensiva).
  Quando un processo perde il controllo del processore, il suo <ins>contesto</ins> (ovvero il *contenuto dei registri del processore*) viene salvato nel <ins>descrittore</ins> (un'*area di memoria associata al processo*).
  
  Le funzioni fondamentali del nucleo riguardano la gestione delle transizioni di stato dei processi, in particolare:
  1. Gestire il <ins>salvataggio ed il ripristino dei contesti dei processi</ins>, ovvero trasferire le informazioni dai registri della CPU al descrittore, quando esso passa dallo stato di esecuzione allo stato di pronto o bloccato.
  2. Effettuare lo <ins>scheduling della CPU</ins>, ovvero scegliere a quale processo assegnare l'unità di elaborazione.
  3. Gestire le <ins>interruzioni dei dispositivi</ins> esterni, traducendo i processi coinvolti da bloccati a pronti.
  4. Realizzare i <ins>meccanismi di sincronizzazione tra processi</ins>.
</details>

### 2. Spiegare Cos'è il Context Switch e Quali Funzioni deve Implementare il Nucleo per Realizzarlo

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Il cambio di contesto (context switch) è un'operazione realizzata dal kernel del SO, che permette a più processi di condividere la CPU in un sistema multitasking. In particolare, il kernel può essere suddiviso in 2 livelli:
  - *livello inferiore*, implementa i meccanismi per realizzare il cambio di contesto;
  - *livello superiore*, fornisce le funzioni (risposta alle interruzioni, primitive per creazione, eliminazione e sincronizzazione dei processi) direttamente utilizzabili dai processi.
  
  Il cambio di contesto permette di effettuare:
  - **Salvataggio_stato**, prevede il salvataggio del contesto del processo in esecuzione (informazioni contenute nei registri del processore) nel suo descrittore (area di memoria), e l'inserimento del descrittore nella coda dei processi bloccati o dei processi pronti.
  ```C
  void Salvataggio_stato() {
    int j;
    j = processo_in_esecuzione;
    descrittori[j].contesto = <valori_registri_CPU>;
  }
  ```
  - **Assegnazione_CPU**, prevede la rimozione del processo a maggior priorità dalla coda dei processi pronti e il caricamento dell'identificatore di tale processo nel registro contenente il processo in esecuzione.
  ```C
  void Assegnazione_CPU() {// scheduling: algoritmo con priorità
    int k = 0, j;
    while (coda_processi_pronti[k].primo == -1) { // -1 se l'elemento è l'ultimo (o la coda è vuota)
		  k++;
	}
    j = Prelievo(coda_processi_pronti[k]);
    processo_in_esecuzione = j;
  }
  ```
  - **Ripristino_stato**, prevede il caricamento del contesto del nuovo processo nei registri del processore.
  ```C
  void Ripristino_stato() {
    int j;
    j = processo_in_esecuzione;
    <registro-temp> = descrittori[j].servizio.delta_t;
    <registri-CPU> = descrittori[j].contesto;
  }
  ```
  Dunque il meccanismo di **cambio di contesto** si presenta come segue:
  ```C
  void Cambio_di_Contesto() {
	int j, k;
	Salvataggio_stato();
	j = processo_in_esecuzione;
	k = descrittori[j].servizio.priorità;
	Inserimento(j, coda_processi_pronti[k]);
	Assegnazione_CPU();
	Ripristino_stato();
  }
  ```
  NB: per consentire la modalità di servizio a divisione di tempo è necessario che il nucleo gestisca un *dispositivo temporizzatore*, che ad intervalli prefissati esegua il cambio di contesto.
</details>

### 3. Implementazione del Semaforo in Sistemi Monoprocessore

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Nel nucleo di un sistema monoprocessore, il semaforo può essere implementato tramite una *variabile intera* (≥0) ed un *puntatore ad una coda* (FIFO) di *descrittori di processi in attesa* sul semaforo. Ipotesi: gestione dei processi basata su priorità, ovvero al semaforo viene associato un insieme di code (una per priorità).
  ```C
  typedef struct {
	int contatore;
	coda_a_livelli coda;
  } descr_semaforo;
  
  descr_semaforo semafori[N_max_semafori];
  
  typedef int semaforo; // ID del semaforo nella lista 'semafori'
  
  void P(semaforo s) {
	int j, k;
	if (semafori[s].contatore == 0) {
		Salvataggio_stato();
		j = processo_in_esecuzione;
		k = descrittori[j].servizio.priorità;
		Inserimento(j, semafori[s].coda[k]);
		Assegnazione_CPU();
		Ripristino_stato();
	}
	else semafori[s].contatore--;
  }
  
  void V(semaforo s) {
	int j, k, p, q = 0; // j, k: processi; p, q: indici priorità
	while (semafori[s].coda[q].primo == -1 && q < min_priorità)
		q++;
	if (semafori[s].coda[q].primo != -1) {
		k = Prelievo(semafori[s].coda[q];
		j = processo_in_esecuzione;
		p = descrittori[j].servizio.priorità;
		if (p < q) // il processo in esecuzione è prioritario
			Inserimento(k, coda_processi_pronti[q]);
		else { // preemption
			Salvataggio_stato();
			Inserimento(j, coda_processi_pronti[p]);
			processo_in_esecuzione = k;
			Ripristino_stato();
		}
	}
	else semafori[s].contatore++;
  }
  ```
</details>

### 4. Spiegare le Possibili Realizzazioni del Nucleo in Sistemi Multiprocessore (SMP e Nuclei Distinti) e Confrontarle

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Il SO che esegue in un'architettura multiprocessore deve gestire una molteplicità di CPU, ognuna delle quali può accedere alla stessa memoria condivisa. Vi sono 2 modelli: *SMP* e a *Nuclei Distinti*.
  
  ##### Modello SMP
  Nel modello SMP (Symmetric Multi Processing) vi è un'<ins>unica copia del nucleo</ins> del Sistema Operativo, allocata <ins>nella memoria comune</ins>, che si occupa di tutte le risorse disponibili, comprese le CPU. Ogni processo può essere allocato su una qualunque CPU. È possibile che processi che eseguono su CPU diverse richiedano contemporaneamente funzioni del nucleo, ovvero vi è *competizione tra CPU* nell'esecuzione del nucleo, dunque vi è <ins>necessità di sincronizzazione</ins>.
  Soluzioni:
  - **Un solo lock**, ovvero viene associato al nucleo un lock per garantire la mutua esclusione delle funzioni del nucleo da parte di processi diversi. Tuttavia, in questo modo si <ins>limita il grado di parallelismo</ins>, escludendo a priori ogni possibilità di esecuzione contemporanea di funzioni del nucleo, che operano su strutture dati distinte (es: due semafori diversi).
  - **Lock multipli**, ovvero si individuano all'interno del nucleo diverse classi di sezioni critiche, ognuna associata ad una struttura dati separata e sufficientemente indipendente dalle altre (es: coda processi pronti, singoli semafori), e a ciascuna viene associato un lock distinto. In questo modo si <ins>incrementa il grado di parallelismo</ins>.

  Il modello SMP consente il <ins>load balancing</ins>, permettendo di <ins>schedulare equamente i processi su processori diversi</ins>. Tuttavia, in alcuni casi può essere vantaggioso assegnare un processo ad un determinato processore (usando la memoria privata del processore, in quanto se questa contiene già il codice del processo, il ripristino del contesto diventa più rapido), richiedendo però in questo caso una *coda dei processi pronti per nodo*, invece di una sola.
  
  ##### Modello a Nuclei Distinti
  Il modello a nuclei distinti prevede che vi siano più istanze del nucleo, raggruppate in una collezione, che eseguono in modo concorrente. Secondo questo modello, i processi che eseguono si possono dividere fra <ins>più nodi virtuali con poche interazioni reciproche</ins>. Ogni nodo virtuale è mappato su un nodo fisico (tutte le strutture del nucleo relative al *nodo virtuale* sono allocate nella memoria privata del nodo fisico) e tutte le interazioni locali ad un nodo virtuale possono essere eseguite indipendentemente e concorrentemente a quelle locali degli altri nodi. La memoria comune dell'architettura viene utilizzata solo per permettere a processi di nodi virtuali diversi di interagire.<br/>
  Nel modello a kernel distinti <ins>un processo può essere schedulato solo sul nodo contenente il relativo descrittore</ins>, rendendo impossibile l'attuazione di politiche di load balancing.
  
  ##### Confronto SMP e Nuclei Distinti
**SMP**:
- *Vantaggi*: permette una <ins>gestione ottimale delle risorse computazionali</ins>, in quanto consente il bilanciamento del carico fra le CPU dei vari nodi. Infatti, secondo questo modello lo scheduler può decidere su quale CPU (fra tutte) allocare un processo.
- *Svantaggi*: il grado di parallelismo tra CPU è sfavorito.

**Nuclei Distinti**:
- *Vantaggi*: favorisce il <ins>grado di parallelismo tra CPU</ins>, in quanto il grado di accoppiamento tra queste è minore. Ciò rende questo modello <ins>più scalabile</ins>.
- *Svantaggi*: vincola ogni processo ad essere schedulato sempre sullo stesso nodo, impedendo il bilanciamento di carico.
</details>

### 5. Implementazione del Semaforo in Sistemi Multiprocessore + Implementazione del Semaforo, delle Relative Operazioni ed il Meccanismo di Segnalazione tra i Nuclei nel caso di Context Switch
<!-- + Esempio di Interazione chiesto dalla prof -->

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  ##### Modello SMP
  In SMP i semafori sono realizzati proteggendo gli accessi ai contatori e alla coda dei processi pronti mediante lock. Se si usa un lock per ogni risorsa, due operazioni <ins>P su semafori diversi possono operare contemporaneamente solo se non sono sospensive</ins>, in quanto i semafori hanno *lock diversi*, ma la *coda dei processi pronti* è una risorsa *condivisa*, altrimenti devono operare in sequenza.<br/>
  Esempio: se vi è scheduling pre-emptive con priorità, una V può portare in esecuzione un processo con priorità superiore a quella di uno dei tanti in esecuzione (anche in altre CPU). Dunque, occorre che il nucleo revochi l'accesso alla CPU di uno di questi ultimi e la assegni al processo più prioritario appena riattivato. È quindi necessario che il nucleo mantenga l'informazione del processo a più bassa priorità in esecuzione e su quale CPU esso operi, rendendo inoltre necessario l'invio di interrupt HW alle varie CPU.
  
  ##### Modello a Nuclei Distinti
  Nel modello a Nuclei Distinti, poiché solo le interazioni tra processi appartenenti a nodi virtuali diversi utilizzano la memoria comune, si distinguono i semafori tra:
  - **semafori privati** di un nodo virtuale, utilizzati solo dai <ins>processi appartenenti al nodo</ins>;
  - **semafori condivisi** tra nodi virtuali, utilizzati da processi appartenenti a nodi diversi, e le cui <ins>informazioni</ins> sono contenute <ins>in memoria comune</ins>.
  
  Ogni semaforo condiviso è rappresentato come:
  - un <ins>intero in memoria comune</ins>, protetto da un lock;
  - una <ins>coda locale per ogni nodo</ins>, contenente i descrittori dei processi locali sospesi nel semaforo;
  - una <ins>coda globale di tutti i *rappresentanti* dei processi sospesi sul semaforo</ins> (il rappresentante di un processo identifica il nodo fisico su cui opera e il pid del processo).
  
  Una P sospensiva su un semaforo condiviso porta a inserire il rappresentante del processo chiamante nella coda globale ed il descrittore nella coda locale; una V, invece, estrae un processo dalla coda globale, ne comunica l'identità al nodo virtuale relativo (tramite interruzione, per garantire il rispetto della priorità), il quale risveglia il processo estraendo il descrittore dalla propria coda locale.
  
  Implementazione in pseudo-C:
  ```C
  void P(semaphore s) {
	if (is_private(s)) {
		// P come nel caso monoprocessore
	} else {
		lock(s.common_lock);
		// P
		// se necessario sospende il rappresentante nel processo in s.q
		unlock(s.common_lock);
	}
  }
  
  void V(semaphore s) {
	if (is_private(s)) {
		// P come nel caso monoprocessore
	} else {
		lock(s.common_lock);
		if (!empty(s.q)) {
			if (s.node == current_node) {
				// P come nel caso monoprocessore
			} else {
				// estrae p da s.q
				int ch = get_buffer(p.node);
				while (busy(ch)) {}
				send(ch, p.id);
				interrupt(p.cpu);
			}
		} else {
			p.c++;
		}
		unlock(s.common_lock);
	}
  }
  ```
</details>
