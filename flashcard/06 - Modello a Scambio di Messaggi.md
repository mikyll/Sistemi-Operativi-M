
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
  4. <ins>tipo di sincronizzazione</ins> fra i processi comunicanti (comunicazione *asincrona*, *sincrona* o con *sincronizzazione estesa*).
</details>

### 2. 

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

### Semantiche per primitive nel modello a scambio di messaggi: confronto tra send sincrone e asincrone (vantaggi e svantaggi).

### receive con semantica bloccante/non bloccante -> canali molti-a-uno (non bloccante->attesa attiva) -> spiegazione di comando con guardia

### modello a scambio di messaggi: quali sono le semantiche di ricezione

### scambio di mess: panoramica di possibili semantiche send e receive (domanda generale da approfondire molto)

### Utilità di associare ad una guardia l'accept (=differenziazione delle varie richieste così che la receive non sia bloccate, possibilità di realizzare server pronti a ricevere ogni richiesta) 

### Quali strumenti può utilizzare un processo nel modello a scambio di messaggi per ricevere i messaggi (primitiva receive e discussione sul comportamento bloccante e i comandi con guardia composta da booleano, receive e comandi da eseguire. Valori che può assumere la guardia: ritardata, attiva e fallita)
