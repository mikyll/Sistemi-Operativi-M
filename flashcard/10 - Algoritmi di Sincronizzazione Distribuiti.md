<h1 align="center">ALGORITMI DI SINCRONIZZAZIONE DISTRIBUITI</h1>

### 1. Caratteristiche di un Sistema Distribuito e Proprietà Desiderate

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  In un sistema distribuito i processi eseguono su nodi fisicamente separati, collegati tra loro da una rete di interconnessione, ed il modello a scambio di messaggi è la sua naturale astrazione.<br/>
  Caratteristiche: concorrenza/parallelismo delle attività dei nodi; assenza di risorse condivise tra nodi; assenza di un clock globale; possibilità di malfunzionamenti indipendenti dei nodi (crash, attacchi, ...), o della rete di comunicazione (latenza, packet loss).
  
  **Proprietà Desiderate**:
  - **scalabilità**, nell'applicazione distribuita le prestazioni dovrebbero crescere al crescere del numero di nodi utilizzati;
  - **tolleranza ai guasti**, l'applicazione dev'essere in grado di funzionare anche in presenza di guasti (crash dei nodi, problemi di rete, ...).
  
  **Speedup**: indicatore per misurare le *prestazioni* di un sistema parallelo/distribuito. Lo speedup per N nodi è dato dal rapporto tra il tempo di esecuzione dell'applicazione ottenuto con un solo nodo e quello ottenuti con N nodi, ovvero: ```speedup(N) = tempo(1) / tempo(N)```. Il caso ideale (sistema scalabile al 100%) è ```speedup(N) = N```.
  
  **Tolleranza ai Guasti**: un sistema distribuito si dice tollerante ai guasti se riesce ad *erogare i propri servizi anche in presenza di guasti* (temporanei, intermittenti o persistenti) in uno o più nodi. Un sistema tollerante ai guasti deve nascondere i problemi agli altri processi, ad esempio tramite ridondanza.
</details>

### 2. Spiegare il Problema del Tempo e della Sincronizzazione nei Sistemi Distribuiti ed una Possibile Soluzione

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  
</details>

### 3. 

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  
</details>

### 4. 

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