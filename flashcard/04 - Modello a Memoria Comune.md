<h1 align="center">MODELLO A MEMORIA COMUNE</h1>

### 1. Aspetti Caratterizzanti del Modello a Memoria Comune e Regione Critica Condizionale

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Nel modello a memoria comune ogni interazione tra i processi avviene tramite oggetti contenuti in memoria comune. Ogni applicazione è vista come un insieme di componenti attivi (processi) e componenti passivi (risorse). I processi possono avere diritto di accesso sulle risorse, di cui necessitano per portare a termine il loro compito. Una risorsa può essere dedicata o condivisa, ed allocata staticamente o dinamicamente.
  
  **Regione Critica Condizionale**: formalismo che consente di *esprimere* qualunque vincolo di sincronizzazione. Si esprime come: ```region R << Sa; when(C) Sb; >>```, dove R è la risorsa condivisa, Sa ed Sb sono istruzioni, e C una condizione da verificare.<br/>
  Il corpo (tra virgolette) rappresenta una sezione critica che dev'essere eseguita in mutua esclusione, e consiste in un'operazione su R. Una volta terminata Sa viene valutata la condizione C:
  - se è *vera* si prosegue con Sb;
  - se è *falsa* si <ins>attende</ins> che C diventi vera.
</details>

### 2. Semaforo: Definizione e Proprietà

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Il semaforo è uno strumento linguistico di basso livello che consente di *risolvere* qualunque problema di sincronizzazione nel modello a memoria comune.
  Il suo meccanismo è realizzato dal kernel della macchina concorrente, e l'attesa può essere implementata mediante i meccanismi di gestione dei thread offerti dal kernel. Viene utilizzato per implementare strumenti di sincronizzazione di più alto livello (ad esempio le condition).
  
  **Definizione Semaforo**: il semaforo ```S``` è una variabile intera non negativa ```val ≥ 0```, alla quale è possibile accedere solo mediante due operazioni mutuamente esclusive ```P``` e ```V```:
  - ```void P(sem S): region S << when(C > 0) S.val-- >>```
  - ```void V(sem S): region S << S.val++ >>```
  Il semaforo viene associato ad una risorsa e, quando un processo vuole operare su tale risorsa, esso chiama una P (down/richiesta):
  - se il valore del semaforo è positivo, il processo lo decrementa, esegue le sue operazioni, dopodiché chiama una V (up/rilascio);
  - altrimenti (se il valore del semaforo è 0), si mette in attesa finché un altro processo, che sta attualmente usando la risorsa gestita dal semaforo, non chiama una V, incrementandone il valore.
  
  **Proprietà del Semaforo**: dato un semaforo ```S```, siano ```val``` il suo valore (intero non negativo), ```I``` il valore ≥0 a cui viene inizializzato, ```nv``` il numero di volte che l'operazione V(S) è stata eseguita, ```np``` il numero di volte che P(S) è stata eseguita.
  
  **Relazione di Invarianza**: ad ogni istante è possibile esprimere il valore del semaforo come ```val = I + nv - np```, da cui (poiché val ≥0) I + nv - np ≥ 0, dunque: ```I + nv ≥ np``` (Relazione di Invarianza).<br/>
  La relazione di invarianza è <ins>sempre soddisfatta</ins> per ogni semaforo.
</details>

