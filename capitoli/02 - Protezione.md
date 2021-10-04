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
##### Diritto Owner
Il diritto *owner* realizza il concetto di "proprietario di una risorsa" (oggetto). Il soggetto che possiede tale diritto di accesso, nei sistemi che lo prevedono, ha la possibilità di concedere/revocare un qualunque diritto di accesso sull'oggetto che gli appartiene (ovvero possiede il diritto owner su tale oggetto) ad un qualunque altro soggetto.<br/>
In una matrice degli accessi, ciò si traduce nella presenza, in ciascuna colonna, di una ed una sola cella nella quale è presente un diritto owner. Per ogni risorsa (dunque per ogni colonna) ci dev'essere un solo soggetto che ne è il proprietario. Ciò significa che tale soggetto ha un ruolo privilegiato nei confronti di quella risorsa ed è l'unico soggetto capace di revocare o concedere diritti di accesso su quella risorsa ad altri soggetti.
Ad esempio, se *S2* ha il diritto 'owner' su *O2* allora può revocare il diritto 'execute' su *O* al soggetto *S1*.

*foto slide 30*

##### Diritto Control
Il diritto *control* permette di revocare un qualunque diritto di accesso, riferendosi non ad un oggetto ma ad un altro soggetto.
Se la cella della colonna di un soggetto ha il diritto control, autorizza l'owner a modificare la riga associata a tale soggetto.<br/>
Ad esempio, se *S1* ha il diritto di 'control' su *S2* e *S2* ha il diritto di 'write' su *O3*, allora *S1* può revocare il diritto di write di *S2* su *O3*.

*foto slide 31*

```
NB: In un sistema MAC, avremmo l'entità centrale che è l'unica autorizzata a stabilire
che cosa si può o non si può fare sulle risorse del sistema. Questa autorità è un sog-
getto che ha il diritto di control su tutti gli altri soggetti (ha il potere assoluto,
nel senso che può modificare il contenuto della riga associata ad ogni altro soggetto).
```

Copy flag, owner e control sono strumenti con cui possiamo *modificare* il contenuto della matrice degli accessi (possiamo aggiungere o togliere diritti nelle varie caselle). Se ci pensiamo bene, una riga della matrice altro non è che il dominio di protezione associato ad un soggetto.

##### Diritto Switch
Uno dei modi possibili per garantire il rispetto del Principio del Privilegio Minimo è prevedere la possibilità di un cambio di dominio a runtime. Un processo che esegue nel dominio di un soggetto, se ha bisogno di diritti differenti, può spostarsi nel dominio specifico che gli garantisce questi diritti nuovi di cui ha bisogno.<br/>
A tal proposito, esiste un diritto speciale chiamato *switch*: esercitando questo diritto, il processo che esegue nel dominio di un certo soggetto, può passare ad un nuovo dominio.
Ad esempio, se il soggetto *S2* ha bisogno del diritto 'write' sull'oggetto O3, ma possiede solo 'read', se possiede il diritto di 'switch' *S1* può passare nel dominio di *S1* per acquistare il diritto di 'write' e accedere a tale risorsa in scrittura.

*foto slide 32*

Nel mondo UNIX il diritto di switch è implementato tramite il bit set-uid: se tale bit è settato, il processo che esegue il file può ottenere il dominio del soggetto proprietario.


#### Realizzazione della Matrice degli Accessi
La matrice degli accessi è una notazione astratta che rappresenta lo *stato* di protezione. Parliamo di stato in quanto sia la forma della matrice che il suo contenuto cambiano dinamicamente nel tempo: ogni colonna rappresenta una risorsa (es: un file), e le risorse vengono create, distrutte e modificate continuamente nel corso dell'esecuzione del sistema operativo.

Nella rappresentazione concreta è necessario considerare:
- la **dimensione** della matrice, ovvero quale può essere la sua dimensione massima (poiché ogni colonna è un file, ed il numero di file in un file system può raggiungere un ordine di centinaia di migliaia di oggetti, ovviamente il numero delle colonne sarà molto elevato);
- il fatto che sia una matrice **sparsa**, ovvero non è detto che in tutte le celle vi siano delle informazioni. Anzi, al contrario, specialmente se si hanno molti file, possiamo trovare diverse celle vuote, in quanto non è detto che tutti i soggetti abilitati al sistema abbiano qualche diritto su ciascun file, dunque tipicamente vi sarà una prevalenza di celle vuote.

