<h1 align="center"> Sistemi Operativi M</h1>

<p align="center">
  Appunti del corso Sistemi Operativi M (<a href="https://www.unibo.it/it/didattica/insegnamenti/insegnamento/2021/468009">72947</a>), anno 2021-2022<!--. A cura di <a href="https://github.com/mikyll">Michele Righi</a> e <a href="https://github.com/TryKatChup">Karina Chichifoi</a>.-->
</p>

<details open="open">
  <summary><h2 style="display: inline-block">Indice</h2></summary>
  <ol>
    <li>
      <a href="#01---virtualizzazione">Virtualizzazione</a>
      <!--<ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>-->
    </li>
    <li>
      <a href="#02---protezione">Protezione</a>
    </li>
  </ol>
</details>

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

#### Tipi (Livelli) di Virtualizzazione
- Livello applicativo - applicazioni virtuali / supporto a tempo di esecuzione (es: JVM, .NET CLR)
- Livello di librerie (API a livello utente) - librerie virtuali (es: WINE, WABI, vCUDA)
- Livello di Sistema Operativo - container (Jail, Virtual Environment, Docker)
- Livello di Astrazione Hardware (HAL) - macchine virtuali (VMware, Virtual PC, Xen, User mode Linux). Questo è il tipo di virtualizzazione che astrae l'hardware ed è quello che ci interessa maggiormente in questo corso. Al contrario dei container, le VM non condividono lo stesso SO.
- Livello di Instruction Set Architecture(ISA) - ISA Virtuale (Bochs, QEMU)

#### Cenni Storici
La virtualizzazione non è un concetto nuovo, bensì nasce negli anni '60 coi sistemi CP/CMS di IBM, dove il Control Program (CP) esegue direttamente sull'hardware svolgendo il ruolo di VMM ed il Conversational Monitor System (CMS) è il sistema operativo, replicato per ogni VM. Con la diffusione del consolidamento dell'hardware, si è passati dal paradigma "one application, one server" (tipico degli anni '80/'90, dovuto al crollo dei costi dell'hardware), ad avere, dagli anni 2000, un unico server bello grosso e potente, su cui installare 20/30 VM, ciascuna delle quali svolge un certo servizio: una soluzione molto più razionale, in quanto avere un numero di macchine fisiche ristrette permette di semplificare la configurazione, la gestione e la manutenzione. Per poi arrivare negli anni 2010 al Cloud Computing.

#### Vantaggi della Virtualizzazione
La virtualizzazione comporta numerosi vantaggi:
- possibilità di avere più SO, anche differenti, sulla stessa architettura fisica;
- isolamento degli ambienti d'esecuzione, utile specialmente per eseguire e testare software dalla sicurazza e affidabilità non certa (nel caso peggiore la singola VM va in crash);
- abbattimento dei costi hardware, in quantosi possono concentrare più macchine (es. server) su un'unica architettura hardware, ed abbattimento dei costi di amministrazione;
- gestione facilitata delle macchine (creazione, installazione -esistono template già preimpostati-, amministrazione, migrazione "a caldo", possibilità di adottare politiche di bilanciamento del carico e robustezza -disaster recovery-)

#### Realizzazione del VMM
Il VMM generalmente deve fornire ad ogni VM le risorse che gli servono per funzionare (CPU, memoria, dispositivi I/O).
I requisiti fondamentali sono i tre seguenti:
- l'ambiente d'esecuzione fornito dev'essere identico a quello della macchina reale (come se non fosse un sistema virtualizzato, ma girasse direttamente sull'architettura hardware);
- dev'essere garantita un'elevata efficienza (che sia accettabile) nell'esecuzione dei programmi;
- dev'essere garantita la stabilità e la sicurezza dell'intero sistema.
Due concetti molto importanti (che fungono anche da parametri per classificarlo) nella realizzazione del VMM sono: il **livello**, ovvero dove è collocato il VMM (può essere un *VMM di sistema* o un *VMM ospitato*); la **modalità di dialogo**, ovvero il modo in cui il VMM accede alle risorse (*virtualizzazione pura* o *paravirtualizzazione*)

In un sistema di virtualizzazione esistono due tipi di "componenti": l'**host** è la piattaforma sulla quale si realizzano le VM, ovvero il livello sottostante che comprende la macchina fisica ed il VMM; il **guest** è la VM vera e propria che comprende il Sistema Operativo e le applicazioni.

##### VMM di Sistema
Si trova direttamente sopra l'hardware e consiste in un Sistema Operativo molto leggero che realizza le funzionalità di virtualizzazione (es: kvm, xen). A meno che non ci sia abbastanza spazio libero sul disco e vi sia la possibilità di impostare un multiboot, per installare un VMM di sistema è necessario eliminare il Sistema Operativo preesistente.