### 3. Semaforo di Mutua Esclusione + Dimostrazione

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Il Semaforo di Mutua Esclusione (o semaforo binario), viene <ins>inizializzato a 1</ins> e viene utilizzato per realizzare le sezioni critiche di una stessa classe, seguendo il <ins>protocollo: prima viene eseguita una P, poi una V</ins>, ovvero ```P(mutex); <sezione_critica>; V(mutex);```, dove mutex è un semaforo inizializzato a 1.
  
  **Ipotesi**: Il semaforo è inizializzato a 1, e vengono eseguite prima la P poi la V.<br/>
  **Tesi**:
  1. le sezioni critiche della stessa classe vengono eseguite in mutua esclusione;
  2. non devono verificarsi deadlock;
  3. un processo che non sta eseguendo una sezione critica non deve impedire agli altri di eseguire la stessa sezione critica (o sezioni della stessa classe).
  
  ###### Dimostrazione di 1
  La tesi di mutua esclusione equivale a dire che il <ins>numero di processi nella sezione critica</ins> Nsez è maggiore o uguale a 0, e minore o uguale a 1, ovvero ```Nsez ≥ 0 e 1 ≥ Nsez```.
  
  Dato che è necessaria una P per entrare nella sezione critica, ed una V per uscire, si ha che il numero dei processi nella sezione critica è dato dal numero di volte in cui è stata eseguita una P, meno il numero di volte in cui è stata eseguita una v, ovvero: ```Nsez = np - nv```.<br/>
  Ma dalla Relazione di Invarianza sappiamo che (I = 1): 1 + nv ≥ np, dunque 1 ≥ np - nv, ovvero ```1 ≥ Nsez```.<br/>
  Inoltre, poiché il protocollo impone che P(mutex) preceda V(mutex), sappiamo che in qualunque istante dell'esecuzione ```np ≥ nv```, dunque np - nv ≥ 0, ovvero ```Nsez ≥ 0```. □
  
  ###### Dimostrazione di 2
  La tesi è l'assenza di deadlock, che dimostriamo per <ins>assurdo</ins>. Se ci fosse un deadlock:
  1. tutti i processi sarebbero in attesa su P(mutex), portando il contatore del semaforo a 0, dunque ```val = 0```;
  2. nessun processo sarebbe nella sezione critica, ovvero ```Nsez = np - nv = 0```.
  
  Sapendo che val = I + nv - np, sostituendo otteniamo val = 1 - (np - nv), ovvero ```val = 1 - Nsez```, ma se val = 0 e Nsez = 0, otteniamo ```0 = 1 - 0```, che è impossibile (assurdo). □
  
  ###### Dimostrazione di 3
  La tesi prevede che non ci siano processi in sezione critica, ovvero ```Nsez = 0```.
  
  Sostituendo nella relazione di invarianza otteniamo che: ```val = 1 - 0 = 1```, ovvero <ins>P non è bloccante</ins> (in quanto la P si blocca solo se val = 0). □
</details>

### 4. Semaforo Evento + Dimostrazione

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Il semaforo evento è un semaforo binario utilizzato per imporre un <ins>vincolo di precedenza</ins> tra le operazioni dei processi.
  Dato un processo *p* che esegue un'operazione *a*, si vuole che *a* possa essere eseguita solo dopo che un altro processo *q* abbia eseguito un'operazione *b*.
  Il semaforo evento S è <ins>inizializzato a 0</ins> e segue il <ins>protocollo: prima di eseguire *a* il processo *p* esegue P(S); il processo *q* dopo aver eseguito *b* esegue V(S)</ins>.
  
  **Ipotesi**: il semaforo è inizializzato a 0, ed i 2 processi seguono il protocollo definito ```p: P(S); a;  q: b; V(S);```.
  **Tesi**: *a* viene eseguita sempre prima di *b*.
  
  ###### Dimostrazione
  Dimostriamo la tesi per assurdo. Supponiamo che sia possibile che *a* venga eseguita in un istante precedente a quello in cui viene eseguita *b*. In questo modo avremmo che è stata eseguita una V(S) ma non una P(S), ovvero ```nv = 1``` e ```np = 0```.<br/>
  Ma per la relazione di invarianza, sappiamo che I + nv ≥ np, ovvero ```0 + 0 ≥ 1```, che è impossibile (assurdo). □
</details>

