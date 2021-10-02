<h1 align="center"> Sistemi Operativi M</h1>

<p align="center">
  Appunti del corso Sistemi Operativi M (<a href="https://www.unibo.it/it/didattica/insegnamenti/insegnamento/2021/468009">72947</a>), anno 2021-2022. A cura di <a href="https://github.com/mikyll">Michele Righi</a> e <a href="https://github.com/TryKatChup">Karina Chichifoi</a>.
  <br/>
	<br/>
	<a href="https://github.com/mikyll/Sistemi-Operativi-M/tree/main/flashcard">Flashcards</a>
	·
  <a href="https://github.com/mikyll/Sistemi-Operativi-M/issues">Report a bug</a>
</p>

<details open="open">
  <summary><h2 style="display: inline-block">Indice</h2></summary>
  <ol>
    <li>
      <a href="#01---virtualizzazione">Virtualizzazione</a>
      <ul>
        <li><a href="#virtualizzazione-di-un-sistema-di-elaborazione">Virtualizzazione di un sistema di elaborazione</a></li>
        <li><a href="#emulazione">Emulazione</a>
          <ul>
            <li><a href="#interpretazione">Interpretazione</a>
            <li><a href="#compilazione-dinamica">Compilazione Dinamica</a>
          </ul>
        </li>
        <li><a href="#tipi-livelli-di-virtualizzazione">Tipi (Livelli) di Virtualizzazione</a></li>
        <!--<li><a href="#cenni-storici">Cenni Storici</a></li>-->
        <li><a href="#vantaggi-della-virtualizzazione">Vantaggi della Virtualizzazione</a></li>
        <li><a href="#realizzazione-del-vmm">Realizzazione del VMM</a>
          <ul>
            <!--<li><a href="#vmm-di-sistema">VMM di Sistema</a></li>
            <li><a href="#vmm-ospitato">VMM Ospitato</a></li>
            <li><a href="#ring-di-protezione">Ring di Protezione</a>
              <ul>
                <li><a href="#ring-deprivileging">Ring Deprivileging</a></li>
                <li><a href="#ring-compression">Ring Compression</a></li>
                <li><a href="#ring-aliasing">Ring Aliasing</a></li>
              </ul>
            </li>
            <li><a href="#supporto-hardware-alla-virtualizzazione">Supporto Hardware alla Virtualizzazione</a></li>
            <li><a href="#realizzazione-del-vmm-in-architetture-non-virtualizzabili">Realizzazione del VMM in Architetture Non Virtualizzabili</a></li>-->
            <li><a href="#fast-binary-translation-ftb">Fast Binary Translation (FTB)</a></li>
            <li><a href="#paravirtualizzazione">Paravirtualizzazione</a></li>
          </ul>
        </li>
        <!--<li><a href="#architetture-virtualizzabili">Architetture Virtualizzabili</a>
          <ul>
            <li><a href="#protezione-nellarchitettura-x86">Protezione nell'architettura x86</a></li>
            <li><a href="#funzionamento-dei-vmm-nellarchitettura-x86-classica">Funzionamento dei VMM nell'architettura x86 classica</a></li>
          </ul>
        </li>-->
        <li><a href="#gestione-di-vm">Gestione di VM</a></li>
        <ul>
            <li><a href="#stati-di-una-vm">Stati di una VM</a></li>
            <li><a href="#migrazione-di-una-vm">Migrazione di una VM</a></li>
        </ul>
      </ul>
    </li>
    <li>
      <a href="#02---protezione">Protezione</a>
    </li>
  </ol>
</details>

<details>
  <summary><h2 style="display: inline-block">Indice per lezioni</h2></summary>
  <ul>
    <li><a href="#01---virtualizzazione">2021/09/21</a></li>
  </ul>
</details>

