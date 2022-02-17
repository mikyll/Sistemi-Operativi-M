[[Index](https://github.com/mikyll/Sistemi-Operativi-M/tree/main/flashcard)]&nbsp;&nbsp;
[[<<](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/01%20-%20Virtualizzazione.md)]
[[&nbsp;<&nbsp;](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/08%20-%20Comunicazione%20con%20Sincronizzazione%20Estesa.md)]
[[&nbsp;>&nbsp;](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/11%20-%20HPC.md)]
[[>>](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/11%20-%20HPC.md)]

<h1 align="center">ALGORITMI DI SINCRONIZZAZIONE DISTRIBUITI</h1>

### 1. Caratteristiche di un Sistema Distribuito, Proprietà Desiderate e Sincronizzazione

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  In un sistema distribuito i processi eseguono su nodi fisicamente separati, collegati tra loro da una rete di interconnessione, ed il modello a scambio di messaggi è la sua naturale astrazione.<br/>
  Caratteristiche: concorrenza/parallelismo delle attività dei nodi; assenza di risorse condivise tra nodi; assenza di un clock globale; possibilità di malfunzionamenti indipendenti dei nodi (crash, attacchi, ...), o della rete di comunicazione (latenza, packet loss).
  
  **Proprietà Desiderate**:
  - **scalabilità**, nell'applicazione distribuita le prestazioni dovrebbero crescere al crescere del numero di nodi utilizzati;
  - **tolleranza ai guasti**, l'applicazione dev'essere in grado di funzionare anche in presenza di guasti (crash dei nodi, problemi di rete, ...).
  
  **Speedup**: indicatore per misurare le *prestazioni* di un sistema parallelo/distribuito. Lo speedup per N nodi è dato dal rapporto tra il tempo di esecuzione dell'applicazione ottenuto con un solo nodo e quello ottenuti con N nodi, ovvero: ```speedup(N) = tempo(1) / tempo(N)```. Il caso ideale (sistema scalabile al 100%) è ```speedup(N) = N```.
  
  **Tolleranza ai Guasti**: un sistema distribuito si dice tollerante ai guasti se riesce ad *erogare i propri servizi anche in presenza di guasti* (temporanei, intermittenti o persistenti) in uno o più nodi. Un sistema tollerante ai guasti deve nascondere i problemi agli altri processi, ad esempio tramite ridondanza.
  
  **Algoritmi di Sincronizzazione**: come nel modello a memoria comune, anche nel modello a scambio di messaggi è importante poter disporre di algoritmi di sincronizzazione tra i processi concorrenti, che permettano di <ins>coordinare opportunamente i vari processi</ins>:
  - *timing*, sincronizzazione dei clock e tempo logico;
  - *mutua esclusione* distribuita;
  - *elezione di coordinatori* di gruppi di processi.
  
  In ogni caso, è sempre desiderabile che tali algoritmi godano di scalabilità e tolleranza ai guasti (e vengono valutati anche in base a tali parametri).
</details>

### 2. Spiegare il Problema della Gestione del Tempo nei Sistemi Distribuiti ed una Possibile Soluzione

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  In un sistema distribuito gli orologi di ogni nodo non sempre sono sincronizzati, dunque è possibile che l'ordine nel quale due eventi vengono registrati sia diverso da quello in cui sono effettivamente accaduti, e questo può generare problemi.<br/>
  Per questo motivo, gli orologi utilizzati in applicazioni distribuite si dividono in **fisici**, che forniscono l'ora esatta, e **logici**, che permettono di associare un timestamp coerente con l'ordine in cui gli eventi si sono effettivamente verificati.
  
  **Orologi Logici**: per implementare gli orologi logici, si definisce una relazione *happened-before* "→", tale che:
  1. *A* e *B* sono eventi in uno stesso processo ed *A* si verifica prima di *B*, allora *A* → *B*;
  2. *A* è l'evento di invio di un messaggio e *B* è l'evento di ricezione dello stesso, allora *A* → *B*;
  3. vale la proprietà transitiva, ovvero se *A* → *B*, e *B* → *C*, allora *A* → *C*.
  
  Assumiamo quindi che ad ogni evento *e* venga associato un timestamp *C(e)* e che tutti i processi concordino su questo, per cui vale la proprietà *: ```A → B ⟺ C(A) < C(B)```. Dunque, se all'interno di un processo *A* precede *B*, avremo che *C(A)* < *C(B)*; se *A* è l'evento di invio e *B* l'evento di ricezione dello stesso messaggio, allora *C(A)* < *C(B)*.
  
  **Algoritmo di Lamport**: Per garantire il rispetto della proprietà *:
  1. ogni processo *Pi* gestisce localmente un <int>contatore</int> *Ci* del tempo logico;
  2. ogni evento del processo fa incrementare il contatore di 1 (*Ci*++);
  3. ogni volta che il processo Pi invia un messaggio *m*, il contatore viene incrementato (*Ci*++) e successivamente al messaggio viene assegnato il timestamp *ts*(*m*)=*Ci*;
  4. quando un processo *Pj* riceve un messaggio *m*, assegna al proprio contatore *Cj* un valore dato dal massimo tra *Cj* e *ts*(*m*), ovvero ```Cj = max{Cj, ts(m)}```, e successivamente lo incrementa di 1 (*Cj*++).
</details>

### 3. Algoritmi di Sincronizzazione Distribuiti: Mutua Esclusione e Soluzioni Possibili (PRO e CONTRO)

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Nei sistemi distribuiti è spesso necessario garantire che due processi non possano eseguire contemporaneamente alcune attività, ad esempio quelle che prevedono accesso a risorse condivise. Questo problema può essere risolto in maniera:
  - *centralizzata*, delegando la gestione ad un processo <ins>coordinatore</ins> al quale tutti gli altri processi si rivolgono per l'utilizzo della risorsa;
  - *decentralizzata*, sincronizzando i prrocessi mediante algoritmi la cui logica è distribuita tra i processi stessi (questo è generalmente un approccio più scalabile, in quanto avere un coordinatore singolo costituisce un collo di bottiglia).
  
  Le soluzioni al problema della mutua esclusione distribuita si dividono inoltre in:
  - *permission-based* (centralizzati o decentralizzati), nelle quali ogni processo che vuole eseguire la sezione critica (operazione mutuamente esclusiva) "<ins>richiede il permesso</ins>" di eseguire ad uno o più altri processi;
  - *token-based* (sempre decentralizzati), in cui i processi si passano un <ins>token</ins> che concede l'autorizzazione ad eseguire la propria sezione critica.
  
  ##### Soluzione Centralizzata
  La soluzione *centralizzata* prevede la presenza di un processo coordinatore che espone due primitive di <ins>richiesta</ins> e <ins>rilascio</ins> della risorsa. Ogni processo che vuole eseguire la propria sezione critica si rivolge al coordinatore per ottenere il *permesso*. Il coordinatore gestisce una richiesta alla volta, utilizzando una <ins>coda FIFO</ins>: se un processo richiede una risorsa che è attualmente utilizzata da un altro processo, viene messo in attesa in una coda, e risvegliato dal coordinatore stesso, quando la risorsa si libera di nuovo (ed è il suo turno nella coda).
  - **Vantaggi**: è un <ins>algoritmo equo</ins> (è privo di *starvation*), ed è implementabile utilizzando <ins>solo 3 messaggi</ins> (richiesta, autorizzazione, rilascio) per ciascuna sezione critica.
  - **Svantaggi**: è <ins>poco scalabile</ins>, in quanto al crescere del numero dei nodi il *coordinatore* può diventare un *collo di bottiglia*; è <ins>poco tollerante ai guasti</ins> e prevede un *Single Point of Failure*, in quanto se si guasta il coordinatore, l'intero sistema si blocca, e inolte, se un processo non ottiene una risposta, non può distinguere il motivo (autorizzazione non concessa o guasto).
  
  ##### Algoritmo di Ricart-Agrawala
  L'algoritmo di *Ricart-Agrawala* è una soluzione *decentralizzata permission-based* che richiede, come requisito per il suo funzionamento, la presenza di un <ins>orologio logico sincronizzato (timestamp)</ins>. Ad ogni processo sono associati 2 thread concorrenti: **main**, che esegue la sezione critica, e **receiver** che riceve le autorizzazioni.
  
  **Main**: quando un main vuole entrare nella sezione critica:
  1. manda una ```RICHIESTA``` d'autorizzazione (con il proprio PID e timestamp) a tutti gli altri nodi;
  2. attende le autorizzazioni (```OK```) di tutti gli altri nodi;
  3. esegue la sezione critica;
  4. invia un ```OK``` a tutte le richieste in attesa.
  
**Receiver**: quando un receiver riceve una richiesta, esso può trovarsi in 3 possibili stati:
1. **RELEASED**, se il processo non è interessato ad eseguire la sezione critica (ed il proprio main non ha inviato richieste), dunque <ins>risponde ```OK```</ins>;
2. **WANTED**, se il processo vuole entrare nella sezione critica (dunque il proprio main è in attesa dell'autorizzazione ```OK```), allora <inst>confronta il timestamp</ins> della richiesta ricevuta *Tr* con quello della richiesta inviata *Ts*:
	- se *Tr* < *Ts*, <ins>risponde con ```OK```</ins>;
	- altrimenti (*Tr* ≥ *Ts*), non risponde e <ins>mette la richiesta ricevuta in coda</ins>;
3. **HELD**, se sta eseguendo la sezione critica, nel qual caso <ins>la richiesta viene messa in coda</ins>.
  
  Esempio:<br/>
  Abbiamo 5 processi: *P1* in stato **HELD**, *P2* in **RELEASED**, *P3* in **WANTED** con ```ts(m) = 3```, *P4* in **RELEASED**, *P5* in **WANTED** con ```ts(m) = 5```.
  1. *P1* riceve le richieste di *P3* e *P5* e le mette in coda, in quanto si trova in stato HELD (sta attualmente eseguendo la sezione critica); *P2* risponde ```OK``` a *P3* e *P5*; *P3*

<table>
	<tr>
		<td width="5%" align="center"><b>Stato</b></td>
		<td width="19%" align="center"><b><i>P1</i></b></td>
		<td width="19%" align="center"><b><i>P2</i></b></td>
		<td width="19%" align="center"><b><i>P3</i></b></td>
		<td width="19%" align="center"><b><i>P4</i></b></td>
		<td width="19%" align="center"><b><i>P5</i></b></td>
	</tr>
	<tr>
		<td align="center">(1)</td>
		<td align="center"><b>HELD</b><br/>sta eseguendo la propria sezione critica</td>
		<td align="center"><b>RELEASED</b></td>
		<td align="center"><b>WANTED</b><br/>
			<ins>invia</ins> <code>RICHIESTA</code> con <b><code>ts(m) = 3</code></b></td>
		<td align="center"><b>RELEASED</b></td>
		<td align="center"><b>WANTED</b><br/>
			<ins>invia</ins> <code>RICHIESTA</code> con <b><code>ts(m) = 5</code></b></td>
	</tr>
	<tr>
		<td align="center">(2)</td>
		<td align="center"><b>HELD</b><br/>
			riceve le richieste di <i>P3</i> e <i>P5</i> e le <ins>mette in coda</ins></td>
		<td align="center"><b>RELEASED</b><br/>
			riceve le richieste di <i>P3</i> e <i>P5</i> e <ins>risponde</ins> <code>OK</code> a entrambi</td>
		<td align="center">
			<b>WANTED <code>ts(m) = 3</code></b><br/>
			1) riceve <code>OK</code> da <i>P2</i>, <i>P4</i>, <i>P5</i><br/>
			2) riceve la richiesta di <i>P5</i> e la <ins>mette in coda</ins> poiché 3 < 5
		</td>
		<td align="center"><b>RELEASED</b><br/>
			riceve le richieste di <i>P3</i> e <i>P5</i> e <ins>risponde</ins> <code>OK</code> a entrambi</td>
		<td align="center">
			<b>WANTED <code>ts(m) = 5</code></b><br/>
			1) riceve <code>OK</code> da <i>P2</i>, <i>P4</i><br/>
			2) riceve la richiesta di <i>P3</i> e <ins>risponde</ins> <code>OK</code> poiché 3 < 5
		</td>
	</tr>
	<tr>
		<td align="center">(3)</td>
		<td align="center"><b>RELEASED</b><br/><ins>estrae dalla coda</ins> tutti i processi e <ins>invia</ins> <code>OK</code> a <i>P3</i> e <i>P5</i></td>
		<td align="center"><b>RELEASED</b></td>
		<td align="center">
			<b>HELD</b><br/>
			1) riceve <code>OK</code> da <i>P1</i><br/>
			2) ha ottenuto tutte le autorizzazioni, entra in sezione critica
		</td>
		<td align="center"><b>RELEASED</b></td>
		<td align="center">
			<b>WANTED</b><br/>
			1) riceve <code>OK</code> da <i>P1</i><br/>
			2) gli manca ancora l'<code>OK</code> di <i>P3</i>
		</td>
	</tr>
	<tr>
		<td align="center">(4)</td>
		<td align="center"><b>RELEASED</b></td>
		<td align="center"><b>RELEASED</b></td>
		<td align="center"><b>RELEASED</b><br/><ins>estrae dalla coda</ins> tutti i processi e <ins>invia</ins> <code>OK</code> a <i>P5</i>
		</td>
		<td align="center"><b>RELEASED</b></td>
		<td align="center">
			<b>HELD</b><br/>
			1) riceve <code>OK</code> da <i>P3</i><br/>
			2) ha ottenuto tutte le autorizzazioni, entra in sezione critica
		</td>
	</tr>
  </table>
  ##### Algoritmo Token-Ring
  + Esempio