La rappresentazione più intuitiva sarebbe riservare dello spazio in memoria pari al numero totale di soggetti per il numero totale di oggetti, ma questa soluzione non è ottimale (fa schifo).
La rappresentazione concreta dello stato di protezione dev'essere ottimizzata sia per quanto riguarda l'occupazione della memoria, sia rispetto all'efficienza nell'accesso e nella gestione delle informazioni di protezione. Per questo motivo esistono 2 approcci:
- **Access Control List (ACL)**, si basa su una *rappresentazione per colonne*. Per ogni oggetto (risorsa) si mantiene una struttura dati che concettualmente rappresenta una colonna della matrice, ma tenendo conto che la colonna, come la matrice, è sparsa, in questa linea si mantengono solo gli elementi significativi, ovvero quello non vuoti. La struttura associata ad un oggetto è dunque una *lista in cui ogni elemento è un soggetto che ha un qualche diritto su quell'oggetto*.
- **Capability List (CL)**, si basa su una *rappresentazione per righe*. Ad ogni soggetto è associta una lista che indica quali sono gli oggetti al quale il soggetto può accedere.

##### ACL: Lista degli Accessi
Ne viene assegnata una a ciascun oggetto ed ha una struttura composta da un insieme di elementi, ognuno dei quali contiene la coppia <soggetto, insieme dei diritti> limitatamente ai soggetti con un insieme non vuoto di diritti per l'oggetto.
Quando un qualunque soggetto *S* tenta un'operazione *M* su un oggetto *O*, il sistema di protezione va a verificare nella ACL associata ad *O* se è presente un elemento riferito al soggetto che sta tentando l'accesso e, se esiste, controlla che contenga il diritto per eseguire *M* (ovvero sia presenta la coppia <*S*, *Rk*> con *M* appartenente a *Rk*).<br/>
In certi casi, per velocizzare l'accesso, viene prevista una lista di default: se è prevista, esistono dei diritti comunia a tutti i soggetti, dunque si va a vedere prima nella lsita di default e, se la ricerca non va a buon fine, si va a vedere nello specifico, elemento per elemento. Chiaramente, se la ricerca non ha successo, l'accesso viene negato.

**Utenti e Gruppi**: molti sistemi, per identificare un soggetto, prevedono non solo il nome utente (UID - User IDentifier), ma anche il gruppo (GID - Group IDentifier) a cui appartiene. Un gruppo aggrega un insieme di utenti, ha un nome, e può essere incluso nelle ACL. È importante sapere che un utente può appartenere anche a più gruppi, ma in un certo istante può appartenere ad un solo gruppo alla volta. Nel caso siano presenti i gruppi, una entry della ACL ha la forma UID, GID: \<insieme di diritti\>.
Quando un soggetto <utente, gruppo> prova ad accedere ad una risorsa in un certo modo, il tentativo di accesso comporta una ricerca nell'ACL dell'oggetto: se compare il modo con il quale l'utente sta cercando di accedere alla risorsa, allora l'operazione viene consentita.

Tipicamente i gruppi esistono per una questione di differenziazione dei ruoli, ciascuno dei quali ha diritti diversi. In generale, nei sistemi che prevedono i gruppi, è comunque possibile svincolare i soggetti dai gruppi, ovvero senza assegnargli alcun gruppo (UID, * \<insieme diritti\>).

##### CL: Capability List
Ne viene associata una a ciascun soggetto ed ha una struttura composta da un insieme di elementi, ognuno dei quali contiene l'indicazione dell'oggetto ed i diritti di accesso che quel soggetto, al quale la CL è associata, può esercitare su quell'oggetto. Chiaramente non avrà tanti elementi quanti sono le colonne della matrice degli accessi, in quanto come detto è una matrice sparsa.
Nella pratica, spesso l'oggetto viene identificato tramite un descrittore ed ovviamente i diritti, che spesso vengono rappresentati in modo compatto tramite una sequenza di bit.

*foto dimostrativa slide 40* - struttura tipica CL

In generale è importante (anche per le ACL), che le informazioni relative alla protezione vengano protette a loro volta da manomissioni. Ci sono vari modi:
- limitazione degli accessi in scrittura al solo kernel del Sistema Operativo (si sfrutta la protezione a livello hardware, oppure i modi di esecuzione del processore: kernel mode / user mode). In questo modo l'utente fa riferimento ad un puntatore (capability) che identifica la sua posizione nella lista appartenente allo spazio del kernel (soluzione simile all'uso di file descriptor in UNIX);
- se l'hardware lo supporta si può utilizzare un'architettura etichettata: a livello hardware ogni singola parola ha bit extra (tag) che esprimono la protezione su quella cella di memoria. In questo modo si può proteggere la protezione proprio a livello basso, a livello di memoria. In architetture di questo tipo il processore sa che, durante lo svolgimento delle sue operazioni di aritmetica, i bit di etichetta devono essere ignorati (non vengono proprio considerati).