<!-- Lezione 2021/09/21-->
## 01 - Virtualizzazione
La virtualizzazione è una tecnologia oggi usatissima. Virtualizzare un sistema di elaborazione (costituito da un insieme di risorse hardware e software) significa presentare all'utilizzatore una visione delle risorse diversa da quella attuale (ad esempio duplicazione della memoria). Questo obbiettivo viene raggiunto mediante un livello intermedio, un layer che svolge appunto il ruolo di intermediario tra utilizzatore (vista logica) e sistema (vista fisica). Solitamente l'intermediario è software, ma talvolta può avere un supporto hardware specifico. Esso permette di eseguire più macchine virtuali su una stessa architettura e ognuna di queste vede le proprie risorse, indipendentemente dalle altre, e da quelle effettive ("reali").
Poiché le Macchine Virtuali (VM) devono funzionare in modo indipendente senza causare problemi al sistema, la gestione delle risorse dev'essere realizzata in modo appropriato: questo compito è affidato al Virtual Machine Monitor (VMM, detto anche Hypervisor), cha ha compiti molto simili a quelli di un Sistema Operativo (SO), motivo per cui viene trattato in questo corso).

Esempi di virtualizzazione:
- **virtualizzazione a livello di processo** - i SO multitasking consentono l'esecuzione parallela di più processi, ognuno dei quali dispone di una VM (CPU, memoria, dispositivi) dedicata. Questo tipo è realizzato dal kernel;
- **virtualizzazione della memoria** - con memoria virtuale ogni processo vede uno spazio di indirizzamento di dimensioni indipendenti dalle reali dimensioni e dallo spazio effettivamente a disposizione. Anche questo è realizzato dal kernel del SO;
- **astrazione** - un oggetto astratto (risorsa virtuale) come rappresentazione semplificata di un oggetto (risorsa fisica). Solitamente, l'astrazione serve per fornire una vista delle sole proprietà significative, nascondendo i dettagli realizzativi non necessari, ad esempio i tipi di dato ad alto livello (numeri intero, numeri in virgola mobile, ecc.) rispetto alla loro rappresentazione binaria nella cella di memoria. Un altro esempio sono i linguaggi di programmazione, in cui un'istruzione di alto livello è un'astrazione di ciò che avviene apiù basso livello, in linguaggio macchina. Il disaccoppiamento è realizzato dalle operazioni (interfaccia) con le quali è possibile utilizzare l'oggetto;
- **linguaggi di programmazione** - la capacità di portare lo stesso programma (scritto in un linguaggio di alto livello) su architetture diverse è possibile grazie alla definizione di una VM in grado di interpretare ed eseguire ogni istruzione del linguaggio, indipendentemente dall'architettura del sistema (SO e hardware). Possibili esempi sono gli interpreti (ad esempio Java Virtual Machine, cacpace di eseguire il bytecode Java), ed i compilatori.

### Virtualizzazione di un sistema di elaborazione
Una singola piattaforma hardware viene condivisa da più elaboratori virtuali, ognuno gestito da un proprio sistema operativo. Il disaccoppiamento viene realizzato dal VMM, che è il mediatore unico tra le VM e l'hardware. I suoi compiti sono *consentire la condivisione da parte di più macchine virtuali di una singola piattaforma hardware*, garantendo *isolamento* tra di esse e *stabilità* del sistema. Il VMM deve realizzare una specie di sandbox per ciascuna VM (se ad esempio una va in crash, le altre non devono risentirne).

### Emulazione
Consiste nell'esecuzione di programmi compilati per una particolare architettura (e quindi con un certo set di istruzioni) su un sistema di elaborazione dotato di un diverso insieme di istruzioni. Vengono emulate interamente le singole istruzioni dell'architettura ospitata, consentendo completa interoperabilità tra ambienti eterogenei. Tramite l'emulazione, sistemi operativi o applicazioni pensati per determinate architetture, possono eseguire senza essere modificati, su architetture completamente differenti.
**Vantaggi**: interoperabilità tra ambienti eterogenei.
**Svantaggi**: problemi di efficienza (basse performance).
Nel tempo questo approccio si è ramificato seguendo due strade: interpretazione e ricompilazione dinamica.

#### Interpretazione
Consiste nella lettura di ogni singola istruzione del codice macchina che dev'essere eseguito e nell'esecuzione di più istruzioni sull'host virtualizzante per ottenere semanticamente lo stesso risultato.
**Vantaggi**: è un metodo generale e potente che presenta una grande flessibilità nell'esecuzione perché consente di emulare e riorganizzare i meccanismi propri delle varie architetture.
**Svantaggi**: produce un sovraccarico generalmente elevato poiché possono essere necessarie molte istruzioni dell'host per interpretare una singola istruzione sorgente.