</details>

### 4. Algoritmi di Sincronizzazione Distribuiti: Elezione del Coordinatore

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  
</details>

### 5. 

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  
</details>

### 6. 

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  
</details>


<!--
### Come può risolversi il problema della mutua esclusione in un sistema distribuito (dopo aver parlato un po’ può chiedere di risolvere un esercizio descritto al momento)

### algoritmi di sincronizzazione nei sistemi distribuiti: problema del tempo, come fare ad avere un unico riferimento temporale (orologio logico , algoritmo di Lamport)

### Come può essere risolto il problema della mutua esclusione in un contesto distribuito?

### sincronizzazione dei sistemi distribuiti in particolare come si risolve l’autoesclusione (soluzioni: centralizzata basata su processi, non centralizzata basata su processi, ?)

### ESEMPIO: supponiamo di avere un gruppo di 5 nodi p1, p2, p3, p4, p5. p3 si trova in stato di held e p2 e p4 si trovano nello stato di release e gli altri nello stato di wanted). cosa succede? le richieste viaggiano a un tempo t, come reagiscono i processi che ricevono il segnale. 

### quali sono i pro e i contro di uno decentralizzato e di uno centralizzato

### sistemi distribuiti: Algoritmi di Elezione(bully)

### Come può essere trattato il problema della mutua esclusione in un sistema distribuito (mutua esclusione distribuita, orologi logici, coordinamento attraverso un coordinatore eletto. approccio centralizzato, decentralizzato. permission based e token based. esempi centralizzato, Ricart Agrawala (con spiegazione di Lamport), Token ring)

### Algoritmi di elezione cosa sono e descriverli (bully e ring) 

### Mutua esclusione in un sistema distribuito (accesso contemporaneo a una risorsa da evitare, e utilizzare un sistema di tempo utilizzabile da tutti i nodi -> lamport. Alg permission based o token based, approccio centralizzato, ricart agrawala e token ring. esempio con un caso specifico richiesto da lei)
-->