#### Revoca dei Diritti di Accesso
In un sistema di protezione dinamica può essere necessario *revocare* i diritti di accesso per un oggetto. Revocare significa negare dei diritti precedentemente concessi. Ne esistono diversi tipi:
- revoca **generale** o **selettiva**, ovvero rispettivamente da un certo momento in poi nessuno potrà accedere ad un determinato oggetto, oppure solo alcuni soggetti non potranno più accedervi (ad esempio quelli appartenenti ad un gruppo);
- revoca **totale** o **parziale**, ovvero rispettivamente riguardante tutti i diritti per l'oggetto, oppure solo un particolare sottoinsieme di diritti;
- revoca **permanente** o **temporanea**, ovvero rispettivamente il diritto di accesso revocato non sarà più disponibile, oppure potrà essere successivamente riacquistato.

**Revoca per un oggetto con ACL**: la revoca in questo caso risulta semplice, in quanto si fa riferimento alla ACL associata all'oggetto in questione e si cancellano i diritti di accesso che si vogliono revocare.

**Revoca per un oggetto con CL**: l'operazione risulta più complessa, in quanto è necessario verificare, per ogni dominio, se contiene la capability con riferimento all'oggetto considerato.
```
NB: cancellare un file significa togliere i diritti su quel file a tutti gli utenti che
li possiedono. In ACL cancello semplicecmente l'ACL del sistema, mentre con CL bisogna
fare ricerca in tutte le CL per vedere se esiste un elemento riferito a quel file e in
caso cancellarlo.
Sicuramente queste operazioni che riguardano un solo oggetto risultano più costose in
sistemi con CL.
```

#### ACL vs CL
Un sistema realizzato esclusivamente con CL può soffrire di un appesantimento dovuto a operazioni che riguardano revoche che interessano più soggetti e quindi causano overhead (costo computazione maggiore).
Naturalmente si può fare anche il discorso duale: se si ha la necessità di fare una modifica allo stato di protezione che interessa un particolare soggetto. Il caso più banale è eliminare il soggetto dal sistema. Ciò questo comporta una modifica allo stato di protezione, in CL si cancella semplicemente la lista associata al soggetto; in ACL bisogna fare ricerca in ogni ACL.<br/>
Dunque non c'è soluzione assoluta, o generalmente migliore dell'altra: nella realtà, nella maggior parte dei sistemi solitamente si usa una soluzione ibrida che combina i due metodi.

Ad esempio, in UNIX, per ogni risorsa (file, in quanto UNIX è file-centrico, ovvero tutte le risorse sono presenti nel filesystem come file) per ogni oggetto viene mantenuta una struttura contenente 12 bit "di protezione". Sono memorizzati sul disco e fanno parte dell'i-node, che è rappresentato sulla memoria di massa all'interno dell'i-list.
Se ci pensiamo bene sono una forma semplificata di ACL: i 12 bit (di cui 9 per i diritti di utente, gruppo e altri) sono una forma semplificata di ACL, in quanto esprimono cosa gli utenti possono o non possono fare. È semplificata in quanto non si ha il dettaglio di un particolare utente, a differenza di Sistemi Operativi come Windows, ma solo dell'utente proprietario (oltre che del gruppo e degli altri utenti).

**Soluzione ibrida**: ogni volta che un processo cerca di accedere ad un oggetto (file), questo viene aperto e nella tabella dei file aperti viene caricato un elemento che altro non è che la capability che il soggetto che l'ha aperto ha nei confronti di quell'oggetto.

