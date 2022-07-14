[[Index](https://github.com/mikyll/Sistemi-Operativi-M/tree/main/flashcard)]&nbsp;&nbsp;
[[<<](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/01%20-%20Virtualizzazione.md)]
[[&nbsp;<&nbsp;](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/05%20-%20Nucleo%20di%20un%20Sistema%20Multiprogrammato%20(Memoria%20Comune).md)]
[[&nbsp;>&nbsp;](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/08%20-%20Comunicazione%20con%20Sincronizzazione%20Estesa.md)]
[[>>](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/flashcard/11%20-%20HPC.md)]

<h1 align="center">MODELLO A SCAMBIO DI MESSAGGI</h1>

### 1. Definire le Caratteristiche del Modello a Scambio di Messaggi ed il Concetto di Canale di Comunicazione

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Nel modello a scambio di messaggi:
  - ogni processo può accedere esclusivamente alle <ins>risorse allocate nella propria memoria locale/privata</ins>;
  - ogni risorsa del sistema è accessibile direttamente da un solo processo (<ins>gestore</ins>);
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
  - <ins>serve un buffer di capacità limitata</ins>, in quanto un buffer di dimensioni illimitate non è concretamente realizzabile e, per mantenere inalterata la semantica, bisogna sospendere il processo mittente se il buffer è pieno.
  
  **Comunicazione Sincrona** (o rendez-vous semplice): <ins>il primo processo</ins> che esegue l'operazione di comunicazione (invio o ricezione) <ins>si sospende</ins>, in attesa che l'altro sia pronto ad eseguire l'operazione corrispondente.<br/>
  Proprietà:
  - favorisce l'<ins>espressività</ins>, in quanto l'invio di un messaggio è un punto di sincronizzazione, ed il messaggio ricevuto contiene informazioni attribuibili allo stato attuale del processo mittente (verifica dei programmi semplificata);
  - il <ins>grado di parallelismo è minore</ins>, rispetto alla comunicazione asincrona;
  - <ins>non servono buffer</ins>, in quanto un messaggio può essere inviato solo se il destinatario è pronto a riceverlo.
  
  **Comunicazione con Sincronizzazione Estesa** (o rendez-vous esteso): è semanticamente analogo alla chiamata di procedura remota (<ins>RPC</ins>), in quanto ogni messaggio inviato rappresenta una richiesta al destinatario dell'esecuzione di una certa azione. Il mittente rimane in attesa dopo l'invio della richiesta e si sblocca quando riceve la risposta (con gli eventuali risultati).<br/>
  Proprietà:
  - <ins>sfrutta il modello client/server</ins>;
  - <ins>elevata espressività</ins> (verificabilità dei programmi);
  - <ins>riduzione del grado di parallelismo</ins>.
</details>

### 3. Descrivere le Primitive Send e Receive e Spiegare com'è possibile Realizzare una Send Sincrona con una Asincrona e Viceversa

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  La **Send** è una primitiva di comunicazione che esprime l'invio di un messaggio ad un canale identificato univocamente. Può essere <ins>asincrona</ins> (canale bufferizzato) o <ins>sincrona</ins> (buffer di capacità nulla), ovvero il processo attende che il ricevente esegua la receive prima di proseguire la propria esecuzione.
  
  La **Receive** è una primitiva di comunicazione che esprime la lettura di un messaggio da un canale identificato univocamente, salvandone il contenuto in una variabile. È <ins>bloccante</ins> (sospende il processo che la esegue) se sul canale non ci sono messaggi da leggere.
  
  L'istruzione di più basso livello è la send asincrona. Per implementare una send sincrona si può inviare un messaggio e rimanere in attesa, su un altro canale, di un messaggio <ins>ACK</ins>, sfruttando la semantica bloccante della receive.
  
  Per costruire invece una send asincrona, con buffer di capacità N, da una send sincrona, è possibile utilizzare una mailbox concorrente. Ovvero, si possono utilizzare <ins>N processi concorrenti collegati in cascata</ins>, ciascuno dei quali esegue una receive dal precedente ed una send verso il successivo (il primo riceve dal processo mittente "produttore", l'ultimo invia al processo destinatario "consumatore").
</details>

### 4. Spiegare la Semantica della Receive, quali Problemi può dare e Come si Possono Risolvere

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Supponiamo di avere un processo server che fornisce diversi servizi, ognuno dei quali viene attivato in seguito alla ricezione di un messaggio su un canale diverso, da parte di un processo client.

  **Problema**: poiché ci sono più canali, il server deve ciclicamente eseguire una receive su ciascun canale, per verificare lo stato delle richieste. Tuttavia, poiché la receive ha semantica bloccante, il <ins>server potrebbe bloccarsi</ins> leggendo da un canale, <ins>mentre sono presenti messaggi in attesa di essere letti su altri canali</ins>.
  
  **Soluzione**: si potrebbe realizzare una <ins>receive con semantica non bloccante</ins>. Il server, prima di eseguire la receive da un canale, ne controlla lo stato:
  - se sono presenti messaggi, ne legge uno;
  - altrimenti, se nel canale non sono presenti messaggi, passa al successivo.
  <!-- ad esempio, in Go si potrebbe realizzare una funzione non_blocking_receive, che verifica con len() lo stato del canale e, se c'è almeno un messaggio, ovvero len() > 0, effettua la receive e resituisce il valore; altrimenti, restituisce un valore nullo oppure un errore. -->
  
  In questo modo la receive non sospende mai il processo server, generando però un **ulteriore problema**: l'<ins>attesa attiva</ins> (se tutti i canali sono vuoti, il server continua ad iterare).
  
  **Meccanismo di Ricezione Ideale**:
  - consente al server di <ins>verificare contemporaneamente la disponibilità di messaggi su più canali</ins>;
  - abilita la <ins>ricezione di un messaggio da un qualunque canale contenente messaggi</ins>;
  - <ins>quando tutti i canali sono vuoti, blocca il processo in attesa che arrivi un messaggio</ins>, qualunque sia il canale su cui arriva.
  
  Questo meccanismo è realizzabile tramite i *comandi con guardia*.
</details>

### 5. Discutere il Comando con Guardia

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Il comando con guardia permette di realizzare un meccanismo di ricezione ideale.<br/>
  Sintassi: ```<guardia> -> <istruzione>;```<br/>
  dove ```<guardia>``` è costituita dalla coppia ```(<espressione_booleana> ; <receive>)```.<br/>
  L'espressione boooleana viene detta **guardia logica**, mentre l'operazione di receive viene detta **guardia d'ingresso** ed ha semantica <ins>bloccante</ins>.
  
  La valutazione di una guardia può fornire 3 diversi valori:
  1. **guardia fallita**, se l'espressione booleana ha il valore <ins>false</ins>;
  2. **guardia ritardata**, se l'espressione booleana ha valore <ins>true</ins> e nel canale su cui viene eseguita <ins>non ci sono messaggi</ins>;
  3. **guardia verificata**, se l'espressione booleana ha valore <ins>true</ins> e nel canale <ins>c'è almeno un messaggio</ins> (dunque la receive può essere eseguita subito).
  
  L'esecuzione di un comando con guardia determina un effetto diverso in base alla valutazione della guardia:
  1. in caso di *guardia fallita*, il comando termina senza produrre alcun effetto;
  2. in caso di *guardia ritardata*, il processo si sospende finché non arriva un messaggio sul canale, dopodiché verrà eseguita la receive e successivamente l'istruzione;
  3. in caso di *guardia valida*, il processo esegue la receive e successivamente l'istruzione.
  
  **Comando con Guardia Alternativo** (`select`): racchiude un numero arbitrario di comandi con guardia semplici. Esso valuta le guardie di tutti i rami e si possono verificare 3 casi:
  1. se *tutte le guardie sono fallite*, il comando termina;
  2. se *tutte le guardie non fallite sono ritardate*, il processo in esecuzione si sospende in attesa che arrivi un messaggio, dopodiché verrà eseguita la receive relativa e successivamente l'istruzione;
  3. se *una o più guardie sono valide*: 
     - viene scelto in modo <ins>non deterministico</ins> uno dei rami con guardia valida;
     - viene eseguita la relativa receive;
     - viene eseguita successivamente l'istruzione;
     - l'esecuzione dell'intero comando alternativo termina.
  
  NB: la scelta del ramo fra quelli con guardia valida è non deterministica per non imporre una politica preferenziale tra i vari casi.

  Sintassi del comando con guardia alternativo:
  ```C
  select {
    [ ] <guardia_1> -> <istruzione_1>;
    [ ] <guardia_2> -> <istruzione_2>;
    ...
    [ ] <guardia_n> -> <istruzione_n>;
  }  
  ```
  
  **Comando con Guardia Ripetitivo**: ha un comportamento analogo al caso "alternativo", ma il ciclo ricomincia tutte le volte che viene eseguita un'istruzione, terminando solo se tutte le guardie sono fallite.
  
  Sintassi del comando con guardia ripetitivo:
  ```C
  do {
    [ ] <guardia_1> -> <istruzione_1>;
    ...
    [ ] <guardia_n> -> <istruzione_n>;
  }
  ```
</details>
