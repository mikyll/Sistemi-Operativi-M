
<h1 align="center">PROTEZIONE</h1>

### 1. Cosa si Intende per Sistema di Protezione. Modello, Politiche, Meccanismi

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Un sistema di protezione permette di definire delle tecniche di controllo degli accessi. Questo si esprime tramite 3 concetti fondamentali:
  ##### Modello
  Definisce:
  - **oggetti**, ovvero la <u>parte passiva</u>, le risorse fisiche e logiche, ad esempio i file;
  - **soggetti**, ovvero la <u>parte attiva</u>, le entità che possono richiedere l'accesso alle risorse, ad esempio utenti e processi;
  - **diritti di accesso**, ovvero le <u>operazioni</u> con cui i soggetti possono operare sugli oggetti, ad esempio lettura e scrittura.
  NB: un soggetto può avere diritti di accesso sia per oggetti che per soggetti.
  ##### Politiche
  Definiscono le *regole* con cui i soggetti possono accedere agli oggetti. Si classificano in 3 tipologie:
  - **Discretionary Access Control (DAC)**, prevede che il proprietario di un oggetto ne controlli i diritti di accesso (gestione delle politiche decentralizzata, come accade in UNIX);
  - **Mandatory Access Control (MAC)**, prevede che i diritti vengano decisi in modo centralizzato (tipico dei sistemi ad alta sicurezza, ad esempio enti governativi);
  - **Role Based Access Control (RBAC)**, prevede che i diritti di accesso alle risorse vengano assegnati in base al ruolo, che viene assegnato in modo centralizzato (gli utenti possono appartenere a diversi ruoli).
  ##### Meccanismi
  Sono gli strumenti messi a disposizione dal sistema di protezione per imporre una determinata politica e vanno realizzati per rispettare:
  - **flessibilità** del sistema di protezione, ovvero devono essere abbastanza generali da permettere l'applicazione di diverse politiche;
  - **separazione tra meccanismi e politiche**, secondo cui la politica definisce "cosa va fatto" ed i meccanismi "come va fatto".
  
  ###### Esempio UNIX
  Politica DAC: l'utente definisce la *politica*, ovvero il valore dei bit di protezione per ogni oggetto di sua proprietà, ed il SO fornisce un *meccanismo* per definire ed interpretare per ciascun oggetto i bit di protezione.
</details>

### 2. Cos'è il Principio del Minimo Privilegio e Come si Può Implementare

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Se il érincipio del Minimo Privilegio (Principle Of Least Authority - POLA) viene rispettato, ad ogni soggetto sono garantiti i diritti di accesso dei soli oggetti strettamente necessari alla sua esecuzione. Il rispetto di questo principio è desiderabile a prescindere dalla politica adottata.
  
  Per implementarlo è possibile adottare un'associazione processo-dominio dinamica, che permetta di effettuare, a tempo di esecuzione del processo, il passaggio da un dominio ad un altro, in base alle risorse ad esso necessario in qualsiasi istante della sua esecuzione.
</details>

### 3. Cos'è il Dominio di Protezione. Dominio statico/dinamico e Cambio di Dominio

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Un dominio di protezione definisce un insieme di coppie <oggetto, diritti di accesso>, che rappresenta l'ambiente di protezione nel quale un certo soggetto esegue. Il dominio di protezione è infatti univoco per ciascun soggetto, e in ogni istante della sua esecuzione, un soggetto (processo) è associato ad uno ed un solo dominio, e può accedere solo agli oggetti specificati nel suo dominio, con i relativi diritti.
  
  L'associazione tra processo e dominio può essere:
  - **statica**, se rimane fissa durante l'intera esecuzione del processo. 
  - **dinamica**, se può variare nel corso dell'esecuzione del processo.
  Poiché a tempo di esecuzione, l'insieme globale delle risorse che un processo potrà usare può non essere conosciuto a priori, e l'insieme minimo delle risorse a lui necessarie cambia dinamicamente durante l'esecuzione, l'associazione statica non permette di realizzare il Principio del Minimo Privilegio. Al contrario, ciò è possibile con l'associazione dinamica, tuttavia occorre un meccanismo di cambio di dominio.
  
  Esempio: cambio di dominio relativo all'esecuzione di system call (2 ring, protezione tra kernel e utente, ma non tra diversi utenti); cambio di dominio in UNIX, realizzato tramite il bit set-uid che, se abilitato, permette al processo che esegue il file di passare nel dominio del proprietario del file.