Esempio: quando cerco di aprire un file in scrittura, viene fatta una verifica sulla ACL (che controlla se posso effettuare quell'operazione); se la verifica va a buon fine, il file viene aperto e nella tabella dei file aperti viene caricato un elemento che concettualmente rappresenta il diritto di quel processo ad accedere in scrittura sul file; quindi viene aggiunta una capability alla tabella dei file aperti. Di fatto la tabella dei file aperti è una CL.

**Differenze**: la ACL si trova in memoria secondaria in modo persistente, la CL è in memoria volatile (e solitamente ha vita più breve): quando il processo finisce di operare sul file, l'elemento viene rimosso e quando il processo termina la tabella dei file aperti viene distrutta.

**Vantaggio**: una volta verificato preliminarmente che sia presente il diritto d'accesso, non c'è più bisogno di consultare la ACL, ma si va a guardare la CL.

### Protezione Multilivello
Come detto, la protezione riguarda il controllo degli accessi alle risorse interne al sistema; la sicurezza riguarda il controllo degli accessi al sistema. Poiché la protezione di un sistema può essere inefficace, se un utente non autorizzato riesce a far eseguire programmi che agiscono sulle risorse del sistema (es: Trojan, o Cavalli di Troia <!-- indotti, con intenzioni malevole, in qualche modo nel filesystem, una volta qui inducono un utente autorrizzato ad eseguire quel programma e provocano dei danni -->), è necessario affiancarvi un sistema di sicurezza, che normalmente ha una *struttura multilivello*.

Il sistema di sicurezza stabilisce delle regole più generali rispetto al sistema di protezione, in cui prima di tutto si classificano gli utenti (ad esempio in funzione del loro ruolo), dopodiché gli oggetti (le risorse). In funzione della confidenzialità dell'oggetto, vengono collocate ad un livello diverso del sistema. In un sistema di questo tipo l'approccio è quello di tipo MAC.

#### Modelli di Sicurezza Multilivello
I modelli di sicurezza multilivello più usati sono due:
- **Bell-La Padula** - obbiettivo di garantire la confidenzialità delle informazioni;
- **Biba** - è antitetico al precedente, e ha l'obbiettivo di garantire l'integrità delle informazioni.
Entrambi aderiscono allo stesso modello multilivello.

In un modello di sicurezza multilivello:
i soggetti (utenti) e gli oggetti (risorse) sono classificati in livelli (classi di accesso):
- livelli per i soggetti (**clearance levels**);
- livelli per gli oggetti (**sensitivity levels**).
```
NB: le regole di sicurezza fissano le regole di interazione tra livelli diversi.
```

##### Modello Bell-La Padula
Nato in ambito militare, ha come obbiettivo primario garantire la **confidenzialità** delle informazioni.
Abbiamo un sistema di protezione (matrice accessi) a cui viene affiancato un modello multilivello che viene gestito con approccio di tipo MAC.
Vi sono 4 diversi livelli di sensibilità degli oggetti:
1. non classificato (+ basso);
2. confidenziale
3. segreto
4. top secret (+ alto)

Questi sono i livelli in cui verranno classificati i documenti.
Se voglio che un documento sia disponibile solo a chi si trova ai vertifici della gerarchia (es. Generale) lo metterò top secret.

Vi sono 2 regole di sicurezza, che caratterizzano il modello, che stabiliscono il verso di propagazione delle informazioni nel sistema:
1. **proprietà di semplice sicurezza**: un processo in esecuzione ad un livello di sicurezza k può **leggere** oggetti a suo livello o a livelli inferiori;
2. **proprità \* (star)**: un processo in esecuzione a livello di sicurezza k può **scrivere** solo oggetti al suo livello o superiori.
Il flusso delle informazioni è dunque dal basso verso l'alto.

*foto dimostrativa slide 53*

**Esempio di difesa contro Trojan per modello Bell-La Padula**: Bell-La Padula serve a impedire attacchi come questo. Supponiamo che i livelli di sicurezza siano 2: riservato e pubblico.
Se facciamo in modo che gli utenti siano classificati, nonostante l'ACL consenta l'accesso in scrittura, la politica di sicurezza lo impedisce (NB: la politica di sicurezza ha precedenza sui meccanismi di protezione).
Tuttavia, il modello Bell-La Padula è stato concepito per mantenere i segreti, non per garantire l'integrità dei dati.

##### Modello Biba
Ha come obbiettivo l'**integrità** dei dati. Serve per garantire che l'integrità delle informazioni di livello superiore venga in qualche modo preservata.
Anche in questo caso prevede 2 regole:
- **proprietà di semplice sicurezza**: è l'opposto di Bell-La Padula, in quanto stabilisce che un processo in esecuzione al livello di sicurezza k può scrivere solo oggetti al suo livello o a quelli inferiori (nessuna scrittura verso l'alto).
- **proprietà di integrità \* **: un processo in esecuzione a livello k può leggere solo oggetti al suo livello o a quelli superiori (nessuna lettura verso il basso).
Il flusso delle informazioni è dunque l'opposto del precedente: dall'alto verso il basso.
```
NB: chiaramente i modelli B-LP e BIBA sono in conflitto tra loro, quindi non pos-
sono essere combinati.
```