### 5. Problema del Rendez-Vous + Barriera di Sincronizzazione

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  **Problema del Rendez-Vous**: si considerino due processi *A* e *B* che devono eseguire rispettivamente *a1*, *a2* e *b1*, *b2*, con il vincolo che l'esecuzione di *a2* e *b2* richieda che siano state completate sia *a1* che *b1*.
  
  **Soluzione**: per risolvere questo problema si possono introdurre due semafori evento (ovvero inizializzati a val = 0) S1 e S2. Il processo *A* esegue in sequenza ```a1; P(S1); V(S2); a2;```, mentre il processo B esegue in sequenza ```b1; P(S2); V(S1); b2;```. In questo modo il processo che termina per primo si blocca sulla P in attesa dell'altro processo, rispettando i vincoli di precedenza.
  
  **Generalizzazione del Problema del Rendez-Vous**: se i processi sono N > 2, è necessaria una struttura più complessa, chiamata *barriera di sincronizzazione*.
  
  **Barriera di Sincronizzazione**: strumento che permette di subordinare l'esecuzione di una serie di operazioni *Pib* (i = 1, ..., N) al completamento di una serie di operazioni *Pia* (i = 1, ..., N).<br/>
  La barriera è composta da:
  - un semaforo binario <ins>mutex, inizializzato a 1</ins>;
  - un semaforo evento <ins>barrier, inizializzato a 0</ins>;
  - un <ins>contatore done, inizializzato a 0</ins>, che rappresenta il numerod i processi che hanno completato la prima operazione (*Pia*).
  
  Ogni processo che termina l'operazione *Pia* richiede il mutex. Una volta ottenuto, incrementa done e, se done == N (ovvero tutti i processi hanno completato le rispettive operazioni *Pia*), chiama V(barrier). In seguito compliche termina attende la V(barrier) eseguita dall'ultimo proecsso che ha completato la propria operazione, prima di chiamare le rispettive V(barrier
  
  Implementazione in pseudo-C del processo i-esimo:
  ```C
  <operazione Pia>
  P(mutex);
  done++;
  if(done == N)
	V(barrier);
  V(mutex);
  P(barrier);
  V(barrier);
  <operazione Pib>
  ```
  In questo modo ogni processo attende la V(barrier) eseguita dall'ultimo processo (N-esimo) che completa la propria operazione *Pia*, prima di chiamare le rispettive V e iniziare la sequenza di risveglio degli N processi, facendo tornare il semaforo barrier a 0.
</details>

### 6. Descrivere l'Implementazione di un Semaforo nel Kernel di un Sistema Monoprocessore

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Un semaforo può essere rappresentato come una struttura dati contenente un contatore *c* ed una coda *q* (politica FIFO). Una *P* su un semaforo con *c* == 0 sospende il processo corrente *p* e lo inserisce in *q* mediante una push; altrimenti, se *c* > 0, il contatore *c* viene decrementato. Una *V* su un semaforo con la coda *q* vuota incrementa il contatore, mentre se *q* non è vuota estra un processo *p* da *q* mediante una pop.
  
  Implementazione in pseudo-C, supponendo che le <ins>interruzioni</ins> siano <ins>disabilitate</ins> durante l'esecuzione di *P* e *V*, in modo da garantire l'atomicità:
  ```C
  typedef struct {
	int c;
	queue q;
  } semaphore;
  
  void P(semaphore s)
  {
	if (s.c > 0)
	{
		s.c--;
	} else {
		// sospensione del processo corrente p, nella coda s.q
	}
  }
  void V(semaphore s)
  {
	if (!isEmpty(s.q))
	{
		// estrazione del primo processo p in attesa, dalla coda s.q
		// risveglio del processo p
	} else {
		s.c++;
	}
  }
  ```
  NB: l'implementazione di *P* e *V* è realizzata dal kernel della macchina concorrente e dipende dal tipo di architettura HW (monoprocessore, multiprocessore, ...) e da come il kernel rappresenta e gestisce i processi concorrenti.
</details>