<img width="60%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/01%20-%20Virtualizzazione/VMM%20di%20Sistema.png" alt="VMM di Sistema"/>

##### VMM Ospitato
Viene installato come una normale applicazione sul Sistema Operativo preesistente, opera nello spazio utente ed accede all'hardware tramite le system call del SO (es. VirtualBox). È più semplice da installare e per la gestione delle periferiche può fare riferimento al Sistema Operativo sottostante, ma ha performance peggiori.

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/01%20-%20Virtualizzazione/VMM%20Ospitato.png" alt="VMM Ospitato"/>

NB: d'ora in poi faremo sempre riferimento a VMM di sistema.

##### Ring di Protezione
L'architettura delle CPU prevede almeno due livelli di protezione (0 "modi di esecuzione"): supervisore/kernel (livello 0) e utente (livello >0). Ogni ring corrisponde ad una diversa modalità di funzionamento del processore:
- a livello 0 è possibile eseguire istruzioni privilegiate della CPU;
- a livello superiore a 0 non possono essere eseguite.

Alcuni programmi, come il kernel del SO, sono progettati per eseguire nel ring 0 (in cui si ha il pieno controllo dell'hardware).
In un sistema virtualizzato, il VMM dev'essere l'unica componente in grado di mantenere il pieno controllo dell'hardware (di conseguenza si troverà a ring 0, mentre le VM a ring >0).
Ma nella VM c'è il Sistema Operativo, il cui kernel non può eseguire a livello 0 come dovrebbe. Di conseguenza sorgono due principali problemi: *ring deprivileging* e *ring compression*.

###### Ring Deprivileging
Il SO della VM si trova ad eseguire in un ring che non gli è proprio e le istruzioni privilegiate richieste dal sistema operativo nell'ambiente guest non possono essere eseguite.
Possibile soluzione **trap & emulate**: se il guest tenta di eseguire un'istruzione prvilegiata (ad esempio 'popf', ovvero la disabilitazione delle interruzioni), la CPU lancia un'eccezione, che viene rilevata dal VMM (trap), al quale trasferisce il controllo; dopodiché il VMM controlla la correttezza dell'operazione e ne emula (emulate) il comportamento.
NB: se la popf potesse essere eseguita direttamente sul processore (a ring 0), verrebbero disabilitati gli interrupt per tutto il sistema ed il VMM non potrebbe riguadagnare il controllo della CPU, mentre il comportamento desiderato sarebbe che gli interrupt venissero sospesi solo per la VM in questione.

###### Ring Compression
Se ad esempio l'architettura ha solo 2 ring, poiché il primo (0) è assegnato al kernel del Sistema Operativo, applicazioni e SO della macchina virtuale eseguono allo stesso livello, con conseguente mancanza di isolamento e protezione.

###### Ring Aliasing
Alcune istruzioni non privilegiate, eseguite a livello user, permettono di accedere in lettura ad alcuni registri la cui gestione dovrebbe essere riservata al VMM, con conseguenti possibili inconsistenze. Ad esempio, il registro CS contiene il livello di privilegio corrente (se la VM pensa di essere in un certo ring, ma leggendo lo stato del registro vede che è sbagliato, potrebbero esserci dei problemi).

###### Supporto nativo alla virtualizzazione
L'architettura della CPU si dice **naturalmente virtualizzabile** (o con supporto nativo alla virtualizzazione) se prevede l'invio di trap allo stato supervisore (0) per ogni istruzione privilegiata invocata da un ring >0. In questi casi è possibile realizzare un approccio "trap & emulate" e si ha supporto nativo all'esecuzione diretta.
Tuttavia, non tutte le architetture sono naturalmente virtualizzabili (es. Intel IA32) e alcune istruzioni privilegiate non provocano una trap, anzi, in alcuni casi causano il crash del sistema.

Nel caso in cui un processore non fornisca supporto nativo alla virtualizzazione, è necessario ricorrere a soluzioni software. Alcune possibili sono: *fast binary translation*, *paravirtualizzazione*.

###### Fast Binary Translation (FTB)
Sfrutta un'idea simile alla compilazione dinamica: il VMM scansiona il codice dei SO guest prima dell'esecuzione, per sostituire a runtime i blocchi contenenti istruzioni privilegiate con blocchi equivalenti dal punto di vista funzionale, ma che contengano invece chiamate al VMM. I blocchi tradotti vengono salvati in cache per eventuali riutilizzi futuri.
**Vantaggi**: ogni VM è una esatta replica della macchina fisica, dunque è possibile installare gli stessi SO di architetture senza virtualizzazione nativa.
**Svantaggi**: la traduzione dinamica è costosa.


###### Paravirtualizzazione

## 02 - Protezione




