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
  
In un sistema distribuito gli orologi di ogni nodo <ins>non sempre sono sincronizzati</ins>, dunque è possibile che l'ordine nel quale due eventi vengono registrati sia diverso da quello in cui sono effettivamente accaduti, e questo può generare problemi.<br/>
  Per questo motivo, gli orologi utilizzati in applicazioni distribuite si dividono in **fisici**, che forniscono l'ora esatta, e **logici**, che permettono di associare un <ins>timestamp</ins> coerente con l'ordine in cui gli eventi si sono effettivamente verificati.
  
  **Orologi Logici**: per implementare gli orologi logici, si definisce una relazione *happened-before* "→", tale che:
  1. *A* e *B* sono eventi in uno stesso processo ed *A* si verifica prima di *B*, allora *A* → *B*;
  2. *A* è l'evento di invio di un messaggio e *B* è l'evento di ricezione dello stesso, allora *A* → *B*;
  3. vale la proprietà transitiva, ovvero se *A* → *B*, e *B* → *C*, allora *A* → *C*.
  
  Assumiamo quindi che ad ogni evento *e* venga associato un timestamp *C(e)* e che tutti i processi concordino su questo, per cui vale la proprietà **[*]**: ```A → B ⟺ C(A) < C(B)```. Dunque, se all'interno di un processo *A* precede *B*, avremo che *C(A)* < *C(B)*; se *A* è l'evento di invio e *B* l'evento di ricezione dello stesso messaggio, allora *C(A)* < *C(B)*.
  
  ##### Algoritmo di Lamport
	L'algoritmo di *Lamport* fornisce una soluzione al problema della sincronizzazione dei processi in un contesto distribuito, basata sull'utilizzo di orologi logici, implementati tramite timestamp. In particolare, per garantire il rispetto della proprietà [*], l'algoritmo afferma che:
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

- **Vantaggi**: è <ins>molto scalabile</ins>.
- **Svantaggi**: ha un <ins>maggiore costo di comunicazione</ins> per singolo partecipante, in quanto sono necessari 2\*(N-1) messaggi per ciascuna sezione critica (il processo in stato *WANTED* invia ```RICHIESTA``` e riceve ```OK``` da parte di tutti gli altri nodi); presenta <ins>poca tolleranza ai guasti</ins> in quanto presenta *N Points of Failure*, in quanto se un nodo va in crash, questo non risponderà più alle richieste, facendo rimanere i processi in attesa.

**Soluzione al Problema dei Guasti**: si può modificare il protocollo, prevedendo un messaggio dopo l'invio della risposta:
- ```OK```, in caso di autorizzazione;
- ```ATTESA```, in caso il processo opposto si trovi in stato di *HELD*.

