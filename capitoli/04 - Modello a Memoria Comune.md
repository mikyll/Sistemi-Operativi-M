
<h1 align="center">Capitolo 4: Modello a Memoria Comune</h1>

Un'app concorrente è vista come un insieme di processi che possono eventualmente accedere a degli oggetti, che sono le risorse (vedi figura slide 3). Processi e oggetti possono essere connessi tramite archi orientati che esprimono il diritto di accesso.
In linea di principio, ci sono dunque delle risorse comuni, ovvero che possono essere accedute da più processi (in questo caso O2/O3, oppure risorse "private" acceedibilli solo da un solo processo (O1, O4).

Nel modello a memoria comune, il tipo di interazione è **competitivo**.
In un sistema concorrente si hanno componenti passivi (oggetti o risorse) e componenti attive, che sono i processi, che compiono accesso a componenti passive.

Risorsa: qualunque oggetto fisico o logico, di cui un processo necessita per portare a termine il suo compito. Sono raggruppate in classi, le quali identificano 
Le risorse possono essere raggruppate in classi, che identificano l'insieme di tutte e sole le operazioni che un processo può eseguire sulle risorse che appartengono a quella classe.
Nel modello a memoria comune, il termine risorsa si identifica come una struttura dati (insieme di informazioni) allocata nella memoria comune. Ciò vale sia per risorse fisiche che logiche.

#### Meccanismo di Controllo degli Accessi

In qualunque sistema operativo è comunque previsto un meccanismo di controllo degli accessi, secondo cui quando un processo richiede l'accesso a una risorsa, tale accesso viene verificato (che avvenga correttaemente) e controllato da tale meccanismo.
Ad ogni risorsa R è associato un gestore. Il **gestore** ha il compito di definire ad ogni istante t dell'esecuzione l'*insieme dei processi SR(t) che nell'istante t hanno il diritto di operare su R*.

Noto il gestore ed SR(t), possiamo introdurre una classificazione che distingue tra risorse dedicate e risorse condivise, allocate staticamente e dinamicamente.
- Risorsa **dedicata**. Se SR(t) ha cardinalità SEMPRE <=1.
- Risorsa **condivisa**. Se la risorsa non è dedicata.
- Risorsa **allocata staticamente**. Se SR(t) = SR(t0), per ogni t, ovvero SR non dipende da t.
- Risorsa **allocata dinamicamente**. 
Altra possibile distinzione: allocazione statica vs dinamica.
R è alloc staticamente se SR(t) non dipende da t, ovvero SR(t) = SR(t0) per ogni t.
R è alloc dinamicamente se è tempo dipendente.

immagine slide 9

caso A è l'unico in cui la risorsa è privata, negli altri la risorsa è comune; nel caso C ad ogni istante non ci sarà mai più di un processo che può esercitare il diritto di accesso su quella risorsa. 
caso B ad ogni istante ci possono essere più di un processo che esercitano il diritto di accesso alla stessa risorsa. 
D più processi che possono esercitare contemporaneamente il diritto di accesso alla risorsa, è forse il caso più generale.
La colonna di destra individua il caso del modello a memoria comune: SR(t) ha cardinalità >1 e bisogna implementare un meccanismo di sincronizzazione e di comunicazione.

Nel caso di risorsa alloc staticamente, l'insieme SR(t) è stabilito dal programmatore a tempo di compilazione.
nel caso di risorsa alloc dinamicamente, il gestore opera a runtime, dunque viene essere chiamato in causa

##### Compiti del gestore di una risorsa
Se siamo nel modello a memoria comune, il gestore è a sua volta una risorsa condivisa.
Ricordare il Monitor: un oggetto che regola opportunamente gli accessi a risorse condivise. Il Monitor è un esempio di gestore:
1. Gestisce l'insieme dei processi e mantiene aggiornato SR(t), ovvero lo stato di allocazione della risorsa;
2. Fornisce i meccanismi che un processo può utilizzare per acquisire il diritto di accesso alla risorsa, e rilasciare tale diritto;
3. implementa la stategia di allocazione della risorsa, definendo quando, a chi e per quanto tempo allocare la risorsa.

Nel modello a scambio di messaggi, invece, il gestore è un processo a sua volta. Gestisce la risorsa per farla sembrare condivisa, ma in realtà è privata al gestore (al processo del gestore), che scambia messaggi con gli altri processi quando ne richiedono l'accesso.


##### Accesso a risorse
Considerando un processo P che deve operare ad un certo istante t su una risorsa R:
- Se R è allocata *staticamente* a P (modalità A e B), se P appartiene a SR, esso possiede il diritto di operare su R in qualunque istante.

```
R.op();	// esecuzione dell'operazione op() su R
```

- Se R è allocata *dinamicamente* a P (modalità C e D), è necessario prevedere un gestore (della risorsa R) GR che implementa le funzioni di richeista e rilascio della risorsa.

```
GR.richiesta();	// richiesta accesso alla risorsa R
R.op();			// esecuzione dell'operazione op su R
GR.rilascio();	// rilascio della risorsa R
```

- Se R è allocata come *risorsa condivisa* (modalità B e D), è necessario assicurare che gli accessi avvengano in modo non divisibile, ovvero ci dev'essere mutua esclusione.

```
NB: le funzioni di accesso alla risorsa devono essere programmate come una classe di sezioni critiche (utilizzando i meccanismi di SINCRONIZZAZIONE offerti dal linguaggi di programmazione e supportati dalla macchina concorrente).
```

- Se R è allocata come risorsa dedicata (modalità A e C), essendo P l'unico processo che accede alla risorsa, non è necessario prevedere alcuna forma di sincronizzazione (non c'è competizione).

