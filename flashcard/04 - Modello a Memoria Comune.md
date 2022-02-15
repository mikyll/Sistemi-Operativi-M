

<h1 align="center">MODELLO A MEMORIA COMUNE</h1>

### 1. Aspetti Caratterizzanti del Modello a Memoria Comune e Regione Critica Condizionale

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Nel modello a memoria comune ogni interazione tra i processi avviene tramite oggetti contenuti in memoria comune. Ogni applicazione è vista come un insieme di componenti attivi (processi) e componenti passivi (risorse). I processi possono avere diritto di accesso sulle risorse, di cui necessitano per portare a termine il loro compito. Una risorsa può essere dedicata o condivisa, ed allocata staticamente o dinamicamente.
  
  **Regione Critica Condizionale**: formalismo che consente di *esprimere* qualunque vincolo di sincronizzazione. Si esprime come: ```region R << Sa; when(C) Sb; >>```, dove R è la risorsa condivisa, Sa ed Sb sono istruzioni, e C una condizione da verificare.<br/>
  Il corpo (tra virgolette) rappresenta una sezione critica che dev'essere eseguita in mutua esclusione, e consiste in un'operazione su R. Una volta terminata Sa viene valutata la condizione C:
  - se è *vera* si prosegue con Sb;
  - se è *falsa* si <u>attende</u> che C diventi vera.
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
  La relazione di invarianza è <u>sempre soddisfatta</u> per ogni semaforo.
</details>

### 3. Semaforo di Mutua Esclusione + Dimostrazione

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Il Semaforo di Mutua Esclusione (o semaforo binario), viene <u>inizializzato a 1</u> e viene utilizzato per realizzare le sezioni critiche di una stessa classe, seguendo il <u>protocollo: prima viene eseguita una P, poi una V</u>, ovvero ```P(mutex); <sezione_critica>; V(mutex);```, dove mutex è un semaforo inizializzato a 1.
  
  **Ipotesi**: Il semaforo è inizializzato a 1, e vengono eseguite prima la P poi la V.<br/>
  **Tesi**:
  1. le sezioni critiche della stessa classe vengono eseguite in mutua esclusione;
  2. non devono verificarsi deadlock;
  3. un processo che non sta eseguendo una sezione critica non deve impedire agli altri di eseguire la stessa sezione critica (o sezioni della stessa classe).
  
  ###### Dimostrazione di 1
  La tesi di mutua esclusione equivale a dire che il <u>numero di processi nella sezione critica</u> Nsez è maggiore o uguale a 0, e minore o uguale a 1, ovvero ```Nsez ≥ 0 e 1 ≥ Nsez```.
  
  Dato che è necessaria una P per entrare nella sezione critica, ed una V per uscire, si ha che il numero dei processi nella sezione critica è dato dal numero di volte in cui è stata eseguita una P, meno il numero di volte in cui è stata eseguita una v, ovvero: ```Nsez = np - nv```.<br/>
  Ma dalla Relazione di Invarianza sappiamo che (I = 1): 1 + nv ≥ np, dunque 1 ≥ np - nv, ovvero ```1 ≥ Nsez```.<br/>
  Inoltre, poiché il protocollo impone che P(mutex) preceda V(mutex), sappiamo che in qualunque istante dell'esecuzione ```np ≥ nv```, dunque np - nv ≥ 0, ovvero ```Nsez ≥ 0```. □
  
  ###### Dimostrazione di 2
  La tesi è l'assenza di deadlock, che dimostriamo per <u>assurdo</u>. Se ci fosse un deadlock:
  1. tutti i processi sarebbero in attesa su P(mutex), portando il contatore del semaforo a 0, dunque ```val = 0```;
  2. nessun processo sarebbe nella sezione critica, ovvero ```Nsez = np - nv = 0```.
  
  Sapendo che val = I + nv - np, sostituendo otteniamo val = 1 - (np - nv), ovvero ```val = 1 - Nsez```, ma se val = 0 e Nsez = 0, otteniamo ```0 = 1 - 0```, che è impossibile (assurdo). □
  
  ###### Dimostrazione di 3
  La tesi prevede che non ci siano processi in sezione critica, ovvero ```Nsez = 0```.
  
  Sostituendo nella relazione di invarianza otteniamo che: ```val = 1 - 0 = 1```, ovvero <u>P non è bloccante</u> (in quanto la P si blocca solo se val = 0). □
</details>

### 4. Semaforo Evento + Dimostrazione

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Il semaforo evento è un semaforo binario utilizzato per imporre un <u>vincolo di precedenza</u> tra le operazioni dei processi.
  Dato un processo *p* che esegue un'operazione *a*, si vuole che *a* possa essere eseguita solo dopo che un altro processo *q* abbia eseguito un'operazione *b*.
  Il semaforo evento S è <u>inizializzato a 0</u> e segue il <u>protocollo: prima di eseguire *a* il processo *p* esegue P(S); il processo *q* dopo aver eseguito *b* esegue V(S)</u>.
  
  **Ipotesi**: il semaforo è inizializzato a 0, ed i 2 processi seguono il protocollo definito ```p: P(S); a;  q: b; V(S);```.
  **Tesi**: *a* viene eseguita sempre prima di *b*.
  
  ###### Dimostrazione
  Dimostriamo la tesi per assurdo. Supponiamo che sia possibile che *a* venga eseguita in un istante precedente a quello in cui viene eseguita *b*. In questo modo avremmo che è stata eseguita una V(S) ma non una P(S), ovvero ```nv = 1``` e ```np = 0```.<br/>
  Ma per la relazione di invarianza, sappiamo che I + nv ≥ np, ovvero ```0 + 0 ≥ 1```, che è impossibile (assurdo). □
</details>

### 5. 

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  
</details>

### 6. 

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  
</details>

Semaforo: definizione formale e proprietà.
Dimostrazioni delle sue proprietà (strumento di sincronizzazione generale che può risolvere tutti i problemi di sincronizzazione)