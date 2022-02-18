[[Index](https://github.com/mikyll/Sistemi-Operativi-M/tree/main/flashcard)]&nbsp;&nbsp;
[[<<](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/01%20-%20Virtualizzazione.md)]
[[&nbsp;<&nbsp;](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/10%20-%20Algoritmi%20di%20Sincronizzazione%20Distribuiti.md)]
[&nbsp;>&nbsp;]
[>>]

<h1 align="center">HIGH PERFORMANCE COMPUTING (HPC)</h1>

### 1. Evoluzione delle Architetture fino al Modello Distribuito

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  La principale motivazione per lo sviluppo delle tecnologie HW e SW per il parallel computing è l'<ins>aumento di performance</ins>: risolvere problemi di complessità elevata in tempi contenuti; risolvere gli stessi problemi in tempi più bassi.
  
  **Evoluzione delle Architetture**: fino ai primi anni 2000 l'evoluzione dei sistemi di calcolo è stata "governata" dalla *Legge di Moore*, secondo cui le performance dei processori crescono costantemente (raddoppiando la densità di transistor all'interno dei chip ogni 18 mesi, con conseguente aumento di *capacità di elaborazione* del chip e aumento della *velocità di calcolo*). A partire dai primi anni 2000, ci si è trovati sempre più prossimi ai <ins>limiti fisici</ins> dei componenti: a causa dell'*effetto joule* (lega la produzione di calore al passaggio di corrente elettrica nei circuiti integrati), non è stato più possibile ad esempio aumentare la frequenza di clock, rendendo <ins>necessario l'aumento di capacità di calcolo a parità di frequenza</ins>. Questo obbiettivo è stato raggiunto grazie all'introduzione di diverse forme di parallelismo a livello HW. Infatti, se l'HW è in grado di svolgere più operazioni per ciclo, la velocità di elaborazione dell'intero sistema aumenta (più processori su singolo chip, più processori su più chip).
</details>

### 2. Cos'è il Von Neumann Bottleneck e Come si può Mitigare. Quali Architetture della Tassonomia di Flynn possono superare questo problema? 

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Il modello di Von Neumann descrive lo schema funzionale di un tradizionale sistema sequenziale. L'unica CPU è collegata alla memoria centrale da un mezzo di interconnessione (es: bus), e questa separazione costituisce una limitazione nella velocità di accesso a dati e istruzioni, che influisce sulla velocità di elaborazione del sistema.
  
  **Von Neumann Bottleneck**: il <ins>bus che collega CPU e memoria centrale</ins> costituisce un collo di bottiglia. Infatti, questo limita la velocità di fetching di istruzioni e dati, che dipendono dalla velocità del bus, e limita conseguentemente la velocità di esecuzione.<br/>
  Per mitigare questo problema, è stato introdotto l'utilizzo di:
  - memorie *cache*;
  - *parallelismo di basso livello* (*ILP* e *HW multithreading*).
  
  **Cache**: è una memoria *associativa* ad <ins>accesso veloce</ins>, in quanto risiede sul chip del processore e si colloca ad un livello intermedio tra memoria centrale e registri del processore; di <ins>capacità limitata</ins>, in quanto non può contenere tutte le istruzioni ed i dati necessari al programma in esecuzione.<br/>Viene gestita con criteri basati sul [*Principio di Località*](https://it.wikipedia.org/wiki/Principio_di_localit%C3%A0_(informatica)), secondo cui: "durante l'esecuzione di una data istruzione presente in memoria, con molta probabilità le successive istruzioni saranno ubicate nelle vicinanze di quella in corso" (località spaziale e/o temporale).<br/>
  Dunque, si potranno avere dei *cache hit*, se l'informazione richiesta è presente in cache, oppure *cache miss*, se non è presente e va caricata dalla memoria centrale. Se la gestione è tale da mantenere un hit-rate sufficientemente elevato, gli effetti del Von Neumann Bottleneck possono essere mitigati.
  
  **Parallelismo di Basso Livello (ILP)**: le istruzioni per essere eseguite seguono una sequenza di fasi (fetching operandi, confronto esponenti e/o shift, somma, normalizzazione risultato e memorizzazione del risultato). Ciascuna di queste può essere separata ed affidata ad un'<ins>unità funzionale</ins> indipendente che opera in parallelo alle altre. Le unità funzionali sono collegate tra di loro mediante una <ins>pipeline</ins> <br/>
  Problema: non sempre questa operazione è fattibile, ad esempio se in un programma è presente una lunga serie di istruzioni tra loro dipendenti (tipo la Serie di Fibonacci).
  
  **HW Multithreading**: i processori moderni offrono parallelismo di alto livello (a livello di thread), mediante HW multithreading, che permette a più thread di condividere la stessa CPU usando una <ins>tecnica di sovrapposizione</ins> (duplicazione registri e context switch efficiente con supporto HW). Esistono 2 approcci:
  - **multithreading a grana fine** (fine-grained), secondo cui viene eseguito <ins>un context switch dopo ogni istruzione</ins>.
	- *Vantaggio*: velocità thread bassa;
	- *Svantaggio*: throughtput alto (ovvero si trasmettono più dati);
  - **multithreading a grana grossa** (coarse-grained), secondo cui <ins>il context switch avviene quando il thread corrente si trova in attesa</ins> (es: attesa del caricamento di informazioni dalla memoria centrale in seguito ad un cache miss).
	- *Vantaggio*: velocità thread alta;
	- *Svantaggio*: throughtput basso.
  
  ILP e HW Multithreading hanno permesso un miglioramento delle prestazioni dei processori, tuttavia tali meccanismi sono trasparenti ai programmatori (Modello *Von Neumann Esteso*). Nei sistemi HPC, invece, il parallelismo disponibile è visibile ai programmatori, che progetta il software sfruttando al meglio le risorse computazionali: **architetture non Von Neumann**.
  
  Nella *Tassonomia di Flynn*, i sistemi HPC riguardano le classi SIMD e MIMD. In particolare le architetture MIMD prevedono l'asincronicità delle attività nei diversi nodi, permettendo ad ogni CPU di eseguire una sequenza di istruzioni diversa dagli altri nodi. Sistemi HPC si dividono in due modelli: a Shared Memory o Distributed Memory. La maggior parte dei sistemi HPC al giorno d'oggi presenta un <ins>modello ibrido</ins> che combina il modello a *memoria distribuita* col modello a *memoria comune*.
</details>

### 3. Quali sono le Metriche per Valutare le Prestazioni di Applicazioni Parallele. Qual è la Differenza Tra Scalabilità Weak e Scalabilità Strong

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Per misurare le prestazioni di un sistema HPC si utilizza l'unità di misura del *FLOPS* (FLoating-point Operations Per Second, "operazioni in virgola mobile al secondo").<br/>
  Per valutare il vantaggio derivante dall'esecuzione di programmi paralleli in sistemi HPC si utilizzano alcune metriche: *speedup* ed *efficienza*.
  
  Lo **Speedup** misura quanto è più veloce la versione parallela rispetto alla versione sequenziale (esprime il <ins>guadagno di un'applicazione parallela rispetto alla versione sequenziale</ins>). È pari a ```S = Tseq / Tpar```, dove *Tseq* è il tempo di esecuzione del programma nella versione sequenziale (su un solo nodo), e *Tpar* nella sua versione parallela.<br/>
  Il caso ideale è che ```S = p```, dove *p* è il numero di processori. Tuttavia, solitamente ci sono altri fattori da considerare nell'equazione, quali ad esempio lo scarto di tempo dovuto all'*overhead*, pertanto in generale si ha che ```S < p```.
  
  L'**Efficienza** misura lo <ins>speedup per numero di processori utilizzati</ins>. È pari a ```E = S / p```, dove *S* è lo speedup, e *p* il numero di processori utilizzati. Il caso ideale è che ```E = 1```, mentre nei casi reali si ha che ```E < 1```.
  
  Un sistema si dice **scalabile** se mantiene la stessa efficienza al variare del numero di processori utilizzati e/o al variare della quantità di dati da elaborare.
  
  La **Legge di Amdahl** considera che <ins>in generale non tutto il programma può essere parallelizzabile</ins>, dunque *Tpar* è dato da ```Tpar = r * Tseq + (1 - r) * Tseq / p```, dove *r* ∈ [0, 1] è una percentuale che esprime la frazione di tempo totale di esecuzione speso nella parte *non parallelizzabile* del programma. La Legge di Amdahl esprime lo speedup *S* come: ```S = Tseq / Tpar = 1 / (r + (1 - r) / p)``` e descrive l'andamento dello speedup al variare del numero di processori impiegati per la soluzione dello stesso problema. Se il numero dei processori tende a infinito, vediamo che la Legge di Amdahl ha un comportamento asintotico (per lim di *p* → ∞, *S* tende a 1/*r* senza mai toccarlo).
  
  **Scalabilità Strong**: valuta l'<ins>efficienza al crescere del numero dei nodi</ins> (<ins>mantenendo costante la dimensione del problema</ins>). Lavoro totale da eseguire costante, ma lavoro da eseguire sul singolo nodo diminuisce al crescere del numero dei nodi (bilanciamento del lavoro sui nodi).
  
  **Scalabilità Weak**: valuta l'<ins>efficienza al variare al crescere delle dimensioni del problema</ins> (<ins>mantenendo costante il carico di lavoro per singolo nodo</ins>). Per valutare la scalabilità weak si usano speedup scalato e efficienza scalata.
  
  La **Legge di Gustafson** afferma che, <ins>assegnando ad ogni processore un workload costante</ins> ```(1 - r)```, lo <ins>speedup cresce linearmente con il numero dei processori</ins>, dunque lo speedup *S* è dato da: ```S = r + (1 - r) * p```.
</details>

### 4. Confrontare MPI ed OpenMP

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Per ottenere i vantaggi del parallelismo, sfruttando efficacemente l'HW a disposizione, il programmatore deve trasformare i propri programmi seriali in codice parallelo. Per farlo è possibile utilizzare 2 approcci: <ins>parallelizzazione automatica</ins> (sfruttando ad esempio dei compilatori), che normalmente permettono di ottenere prestazioni non troppo soddisfacenti; <ins>parallelizzazione esplicita</ins>, utilizzando ad esempio dei linguaggi e librerie appositi per il calcolo parallelo.
  A tal proposito esistono due modelli di interazione: scambio di messaggi (*MPI*) e memoria condivisa (*OpenMP*).
  
  **MPI**: è uno <ins>standard</ins> che stabilisce un protocollo per la comunicazione fra processi in sistemi paralleli (*senza memoria condivisa*). Permette di eseguire più istanze di un programma in parallelo su più nodi.
  
  **OpenMPI**: è una <ins>libreria</ins> per applicazioni parallele in sistemi a *memoria condivisa*.
  
  <table>
	<tr>
		<td align="center" width="50%"><b>MPI</b></td>
		<td align="center" width="50%"><b>OpenMPI</b></td>
	</tr>
	<tr>
		<td align="center">Interazione basata su scambio di messaggi</td>
		<td align="center">Interazione basata su memoria condivisa</td>
	</tr>
	<tr>
		<td align="center">complessità d'uso</td>
		<td align="center">semplicità d'uso</td>
	</tr>
	<tr>
		<td align="center">load balancing a carico del programmatore</td>
		<td align="center">load balancing semplice da realizzare</td>
	</tr>
	<tr>
		<td align="center">elevata scalabilità</td>
		<td align="center">scalabilità limitata al numero di CPU disponibili sul nodo utilizzato</td>
	</tr>
	<tr>
		<td align="center">elevata portabilità (funziona anche su sistemi a memoria condivisa)</td>
		<td align="center">compatibile solo con sistemi a memoria condivisa (multicore/multiprocessors)</td>
	</tr>
  </table>
  
</details>