#### Compilazione Dinamica
Invece una singola istruzione del sistema ospitato, vengono letti interi blocchi di istruzioni, che vengono tradotti (per la nuova architettura), ottimizzati e messi in esecuzione.
**Vantaggi**: migliori prestazioni rispetto al metodo precedente, in quanto si leggono interi blocchi di codice, che vengono tradotti ed ottimizzati, consentendo di sfruttare tutte le possibilità offerte dalla nuova architettura; inoltre, le parti di codice usate frequentemente possono essere bufferizzate per evitare di doverle ricompilare in seguito.
Tutti i più noti emulatori utilizzano questa tecnica per implementare l'emulazione.

### Tipi (Livelli) di Virtualizzazione
- Livello applicativo - applicazioni virtuali / supporto a tempo di esecuzione (es: JVM, .NET CLR)
- Livello di librerie (API a livello utente) - librerie virtuali (es: WINE, WABI, vCUDA)
- Livello di Sistema Operativo - container (Jail, Virtual Environment, Docker)
- Livello di Astrazione Hardware (HAL) - macchine virtuali (VMware, Virtual PC, Xen, User mode Linux). Questo è il tipo di virtualizzazione che astrae l'hardware ed è quello che ci interessa maggiormente in questo corso. Al contrario dei container, le VM non condividono lo stesso SO.
- Livello di Instruction Set Architecture(ISA) - ISA Virtuale (Bochs, QEMU)

### Cenni Storici
La virtualizzazione non è un concetto nuovo, bensì nasce negli anni '60 coi sistemi CP/CMS di IBM, dove il Control Program (CP) esegue direttamente sull'hardware svolgendo il ruolo di VMM ed il Conversational Monitor System (CMS) è il sistema operativo, replicato per ogni VM. Con la diffusione del consolidamento dell'hardware, si è passati dal paradigma "one application, one server" (tipico degli anni '80/'90, dovuto al crollo dei costi dell'hardware), ad avere, dagli anni 2000, un unico server bello grosso e potente, su cui installare 20/30 VM, ciascuna delle quali svolge un certo servizio: una soluzione molto più razionale, in quanto avere un numero di macchine fisiche ristrette permette di semplificare la configurazione, la gestione e la manutenzione. Per poi arrivare negli anni 2010 al Cloud Computing.

### Vantaggi della Virtualizzazione
La virtualizzazione comporta numerosi vantaggi:
- possibilità di avere più SO, anche differenti, sulla stessa architettura fisica;
- isolamento degli ambienti d'esecuzione, utile specialmente per eseguire e testare software dalla sicurazza e affidabilità non certa (nel caso peggiore la singola VM va in crash);
- abbattimento dei costi hardware, in quantosi possono concentrare più macchine (es. server) su un'unica architettura hardware, ed abbattimento dei costi di amministrazione;
- gestione facilitata delle macchine (creazione, installazione -esistono template già preimpostati-, amministrazione, migrazione "a caldo", possibilità di adottare politiche di bilanciamento del carico e robustezza -disaster recovery-)

### Realizzazione del VMM
Il VMM generalmente deve fornire ad ogni VM le risorse che gli servono per funzionare (CPU, memoria, dispositivi I/O).
I requisiti fondamentali sono i tre seguenti:
- l'ambiente d'esecuzione fornito dev'essere identico a quello della macchina reale (come se non fosse un sistema virtualizzato, ma girasse direttamente sull'architettura hardware);
- dev'essere garantita un'elevata efficienza (che sia accettabile) nell'esecuzione dei programmi;
- dev'essere garantita la stabilità e la sicurezza dell'intero sistema.
Due concetti molto importanti (che fungono anche da parametri per classificarlo) nella realizzazione del VMM sono: il **livello**, ovvero dove è collocato il VMM (può essere un *VMM di sistema* o un *VMM ospitato*); la **modalità di dialogo**, ovvero il modo in cui il VMM accede alle risorse (*virtualizzazione pura* o *paravirtualizzazione*)

In un sistema di virtualizzazione esistono due tipi di "componenti": l'**host** è la piattaforma sulla quale si realizzano le VM, ovvero il livello sottostante che comprende la macchina fisica ed il VMM; il **guest** è la VM vera e propria che comprende il Sistema Operativo e le applicazioni.

