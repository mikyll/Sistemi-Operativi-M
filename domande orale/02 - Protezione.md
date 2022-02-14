
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

### 5. Diritti di Accesso: Copy Flag (*), Owner, Control, Switch. È Possibile Capire Quale Politica Si Sta Utilizzando?

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  **Copy Flag** (\*): è un diritto di accesso esercitato da un soggetto su un particolare diritto di accesso per un oggetto, che permette la propagazione di tale diritto ad altri soggetti. La propagazione può essere realizzata in due modi: *trasferimento* (il soggetto iniziale perde il diritto), e *copia* (il soggetto iniziale mantiene il diritto).
  
  **Owner**: possedere tale diritto su un oggetto permette di assegnare e revocare un qualunque diritto di accesso su tale oggetto ad altri soggetti.
  
  **Control**: se un soggetto S1 possiede tale diritto su un altro soggetto S2, S1 può revocare a S2 un qualunque diritto di accesso per oggetti nel suo dominio (di S2).
  
  **Switch**: se un soggetto possiede tale diritto su un altro soggetto, può spostarsi nel dominio di quest'ultimo.
  
  In certi casi è possibile capire quale politica il sistema di protezione sta adottando in base alla presenza di certi diritti. Copy flag e owner, ad esempio, indicano esplicitamente l'utilizzo di una politica DAC (decentralizzata).
</details>

### 6. Sicurezza Multilivello: Modello Bell-La Padula e BIBA + Esempio Cavallo di Troia

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  In alcuni ambienti è necessario un controllo più stretto sulle regole di accesso alle risorse (es: militare). I sistemi di sicurezza multilivello prevedono che vengano stabilite regole generali non modificabili senza aver ottenuto dei permessi speciali (basato su politica MAC, ovvero controllo degli accessi obbligatorio). In un sistema di sicurezza multilivello, i soggetti e gli oggetti sono classificati in livelli (classi di accesso) e vengono imposte delle regole di sicurezza che controllano il flusso delle informazioni tra i livelli.
  
  ##### Modello Bell-La Padula
  È progettato per garantire la segretezza (confidenzialità) dei dati, ma non l'integrità. Associa al sistema di protezione (matrice degli accessi), un modello di sicurezza multilivello, che prevede 2 regole:
  1. **Semplice sicurezza**, permette ad un processo in esecuzione ad un determinato livello, di <u>leggere solo oggetti di livello pari o inferiore</u>;
  2. **Star (o di integrità)**, permette ad un processo in esecuzione ad un determinato livello, di <u>scrivere solo oggetti di livello pari o superiore</u>.
  
  <img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/02%20-%20Protezione/Flusso%20Modello%20Bell-La%20Padula.png"/>
  
  Esempio di **difesa da un Trojan**, con modello Bell-La Padula:
  S1 possiede un file F1 da proteggere, con permessi di lettura/scrittura che appartengono solo a lui (S1);
  S2 è ostile e vuole rubarli, e possiede un file eseguibile CT (Cavallo di Troia), che ha installato nel sistema, assieme ad un file F2 che usa come "tasca posteriore".
  ACL:
  - S2 ha permessi di lettura/scrittura per F2 (tasca posteriore);
  - S2 dà a S1 il permesso di scrittura su F2;
  - S2 dà a S1 il permesso di esecuzione su CT;
  - F1 (file da proteggere) è leggibile solo da S1.
  S2 induce S1 ad eseguire CT che, essendo eseguito a nome di S1, può leggere F1 e scrivere su F2. In quanto sia lettura che scrittura soddisfano i vincoli della ACL.<br/>
  Tuttavia, se il sistema prevedesse il modello di sicurezza multilivello Bell-La Padula, e ci fossero ad esempio 2 livelli (*riservato*, per processi e file di S1, e *pubblico*, per processi e file di S2), il processo che esegue CT assumerebbe il livello di S1 (riservato), dunque potrebbe leggere il file F1 da proteggere, in quanto di pari livello (proprietà di semplice sicurezza rispettata), ma non potrebbe scrivere sul file F2, in quanto di livello inferiore (proprietà star violata). Dunque l'accesso, nonostante è consentito dalla ACL, viene negato.
  
  ##### Modello BIBA
  È progettato per garantire l'integrità dei dati, ma non la segretezza. Prevede anch'esso 2 regole:
  1. **Semplice sicurezza**, permette ad un processo in esecuzione ad un determinato livello, di <u>scrivere solo oggetti di livello pari o inferiore</u>;
  2. **Star (o di integrità)**, permette ad un processo in esecuzione ad un determinato livello, di <u>leggere solo oggetti di livello pari o superiore</u>.
  
  I modelli Bell-La Padula e BIBA sono in conflitto e non possono essere utilizzati contemporaneamente. Le politiche di sicurezza multilivello coesistono con le regole imposte dal sistema di protezione (ACL/CL) e hanno la *priorità* su quest'ultime.
</details>

### 7. Cosa sono i Sistemi Trusted, quali sono i componenti principali e che proprietà devono avere

<details>
  <summary><b>Visualizza risposta</b></summary>
  
  Un sistema trusted è un sistema per il quale è possibile definire formalmente dei requisiti di sicurezza. L'architettura di tale sistema prevede 2 componenti fondamentali:
  - **Reference Monitor (RM)**, è un elemento di controllo realizzato dall'HW e dal SO, che <u>regola l'accesso</u> dei soggetti agli oggetti <u>in base alle regole di sicurezza</u> (ad esempio fornite da un modello di sicurezza multilivello, tipo Bell-La Padula).
  - **Trusted Computing Base (TCB)**, è un elemento che <u>contiene i livelli di sicurezza</u> di soggetti (privilegi di sicurezza) e oggetti (classificazione rispetto alla sicurezza).
  
  I Sistemi Trusted devono rispettare le seguenti proprietà:
  - **mediazione completa**, ovvero le regole di sicurezza devono essere applicate ad ogni accesso alle risorse, e non solo. Dunque, essendo questa un'operazionee piuttosto frequente, per motivi di efficienza è necessario che la soluzione venga implementata (almeno parzialmente) via HW;
  - **isolamento**, ovvero sia RM che TCB devono essere isolati e protetti rispetto a modifiche non autorizzate (anche ad esempio da parte del kernel del SO);
  - **verificabilità**, ovvero dev'essere possibile dimostrare formalmente che il RM esegua correttamente il suo compito (imponendo il rispetto delle regole di sicurezza, e fornendo mediazione completa ed isolamento). Questo solitamente è un requisito difficile da soddisfare in un sistema general-purpose.
  
  **Audit File**: è una specie di file di log, che mantiene tutte le informazioni sulle operazioni eseguite più importanti e di interesse dal punto di vista della sicurezza del sistema, ad esempio modifiche autorizzate alla TCB o tentativi di violazione.
  
  ###### Classificazione della Sicurezza dei Sistemi di Calcolo
  Secondo l'Orange Book (documento pubblicato dal Dipartimento della Difesa americano), la sicurezza di un sistema viene classificata in base a 4 categorie:
  - Categoria **D: Minimal Protection**. Non prevede sicurezza. Es: MS-DOS.
  - Categoria **C: Discretionary Protection**. Es: Unix.
  - Categoria **B: Mandatory Protection**. Introduzione di livelli di sicurezza (es: Bell-La Padula).
  - Categoria **A: Verified Protection**.
</details>