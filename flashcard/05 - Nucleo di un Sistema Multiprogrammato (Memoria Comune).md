<h1 align="center">NUCLEO DI UN SISTEMA MULTIPROGRAMMATO (MEMORIA COMUNE)</h1>

### 1. Spiegare Cos'è il Nucleo e quali sono le sue Funzioni Fondamentali

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Il nucleo (o kernel) è il modulo realizzato in SW, HW o FW che supporta il concetto di processo e realizza gli strumenti necessari per la loro gestione. È il livello più interno di qualunque sistema basato su processi ed è l'unico conscio dell'esistenza delle interruzioni (sono invisibili ai processi).<br/>
  Caratteristiche fondamentali del nucleo:
  - **efficienza**, in quanto condiziona l'intera struttura a processi;
  - **dimensioni ridotte**, in quanto le funzioni richieste al nucleo sono estremamente semplici;
  - **separazione meccanismi e politiche**, il nucleo deve il più possibile contenere solo *meccanismi*, consente scegliere ed applicare diverse politiche di gestione a seconda del tipo di applicazione.
  
  Stati di un processo (in un sistema in cui il numero di processi è maggiore del numero delle unità di elaborazione):
  - **esecuzione**, quando al processo è assegnata l'unità di elaborazione;
  - **pronto**, quando al processo è revocata l'unità di elaborazione;
  - **bloccato**, quando il processo non è attivo (P sospensiva).
  Quando un processo perde il controllo del processore, il suo <u>contesto</u> (ovvero il *contenuto dei registri del processore*) viene salvato nel <u>descrittore</u> (un'*area di memoria associata al processo*).
  
  Le funzioni fondamentali del nucleo riguardano la gestione delle transizioni di stato dei processi, in particolare:
  1. Gestire il <u>salvataggio ed il ripristino dei contesti dei processi</u>, ovvero trasferire le informazioni dai registri al descrittore, quando esso passa dallo stato di esecuzione allo stato di pronto o bloccato.
  2. Effettuare lo <u>scheduling della CPU</u>, ovvero scegliere a quale processo assegnare l'unità di elaborazione.
  3. Gestire le <u>interruzioni dei dispositivi</u> esterni.
  4. Realizzare i <u>meccanismi di sincronizzazione</u>.
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