#### VMM di Sistema
Si trova direttamente sopra l'hardware e consiste in un Sistema Operativo molto leggero che realizza le funzionalità di virtualizzazione (es: kvm, xen). A meno che non ci sia abbastanza spazio libero sul disco e vi sia la possibilità di impostare un multiboot, per installare un VMM di sistema è necessario eliminare il Sistema Operativo preesistente.

<img width="60%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/01%20-%20Virtualizzazione/VMM%20di%20Sistema.png" alt="VMM di Sistema"/>

#### VMM Ospitato
Viene installato come una normale applicazione sul Sistema Operativo preesistente, opera nello spazio utente ed accede all'hardware tramite le system call del SO (es. VirtualBox). È più semplice da installare e per la gestione delle periferiche può fare riferimento al Sistema Operativo sottostante, ma ha performance peggiori.

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/01%20-%20Virtualizzazione/VMM%20Ospitato.png" alt="VMM Ospitato"/>

NB: d'ora in poi faremo sempre riferimento a VMM di sistema.

#### Ring di Protezione
L'architettura delle CPU prevede almeno due livelli di protezione (0 "modi di esecuzione"): supervisore/kernel (livello 0) e utente (livello >0). Ogni ring corrisponde ad una diversa modalità di funzionamento del processore:
- a livello 0 è possibile eseguire istruzioni privilegiate della CPU;
- a livello superiore a 0 non possono essere eseguite.

Alcuni programmi, come il kernel del SO, sono progettati per eseguire nel ring 0 (in cui si ha il pieno controllo dell'hardware).
In un sistema virtualizzato, il VMM dev'essere l'unica componente in grado di mantenere il pieno controllo dell'hardware (di conseguenza si troverà a ring 0, mentre le VM a ring >0).
Ma nella VM c'è il Sistema Operativo, il cui kernel non può eseguire a livello 0 come dovrebbe. Di conseguenza sorgono due principali problemi: *ring deprivileging* e *ring compression*.

##### Ring Deprivileging
Il SO della VM si trova ad eseguire in un ring che non gli è proprio e le istruzioni privilegiate richieste dal sistema operativo nell'ambiente guest non possono essere eseguite.
Possibile soluzione **trap & emulate**: se il guest tenta di eseguire un'istruzione prvilegiata (ad esempio 'popf', ovvero la disabilitazione delle interruzioni), la CPU lancia un'eccezione, che viene rilevata dal VMM (trap), al quale trasferisce il controllo; dopodiché il VMM controlla la correttezza dell'operazione e ne emula (emulate) il comportamento.
NB: se la popf potesse essere eseguita direttamente sul processore (a ring 0), verrebbero disabilitati gli interrupt per tutto il sistema ed il VMM non potrebbe riguadagnare il controllo della CPU, mentre il comportamento desiderato sarebbe che gli interrupt venissero sospesi solo per la VM in questione.

##### Ring Compression
Se ad esempio l'architettura ha solo 2 ring, poiché il primo (0) è assegnato al kernel del Sistema Operativo, applicazioni e SO della macchina virtuale eseguono allo stesso livello, con conseguente mancanza di isolamento e protezione.

##### Ring Aliasing
Alcune istruzioni non privilegiate, eseguite a livello user, permettono di accedere in lettura ad alcuni registri la cui gestione dovrebbe essere riservata al VMM, con conseguenti possibili inconsistenze. Ad esempio, il registro CS contiene il livello di privilegio corrente (se la VM pensa di essere in un certo ring, ma leggendo lo stato del registro vede che è sbagliato, potrebbero esserci dei problemi).

#### Supporto Hardware alla Virtualizzazione
L'architettura della CPU si dice **naturalmente virtualizzabile** (o con supporto nativo alla virtualizzazione) se prevede l'invio di trap allo stato supervisore (0) per ogni istruzione privilegiata invocata da un ring >0. In questi casi è possibile realizzare un approccio "trap & emulate" e si ha supporto nativo all'esecuzione diretta.
Tuttavia, non tutte le architetture sono naturalmente virtualizzabili (es. Intel IA32) e alcune istruzioni privilegiate non provocano una trap, anzi, in alcuni casi causano il crash del sistema.

#### Realizzazione del VMM in Architetture Non Virtualizzabili
Nel caso in cui un processore non fornisca supporto nativo alla virtualizzazione, è necessario ricorrere a soluzioni software. Alcune possibili sono: *fast binary translation*, *paravirtualizzazione*.

##### Fast Binary Translation (FTB)
Sfrutta un'idea simile alla compilazione dinamica: il VMM scansiona il codice dei SO guest prima dell'esecuzione, per sostituire a runtime i blocchi contenenti istruzioni privilegiate con blocchi equivalenti dal punto di vista funzionale, ma che contengano invece chiamate al VMM. I blocchi tradotti vengono salvati in cache per eventuali riutilizzi futuri.
**Vantaggi**: ogni VM è una esatta replica della macchina fisica, dunque è possibile installare gli stessi SO di architetture senza virtualizzazione nativa.
**Svantaggi**: la traduzione dinamica è costosa.

<img width="40%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/01%20-%20Virtualizzazione/Fast%20Binary%20Translation.png" alt="Fast Binary Translation"/>

##### Paravirtualizzazione
È l'approccio più diffuso al giorno d'oggi, oltre a FBT. Il VMM (hypervisor) offre ai SO guest un'interfaccia virtuale (hypercall API) alla quale i SO guest devono fare riferimento per avere accesso alle risorse. (NB: così come il *Sistema* Operativo fornisce delle *system* call, l'*Hyper*visor fornisce delle *hyper* call). Ciò consente  di eseguire istruzioni privilegiate chiamando direttamente la relativa hyper call, senza dover generare interrupt al VMM. Xen utilizza questa tecnica.
**Vantaggi**: la struttura del VMM è semplificata e si ottengono prestazioni migliori, in quanto non si ha il ritardo dovuto alla compilazione di FTB.
**Svantaggi**: vi è la necessità di porting dei SO guest (i kernel devono essere resi compatibili, soluzione che è preclusa a molti sistemi operativi proprietari, fra cui Windows).

