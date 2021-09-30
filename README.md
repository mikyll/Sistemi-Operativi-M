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



