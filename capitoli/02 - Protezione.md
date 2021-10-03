## 02 - Protezione
**Sicurezza**: riguarda l'insieme delle *tecniche per regolamentare l'accesso* degli utenti al sistema di elaborazione. La sicurezza impedisce accessi non autorizzati al sistema e i conseguenti tentativi dolosi di alterazione e distruzione dei dati. La sicurezza riguarda l'interfaccia del sistema verso il mondo esterno. Le tecnologie di sicurezza di un sistema informatico realizzano meccanismi per l'identificazione, l'autenticazione e l'autorizzazione di utenti "fidati".<br/>
**Protezione**: *insieme di attività volte a garantire il controllo dell'accesso* alle risorse logiche e fisiche da parte degli utenti autorizzati all'uso di un sistema di calcolo. Rispetto alla sicurezza ha un campo d'azione più interno al sistema. Per rendere un sistema "sicuro" è necessario stabilire per ogni utente autorizzato quali siano le risorse a cui può accedere e con quali operazioni può farlo. Ciò viene stabilito dal sistema di protezione tramite le tecniche di controllo degli accessi.

### Protezione: Modelli, Politiche e Meccanismi
In un sistema il controllo degli accessi si esprime tramite la definizione di tre livelli concettuali:
modelli, politiche e meccanismi.

#### Modelli
Un modello di protezione definisce i *soggetti*, gli *oggetti* ai quali i soggetti hanno accesso ed i *diritti* di accesso:
- **soggetti** - rappresentano la parte attiva di un sistema, ovvero le entità che possono richiedere l'accesso alle risorse. Ad esempio: gli utenti o i processi che eseguono per conto degli utenti;
- **oggetti** - costituiscono la parte passiva di un sistema, ovvero le risorse fisiche e logiche alle quali si può accedere e su cui si può operare. Ad esempio i file, o i processi intesi come risorsa (solitamente in un sistema di protezione i soggetti sono gli utenti, e i processo sono oggetti);
- **diritti di accesso** - sono le operazioni con le quali è possibile operare sugli oggetti. Ad esempio, in Linux i diritti di accesso sono lettura, scrittura, esecuzione.
```
NB: un soggetto può avere diritti di accesso anche per altri soggetti (es: processo che
controlla un altro processo.
```
Ad ogni soggetto è associato un dominio, che rappresenta l'*ambiente di protezione* nel quale il soggetto esegue e specifica i diritti di accesso posseduti da tale soggetto nei confronti di ogni risorsa.
Un dominio di protezione è unico per ogni soggetto, mentre un soggetto (ad esempio un processo) può cambiere dominio durante la sua esecuzione.

#### Politiche
Le politiche di protezione definiscono le regole con le quali i soggetti possono accedere agli oggetti. Mentre il modello è qualcosa di insito nel sistema, le politiche generalmente vengono scelte da chi opera su quel sistema. Si classificano in 3 diverse tipologie:
- **DAC** (Discretional Access Control) - il creatore di un oggetto controolla i diritti di accesso per quell'oggetto (tipologia adottata da UNIX, che fornisce un meccanismo per definire e interpretare per ciascun file i 3 bit di read, write ed execute, per il proprietario, il gruppo e gli altri). La definizione delle politiche è decentralizzata.
- **MAC** (Mandatory Access Control) - i diritti d'accesso vengono definiti in modo centralizzato. Questa soluzione viene utilizzata in sistemi di alta sicurezza per garantire assoluta confidenzialità e i diritti vengono gestiti da un'entità centrale.
- **RBAC** (Role Based Access Control) - ad un ruolo vengono assegnati specifici diritti di accesso sulle risorse. Gli utenti possono appartenere a diversi ruoli. I diritti attribuiti ad ogni ruolo vengono assegnati in modo centralizzato.
**Principio del Privilegio Minimo** (o POLA - Principle Of Least Authority): ad ogni soggetto sono garantiti i diritti di accesso solo agli oggetti strettamente necessari per la sua esecuzione. Questa è una caratteristica desiderabile per tutte le politiche di protezione.

#### Meccanismi
I meccanismi di protezione sono gli strumenti messi a disposizione dal sistema di protezione per imporre una determinata politica.
Principi di realizzazione:
- **flessibilità del sistema di protezione**: i meccanismi di protezione devono essere sufficientemente generali per consentire per consentire l'applicazione di diverse politiche di protezione;
- **separazione tra meccanismi e politiche**: la politica definisce *cosa va fatto* ed il meccanismo *come va fatto*. È desiderabile la massima indipendenza tra le due componenti.

### Dominio di Protezione
Un dominio definisce un insieme di coppie, ognuna contenente l'identificatore di un oggetto e l'insieme delle operazioni che il soggetto associato a tale dominio può eseguire su ciascun oggetto (diritti di accesso).

Es: D(S) = {\<o, diritti\> | o è un oggetto, diritti è un insieme di operazioni}

Ogni dominio è associato univocamente ad un soggetto, mentre un soggetto (ad esempio un processo) può eventualmente cambiare dominio durante la sua esecuzione; il soggetto può accedere solo agli oggetti definiti nel suo dominio, utilizzando i diritti specificati dal dominio.

Domini disgiunti o con diritti di accesso in comune: esiste la possibilità per due o più soggetti di effettuare alcune operazioni comuni su un oggetto condiviso. Le operazioni vengono svolte dai processi che operano per conto di soggetti (a cui sono associati i domini).
```
NB: in ogni istante della sua esecuzione, il processo esegue in uno ed un solo dominio.
```