<img width="34%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/01%20-%20Virtualizzazione/Paravirtualizzazione.png" alt="Paravirtualizzazione"/>

### Architetture virtualizzabili
Con **virtualizzazione pura**, si intende un'architettura che non costringe l'amministratore (o l'utente) a installare nella macchina virtuale un kernel modificato (che non sia l'originale del Sistema Operativo), dunque è il caso di architetture con supporto nativo alla virtualizzazione, ma anche FTB, in quanto anche lì non c'è bisogno di modificare il kernel.
**Vantaggi**: non c'è ring compression né ring aliasing (il guest esegue in un ring separato -intermedio- diverso da quello delle applicazioni; il ring deprivileging è risolto tramite trap & emualte (gestione tramite VMM); trasparenza (l'API presentata dall'hypervistor è la stessa offerta dal processore).
Prodotti virtualizzabili: xen, vmware, kvm.

#### Protezione nell'architettura x86
**Prima generazione**: non avevano nessuna capacità di protezione e non facevano distinzione tra SO e applicazioni, facendo girare entrambe con i medesimi privilegi. Il problema principale era dato dal fatto che, in questo modo, le applicazioni potevano accedere direttamente ai sottosistemi di I/O, allocare memoria senza alcun intervento del SO, con conseguenti problemi di sicurezza e stabilità (es. MS-DOS).

**Seconda generazione**: viene introdotto il concetto di protezione, con la distinzione tra SO, che possiede controllo assoluto sulla macchina fisica sottostante, e le applicazioni, che possono interagire con le risorse fisiche solo facendone richiesta al SO (concetto di ring di protezione).

**Registro CS**: i due bit meno significativi vengono riservati per rappresentare il Livello Corrente di Privilegio (CPL) => 4 possibili ring, a priorità crescente (0 maggiori privilegi, destinato al kernel del SO, ... , 3 minori privilegi, destinato alle applicazioni utente).
In questo modo si ottiene **protezione della CPU**: non è permesso a ring diversi dallo 0 di eseguire le istruzioni privilegiate, che son destinate solo al kernel del SO, in quanto considerate critiche e potenzialmente pericolose. Una qualsiasi violazione può provocare un'eccezione, gestita immediatamente dal SO, che può reagire, ad esempio, terminando l'applicazione in esecuzione.