In questo modo, basterà impostare un <ins>timeout</ins> nel richiedente per rilevare la presenza di guasti nel destinatario.
  
  ##### Algoritmo Token-Ring
  L'algoritmo *Token-Ring* è una soluzione *decentralizzata token-based* che prevede che i processi siano collegati tra di loro secondo una <ins>topologia ad anello orientato</ins>, in cui ciascun processo conosce i suoi vicini, e si scambiano un messaggio (token) nel verso relativo all'ordine dei processi. Il token rappresenta il permesso unico di eseguire sezioni critiche.<br/>
  Quando un processo riceve il token:
  1. se si trova in stato **WANTED**, allora <ins>trattiene il token</ins> ed esegue la propria sezione critica, dopodiché (una volta terminata l'operazione) passa il token al processo successivo;
  2. se si trova in stato **RELEASED**, <ins>passa direttamente il token</ins> al processo successivo nell'anello.
  
  - **Vantaggi**: è <ins>molto scalabile</ins>;
  - **Svantaggi**: ha un <ins>costo di comunicazione variabile</ins> (il numero di messaggi per ogni sezione critica dipende dal numero dei nodi presenti, dunque è *potenzialmente infinito*); come per Ricart-Agrawala, <ins>non è tollerante ai guasti</ins> e presenta *N Points of Failure* e vi è la possibilità di perdere il token se il nodo che lo detiene va in crash.
  
  **Soluzione al Problema dei Guasti**: come per Ricart-Agrawala, si può modificare il protocollo per prevedere che ad ogni invio del token, venga restituita una <ins>risposta</ins> e, in caso questa non arrivi entro un <ins>timeout</ins>, il nodo viene considerato guasto, escluso dall'anello e si passa il token al successivo.
</details>

### 4. Algoritmi di Sincronizzazione Distribuiti: Elezione del Coordinatore

<details>
  <summary><b>Visualizza risposta</b></summary>
  
In alcuni algoritmi è previsto che un processo **coordinatore** rivesta un ruolo speciale nella sincronizzazione tra i vari nodi. La designazione del coordinatore può essere *statica*, se viene scelto prima dell'esecuzione, o *dinamica*, mediante un <ins>algoritmo di elezione</ins> a tempo di esecuzione. Quest'ultima permette, di cambiare coordinatore a runtime se quello attuale smette di rispondere (ad esempio a causa di un guasto).<br/>
**Assunzioni di base**: ogni processo è identificato da un ID univoco; ogni processo conosce gli ID di tutti gli altri (ma non il loro stato).
**Obbiettivo**: viene designato vincitore (nuovo coordinatore) il processo attivo con l'ID più alto.

##### Algoritmo Bully
L'algoritmo di elezione *Bully* prevede che quando un processo *Pk* (k = 1, ..., N) rileva che il coordinatore non è più attivo, organizzi un'elezione:
1. *Pk* invia un messaggio ```ELEZIONE``` a tutti i processi con ID più alto del suo;
2. se nessun processo risponde, *Pk* vince l'elezione e diventa il nuovo coordinatore, dunque comunica a tutti gli altri il nuovo ruolo inviando un messaggio ```COORDINATORE```;
3. se un processo *Pj* (j > k) risponde, *Pj* prende il controllo dell'elezione, e *Pk* rinuncia, smettendo di rispondere ai successivi messaggi ```ELEZIONE```. 

Ogni processo attivo risponde ad ogni messaggio ```ELEZIONE``` ricevuto.
  
##### Algoritmo ad Anello
L'algoritmo di elezione ad *Anello* prevede che i processi siano collegati tramite una topologia logica ad anello orientato, in cui i processi sono posizionati in ordine in base al loro ID, che rappresenta anche la loro priorità. Quando un processo *Pk* rileva che il coordinatore non è più attivo (non risponde), organizza un'elezione:
1. *Pk* invia un messaggio ```ELEZIONE``` contenente il suo ID al successore, bypassandolo in caso sia in crash (si presuppone che un processo abbia gli strumenti per farlo);
2. quando un processo *Pi* riceve un messaggio ```ELEZIONE```:
	- se il messaggio non contiene il suo ID (di *Pj*), aggiunge il suo ID al messaggio e lo spedisce al successivo;
	- se il messaggio contiene il suo ID, significa che è stato compiuto un <ins>giro completo dell'anello</ins>, dunque *Pj* designa come coordinatore il processo avente l'ID più alto nel messaggio, e invia al successivo un messaggio ```COORDINATORE```, contenente l'ID del processo designato come nuovo coordinatore;
3. quando un processo riceve un messaggio ```COORDINATORE```, notifica il risultato dell'elezione al successivo, che farà lo stesso con quello dopo, e così via.
</details>

### 5. Spiegare degli Esempi forniti dalla Prof

##### Esempio 1: Algoritmo di Ricart-Agrawala
  Abbiamo 5 processi:
  - *P1* in stato **HELD**;
  - *P2* in **RELEASED**;
  - *P3* in **WANTED** con **```ts(m) = 3```**;
  - *P4* in **RELEASED**;
  - *P5* in **WANTED** con **```ts(m) = 5```**.
  
  Spiegare come e in quale ordine vengono gestite le richieste, secondo l'algoritmo di *Ricart-Agrawala*.

<details>
  <summary><b>Visualizza risposta</b></summary>
  
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
</details>