#### Specifica della Sincronizzazione

Usiamo un formalismo che ci permette di esprimere la specifica di qualunque vincolo di sincronizzazione: la *regione critica condizionale*.

**Regione critica condizionale**: ```region R << Sa; when(C); Sb; >>```\

Il *corpo* della region (racchiuso tra le virgolette) rappresenta un'operazione da eseguire sulla risorsa condivisa R e quindi costituisce una sezione critica che dev'essere eseguita in mutua esclusione con le altre operazioni definite su R. È costituito da due istruzioni da eseguire in sequenza: l'istruzione Sa e, successivamente, l'istruzione Sb. Una volta terminata l'esecuzione di Sa, si valuta la condizione C: se è vera si prosegue con l'operazione Sb. Se è falsa, il processo che ha invocato l'esecuzione della region attende che C diventi vera.

###### Casi Particolari

- **region R <<S;>>**, specifica della sola mutua esclusione senza che sia prevista alcuna forma di sincronizzazione diretta;

- **region R <<when(C)>>**, specifica di un semplice vincolo di sincronizzazione, quindi il processo attende che C sia verificata prima di eseguire;

- **region R <<when(C) S;>>**, specifica il caso in cui la condizione C di sincronizzazione caratterizza lo stato in cui la risorsa R deve trovarsi al fine di poter eseguire l'operazione S.

Esempio: scambio di informazioni fra processi

```
T buffer;
bool pieno;

void inserisci(T dato):
region M <<when(pieno == false) buffer = dato; pieno = true;>>

T estrai():
region M <<when(pieno == true) pieno = false; return buffer;>>
```

#### Mutua Esclusione

Il problema della mutua esclusione nasce quando più di un processo alla volta può avere accesso alle stesse variabili comuni. La *regola dei mutua esclusione* impone che le operazioni con le quali i processi accedono alle variabili comuni non si sovrappongano nel tempo. Nessun vincolo è imposto sull'ordine con il quale le operazioni sulle variabili vengano eseguite.

**Sezione Critica**: sequenza di istruzioni con le quali un processo accede e modifica un insieme di variabili comuni.
Ad un insieme di variabili comuni possono essere associate una sola sezione critica (usata da tutti i processi) o più sezioni critiche (classe di sezioni critiche).

**Regola di Mutua Esclusione**: sezioni critiche appartenenti alla stessa classe devono escludersi mutuamente nel tempo. Oppure: ad ogni istante può essere "in esecuzione" al più una sezione critica di ogni classe.

Protocollo per la mutua esclusione: per specificare una sezione critica S che opera su una risorsa condivisa R, utilizziamo il seguente blocco:

```
<prologo>
S;
<epilogo>
```