#### Associazione tra Processo e Dominio
L'associazione tra processo e dominio può essere statica o dinamica.<br/>
**Statica**: l'insieme delle risorse disponibili ad un processo rimane fisso durante il suo tempo di vita. Osservazioni: questo tipo di associazione non è adatta al Principio del Privilegio Minimo, in quanto l'insieme globale delle risorse che un processo potrà usare può non essere un'informazione disponibile prima della sua esecuzione; inoltre, l'insieme minimo di risorse necessarie ad un processo per garantire tale Principio, può cambiare in modo dinamico durante l'esecuzione.<br/>
**Dinamica**: l'associazione tra processo e dominio varia durante l'esecuzione del processo. In questo modo si può mettere in pratica il *Principio del Privilegio Minimo*, in quanto in ogni sua fase di esecuzione il processo può acquisire diritti diversi (ovvero solo quelli strettamente necessari). Tuttavia in questo caso *occorre un meccanismo per consentire il passaggio da un dominio all'altro* del processo.

##### Esempio di cambio di dominio
**Standard dual mode** (kernel/user mode): ci sono due domini (ring) di protezione, quello dell'utente (user mode) e quello del kernel (monitor o kernel mode). Quando un processo deve eseguire un'istruzione privilegiata, chiama una system call ed avviene il cambio di dominio. Questo non realizza la protezione tra utenti, ma solo tra kernel e utente.

**UNIX**: il dominio è associato all'utente. Ogni *processo* è caratterizzato dall'attributo UserID (UID). Il cambio di dominio corrisponde al cambio temporaneo di identità (UID) del processo.
Ad ogni *file* invece sono associati il proprietario (user-id) ed un bit di dominio (set-uid). Se il un file ha il bit set-uid settato, quando un utente B, diverso dal proprietario A di tale file, ne lancia l'esecuzione, al processo che esegue viene assegnato lo user-id dell'utente B. Così facendo il file entra nel dominio di B.

### Matrice degli Accessi
Un sistema di protezione può essere rappresentato a livello astratto utilizzando il modello della matrice degli accessi: tale matrice consente di rappresentare il modello e le politiche valide nel sistema considerato.
- ogni riga rappresenta un soggetto (utente);
- ogni colonna rappresenta un oggetto (risorsa, file)
- ogni elemento rappresenta i diritti accordati ai soggetti sugli oggetti.
Le informazioni contenute nella matrice possono variare nel tempo: le informazioni contenute nella matrice all'istante t rappresentano lo stato di protezione del sistema in t. La matrice degli accessi offre ai meccanismi di protezione le informazioni che consentono di verificare il rispetto dei vincoli di accesso. Il meccanismo associato al modello:
- ha il compito di verificare se ogni richiesta di accesso che proviene da un processo (che opera in un determinato dominio) è consentita oppure no;
- autorizza l'esecuzione delle richieste consentite e impedisce quelle vietate;
- consente di modificare dinamicamente e in modo controllato il cambiamento dello stato di protezione.

#### Verifica del Rispetto dei Vincoli di Accesso
Il meccanismo consente di assicurare che un processo che opera nel dominio *Di* possa accedere solo agli oggetti specificati nella riga *i* e solo con i diritti di accesso indicati. Quando un'operazione M dev'essere eseguita nel dominio *Di* sull'oggetto *Oj*, il meccanismo consente di controllare che M sia contenuta nella casella *access(i,j)*. In caso affermativo l'operazione può essere eseguita, altrimenti viene restituito un errore.

#### Modifica dello Stato di Protezione
In base alla politica di protezione adottata, lo stato di protezione può essere modificato da entità differenti:
- nella politica DAC (es: sistemi UNIX-based) può essere modificato dai soggetti, ovvero gli utenti;
- nella politica MAC può essere fatto solo dall'entità centrale.

##### Modello Graham-Denning
Questo modello stabilisce quali sono i comandi che consentono una modifica dello stato di protezione, identificando 8 primitive:
- create object, aggiunge una colonna;
- delete object, rimuove una colonna;
- create subject, aggiunge una riga;
- delete subject, rimuove una riga;
- read access right, legge il diritto d'accesso;
- grant access right, assegna il diritto d'accesso;
- delete access right, rimuove il diritto d'accesso;
- transfer access right, propaga il diritto d'accesso.

##### Propagazione dei Diritti di Accesso (Copy Flag)
La possibilità di copiare un diritto di accesso per un oggetto da un dominio ad un altro nella matrice di accesso è indicata con un asterisco * (*copy flag*).
Un soggetto *Sa* può trasferire un diritto di accesso *a* (ad esempio 'read') per un oggetto *Ox* ad un altro soggetto *Sb* solo se *Sa* ha accesso a *Ox* con il diritto *a* e tale diritto ha il copy flag (ovvero solo se nella tabella delle matrici, l'elemento *Sa*\\*Ox* contiene *a*\*, ad esempio read*)

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/02%20-%20Protezione/Matrice%20degli%20Accessi%20Copy%20Flag.png" alt="Matrice degli Accessi e Copy Flag"/>

L'operazione di propagazione può essere realizzata in due modi:
- **trasferimento** del diritto, il soggetto iniziale perde il diritto di accesso, che viene spostato al nuovo soggetto;
- **copia** del diritto, il soggetto iniziale mantiene il diritto di accesso, "duplicandolo" al nuovo soggetto.

<!-- lezione 2021-09-29 -->

