
<h1 align="center">MODELLO A SCAMBIO DI MESSAGGI</h1>

### 1. Definire le Caratteristiche del Modello a Scambio di Messaggi ed il Concetto di Canale di Comunicazione

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Nel modello a scambio di messaggi:
  - ogni processo può accedere esclusivamente alle <ins>risorse allocate nella propria memoria locale/privata</ins>;
  - ogni risorsa del sistema è accessibile direttamente ad un solo processo (<ins>gestore</ins>);
  - se una risorsa è necessaria a più processi, ciascuno di questi (client) dovrà <ins>delegare l'unico processo in grado di operarvi</ins> (server/gestore), il quale restituirà successivamente i risultati;
  - il meccanismo di base per qualunque tipo di interazione fra i processi è lo <ins>scambio di messaggi</ins>.
  
  **Canale di Comunicazione**: collegamento logico mediante il quale 2 o più processi comunicano. L'astrazione del canale è realizzata dal kernel come meccanismo primitivo per lo scambio di informazioni, mentre è compito dei linguaggi di programmazione offrire gli strumenti linguistici di alto livello per istanziarli ed utilizzarli.<br/>
  Caratteristiche:
  1. <ins>direzione del flusso dei dati</ins> che il canale può trasferire (*monodirezionale* o *bidirezionale*);
  2. <ins>designazione</ins> dei processi <ins>mittente e destinatario</ins>:
	  - *link* = uno-a-uno (canale simmetrico);
	  - *port* = molti-a-uno (canale asimmetrico);
	  - *mailbox* = molti-a-molti (canale asimmetrico);
  3. <ins>tipo di sincronizzazione</ins> fra i processi comunicanti (comunicazione *asincrona*, *sincrona* o con *sincronizzazione estesa*).
</details>

### 2. Spiegare la Differenza tra Comunicazione Asincrona, Sincrona e con Sincronizzazione Estesa

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  **Comunicazione Asincrona**: il processo <ins>mittente continua la sua esecuzione</ins> immediatamente dopo l'invio del messaggio.<br/>
  Proprietà:
  - la <ins>carenza espressiva</ins> rende difficile la verifica dei programmi, in quanto la ricezione del messaggio può avvenire in un istante successivo all'invio e, di conseguenza, il messaggio ricevuto non contiene informazioni attribuibili allo stato attuale del mittente;
  - <ins>favorisce il grado di concorrenza/parallelismo</ins>, in quanto l'invio di un messaggio non costituisce un punto di sincronizzazione per mittente e destinatario;
  - serve un buffer di capacità limitata, in quanto un buffer di dimensioni illimitate non è concretamente realizzabile e, per mantenere inalterata la semantica, bisogna sospendere il processo mittente se il buffer è pieno).
  
  **Comunicazione Sincrona** (o rendez-vous semplice): <ins>il primo processo</ins> che esegue l'operazione di comunicazione (invio o ricezione) <ins>si sospende</ins>, in attesa che l'altro sia pronto ad eseguire l'operazione corrispondente.<br/>
  Proprietà:
  - favorisce l'<ins>espressività</ins>, in quanto l'invio di un messaggio è un punto di sincronizzazione, ed il messaggio ricevuto contiene informazioni attribuibili allo stato attuale del processo mittente (verifica dei programmi semplificata);
  - il <ins>grado di parallelismo è minore</ins>, rispetto alla comunicazione asincrona;
  - <ins>non servono buffer</ins>, in quanto un messaggio può essere inviato solo se il destinatario è pronto a riceverlo.
  
  **Comunicazione con Sincronizzazione Estesa** (o rendez-vous esteso): è semanticamente analogo alla chiamata di procedura remota (<ins>RPC</ins>), in quanto ogni messaggio inviato rappresenta una richiesta al destinatario dell'esecuzione di una certa azione. Il mittente rimane in attesa dopo l'invio della richiesta e si sblocca quando riceve la risposta (con gli eventuali risultati).<br/>
  Proprietà:
  - <ins>elevata espressività</ins> (verificabilità dei programmi);
  - <ins>riduzione del grado di parallelismo</ins>.
</details>

### 3. Descrivere le Primitive Send e Receive e Spiegare com'è possibile Realizzare una Send Sincrona con una Asincrona e Viceversa

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  La **Send** è una primitiva di comunicazione che esprime l'invio di un messaggio ad un canale identificato univocamente. Può essere <ins>asincrona</ins> (canale bufferizzato) o <ins>sincrona</ins> (buffer di capacità nulla), nel qual caso il processo attende che il ricevente esegua la receive prima di proseguire la propria esecuzione.
  
  La **Receive** è una primitiva di comunicazione che esprime la lettura di un messaggio da un canale identificato univocamente, salvandone il contenuto in una variabile. È <ins>bloccante</ins> (sospende il processo che la esegue) se sul canale non ci sono messaggi da leggere.
  
  L'istruzione di più basso livello è la send asincrona. Per implementare una send sincrona si può inviare un messaggio e rimanere in attesa, su un altro canale, di un messaggio <ins>ACK</ins>, sfruttando la semantica bloccante della receive.
  
  Per costruire invece una send asincrona, con buffer di capacità N, da una send sincrona, è possibile utilizzare una mailbox concorrente. Ovvero, si possono utilizzare <ins>N processi concorrenti collegati in cascata</ins>, ciascuno dei quali esegue una receive dal precedente ed una send verso il successivo (il primo riceve dal processo mittente "produttore", l'ultimo invia al processo destinatario "consumatore").
</details>

### 4. Spiegare la Semantica della Receive, quali Problemi può dare e Come si Possono Risolvere

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

### Semantiche per primitive nel modello a scambio di messaggi: confronto tra send sincrone e asincrone (vantaggi e svantaggi).

### receive con semantica bloccante/non bloccante -> canali molti-a-uno (non bloccante->attesa attiva) -> spiegazione di comando con guardia

### modello a scambio di messaggi: quali sono le semantiche di ricezione

### scambio di mess: panoramica di possibili semantiche send e receive (domanda generale da approfondire molto)

### Utilità di associare ad una guardia l'accept (=differenziazione delle varie richieste così che la receive non sia bloccate, possibilità di realizzare server pronti a ricevere ogni richiesta) 

### Quali strumenti può utilizzare un processo nel modello a scambio di messaggi per ricevere i messaggi (primitiva receive e discussione sul comportamento bloccante e i comandi con guardia composta da booleano, receive e comandi da eseguire. Valori che può assumere la guardia: ritardata, attiva e fallita)