**Segmentazione**: ogni segmento è rappresentato da un descrittore (in una tabella) in cui sono indicati il Livello di Protezione richiesto (PL) ed i permessi di accesso (r, w, x).
In questo modo si ottiene **protezione della memoria**: una violazione dei vincoli di protezione provoca un'eccezione. Ciò accade, ad esempio, se il valore di CLP è maggiore del PL del segmento di codice contenente l'istruzione invocata.


#### Funzionamento dei VMM nell'architettura x86 classica
Anche in questo caso è presente il problema del ring deprivileging, in quanto viene dedicato il ring 0 alla VMM e conseguentemente i SO guest vengono collocati in ring a privilegi ridotti. Vengono comunemente utilizzate 2 tecniche:
- 0/1/3: il SO viene spostato dal ring 0 (dove nativamente dovrebbe trovarsi) al ring 1, lasciando le applicazioni al ring 3, mentre al ring 0 viene installato il VMM;
- 0/3/3: il SO viene spostato direttamente al ring 3, dove si trovano anche le applicazioni, mentre sul ring 0 viene installato il VMM. In questa modalità non è possibile generare eccezioni, quindi devono essere intrapresi meccanismi molto sofisticati con un controllo continuo da parte del VMM.

### Gestione di VM
Il compito fondamentale del VMM è quello di gestire le VM (creazione, accensione/spegnimento, eliminazione, migrazione live).