### Architetture dei Sistemi ad Elevata Sicurezza
**Sistemi Operativi Fidati**: sistemi per cui è possibile definire (e in certi casi dimostrare formalmente) determinati requisiti o regole di sicurezza.

**Reference Monitor (RM)**: elemento di controllo realizzato nell'hardware del sistema operativo, che regola l'accesso dei soggetti agli oggetti sulla base di parametri di sicurezza del soggetto e dell'oggetto (in pratica ha il compito di imporre il rispetto delle regole di sicurezza). Poiché ad ogni singola iterazione deve fare delle verifiche e ciò ha un costo, spesso si tende ad implementarlo, almeno in parte, a livello hardware.

**Trusted Computing Base (TCB)**: il RM ha accesso ad una base di calcolo fidata, chiamata *TCB*. Questa contiene delle informazioni che tracciano la classificazione dei soggetti e degli oggetti all'interno del sistema.

Il reference monitor deve imporre le regole di sicurezza (che sono ad esempio, nel caso del modello Bell-La Padula, "no read-up" e "no write-down") ed ha le seguenti proprietà:
- **mediazione completa**, ovvero le regole di sicurezza vengono applicate ad ogni singolo accesso da parte di un soggetto ad un particolare oggetto;
```
NB: non è così intuitiva, nei SO comuni, es. derivati da UNIX, le regole di prote-
zione vengono verificate solo a lato apertura del file. Se abbiamo un sistema fi-
dato invece, ad ogni singola operazione (ad esempio di scrittura) viene fatta una
verifica da parte del RM.
```
- **isolamento**, ovvero il RM e la TCB devono essere a loro volta protette da parte di eventuali accessi non autorizzati. Deve essere possibile accederli solo se ci si trova in modalità privilegiata;
- **verificabilità**, ovvero la correttezza del RM dev'essere dimostrata. Dev'essere possibile verificare/dimostrare formalmente che il monitor fa quello per cui è stato progettato. Questa proprietà non è semplicissima da applicare.

### Classificazione della Sicurezza dei Sistemi di Calcolo
*Orange Book* è un documento pubblicato dal Dipartimento della Difesa americano, in cui sono specificate 4 categorie di sicurezza (A, B, C, D, in ordine decrescente), e che consente quindi di etichettare i sistemi di calcolo sulla base delle caratteristiche di sicurezza.
(- sicuro) D < C < B < A (+ sicuro)

- **D - Protezione Minima**, non sono presenti meccanismi che consentono di esercitare sicurezza né protezione. Al giorno d'oggi non ne esistono. Esempio: MS-DOS, all'epoca i sistemi erano concepiti per essere mono-utente;
- **C - Protezione Discreta**, si suddivide in *C1* e *C2*:
    - *C1*, i sistemi prevedono dei meccanismi di:
        - autenticazione degli utenti (i dati di autenticazione sono protetti e non accessibili ad utenti non autorizzati);
        - protezione dei dati e programmi propri di ogni utente;
        - controllo degli accessi a oggetti comuni per gruppi di utenti definiti.
    Esempio: UNIX e sistemi derivati
    - *C2*, il controllo degli accessi è fatto su base individuale, non collettiva, al contrario di UNIX.
    Esempio: Windows.
- **B - Protezione Obbligatoria**, si suddivide in *B1*, *B2* e *B3*:
    - *B1*, come C2 ma con introduzione dei livelli di sicurezza (modello Bell-La Padula), almeno 2.
    - *B2*, si estende l'uso di etichette di riservatezza ad ogni risorsa del sistema, anche canali di comunicazione;
    - *B3*, la TCB consente la creazione di liste di controllo degli accessi in cui sono identificati utenti o gruppi cui non è consentito l'accesso ad un oggetto specifico.
- **A - Protezione Verificata (Massima Sicurezza)**, si suddivide in *A1* e classi superiori ad A1. È equivalente a B3, ma con il vincolo di essere progettato e realizzato utilizzando metodi formali di definizione e verifica (il RM è verificabile). Un sistema appartiene ad una categoria superiore ad A1 se è stato progettato in impianti di produzione affidabili, da persone affidabili.
