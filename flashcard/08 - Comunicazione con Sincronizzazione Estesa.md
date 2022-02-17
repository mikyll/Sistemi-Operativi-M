<h1 align="center">COMUNICAZIONE CON SINCRONIZZAZIONE ESTESA</h1>

### 1. Spiegare in Cosa Consiste la Comunicazione con Sincronizzazione Estesa

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  La sincronizzazione estesa è un meccanismo di comunicazione che prevede che un processo chiamante richieda un servizio ad un altro processo e rimanga sospeso fino al completamento del servizio richiesto. Semanticamente, la sincronizzazione estesa è <ins>analoga ad una chiamata di funzione</ins>, in quanto il programma chiamante prosegue solo dopo che l'esecuzione della funzione è stata completata. La differenza sostanziale è che il servizio richiesto viene eseguito remotamente da un processo differente da quello chiamante. Il server può essere implementato in 2 modi: *Remote Procedure Call* (RPC) oppure *rendez-vous esteso*.
</details>

### 2. Descrivere RPC e Rendez-Vous Esteso e Confrontarli

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  ##### RPC
  Per ogni operazione che il client può richiedere viene dichiarata una procedura lato server. Al momento dell'effettiva richiesta, il <ins>server crea un nuovo processo (**fork**)</ins>, il quale effettua una chiamata all procedura corrispondente e, una volta terminata l'operazione, <ins>invia direttamente lui stesso la risposta</ins> al client.<br/>
  L'insieme delle procedure remote è definito all'interno di un componente software (*modulo*), che contiene anche le variabili locali al server, ed eventuali procedure e processi locali. I singoli moduli operano in spazi di indirizzamento diversi e possono quindi essere allocati su nodi distinti di una rete.
  
  ##### Rendez-Vous Esteso
  Ogni operazione viene specificata come un insieme di istruzioni, preceduto da un'<ins>istruzione **accept** che sospende il processo server</ins> (sincronizzazione) in attesa di una richiesta dell'operazione. All'arrivo della richiesta il processo esegue il relativo insieme di istruzioni ed i risultati ottenuti sono inviati al chiamante.<br/>
  La accept è bloccante se non sono presenti richieste di servizio. Se uno stesso servizio viene richiesto da più processi, le richieste vengono inserite in una coda associata al servizio, gestita con politica FIFO. Ad uno stesso servizio possono essere associate più accept nel codice eseguito dal server, dunque <ins>ad una richiesta possono corrispondere azioni diverse</ins>. Lo schema di comunicazione realizzato dal meccanismo di rendez-vous è di tipo asimmetrico molti-a-uno.<br/>
  Il server può selezionare le richieste da servire in base al suo <ins>stato interno</ins>, utilizzando i comandi con guardia; oppure anche in base ai <ins>parametri di ingresso della richiesta</ins>, anche in questo caso specificando i controlli da effettuare nel comando con guardia. Per utilizzare entrambi contemporaneamente (stato interno e parametri di ingresso).<br/>
  Ada è un linguaggio molto espressivo dal punto di vista della comunicazione fra processi, perché permette ad esempio di eseguire operazioni diverse (accept diverse) per una richiesta dello stesso tipo.
  
  ##### Differenze
  - RPC rappresenta solo un meccanismo di <ins>comunicazione</ins> fra processi, mentre delega al programmatore la gestione della sincronizzazione dei vari processi figli del servitore, permettendo di eseguire più operazioni concorrentemente (es: Java RMI, Distributed Processes).
  - Rendez-vous Esteso combina <ins>comunicazione con sincronizzazione</ins>, in quanto vi è un solo processo server, al cui interno sono definite le istruzioni che consentono di realizzare il servizio richiesto.
</details>