#### Stati di una VM
Una macchina virtuale può trovarsi nei seguenti stati:
- **running** (o attiva): la macchina ha superato la fase di bootstrap ed è stata caricata nella *RAM* del server su cui è allocata;
- **inactive** (powered off): la macchina è spenta ed è rappresentata nel *file system* tramite un file immagine;
- **paused**: la macchina è in *attesa* di un evento (es: I/O richiesto da un processo nell'ambiente guest);
- **suspended**: lo stato correnteviene salvato nel file system dal VMM. L'uscita da tale stato avviene tramite un'operazione di *resume*.

suspend: il VMM salva lo stato della VM in memoria secondaria, mettendola in stand by;
resume: il VMM ripristina lo stato della VM in memoria centrale (lo stato è quello in cui si trovava quando è stata sospesa). Questa ooperazione può avvenire su un nodo diverso da quello della suspend.

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/01%20-%20Virtualizzazione/Stati%20di%20una%20VM.png" alt="Stati di una VM"/>

#### Migrazione di una VM
È una funzionalità necessaria soprattutto nei datacenter, per una gestione agile delle VM, a fronte di:
- variazioni dinamiche del carico (**load balancing**, consolidamento);
- **manutenzione "online"** dei server, senza dover interrompere i servizi forniti;
- gestione finalizzata al **risparmio energetico**;
- **tolleranza ai guasti** e disaster recovery.

Strumento fondamentale per queste procedure è la migrazione, ovvero la possibilità di muovere VM tra server.
**Migrazione live**: possibilità di spostare una VM da un server fisico ad un altro, senza doverla spegnere. È desiderabile minimizzare il downtime, il tempo di migrazione ed il consumo di banda.

<!-- Lezione 2021-09-28 -->
#### Soluzione: precopy
La soluzione più diffusa al giorno d'oggi si basa su un meccanismo di precopia, che viene attuata in una serie di passi:
1. **Pre-migrazione**: fase iniziale in cui si capisce quali sono i nodi interessati, ovvero si individua la VM da migrare (nodo A) e l'host di destinazione (nodo B);
2. **Reservation**: viene riservato un contenitore vuoto nel server di destinazione (reservation del posto per la macchina da migrare);
3. **Pre-copia iterativa delle pagine**: la VM da migrare chiaramente avrà un file immagine (tipicamente un file di stato mappato sui registri CPU). In questa fase viene eseguita una copia nell'host B di tutte le pagine allocate in memoria sull'host A. Poiché le pagine in memoria, soprattutto se la VM è in esecuzione, possono variare, ovviamente non è detto che vengano copiate una volta sola. Alla successiva iterazione, vengono copiate solo le pagine modificate (*dirty pages*), fino a quando il numero di queste è inferiore ad una certa soglia data;
4. **Sospensione della VM**: raggiunta la soglia (quando rimangono poche pagine), si applica la suspend sulla macchina d'origine (in seguito avverrà una resume sulla macchina di destinazione);
5. **Commit**: la copia della VM sul nodo di destinazione è completa, dunque si può procedere con una commit (ovvero ci si affranca completamente dal nodo di origine, dal quale la VM viene eliminata);
6. **Resume**: viene eseguita la resume sul nodo B, in cui si trova una macchina pronta a ripartire, completa sia come immagine sul file system, sia come stato presente nei registri.

Con questa modalità, si ha downtime solo durante la copia delle ultime dirty pages, ovvero quando si è raggiunta la soglia preimpostata.
```
NB: chiaramente la prima iterazione della precopia è quella che richiede più tempo,
quelle successive ne richiedono meno perché salvano solo le pagine modificate.
```
Sebbene la precopia sia la modalità oggi più diffusa, ne esistono anche altre, ad esempio *post-copy*, in cui la macchina viene sospesa e vengono copiate (non iterativamente) pagine e stato. Così facendo si ottiene un tempo totale di migrazione più basso, ma un downtime dei servizi forniti dalla VM molto più elevato.

### XEN (Approfondimento)
XEN è un progetto che nasce in ambito accademico a Cambridge. Nasce come hypervisor (VMM paravirtualizzato), richiede che le VM che girano sopra xen abbiano un kernel adattato all'interfaccia che xen offre ai propri utilizzatorii. Per quanto riguarda il porting di Linux ha coinvolto circa 3000 linee di codice del kernel, per adattarlo in modo che potesse dialogare con le API di XEN.
Dal punto di vista commerciale ha limitato la gamma di kernel installabili, per quanto riguarda i SO proprietari, nonostante un tentativo di porting dei Sistemi Operativi (ad esempio Windows, che non è stato portato a termine).

#### Architettura
XEN è costituito da un VMM *hypervisor*, che si appoggia direttamente sull'hardware (virtualizzazione di sistema - quindi è necessario avere spazio e in caso togliere il SO preesistente) e si occupa della virtualizzazione della CPU, della memoria e dei dispositivi di ogni VM. In XEN le macchine virtuali vengono chiamate *domain* e su ogni sistema XEN c'è una VM speciale chiamata *domain 0* che è privilegiata: a livello architetturale è come tutte le altre ma, tramite un'interfaccia di controllo fornita da XEN, può amministrare tutto il sistema. Questa interfaccia è accessibile solo dal domain 0, ed è separata dall'hypervisor stesso, scelta che permette di ottenere una separazione dei meccanismi dalle politiche: all'interno delle applicazioni che consento la configurazione ed il controllo del sistema abbiamo le politiche (espresse dall'utente), che vengono poi implementate e messe in pratica dall'hypervisor. Infatti, tipicamente nel domain 0 girano applicazioni che consentono all'amministratore di configurare il sistema virtualizzato e operando sulla console di questa VM è possibile creare una VM guest (di domain U - utente), eliminarla, migrarla, ecc.