</details>

### 4. Matrice degli Accessi e Com'è Possibile Rappresentarla Concretamente (PRO e CONTRO, e Possibile Soluzione)

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  La matrice degli accessi permette di rappresentare, a livello astratto, lo <u>stato di protezione</u> di un sistema in un determinato istante, ad esempio utilizzando le righe per indicare i soggetti, e le colonne per gli oggetti, mentre i singoli elementi contengono i vari diritti di accesso. Offre ai meccanismi le informazioni che gli consentono di verificare il rispetto dei vincoli di accesso.
  
  Solitamente il numero dei soggetti e soprattutto degli oggetti tende ad essere molto grande (e i diritti di accesso generalmente sono sparsi). Dunque, la matrice degli accessi non può essere realizzata come un'unica struttura Ns x No, in quanto ciò non sarebbe ottimale per l'occupazione della memoria né per l'efficienza negli accessi. Per questo motivo, la realizzazione concreta dev'essere ottimizzata, ed esistono 2 approcci:
  - **Access Control List (ACL)**, si basa su una rappresentazione per <u>colonne</u> e prevede che ad ogni <u>oggetto</u> sia associata una lista di coppie <soggetto, insieme-dei-diritti> (o <soggetto, gruppo, insieme-dei-diritti>), solo per i soggetti con un insieme non vuoto di diritti per l'oggetto.
  Quando dev'essere fatta un'operazione M su un oggetto Oj dal soggetto Si, il meccanismo di protezione cerca nell'ACL corrispondente all'oggetto Oj l'entry corrispondente al soggetto Si e controlla se è presente il diritto di eseguire M. La ricerca può essere fatta anche su una lista che contiene i diritti di accesso comuni a tutti i soggetti.
  La <u>revoca</u> di un diritto di accesso è molto <u>semplice</u> con ACL, in quanto basta fare riferimento all'oggetto coinvolto.
  - **Capability List (CL)**, si basa su una rappresentazione per <u>righe</u> e prevede che per ogni <u>soggetto</u> si abbia una lista di coppie <oggetto, insieme-dei-diritti>, che prendono nome di *capability* (capability = coppia).
  Per proteggere le CL da manomissioni, esse vengono memorizzate nello spazio del kernel e l'utente può far riferimento solo ad un puntatore che identifica la sua posizione nella lista.
  La <u>revoca</u> di un diritto di accesso è più <u>complesso</u> perché è necessario verificare, per ogni dominio (soggetto), se contiene capability che fanno riferimento all'oggetto considerato.
  
  Entrambe le soluzioni presentano **problemi di efficienza**: con le ACL i diritti di accesso di un particolare soggetto sono sparsi nelle varie ACL; con CL, l'informazione relativa a tutti i diritti di accesso applicabili ad un certo oggetto è sparsa nelle varie CL.
  
  **Soluzione Ibrida**: vengono combinati i due metodi. La ACL viene memorizzata in <u>memoria persistente</u> (secondaria) e, quando un soggetto tenta di accedere ad un oggetto per la prima volta, se il diritto invocato è presente nella ACL, viene restituita la CL relativa al soggetto richiedente, e salvata in <u>memoria volatile</u> (RAM). In questo modo il soggetto può accedere all'oggetto più volte senza dover analizzare nuovamente la ACL. Dopo l'ultimo accesso, la CL viene distrutta dalla memoria volatile.
</details>

### 5. Diritti di Accesso: Copy Flag (*), Owner, Control, Switch

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  
</details>

### 6. È possibile capire quale politica si sta utilizzando?
Owner e Copyflag indicano esplicitamente l'uso di una politica DAC.



### 6. Sicurezza Multilivello: Modello Bell-La Padula e BIBA + Esempio Cavallo di Troia

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  
</details>

### 7. Sistemi Trusted

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  
</details>