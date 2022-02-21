[[Index](https://github.com/mikyll/Sistemi-Operativi-M/tree/main/flashcard)]&nbsp;&nbsp;
[[<<](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/01%20-%20Virtualizzazione.md)]
[[&nbsp;<&nbsp;](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/02%20-%20Protezione.md)]
[[&nbsp;>&nbsp;](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/04%20-%20Modello%20a%20Memoria%20Comune.md)]
[[>>](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/11%20-%20HPC.md)]

<h1 align="center">PROGRAMMAZIONE CONCORRENTE</h1>

### 1. Descrivere la Tassonomia di Flynn

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  La tassonomia di Flynn è la più usata classificazione dei sistemi di calcolo e si basa su 2 concetti: parallelismo a livello di istruzioni (Single Instruction stream, o Multiple Instruction stream) e parallelismo a livello di dati (Single Data stream o Multiple Data stream):
  
  - **Single Instruction, Single Data (SISD)**, riguarda gli elaboratori monoprocessore (es: macchina di Von Neumann);
  - **Single Instruction, Multiple Data (SIMD)**, prevede molte unità di elaborazione che eseguono la stessa istruzione su una moltitudine di dati differenti (es: elaboratori vettoriali: GPU);
  - **Multiple Instruction, Single Data (MISD)**, il sistema è in grado di gestire un unico flusso di dati che ad ogni istante può essere elaborato da più istruzioni differenti (es: pipelined computer);
  - **Multiple Instruction, Multiple Data (MIMD)**, insieme di nodi di elaborazione ognuno dei quali può eseguire flussi di istruzioni diverse, su dati diversi (es: multiprocessori).
</details>

### 2. Descrivere le Possibili Interazioni tra Processi

<details>
  <summary><b>Visualizza risposta</b></summary>
  
Esistono 3 possibili tipi di interazione fra processi:
1. **Cooperazione**, comprende tutte le interazioni <ins>prevedibili e desiderate</ins>, che sono in qualche modo dettate dall'algoritmo (ovvero date dagli archi del grafo di precedenza ad ordinamento parziale). Si può esprimere in 2 modi, entrambi dei quali esprimono un *vincolo di precedenza*:
    - mediante <ins>segnali temporali</ins>, ovvero pura sincronizzazione;
    - mediante <ins>scambio di dati</ins>, ovvero con comunicazione.
2. **Competizione**, consiste in un'interazione <ins>prevedibile ma non desiderata</ins>, in quanto non fa parte dell'algoritmo, ma è imposta dai limiti delle risorse a cui i processi devono accedere, ad esempio una risorsa che può essere acceduta solo in modo mutuamente esclusivo. In questo caso si prevede il concetto di *sezione critica*, ovvero la sequenza di istruzioni con cui un processo accede ad un oggetto condiviso mutuamente esclusivo. Ad una risorsa possono essere associate anche più di una sezione critica, di classi differenti.
3. **Interferenza**, consiste in un'interazione <ins>non prevedibile e non desiderata</ins> solitamente causata da *errori del programmatore* (es: deadlock).
</details>

### 3. Costrutti Linguistici per la Specifica della Concorrenza

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Il linguaggio concorrente deve fornire costrutti che consentano di gestire i processi. Esistono 2 modelli differenti:
  - **Fork/Join**, comprende una primitiva <ins>fork</ins> per la *creazione* e l'*attivazione* di un processo che eseguirà in parallelo, ed una primitiva <ins>join</ins> per la sincronizzazione con la terminazione di un processo. Nel grafo di precedenza, una fork coincide con una biforcazione, mentre una join con un nodo avente due precedenti.
  - **Cobegin/Coend**, comprende una primitiva <ins>cobegin</ins> per la specifica di un *blocco di codice che deve essere eseguito in parallelo*, ed una primitiva <ins>coend</ins> per la specifica della terminazione del blocco. Le istruzioni contenute all'interno vengono eseguite in parallelo ed è possibile innestare dei blocchi uno dentro l'altro.
</details>

### 4. Proprietà dei Programmi Sequenziali e Concorrenti

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Una delle attività più importanti per chi sviluppa programmi è la verifica di correttezza dei programmi realizzati.
  
  Proprietà dei Programmi *Sequenziali*:
  1. **Safety**, ovvero la correttezza del risultato finale (il programma non entrerà mai in uno stato in cui le variabili assumono valori non desiderati).
  2. **Liveness**, ovvero la terminazione del programma (prima o poi il programma entrerà in uno stato in cui le variabili assumono valori desiderati).
  
  Proprietà dei Programmi *Non Sequenziali*:
  1. **Safety**, correttezza del risultato finale.
  2. **Liveness**, terminazione del programma.
  3. **Mutua Esclusione nell'Accesso a Risorse Condivise**, ovvero per ogni esecuzione non si deve mai verificare che più di un processo acceda contemporaneamente ad una stessa risorsa (mutuamente esclusiva).
  4. **Assenza di Deadlock**.
  5. **Assenza di Starvation**, ovvero ciascun processo che richiede l'accesso ad una certa risorsa, prima o poi lo otterrà.
</details>