#### Realizzazione
Un VMM assomiglia per certi versi al kernel di un SO: deve gestire in modo appropriato l'hardware e fornire un accesso particolare agli utilizzatori (che nel caso di un sistema virtualizzato non sono gli utenti ma le VM.
Ogni VM vede una *CPU* come se fosse a lei esclusivamente dedicata, quando in realtà non è così: le risorse vengono condivise grazie all'attività dell'hypervisor tra tutti gli utilizzatori secondo politiche particolari (ad esempio per quanto riguarda la CPU l'hypervisor dovrà mettere in atto politiche di scheduling particolari).
Stessa cosa vale per la *memoria*, anch'essa dev'essere in qualche modo messa a disposizione per gli utilizzatori dal VMM, che deve garantire i criteri di sicurezza opportuni.
Altro compito importantissimo del VMM è quello della gestione dei *dispositivi* (quindi I/O).

qualche cenno sulle caratteristiche di XEN: noi facciamo riferimento a XEN "paravirtualizzato": in questi sistemi necessario separare il kernel dalla macchina virtuale e dalle applicazioni, in quanto XEN adotta una configurazione dei ring 0/1/3 (VMM esegue a ring 0, Sistemi Operativi a ring 1, le applicazioni a ring 3, così non si ha ring compression).
Le app possono utilizzare le system call per comunicare col sistema operativo, i sistemi operativi possono comunicare col VMM tramite delle hyper calls (sono come system call fornite dal nucleo dell'hypervisor per permettere agli SO di eseguire istruzioni particolari).

#### Gestione della Memoria e Paginazione
I SO guest gestiscono la memoria virtuale mediante la paginazione tradizionale: le page table delle VM vengono mappate in memoria fisica da XEN, il quale è l'unico a potervi accedere in *scrittura*, su richiesta delle VM. L'accesso in *lettura*, invece, è permesso anche ai sistemi operativi ospitati.

##### Protezione: Memory Split
Com'è strutturato lo spazio di indirizzamento delle singole VM guest? Si adotta un principio di **memory split**.
```
NB: Consideriamo sempre il parallelo con sistema non virtualizzato/sistema virtualizzato:
in un sistema virtualizzato, ogni utilizzatore è una VM, quindi ogni entità che si inter-
faccia col VMM (equivalente del kernel) è una VM. Così come accade nei sistemi non virtua-
lizzati, in cui ogni processo ha un utilizzatore e un suo spazio di indirizzamento, anche
nei sistemi virtualizzati ogni VM ha un suo spazio di indirizzamento virtuale (perché sia-
mo in presenza di memoria virtuale).
```
Per motivi di efficienza, poiché chiaramente nella commutazione tra una VM e l'altra c'è problema di reperire il codice di XEN, lo spazio di indirizzamento di ogni VM è strutturato a "segmenti": nei primi 64MiB viene allocato XEN (ring 0), poi c'è una parte relativa al Kernel del SO guest (ring 1), poi c'è lo spazio utente, che verrà utilizzato dalle applicazioni (ring 3). I VM guest si occupano delle politiche di gestione della paginazione, mentre i meccanismi, ovvero l'effettiva implementazione della paginazione, sono compito del VMM, in quanto il kernel del SO guest, non può occuparsene, non essendo nel ring privilegiato 0. Ciò garantisce maggiore protezione in quanto si ha separazione tra politiche (a carico dei guest - alto livello) e meccanismi (a carico del VMM - basso livello).
Con questa soluzione, quando viene creato un nuovo processo nello spazio del guest, fra le altre cose dev'essere creata una Tabella delle Pagine (PT) associata a tale processo. Ovviamente, poiché come detto tale operazione non può essere fatta dal kernel del sistema operativo che ospita quel processo (in quanto si trova a ring 1), dev'essere fatta da qualcun'altro. Quindi ciò che succede è che il guest richiede una nuova PT all'hypervisor, il quale la crea e vi aggiunge anche lo spazio riservato a XEN; così facendo XEN registra la tabella e acquisisce il diritto di scrittura esclusivo (i guesto potranno solo leggerle), e ogni volta che il guest di tale TP dovrà aggiornarla, proverà a scriverci generando un trap *protection fault*, che verrà catturata e gestita da XEN, permettendogli di verifcare la correttezza della richiesta ed aggiornare effettivamente, in seguito, la Tabella delle Pagine.

##### Protezione: Balloon Process
Per com'è gestita la protezione in XEN, l'unica componente capace di allocare memoria è il VMM (ring 0), ma può farlo solo in seguito a richeiste delle VM guest, in quanto come detto, le politiche si trovano in alto livello (ring 3), mentre i meccanismi a basso livello (ring 0). Però, in alcuni casi (es: attivazione nuova VM, operazione per la quale serve acquisire memoria necessaria per allocare lo spazio di indirizzamento di quella macchina virtuale), può essere necessario al VMM dover ottenere nuove pagine. Questa possibilità, ovvero di richiedere pagine, il VMM non ce l'ha. Può farlo solo in seguito a richieste da parte dei guest. Per risolvere questo problema, su XEN è stato adottata una soluzione (peculiare per la paravirtualizzazione) chiamata **balloon process**: in ogni guest c'è un processo in costante esecuzione, che è in grado di dialogare direttamente con l'hypervisor. In caso di necessità di pagine, il VMM può chiedere a tali processi di "gonfiarsi", ovvero richiedere al proprio SO ulteriori pagine. Tale richiesta provoca l'allocazione di nuove pagine al balloon process che, una volta ottenute, le cede al VMM.



## 02 - Protezione




