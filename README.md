<h1 align="center">Sistemi Operativi M</h1>

<p align="center">
  Appunti del corso Sistemi Operativi M (<a href="https://www.unibo.it/it/didattica/insegnamenti/insegnamento/2021/468009">72947</a>), anno 2021-2022. A cura di <a href="https://github.com/mikyll">Michele Righi</a>, <a href="https://github.com/TryKatChup">Karina Chichifoi</a> e <a href="https://github.com/lnwor">Lorenzo Guerra</a>.
  <br/>
	<br/>
	<a href="https://github.com/mikyll/Sistemi-Operativi-M/tree/main/capitoli">Capitoli</a>
	·
	<a href="https://github.com/mikyll/Sistemi-Operativi-M/tree/main/prove esame">Prove Esame</a>
	·
	<a href="https://github.com/mikyll/Sistemi-Operativi-M/tree/main/flashcard">Flashcards</a>
	·
	<a href="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/prove esame/2014-07-14%20-%20Airbolo/sol2_airbolo_madonna_che_meme_questo.go">Meme</a>
	·
	<a href="https://github.com/mikyll/Sistemi-Operativi-M/issues">Report a bug</a>
</p>

<details open="open">
  <summary><h2 style="display: inline-block">Indice</h2></summary>
  <ol>
<!-- CAPITOLO 1 -->
    <li><a href="#01---virtualizzazione-">Virtualizzazione</a>
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
        <li><a href="#gestione-di-vm">Gestione di VM</a>
          <ul>
            <li><a href="#stati-di-una-vm">Stati di una VM</a></li>
            <li><a href="#migrazione-di-una-vm">Migrazione di una VM</a></li>
            <li><a href="#soluzione-precopy">Soluzione: precopy</a></li>
          </ul>
				</li>
        <li><a href="#xen">XEN</a>
          <ul>
            <li><a href="#architettura">Architettura</a></li>
            <li><a href="#realizzazione">Realizzazione</a></li>
            <li><a href="#gestione-della-memoria-e-paginazione">Gestione della Memoria e Paginazione</a>
							<ul>
                <li><a href="#protezione-memory-split">Protezione: Memory Split</a></li>
                <li><a href="#protezione-balloon-process">Protezione: Balloon Process</a></li>
              </ul>
						</li>
            <li><a href="#cenni-su-virtualizzazione-della-cpu">Cenni su Virtualizzazione della CPU</a></li>
            <li><a href="#virtualizzazione-dei-dispositivi-io">Virtualizzazione dei Dispositivi (I/O)</a></li>
            <li><a href="#gestione-delle-interruzioni">Gestione delle Interruzioni</a></li>
            <li><a href="#migrazione-live">Migrazione Live</a></li>
					</ul>
				</li>
      </ul>
    </li>
<!-- CAPITOLO 2 -->
    <li><a href="#02---protezione-">Protezione</a>
      <ul>
        <li><a href="#protezione-modelli-politiche-e-meccanismi">Protezione: Modelli, Politiche e Meccanismi</a>
					<ul>
						<li><a href="#modelli">Modelli</a></li>
						<li><a href="#politiche">Politiche</a></li>
						<li><a href="#meccanismi">Meccanismi</a></li>
					</ul>
				</li>
				<li><a href="#dominio-di-protezione">Dominio di Protezione</a>
					<ul>
						<li><a href="#associazione-tra-processo-e-dominio">Associazione tra Processo e Dominio</a>
							<ul>
								<li><a href="#esempio-di-cambio-di-dominio">Esempio di cambio di dominio</a></li>
							</ul>
					</ul>
				</li>
				<li><a href="#matrice-degli-accessi">Matrice degli Accessi</a>
					<ul>
						<!--<li><a href="#verifica-del-rispetto-dei-vincoli-di-accesso">Verifica del Rispetto dei Vincoli di Accesso</a></li>-->
						<li><a href="#modifica-dello-stato-di-protezione">Modifica dello Stato di Protezione</a>
							<ul>
								<!--<li><a href="#modello-graham-denning">Modello Graham-Denning</a></li>-->
								<li><a href="#propagazione-dei-diritti-di-accesso-copy-flag">Propagazione dei Diritti di Accesso (Copy Flag)</a></li>
								<!--<li><a href="#diritto-owner">Diritto Owner</a></li>
								<li><a href="#diritto-control">Diritto Control</a></li>
								<li><a href="#diritto-switch">Diritto Switch</a></li>-->
							</ul>
						<li><a href="#realizzazione-della-matrice-degli-accessi">Realizzazione della Matrice degli Accessi</a>
							<ul>
								<li><a href="#acl-lista-degli-accessi">ACL: Lista degli Accessi</a></li>
								<li><a href="#cl-capability-list">CL: Capability List</a></li>
							</ul>
						</li>
						<!--<li><a href="#revoca-dei-diritti-di-accesso">Revoca dei Diritti di Accesso</a></li>
						<li><a href="#acl-vs-cl">ACL vs CL</a></li>-->
					</ul>
				</li>
				<li><a href="#protezione-multilivello">Protezione Multilivello</a>
					<ul>
						<li><a href="#modelli-di-sicurezza-multilivello">Modelli di Sicurezza Multilivello</a>
							<ul>
								<li><a href="#modello-bell-la-padula">Modello Bell-La Padula</a></li>
								<li><a href="#modello-biba">Modello Biba</a></li>
							</ul>
						</li>
					</ul>
				</li>
				<!--<li><a href="#architetture-dei-sistemi-ad-elevata-sicurezza">Architetture dei Sistemi ad Elevata Sicurezza</a></li>
				<li><a href="#classificazione-della-sicurezza-dei-sistemi-di-calcolo">Classificazione della Sicurezza dei Sistemi di Calcolo</a></li>-->
			</ul>
		</li>
<!-- CAPITOLO 3 -->
		<li><a href="#03---programmazione-concorrente-">Programmazione Concorrente</a>
			<ul>
				<!--<li><a href="#cenni-storici">Cenni Storici</a></li>-->
				<li><a href="#tipi-di-architettura">Tipi di architettura</a></li>
				<li><a href="#classificazione-delle-architetture">Classificazione delle Architetture</a>
					<!--<ul>
						<li><a href="#single-processor">Single Processor</a></li>
						<li><a href="#shared-memory-multiprocessors">Shared-Memory Multiprocessors</a></li>
						<li><a href="#distributed-memory">Distributed-Memory</a>
							<ul>
								<li><a href="#multicomputers">Multicomputers</a></li>
								<li><a href="#network-systems">Network Systems</a></li>
							</ul>
						</li>
					</ul>-->
				</li>
				<li><a href="#tipi-di-applicazioni">Tipi di Applicazioni</a></li>
				<li><a href="#processi-non-sequenziali-e-tipi-di-iterazione">Processi Non Sequenziali e Tipi di Iterazione</a>
					<ul>
						<li><a href="#processo-sequenziale">Processo Sequenziale</a></li>
						<li><a href="#processo-non-sequenziale">Processo Non Sequenziale</a>
							<!--<ul>
								<li><a href="#elaboratore-non-sequenziale">Elaboratore Non Sequenziale</a></li>
								<li><a href="#linguaggi-concorrenti">Linguaggi Concorrenti</a></li>
							</ul>-->
						</li>
						<li><a href="#scomposizione-di-un-processo-non-sequenziale">Scomposizione di un Processo Non Sequenziale</a>
					<ul>
						<li><a href="#interazione-tra-processi">Interazione tra Processi</a>
							<ul>
								<li><a href="#cooperazione">Cooperazione</a></li>
								<li><a href="#competizione">Competizione</a></li>
								<li><a href="#interferenza">Interferenza</a></li>
							</ul>
						</li>
					</ul>
				</li>
			</ul>
		</li>
		<li><a href="#architetture-e-linguaggi-per-la-programmazione-concorrente">Architetture e Linguaggi per la Programmazione Concorrente</a></li>
		<li><a href="#architettura-di-una-macchina-concorrente">Architettura di una Macchina Concorrente</a>
			<ul>
				<li><a href="#architettura-della-macchina-m">Architettura della Macchina M</a></li>
			</ul>
		</li>
		<li><a href="#costrutti-linguistici-per-la-specifica-della-concorrenza">Costrutti Linguistici per la Specifica della Concorrenza</a>
			<ul>
				<li><a href="#forkjoin">Fork/Join</a></li>
				<li><a href="#cobegincoend">Cobegin/Coend</a></li>
			</ul>
		</li>
		<li><a href="#proprietà-dei-programmi">Proprietà dei Programmi</a>
			<ul>
				<li><a href="#verifica-della-correttezza-di-un-programma">Verifica della Correttezza di un Programma</a></li>
				<li><a href="#proprietà-di-safety-e-liveness">Proprietà di Safety e Liveness</a>
					<ul>
						<li><a href="#proprietà-dei-programmi-sequenziali">Proprietà dei Programmi Sequenziali</a></li>
						<li><a href="#proprietà-dei-programmi-concorrenti">Proprietà dei Programmi Concorrenti</a>
							<ul>
								<li><a href="#verifica-di-proprietà-nei-programmi-concorrenti">Verifica di Proprietà nei Programmi Concorrenti</a></li>
							</ul>
						</li>
					</ul>
				</li>
			</ul>
		</li>
  </ol>
</details>

<details>
  <summary><h2 style="display: inline-block">Indice per lezioni</h2></summary>
  <ul>
    <li><a href="#01---virtualizzazione-">2021/09/21</a></li>
    <li><a href="#soluzione-precopy">2021/09/28</a></li>
    <li><a href="#diritto-owner">2021/09/29</a></li>
		<li><a href="#03---programmazione-concorrente-">2021/10/05</a></li>
		<li><a href="#interazione-tra-processi">2021/10/06</a></li>
		<li><a href="#proprietà-dei-programmi-concorrenti">2021/10/12</a></li>
		<li><a href="#">2021/10/13</a></li>
  </ul>
</details>

<!-- Lezione 2021/09/21-->
## 01 - Virtualizzazione [![Vai al Capitolo Singolo](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/icon-document.png)](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/capitoli/01%20-%20Virtualizzazione.md "Vai al Capitolo Singolo")
La virtualizzazione è una tecnologia oggi usatissima. Virtualizzare un sistema di elaborazione (costituito da un insieme di risorse hardware e software) significa presentare all'utilizzatore una visione delle risorse diversa da quella attuale (ad esempio duplicazione della memoria). Questo obbiettivo viene raggiunto mediante un livello intermedio, un layer che svolge appunto il ruolo di intermediario tra utilizzatore (vista logica) e sistema (vista fisica). Solitamente l'intermediario è software, ma talvolta può avere un supporto hardware specifico. Esso permette di eseguire più macchine virtuali su una stessa architettura e ognuna di queste vede le proprie risorse, indipendentemente dalle altre, e da quelle effettive ("reali").\
Poiché le Macchine Virtuali (VM) devono funzionare in modo indipendente senza causare problemi al sistema, la gestione delle risorse dev'essere realizzata in modo appropriato: questo compito è affidato al Virtual Machine Monitor (VMM, detto anche Hypervisor), cha ha compiti molto simili a quelli di un Sistema Operativo (SO), motivo per cui viene trattato in questo corso).

Esempi di virtualizzazione:
- **virtualizzazione a livello di processo** - i SO multitasking consentono l'esecuzione parallela di più processi, ognuno dei quali dispone di una VM (CPU, memoria, dispositivi) dedicata. Questo tipo è realizzato dal kernel;
- **virtualizzazione della memoria** - con memoria virtuale ogni processo vede uno spazio di indirizzamento di dimensioni indipendenti dalle reali dimensioni e dallo spazio effettivamente a disposizione. Anche questo è realizzato dal kernel del SO;
- **astrazione** - un oggetto astratto (risorsa virtuale) come rappresentazione semplificata di un oggetto (risorsa fisica). Solitamente, l'astrazione serve per fornire una vista delle sole proprietà significative, nascondendo i dettagli realizzativi non necessari, ad esempio i tipi di dato ad alto livello (numeri intero, numeri in virgola mobile, ecc.) rispetto alla loro rappresentazione binaria nella cella di memoria. Un altro esempio sono i linguaggi di programmazione, in cui un'istruzione di alto livello è un'astrazione di ciò che avviene a più basso livello, in linguaggio macchina. Il disaccoppiamento è realizzato dalle operazioni (interfaccia) con le quali è possibile utilizzare l'oggetto;
- **linguaggi di programmazione** - la capacità di portare lo stesso programma (scritto in un linguaggio di alto livello) su architetture diverse è possibile grazie alla definizione di una VM in grado di interpretare ed eseguire ogni istruzione del linguaggio, indipendentemente dall'architettura del sistema (SO e hardware). Possibili esempi sono gli interpreti (ad esempio Java Virtual Machine, cacpace di eseguire il bytecode Java), ed i compilatori.

### Virtualizzazione di un sistema di elaborazione
Una singola piattaforma hardware viene condivisa da più elaboratori virtuali, ognuno gestito da un proprio sistema operativo. Il disaccoppiamento viene realizzato dal VMM, che è il mediatore unico tra le VM e l'hardware. I suoi compiti sono *consentire la condivisione da parte di più macchine virtuali di una singola piattaforma hardware*, garantendo *isolamento* tra di esse e *stabilità* del sistema. Il VMM deve realizzare una specie di sandbox per ciascuna VM (se ad esempio una va in crash, le altre non devono risentirne).

### Emulazione
Consiste nell'esecuzione di programmi compilati per una particolare architettura (e quindi con un certo set di istruzioni) su un sistema di elaborazione dotato di un diverso insieme di istruzioni. Vengono emulate interamente le singole istruzioni dell'architettura ospitata, consentendo completa interoperabilità tra ambienti eterogenei. Tramite l'emulazione, sistemi operativi o applicazioni pensati per determinate architetture, possono eseguire senza essere modificati, su architetture completamente differenti.\
**Vantaggi**: interoperabilità tra ambienti eterogenei.\
**Svantaggi**: problemi di efficienza (basse performance).\
Nel tempo questo approccio si è ramificato seguendo due strade: interpretazione e ricompilazione dinamica.

#### Interpretazione
Consiste nella lettura di ogni singola istruzione del codice macchina che dev'essere eseguito e nell'esecuzione di più istruzioni sull'host virtualizzante per ottenere semanticamente lo stesso risultato.\
**Vantaggi**: è un metodo generale e potente che presenta una grande flessibilità nell'esecuzione perché consente di emulare e riorganizzare i meccanismi propri delle varie architetture.\
**Svantaggi**: produce un sovraccarico generalmente elevato poiché possono essere necessarie molte istruzioni dell'host per interpretare una singola istruzione sorgente.

#### Compilazione Dinamica
Invece una singola istruzione del sistema ospitato, vengono letti interi blocchi di istruzioni, che vengono tradotti (per la nuova architettura), ottimizzati e messi in esecuzione.\
**Vantaggi**: migliori prestazioni rispetto al metodo precedente, in quanto si leggono interi blocchi di codice, che vengono tradotti ed ottimizzati, consentendo di sfruttare tutte le possibilità offerte dalla nuova architettura; inoltre, le parti di codice usate frequentemente possono essere bufferizzate per evitare di doverle ricompilare in seguito.
Tutti i più noti emulatori utilizzano questa tecnica per implementare l'emulazione.

### Tipi (Livelli) di Virtualizzazione
- Livello applicativo - applicazioni virtuali / supporto a tempo di esecuzione (es: JVM, .NET CLR)
- Livello di librerie (API a livello utente) - librerie virtuali (es: WINE, WABI, vCUDA)
- Livello di Sistema Operativo - container (Jail, Virtual Environment, Docker)
- Livello di **Astrazione Hardware** (HAL) - **macchine virtuali** (VMware, Virtual PC, Xen, User mode Linux). Questo è il tipo di virtualizzazione che astrae l'hardware ed è quello che ci interessa maggiormente in questo corso. Al contrario dei container, le VM non condividono lo stesso SO.
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

```
NB: d'ora in poi faremo sempre riferimento a VMM di sistema.
```

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
```
NB: se la popf potesse essere eseguita direttamente sul processore (a ring 0), verrebbero
disabilitati gli interrupt per tutto il sistema ed il VMM non potrebbe riguadagnare il 
controllo della CPU, mentre il comportamento desiderato sarebbe che gli interrupt venisse-
ro sospesi solo per la VM in questione.
```

##### Ring Compression
Se ad esempio l'architettura ha solo 2 ring, poiché il primo (0) è assegnato al kernel del Sistema Operativo, applicazioni e SO della macchina virtuale eseguono allo stesso livello, con conseguente mancanza di isolamento e protezione.

##### Ring Aliasing
Alcune istruzioni non privilegiate, eseguite a livello user, permettono di accedere in lettura ad alcuni registri la cui gestione dovrebbe essere riservata al VMM, con conseguenti possibili inconsistenze. Ad esempio, il registro CS contiene il livello di privilegio corrente (se la VM pensa di essere in un certo ring, ma leggendo lo stato del registro vede che è sbagliato, potrebbero esserci dei problemi).

#### Supporto Hardware alla Virtualizzazione
L'architettura della CPU si dice **naturalmente virtualizzabile** (o con supporto nativo alla virtualizzazione) se prevede l'invio di trap allo stato supervisore (0) per ogni istruzione privilegiata invocata da un ring >0. In questi casi è possibile realizzare un approccio "trap & emulate" e si ha supporto nativo all'esecuzione diretta.
Tuttavia, non tutte le architetture sono naturalmente virtualizzabili (es. Intel IA32) e alcune istruzioni privilegiate non provocano una trap, anzi, in alcuni casi causano il crash del sistema.

#### Realizzazione del VMM in Architetture Non Virtualizzabili
Nel caso in cui un processore non fornisca supporto nativo alla virtualizzazione, per la realizzazione del VMM è necessario ricorrere a soluzioni software. Alcune possibili sono: *fast binary translation*, *paravirtualizzazione*.

##### Fast Binary Translation (FTB)
Sfrutta un'idea simile alla compilazione dinamica: il VMM scansiona il codice dei SO guest prima dell'esecuzione, per sostituire a runtime i blocchi contenenti istruzioni privilegiate con blocchi equivalenti dal punto di vista funzionale, ma che contengano invece chiamate al VMM. I blocchi tradotti vengono salvati in cache per eventuali riutilizzi futuri.\
**Vantaggi**: ogni VM è una esatta replica della macchina fisica, dunque è possibile installare gli stessi SO di architetture senza virtualizzazione nativa.\
**Svantaggi**: la traduzione dinamica è costosa.

<img width="40%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/01%20-%20Virtualizzazione/Fast%20Binary%20Translation.png" alt="Fast Binary Translation"/>

##### Paravirtualizzazione
È l'approccio più diffuso al giorno d'oggi, oltre a FBT. Il VMM (hypervisor) offre ai SO guest un'interfaccia virtuale (hypercall API) alla quale i SO guest devono fare riferimento per avere accesso alle risorse. Le istruzioni non privilegiate vengono eseguite direttamente dai SO guest, mentre per le istruzioni privilegiate, eseguono delle hypercall. ```NB: così come il *Sistema* Operativo fornisce delle *system* call, l'*Hyper*visor fornisce delle *hyper* call.``` Ciò consente  di eseguire istruzioni privilegiate chiamando direttamente la relativa hyper call, senza dover generare interrupt al VMM. Xen utilizza questa tecnica.\
**Vantaggi**: la struttura del VMM è semplificata e si ottengono prestazioni migliori, in quanto non si ha il ritardo dovuto alla compilazione di FTB.\
**Svantaggi**: vi è la necessità di porting dei SO guest (i kernel devono essere resi compatibili, soluzione che è preclusa a molti sistemi operativi proprietari, fra cui Windows).

<img width="34%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/01%20-%20Virtualizzazione/Paravirtualizzazione.png" alt="Paravirtualizzazione"/>

### Architetture virtualizzabili
Con **virtualizzazione pura**, si intende un'architettura che non costringe l'amministratore (o l'utente) a installare nella macchina virtuale un kernel modificato (che non sia l'originale del Sistema Operativo), dunque è il caso di architetture con supporto nativo alla virtualizzazione, ma anche FTB, in quanto anche lì non c'è bisogno di modificare il kernel.\
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
- **0/1/3**: VMM (0), SO (1), applicazioni (3). Il SO viene spostato dal ring 0 (dove nativamente dovrebbe trovarsi) al ring 1, lasciando le applicazioni al ring 3, mentre al ring 0 viene installato il VMM;
- **0/3/3**: VMM (0), SO (3), applicazioni (3). Il SO viene spostato direttamente al ring 3, dove si trovano anche le applicazioni, mentre sul ring 0 viene installato il VMM. In questa modalità non è possibile generare eccezioni, quindi devono essere intrapresi meccanismi molto sofisticati con un controllo continuo da parte del VMM.

### Gestione di VM
Il compito fondamentale del VMM è quello di gestire le VM (creazione, accensione/spegnimento, eliminazione, migrazione live).

#### Stati di una VM
Una macchina virtuale può trovarsi nei seguenti stati:
- **running** (o attiva): la macchina ha superato la fase di bootstrap ed è stata caricata nella *RAM* del server su cui è allocata;
- **inactive** (powered off): la macchina è spenta ed è rappresentata nel *file system* tramite un file immagine;
- **paused**: la macchina è in *attesa* di un evento (es: I/O richiesto da un processo nell'ambiente guest);
- **suspended**: lo stato correnteviene salvato nel file system dal VMM. L'uscita da tale stato avviene tramite un'operazione di *resume*.

suspend: il VMM salva lo stato della VM in memoria secondaria, mettendola in stand by;
resume: il VMM ripristina lo stato della VM in memoria centrale (lo stato è quello in cui si trovava quando è stata sospesa). Questa operazione può avvenire su un nodo diverso da quello della suspend.

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/01%20-%20Virtualizzazione/Stati%20di%20una%20VM.png" alt="Stati di una VM"/>

#### Migrazione di una VM
È una funzionalità necessaria soprattutto nei datacenter, per una gestione agile delle VM, a fronte di:
- variazioni dinamiche del carico (**load balancing**, consolidamento);
- **manutenzione "online"** dei server, senza dover interrompere i servizi forniti;
- gestione finalizzata al **risparmio energetico**;
- **tolleranza ai guasti** e disaster recovery.

Strumento fondamentale per queste procedure è la migrazione, ovvero la possibilità di muovere VM tra server.
**Migrazione live**: possibilità di spostare una VM da un server fisico ad un altro, senza doverla spegnere. È desiderabile minimizzare il downtime, il tempo di migrazione ed il consumo di banda.

<!-- Lezione 2021/09/28 -->
<!--
RIPASSINO
Avevamo introdotto il tema della migrazione live, meccanismo che consente lo spostamento della VM da un nodo fisico ad un altro. Questo spostamento, se fatto in modo "live", significa che può essere eseguito senza neanche spegnere la macchina in questione.
In questo modo si ha un downtime davvero trascurabile (poco più di 100ms)

La capacità/caratteristica di indipendenza e isolamento dall'ambiente fisico è il vantaggio principale che ci consente di realizzare in modo molto semplice dal punto di vista tecnico, la migrazione live, ovvero suspend su un nodo e resume su un altro nodo.

Fra gli obbiettivi, sicuramente uno prioritario è minimizzare il downtime, in quanto se la macchina fornisce un servizio, vogliamo che questo rimanga disponibile per il maggior tempo possibile e dunque, appunto, che il downtime venga minimizzato.
Altri aspetti da tenere in considerazione:
- ridurre al minimo il tempo di migrazione, come tempo complessivo che richiede la migrazione;
- occupare meno banda possibile.

Naturalmente ci sono dei vantaggi nella realizzazione della migrazione se le architetture alle quali stiamo facendo riferimento hanno file system in comune (es. cluster), ovvero condividono gli stessi dischi.
-->
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

### XEN
XEN è un progetto che nasce in ambito accademico a Cambridge. Nasce come hypervisor (VMM) paravirtualizzato, richiede che le VM che girano sopra xen abbiano un kernel adattato all'interfaccia che xen offre ai propri utilizzatori. Per quanto riguarda il porting di Linux ha coinvolto circa 3000 linee di codice del kernel, per adattarlo in modo che potesse dialogare con le API di XEN.
Dal punto di vista commerciale ha limitato la gamma di kernel installabili, per quanto riguarda i SO proprietari, nonostante un tentativo di porting dei Sistemi Operativi (ad esempio Windows XP, che non è stato portato a termine).

#### Architettura
XEN è costituito da un VMM *hypervisor*, che si appoggia direttamente sull'hardware (virtualizzazione di sistema - quindi è necessario avere spazio e in caso togliere il SO preesistente) e si occupa della virtualizzazione della CPU, della memoria e dei dispositivi per ogni VM. In XEN le macchine virtuali vengono chiamate *domain* e su ogni sistema XEN c'è una VM speciale chiamata *domain 0* che è privilegiata: a livello architetturale è come tutte le altre ma, tramite un'interfaccia di controllo fornita da XEN, può amministrare tutto il sistema. Questa interfaccia è accessibile solo dal domain 0, ed è separata dall'hypervisor stesso, scelta che permette di ottenere una separazione dei meccanismi dalle politiche: all'interno delle applicazioni che consento la configurazione ed il controllo del sistema abbiamo le politiche (espresse dall'utente), che vengono poi implementate e messe in pratica dall'hypervisor. Infatti, tipicamente nel domain 0 girano applicazioni che consentono all'amministratore di configurare il sistema virtualizzato e operando sulla console di questa VM è possibile creare una VM guest (di domain U - utente), eliminarla, migrarla, ecc.

#### Realizzazione
Un VMM assomiglia per certi versi al kernel di un SO: deve gestire in modo appropriato l'hardware e fornire un accesso particolare agli utilizzatori (che nel caso di un sistema virtualizzato non sono gli utenti ma le VM.
Ogni VM vede una *CPU* come se fosse a lei esclusivamente dedicata, quando in realtà non è così: le risorse vengono condivise grazie all'attività dell'hypervisor tra tutti gli utilizzatori secondo politiche particolari (ad esempio per quanto riguarda la CPU l'hypervisor dovrà mettere in atto politiche di scheduling particolari).
Stessa cosa vale per la *memoria*, anch'essa dev'essere in qualche modo messa a disposizione per gli utilizzatori dal VMM, che deve garantire i criteri di sicurezza opportuni.
Altro compito importantissimo del VMM è quello della gestione dei *dispositivi* (quindi I/O).

Qualche cenno sulle caratteristiche di XEN: noi facciamo riferimento a XEN "paravirtualizzato": in questi sistemi necessario separare il kernel dalla macchina virtuale e dalle applicazioni, in quanto XEN adotta una configurazione dei ring 0/1/3 (VMM esegue a ring 0, Sistemi Operativi a ring 1, le applicazioni a ring 3, così non si ha ring compression).
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
Per motivi di efficienza, poiché chiaramente nella commutazione tra una VM e l'altra c'è problema di reperire il codice di XEN, lo spazio di indirizzamento di ogni VM è strutturato a "segmenti": nei primi 64MiB viene allocato XEN (ring 0), poi c'è una parte relativa al Kernel del SO guest (ring 1), poi c'è lo spazio utente, che verrà utilizzato dalle applicazioni (ring 3).\
I VM guest si occupano delle politiche di gestione della paginazione, mentre i meccanismi, ovvero l'effettiva implementazione della paginazione, sono compito del VMM, in quanto il kernel del SO guest, non può occuparsene, non essendo nel ring privilegiato 0. Ciò garantisce maggiore protezione in quanto si ha separazione tra politiche (a carico dei guest - alto livello) e meccanismi (a carico del VMM - basso livello).
Con questa soluzione, quando viene creato un nuovo processo nello spazio del guest, fra le altre cose dev'essere creata una Tabella delle Pagine (PT) associata a tale processo. Ovviamente, poiché come detto tale operazione non può essere fatta dal kernel del sistema operativo che ospita quel processo (in quanto si trova a ring 1), dev'essere fatta da qualcun'altro. Quindi ciò che succede è che il guest richiede una nuova PT all'hypervisor, il quale la crea e vi aggiunge anche lo spazio riservato a XEN; così facendo XEN registra la tabella e acquisisce il diritto di scrittura esclusivo (i guest potranno solo leggerle), e ogni volta che il guest di tale TP dovrà aggiornarla, proverà a scriverci generando un trap *protection fault*, che verrà catturata e gestita da XEN, permettendogli di verificare la correttezza della richiesta ed aggiornare effettivamente, in seguito, la Tabella delle Pagine.

##### Protezione: Balloon Process
Per com'è gestita la protezione in XEN, l'unica componente capace di allocare memoria è il VMM (ring 0), ma può farlo solo in seguito a richeiste delle VM guest, in quanto come detto, le politiche si trovano in alto livello (ring 3), mentre i meccanismi a basso livello (ring 0). Però, in alcuni casi (es: attivazione nuova VM, operazione per la quale serve acquisire memoria necessaria per allocare lo spazio di indirizzamento di quella macchina virtuale), può essere necessario al VMM dover ottenere nuove pagine. Questa possibilità, ovvero di richiedere pagine, il VMM non ce l'ha. Può farlo solo in seguito a richieste da parte dei guest. Per risolvere questo problema, su XEN è stato adottata una soluzione (peculiare per la paravirtualizzazione) chiamata **balloon process**: in ogni guest c'è un processo in costante esecuzione, che è in grado di dialogare direttamente con l'hypervisor. In caso di necessità di pagine, il VMM può chiedere a tali processi di "gonfiarsi", ovvero richiedere al proprio SO ulteriori pagine. Tale richiesta provoca l'allocazione di nuove pagine al balloon process che, una volta ottenute, le cede al VMM.

#### Cenni su Virtualizzazione della CPU
Il VMM definisce un'architettura virtuale simile a quella del processore, nella quale però, le istruzioni privilegiate sono sostituite da opportune hypercalls:

Il VMM si occupa dello scheduling delle VM, seguendo un algoritmo molto generale (in grado di soddisfare dei vincoli temporali molto stringenti) chiamato *Borrowed Virtual Time*, che si basa sulla nozione di virtual-time: è un tempo che va avanti solo fintanto che la VM è attiva, ovvero se si trova in uno stato di sospensione il tempo si ferma e riprende quando viene attivato. Xen adotta due clock, uno relativo al real-time, l'altro al virtual-time.

#### Virtualizzazione dei dispositivi (I/O)
Le VM devono poter accedere ai dispositivi che sono disponibili a livello hardware. La scelta di XEN è quella, ovviamente, di virtualizzare l'interfaccia di ogni dispositivo, ma farlo tramite due tipi di driver: *back-end driver* e *front-end driver*.\
**Back-end driver** è il driver vero e proprio, che permette, tramite un'interfaccia del VMM chiamata *Safe Hardware Interface*, di comunicare ed utilizzare il dispostivo collegato a livello hardware. Tipicamente viene installato all'interno di una VM particolare che è sempre ancorata al nodo fisico (dominio 0 - solitamente qui vengono installati tutti i driver di ogni dispositivo presente connesso a livello fisico in quel nodo).\
**Front-end driver** è un driver "astratto", generico, non riferito adun dispositivo particolare, che viene installato tipicamente nel kernel del SO di una VM guest. Questo driver, all'occorrenza si collega al back-end driver specifico.
```
NB: non c'è niente che vieti di installare un back-end direttamente su una VM di domain U,
ma può convenire concentrarli tutti nel domain 0, sia perché siamo certi che quella macchi-
na non si sposterà mai da lì, essendo ancorata all'hardware, sia per motivi di portabilità.
```

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/01%20-%20Virtualizzazione/Driver.png" alt="Front-end and Back-end Drivers"/>

Ovviamente, per consentire la comunicazione tra back-end driver e front-end driver, serve un meccanismo che gestica le richieste. Questo viene realizzato tramite delle strutture chiamate asynchronous I/O rings (buffer FIFO circolari) in cui ogni elemento è una specie di descrittore che rappresenta una particolare richiesta. Le richieste di accesso ad un particolare device vengono fatte dal guest tramite il front-end che deposita la richiesta nel ring relativo, mentre dall'altra parte c'è il back-end che le preleva e le gestisce.

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/01%20-%20Virtualizzazione/Asynchronous%20IO%20Ring.png" alt="Structure of Asynchronous I/O Rings"/>

**Vantaggi**: il driver viene scorporato in due parti, svingolando la VM dal particolare server fisico in cui risiede (il front-end driver della VM rimane lo stesso anche se questa viene spostata su un altro nodo), garantendo *portabilità*; inoltre, mantenendo i driver fuori dall'hypervisor, si ha che esso è più semplificato e leggero.\
**Svantaggi**: il meccanismo di comunicazione fra i due tipi di driver appesantisce l'accesso ai dispositivi.

#### Gestione delle Interruzioni
La gestione delle interruzioni viene virtualizzata in modo molto semplice: ogni interruzione viene gestita direttamente dal SO guest, eccezione fatta per la *page fault*, che richiede accesso al registro CR2, il quale contiene l'indirizzo che ha provocato il page fault. Poiché tale registro è accessibile solo a ring 0, la gestione del page fault deve coinvolgere il VMM: la routine di gesstione eseguita da XEN legge CR2, lo copia in una variabile nello spazio del SO ospitato, al quale viene restituito il controllo per poter gestire il page fault.

#### Migrazione Live
Il comando di migrazione viene eseguito da un demone di migrazione che si trova nel domain 0 del server di origine della macchina da migrare. La soluzione è basata sulla precopy e le pagine da migrare vengono compresse per ridurre l'occupazione di banda.


## 02 - Protezione [![Vai al Capitolo Singolo](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/icon-document.png)](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/capitoli/02%20-%20Protezione.md "Vai al Capitolo Singolo")
**Sicurezza**: riguarda l'insieme delle *tecniche per regolamentare l'accesso* degli utenti al sistema di elaborazione. La sicurezza impedisce accessi non autorizzati al sistema e i conseguenti tentativi dolosi di alterazione e distruzione dei dati. La sicurezza riguarda l'interfaccia del sistema verso il mondo esterno. Le tecnologie di sicurezza di un sistema informatico realizzano meccanismi per l'identificazione, l'autenticazione e l'autorizzazione di utenti "fidati".\
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
- **DAC** (Discretional Access Control) - il creatore di un oggetto controlla i diritti di accesso per quell'oggetto (tipologia adottata da UNIX, che fornisce un meccanismo per definire e interpretare per ciascun file i 3 bit di read, write ed execute, per il proprietario, il gruppo e gli altri). La definizione delle politiche è decentralizzata.
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
L'associazione tra processo e dominio può essere statica o dinamica.\
**Statica**: l'insieme delle risorse disponibili ad un processo rimane fisso durante il suo tempo di vita. Osservazioni: questo tipo di associazione non è adatta al Principio del Privilegio Minimo, in quanto l'insieme globale delle risorse che un processo potrà usare può non essere un'informazione disponibile prima della sua esecuzione; inoltre, l'insieme minimo di risorse necessarie ad un processo per garantire tale Principio, può cambiare in modo dinamico durante l'esecuzione.\
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

<!-- lezione 2021/09/29 -->
##### Diritto Owner
Il diritto *owner* realizza il concetto di "proprietario di una risorsa" (oggetto). Il soggetto che possiede tale diritto di accesso, nei sistemi che lo prevedono, ha la possibilità di concedere/revocare un qualunque diritto di accesso sull'oggetto che gli appartiene (ovvero possiede il diritto owner su tale oggetto) ad un qualunque altro soggetto.\
In una matrice degli accessi, ciò si traduce nella presenza, in ciascuna colonna, di una ed una sola cella nella quale è presente un diritto owner. Per ogni risorsa (dunque per ogni colonna) ci dev'essere un solo soggetto che ne è il proprietario. Ciò significa che tale soggetto ha un ruolo privilegiato nei confronti di quella risorsa ed è l'unico soggetto capace di revocare o concedere diritti di accesso su quella risorsa ad altri soggetti.
Ad esempio, se *S2* ha il diritto 'owner' su *O2* allora può revocare il diritto 'execute' su *O* al soggetto *S1*.

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/02%20-%20Protezione/Matrice%20degli%20Accessi%20Owner.png" alt="Matrice degli Accessi e Diritto Owner"/>

##### Diritto Control
Il diritto *control* permette di revocare un qualunque diritto di accesso, riferendosi non ad un oggetto ma ad un altro soggetto.
Se la cella della colonna di un soggetto ha il diritto control, autorizza l'owner a modificare la riga associata a tale soggetto.\
Ad esempio, se *S1* ha il diritto di 'control' su *S2* e *S2* ha il diritto di 'write' su *O3*, allora *S1* può revocare il diritto di write di *S2* su *O3*.

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/02%20-%20Protezione/Matrice%20degli%20Accessi%20Control.png" alt="Matrice degli Accessi e Diritto Control"/>

```
NB: In un sistema MAC, avremmo l'entità centrale che è l'unica autorizzata a stabilire
che cosa si può o non si può fare sulle risorse del sistema. Questa autorità è un sog-
getto che ha il diritto di control su tutti gli altri soggetti (ha il potere assoluto,
nel senso che può modificare il contenuto della riga associata ad ogni altro soggetto).
```

Copy flag, owner e control sono strumenti con cui possiamo *modificare* il contenuto della matrice degli accessi (possiamo aggiungere o togliere diritti nelle varie caselle). Se ci pensiamo bene, una riga della matrice altro non è che il dominio di protezione associato ad un soggetto.

##### Diritto Switch
Uno dei modi possibili per garantire il rispetto del Principio del Privilegio Minimo è prevedere la possibilità di un cambio di dominio a runtime. Un processo che esegue nel dominio di un soggetto, se ha bisogno di diritti differenti, può spostarsi nel dominio specifico che gli garantisce questi diritti nuovi di cui ha bisogno.\
A tal proposito, esiste un diritto speciale chiamato *switch*: esercitando questo diritto, il processo che esegue nel dominio di un certo soggetto, può passare ad un nuovo dominio.
Ad esempio, se il soggetto *S2* ha bisogno del diritto 'write' sull'oggetto O3, ma possiede solo 'read', se possiede il diritto di 'switch' *S1* può passare nel dominio di *S1* per acquistare il diritto di 'write' e accedere a tale risorsa in scrittura.

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/02%20-%20Protezione/Matrice%20degli%20Accessi%20Switch.png" alt="Matrice degli Accessi e Diritto Switch"/>

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
Quando un qualunque soggetto *S* tenta un'operazione *M* su un oggetto *O*, il sistema di protezione va a verificare nella ACL associata ad *O* se è presente un elemento riferito al soggetto che sta tentando l'accesso e, se esiste, controlla che contenga il diritto per eseguire *M* (ovvero sia presenta la coppia <*S*, *Rk*> con *M* appartenente a *Rk*).\
In certi casi, per velocizzare l'accesso, viene prevista una lista di default: se è prevista, esistono dei diritti comunia a tutti i soggetti, dunque si va a vedere prima nella lsita di default e, se la ricerca non va a buon fine, si va a vedere nello specifico, elemento per elemento. Chiaramente, se la ricerca non ha successo, l'accesso viene negato.

**Utenti e Gruppi**: molti sistemi, per identificare un soggetto, prevedono non solo il nome utente (UID - User IDentifier), ma anche il gruppo (GID - Group IDentifier) a cui appartiene. Un gruppo aggrega un insieme di utenti, ha un nome, e può essere incluso nelle ACL. È importante sapere che un utente può appartenere anche a più gruppi, ma in un certo istante può appartenere ad un solo gruppo alla volta. Nel caso siano presenti i gruppi, una entry della ACL ha la forma UID, GID: \<insieme di diritti\>.
Quando un soggetto <utente, gruppo> prova ad accedere ad una risorsa in un certo modo, il tentativo di accesso comporta una ricerca nell'ACL dell'oggetto: se compare il modo con il quale l'utente sta cercando di accedere alla risorsa, allora l'operazione viene consentita.

Tipicamente i gruppi esistono per una questione di differenziazione dei ruoli, ciascuno dei quali ha diritti diversi. In generale, nei sistemi che prevedono i gruppi, è comunque possibile svincolare i soggetti dai gruppi, ovvero senza assegnargli alcun gruppo (UID, * \<insieme diritti\>).

##### CL: Capability List
Ne viene associata una a ciascun soggetto ed ha una struttura composta da un insieme di elementi, ognuno dei quali contiene l'indicazione dell'oggetto ed i diritti di accesso che quel soggetto, al quale la CL è associata, può esercitare su quell'oggetto. Chiaramente non avrà tanti elementi quanti sono le colonne della matrice degli accessi, in quanto come detto è una matrice sparsa.
Nella pratica, spesso l'oggetto viene identificato tramite un descrittore ed ovviamente i diritti, che spesso vengono rappresentati in modo compatto tramite una sequenza di bit.

<img width="60%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/02%20-%20Protezione/Struttura%20Capability%20List.png" alt="Struttura Tipica di una Capability List"/>

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
Naturalmente si può fare anche il discorso duale: se si ha la necessità di fare una modifica allo stato di protezione che interessa un particolare soggetto. Il caso più banale è eliminare il soggetto dal sistema. Ciò questo comporta una modifica allo stato di protezione, in CL si cancella semplicemente la lista associata al soggetto; in ACL bisogna fare ricerca in ogni ACL.\
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

<img width="60%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/02%20-%20Protezione/Flusso%20Modello%20Bell-La%20Padula.png" alt="Flusso delle Informazioni nel Modello Bell-La Padula"/>

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


<!-- lezione 2021/10/05 -->
## 03 - Programmazione Concorrente [![Vai al Capitolo Singolo](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/icon-document.png)](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/capitoli/03%20-%20Programmazione%20Concorrente.md "Vai al Capitolo Singolo")
La *programmazione concorrente* è l'insieme delle tecniche, metodologie e strumenti per il support all'esecuzione di sistemi software composti da *insiemi di attività svolte simultaneamente*.

### Cenni Storici
La programmazione concorrente nasce negli anni '60, proprio nell'ambito dei Sistemi Operativi, quando ci fu l'introduzione dei canali o controllori di dispositivi (hardware): questi consentono l'esecuzione concorrente di operazioni nei dispositivi ed istruzioni nei programmi eseguiti dall'unità di elaborazione centrale.

L'interazione tra dispositivi ed unità centrale di elaborazione (processore) è basata fortemente sul meccanismo delle interruzioni (segnali di interrupt).
Quando la CPU riceve un segnale di interrupt dalla periferica, può tempestivamente gestiree quel particolare evento, che potrebbe essere ad esempio il trasferimento di dati.\
Questo meccanismo di interruzioni è stato poi importato ed utilizzato ampiamente in sistemi multiprogrammati time-sharing, in cui è impiegato il concetto di **quanto  di tempo** che consente di dividere equamente il tempo di CPU tra tutte le applicazioni in esecuzione su quel sistema/ambiente di esecuzione. Il modo per sancire il termine di un quanto di tempo assegnato ad un certo processo, che esegue un'applicazione, è ancora rappresentato dall'interruzione. Si ha lo scatto all'interruzione quando il quanto di tempo è esaurito, e dunque tempestivamente il Sistema Operativo si occupa di gestire il *cambio di contesto* tra un'applicazione e la successiva, secondo le politiche di scheduling che possiede.\
Le interruzioni possono accadere ad istanti impredicibili, dunque in un sistema time-sharing parti di programmi possono essere eseguite in modo non predicibile. Infatti, una delle principali caratteristiche delle applicazioni concorrenti è il *non determinismo*: lo stesso programma eseguito in tempi diversi può comportare risultati diversi anche se il codice non cambia. Questo, ad esempio, si può rilevare quando cci sono parti di programmi che condividono le stesse variabili comuni: in questi casi, se non viene sincronizzato l'accesso a tali variabili, si possono creare delle interferenze.

Successivamente sono stati introdotti i sistemi multiprocessore, ovvero con più unità di elaborazione (parallelismo supportato a livello hardware). Se prima il parallelismo era puramente virtuale, con tali architetture il parallelismo era diventato effettivamente "reale", in quanto si potevano avere fisicamente diversi microprocessori che lavoravano in modo concorrente.
Ciò ha comportato diversi vantaggi, soprattutto in termini di prestazioni: in particolare, vengono abbattuti i tempi di esecuzione.

In un sistema concorrente i principali problemi sono:
- con quale criterio modellare l'applicazione concorrente;
- come suddividerla in attività concorrenti (quanti processi utilizzare);
- come garantire la corretta sincronizzazione delle loro operazioni (in generale le attività nelle quali si scompone l'applicaczione possono aver bisogno di interagire fra di loro, dunque è necessario imporre dei vincoli di precedenza).
Queste decisioni dipendono da:
- tipo di architettura hardware;
- tipo di applicazione.

### Tipi di Architettura

#### Single Processor
Si ha un solo processore che possiede delle memorie ad accesso rapido (tipicamente 2 cache) ed una memoria primaria. Non sono necessari ulteriori layer di comunicazione con altre unità di calcolo, in quanto ne è presente solo una.

<img width="20%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Single%20Processor.png" alt="Single Processor"/>

#### Shared-Memory Multiprocessors
Si tratta di un'architettura costituita da diversi nodi, ciascuno dei quali ha una propria unità di calcolo (microprocessore) e delle memorie ad accesso rapido (cache). Ogni nodo ha la possibilità di accedere a qualunque parte della memoria, grazie alla **rete di interconnessione**. È il più comune al giorno d'oggi.

<img width="45%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Shared-Memory%20Multiprocessors.png" alt="Shared-Memory Multiprocessors"/>

Possiamo distinguere due modelli di sistemi multiprocessore:
**UMA (Uniform Memory Access)**: sistemi a multiprocessore con un numero ridotto di processori (da 2 a circa 30). Sono caratterizzati da un'interconnessione realizzata tipicamente da memory bus o crossbar switch; *tempo di accesso alla memoria uniforme* (indipendentemente dal processore e dalla cella di memoria da accedere, il tempo di accesso rimane costante); sono chiamati anche SMP (Symmetric MultiProcessors).\
**NUMA (Non Uniform Memory Access)**: sistemi con un numero elevato di processori (decine o centinaia). Sono caratterizzati da: memoria organizzata gerarchicamente, per evitare la congestione del bus; rete di interconnessione strutturata anch'essa in modo gerarchico (insieme di switch e memorie strutturato ad albero) ed ogniprocessore ha memorie più vicine ed altre più lontane; tempo di accesso dipendente dalla distanza tra processore e memoria (NUMA).

#### Distributed-Memory
Nelle architetture con memoria distribuita ogni processore accede alla propria memoria che non è condivisa tra i nodi di elaborazione. La memoria è quindi specifica del processore a cui è associata ed un'unità di elaborazione non può fare riferimento alla memoria di un altro nodo. In questo tipo di architettura i nodi possono essere singoli processori o multiprocessori a memoria condivisa.\
Rientrano in questa categoria i *Multicomputers* ed i *Network Systems*.

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Distributed-Memory.png" alt="Distributed-Memory"/>

##### Multicomputers
Modello in cui i nodi e la rete sono *fisicamente vicini*, ovvero nella stessa struttura fisica. La rete di interconnessione offre un cammino di comunicazione tra i processi ad alta velocità e larghezza di banda. Ad esempio i Cluster ed i sistemi ad alto parallelismo (HPC). I multicomputer sono fatti per essere aggregati in una stessa struttura fisica.
```
NB: un Cluster of Computers (CoW), un insieme di nodi, tipicamente chiamati server, fi-
sicamente vicini, in cui ogni nodo è una scheda inserita in una struttura fisica, detta
"rack", dove solitamente la rete di interconnessione è una linea ad alta velocità e 
con larghezza di banda sufficientemente ampia. 
```

##### Network Systems
Sistemi in cui i nodi sono collegati da una rete locale (es: Ethernet) o geografica (es: Internet).

### Classificazione delle Architetture
La classificazione dei sistemi di calcolo più utilizzata è la *Tassonomia di Flynn (1972)*, in cui vengono inquadrate architetture e sistemi di elaborazione secondo due parametri:
1. **parallelismo a livello di istruzioni**
	- **Single Instruction Stream**, può essere eseguito un solo singolo flusso di istruzioni;
	- **Multiple Instruction Stream**, possono essere eseguiti più flussi di istruzioni in parallelo.
2. **parallelismo a livello di dati**
	- **Single Data Stream**, l'architettura è in grado di elaborare un singolo flusso sequenziale di dati;
	- **Multiple Data Streams**, l'architettura è in grado di processare più flussi di dati paralleli.

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Tassonomia%20di%20Flynn%20(1972)%20(1).png" alt="Tassonomia di Flynn (1972) (1)"/><img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Tassonomia%20di%20Flynn%20(1972)%20(2).png" alt="Tassonomia di Flynn (1972) (2)"/>

**SISD - Single Instruction (stream), Single Data (stream)**: sistemi monoprocessore che fanno riferimento all'architettura classica della macchina di Von Newman. Come dice il nome è in grado di gestire un singolo flusso di istruzioni (un programma) alla volta, su un singolo flusso di dati.

**SIMD - Single Instruction, Multiple Data**: architetture tipicamente parallele in cui vi sono diversi processori che, ad ogni istante, possono eseguire la stessa singola istruzione ma su dati diversi. Ad esempio rientrano in questa categoria gli array processors, di cui fanno parte anche le GPU.
```
NB: le GPU sono costituite da un insieme di nodi di elaborazione, a cui è assegnato una
singola control unit. Poiché elaborano dati che sono rappresentati da grandi matrici 
di informazioni (elaborazione di immagini), il modello SIMD risulta particolarmente ef-
ficace.
```
Rientrano in questa categoria anche i vector processors (migliaia di unità di elaborazione, non troppo potenti, ma che messe insieme e se controllate opportunamente, possono risolvere particolari classi di problemi in modo piuttosto efficiente e veloce).

**MIND - Multiple Instruction, Multiple Data**: insieme di nodi di elaborazione ognuno dei quali può eseguire flussi di istruzioni diverse su dati diversi. Ogni nodo può essere utilizzato da un processo che svolge operazioni diverse su dati differenti. Rientrano in questa categoria i sistemi multiprocessore (quelli che probabilmente conosciamo meglio), ma anche i MultiComputers.

**MISD - Multiple Instruction, Single Data**: il sistema è in grado di gestire un unico flusso di dati che ad ogni istante può essere elaborato con molteplici flussi di istruzioni. Non ci sono esempi particolarmente significativi da portare, ma è il caso dei "pipelined computer", dove lee diverse unità di elaborazione sono messe in cascata (pipeline), che lavora su quel flusso di dati, ognuna facendo qualcosa di differente.

### Tipi di Applicazioni
Ricapitolando, il progetto di applicazioni concorrenti dev'essere sviluppato in base al tipo di architettura, ma anche in base ai vincoli dati dal Sistema Operativo.

1. **multithreaded**:
	- si ha un'applicazione strutturata come un insieme di processi (thread) che:
		- permette di dominare la complessità del problema da risolvere;
		- aumentare l'efficienza, in quanto il carico di lavoro viene "scaricato" in parallelo;
		- semplificare la programmazione (secondo un modello di scomposizione dell'algoritmo in più parti che possono procedere contemporaneamente).
	- i processi possono condividere variabili;
	- sono caratterizzati dal fatto che generalmente esistono più processi che processori;
	- i processi sono schedulati ed eseguiti indipendentemente.
2. **sistemi multitasking/sistemi distribuiti**:
	- le componenti dell'applicazione (task) vengono eseguite su nodi (eventualmente virtuali) collegati tramite opportuni mezzi di interconnessione (es: canali);
	- i processi non possono condividere variabili, infatti comunicano scambiandosi messaggi;
	- questa organizzazione è tipica del modello client/server.
	I componenti in un sistema distribuito sono spesso multithreaded.
```
NB: in certi ambiti (sistemi distribuiti) esistono anche sistemi ibridi di applicazioni
in cui alcune parti sono multithreaded, mentre altre interagiscono a scambio di messag-
gio.
```
3. **applicazioni parallele**:
	- possiamo avere sia un modello in cui i processi condividono memoria, sia un modello a scambio di emssaggi;
	- hanno l'obbiettivo di risolvere il problema dato nel modo più veloce possibile, oppure un problema di dimensioni più grandi nello stesso tempo, sfruttando efficacemente il parallelismo disponibile a livello hardware;
	- sono eseguite su sistemi paralleli (es: HPC, array processors), facendo uso di algoritmi paralleli;
	- a seconda del modello architetturale, l'esecuzione è portata avanti da istruzioni/thread/processi paralleli che interagiscono utilizzando librerie specifiche.

### Processi Non Sequenziali e Tipi di Iterazione
**Algoritmo**: procedimento logico che deve essere eseguito per risolvere un determinato problema. È ciò che succede quando mettiamo in esecuzione un programma

**Programma**: descrizione di un algoritmo mediante un opportuno formalismo (linguaggio di programmazione), che rende possibile l'esecuzione dell'algoritmo da parte di un particolare elaboratore.

**Processo**: insieme ordinato degli eventi cui dà luogo un elaboratore quando opera sotto il controllo di un programma.

**Elaboratore**: entità astratta realizzata in hardware e parzialmente in software, in grado di eseguire programmi (descritti in un dato linguaggio).

**Evento**: esecuzione di un'operazione tra quelle appartenenti all'insieme che l'elaboratore sa riconoscere ed eseguire. Ogni evento determina una transizione di stato dell'elaboratore.
```
NB: un programma descrive non un processo, ma un insieme di processi, ognuno dei quali
è relativo all'esecuzione del programma da parte dell'elaboratore per un determinato 
insieme di dati in ingresso.
```

#### Processo Sequenziale
Con *processo sequenziale* si intende il caso in cui l'insieme degli eventi che avvengono all'interno dell'elaboratore quando esegue un dato programma (l'insieme degli eventi che fanno parte dell'esecuzione prende il nome di "traccia del programma"), sia una vera e propria sequenza. Ovvero che gli eventi siano ordinati in modo sequenziale: per ogni evento, tranne il primo e l'ultimo, c'è sempre un solo evento che lo precede ed un solo evento che lo segue.

**Grafo di Precedenza**: è uno schema che permette di rappresentare, tramite un formalismo, la traccia del programma. Ogni nodo rappresenta un singolo evento durante l'esecuzione del programma, ogni arco rappresenta la *precedenza temporale* tra un nodo ed il successivo. Nel caso di un algoritmo strettamente sequenziale, il grafo di precedenza che lo rappresenta si dice ad **ordinamento totale** (qualunque coppia di nodi venga presa nel grafo, questa coppia è sempre ordinata).

<img width="60%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Esempio%20MCD%20(algoritmo).png" alt="Algoritmo MCD"/> <img width="11%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Esempio%20MCD%20(grafo).png" alt="Grafo MCD"/>

#### Processo Non Sequenziale
Con *processo non sequenziale* si intende il caso in cui l'insieme degli eventi che lo descrive è ordinato secondo una relazione d'ordine parziale. In altre parole, un processo si dice non sequenziale se il grafo di precedenza che lo descrive non è ordinato in modo totale, ma è caratterizzato da un **ordinamento parziale**.

<img width="40%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Esempio%20Elaborazione%20File%20(algoritmo).png" alt="Algoritmo Elaborazione File"/> <img width="9%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Esempio%20Elaborazione%20File%20(grafo).png" alt="Grafo Elaborazione File"/>

L'esecuzione di un processo non sequenziale richiede:
- innanzitutto che o a livello software o hardware l'*elaboratore* sia *non sequenziale*, ovvero ci dia la possibilità di eseguire operazioni simultanee;
- un *linguaggio di programmazione non sequenziale*.

###### Elaboratore Non Sequenziale
È in grado di eseguire più operazioni contemporaneamente e si hanno due possibilità:
- sistemei multielaboratori (a)
- sistemi monoelaboratori (b)

<img width="60%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Elaboratori%20Non%20Sequenziali.png" alt="Elaboratori Non Sequenziali"/>

###### Linguaggi Concorrenti
I linguaggi concorrenti (o non sequenziali) hanno la caratteristica comune di consentire, a livello di programma, la descrizione di un insieme di attività concorrenti, tramite moduli che possono essere eseguiti in parallelo (es: processi sequenziali).\
In generale, un linguaggio concorrente permette di esprimere il (potenziale) parallelismo nell'esecuzione di moduli differenti.

Tipicamente ci sono due modi in cui viene realizzato il modulo concorrente di un linguaggio:
- parallelismo espresso a livello di **singola istruzione**, oggi poco usato (es: CSP, Occam);
- parallelismo a livello di **sequenza di istruzioni**, molto più frequente (es: Java, Ada, Go, ...).

#### Scomposizione di un Processo Non Sequenziale
Se il linguaggio concorrente permette di esprimere il parallelismo a livello di sequenza di istruzioni, allora si può scomporre un processo non sequenziale in un insieme di processi sequenziali eseguiti contemporaneamente, e far fronte alla complessità di un algoritmo non sequenziale.\
Una volta noto l'algoritmo non sequenziale si tratta di ricavare dal suo grafo di precedenza una collezione di grafi di processi sequenziali, che chiaramente saranno legati fra di loro da vincoli di precedenza.\
Le attività rappresentate dai processi possono essere:
- **completamente indipententi**, se l'evoluzione del processo non influenza quella degli altri. Di fatto nel grafo abbiamo un unico punto di partenza ed un unico punto di arrivo, ma i nodi potrebbero esprimersi, ad esempio, come una serie di 3 sequenze di nodi, che non sono però legate fra loro da vincoli di precedenza (gli eventi che appartengono ad un processo non sono legati ad altri eventi appartenenti ad altri processi);
- **interagenti**, se sono assoggettati a vincoli di precedenza tra stati che appartengono a processi diversi (vincoli di precedenza fra le operazioni e vincoli di sincronizzazione).

<!-- lezione 2021/10/06 -->
##### Interazione tra Processi
Esistono tre possibili tipi di interazione tra processi: *cooperazione*, *competizione*, *interferenza*.

###### Cooperazione
Comprende tutte le intearazioni *prevedibili* e *desiderate*, che sono in qualche modo dettate dall'algoritmo (date cioè dagli archi del grafo di precedenza ad ordinamento parziale). È insita nella logica che vogliamo rappresentare. Si può esprimere in 2 modi: **segnali temporali**, ovvero sincronizzazione pura, che esprime solo ed unicamente un vincolo di precedenza; **scambio di dati**, ovvero comunicazione vera e propria. In entrambi i casi esiste comunque un vincolo di precedenza tra gli eventi di processi diversi.\
C'è una relazione di causa ed effetto tra l'esecuzione dell'operazione di invio da parte del processo mittente e l'operazione di ricezione da parte del processo ricevente, con un vincolo di precedenza tra questi eventi (*sincronizzazione* di due processi). Il linguaggio di programmazione deve fornire i costrutti linguistici necessari a specificare la sincronizzazione e la eventuale comunicazione tra i processi.\
Esempio di cooperazione: interazione data da vincoli temporali (es: un processo esegue delle operazioni ogni 2 secondi, un altro ogni 3 ed un terzo li coordina attivando periodicamente tali processi).

###### Competizione
Consiste in un'interazione *prevedibile* e *non desiderata* (in quanto non fa parte dell'algoritmo che si vuole implementare, ma è solitamente dato da un limite della risorsa fisica o logica), ma *necessaria*. Infatti, la macchina concorrente, su cui i processi sono eseguiti, mette a disposizione un numero limitato di risorse condivise, disponibili nell'ambiente di esecuzione. Poiché alcune di queste non possono essere accedute o utilizzate contemporaneamente da più processi (o lo sono solo per un numero limitato), è necessario prevedere meccanismi che regolino la competizione, coordinando l'accesso alla risorsa da parte dei vari processi, in modo **mutuamente esclusivo**. Questo può determinare l'imposizione di vincoli di sincronizzazione (se una risorsa può essere usata da un solo processo alla volta, nella fase in cui sta venendo usata da un certo processo, nessun altro deve poterla utilizzare): un processo che tenta di accedere una risorsa già occupata (se non rispetta certi vincoli) dev'essere bloccato.\
**Sezione critica**: indica una sequenza di istruzioni con cui un processo accede ad una risorsa condivisa mutuamente esclusiva. Ad una risorsa possono essere associate, in casi particolari, anche più di una sezione critica. Se su una risorsa vale la mutua esclusione, sezioni critiche appartenenti alla stessa classe non possono eseguire contemporaneamente.\
Esempio di competizione: processi che devono accedere ad una stampante (risorsa mutuamente esclusiva).

###### Interferenza
È un tipo di interazione *non prevista* e *non desiderata*. Solitamente è provocata da errori del programmatore (infatti solitamente si cerca di eliminarle o escluderle), il quale non ha modellato correttamente l'interazione dei propri processi non sequenziali interagenti.\
Può non manifestarsi, in quanto a volte dipende dalla velocità relativa dei processi; gli errori possono manifestarsi nel corso dell'esecuzione del programma, a seconda delle diverse condizioni di velocità di esecuzione dei processi. In questi casi si parla di errori dipendenti dal tempo.\
Esempio tipico: deadlock.

### Architetture e Linguaggi per la Programmazione Concorrente
Avendo a disposizione una *macchina concorrente* **M** (in grado di eseguire più processi sequenziali contemporaneamente) e di un *linguaggio di programmazione* con il quale descrivere algoritmi non sequenziali, è possibile scrivere e far eseguire programmi concorrenti. L'elaborazione complessiva può essere descritta come un insieme di *processi sequenziali interagenti*.\
Le **proprietà di un linguaggio di programmazione concorrente** sono:
- fornire appositi costrutti con i quali sia possibile dichiarare moduli di programma destinati ad essere eseguiti come processi sequenziali distinti;
- non tutti i processi vengono eseguiti contemporaneamente. Alcuni processi vengono svolti se, dinamicamente, si verificano particolari condizioni. È quindi necessario poter specificare quando un processo deve essere attivato e termianto;
- devono essere presenti strumenti linguistici per specificare le interazioni che dinamicamente possono verificarsi tra i vari processi.

### Architettura di una Macchina Concorrente
<img width="70%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Architettura%20Macchina%20Concorrente%20(1).png" alt="Architettura Macchina Concorrente (1)"/>
M offre un certo numero di unità di elaborazione virtuali, che però non sempre sono in numero sufficiente per supportare l'esecuzione contemporanea dei processi di un programma concorrente.\
M è una macchina astratta ottenuta tramite tecniche software (o hardware) basandosi su una macchina fisica M' generalmente più semplice (con un numero di unità di elaborazione solitamente minore del numero dei processi).\

<img width="60%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Architettura%20Macchina%20Concorrente%20(2).png" alt="Architettura Macchina Concorrente (2)"/>

Al proprio interno M contiene ciò che dev'essere messo in atto quando viene richiesta l'esecuzione di processi concorrenti e tutto ciò che riguarda l'interazione (sincronizzazione con scambio di informazioni).\
Il nucleo corrisponde al supporto a tempo di esecuzione del compilatore di un linguaggio concorrente e comprende sempre due funzionalità base:
- meccanismo di **multiprogrammazione**, preposto alla gestione delle unità di elaborazione della macchina M', ovvero le unità reali. Questo meccanismo è realizzato dal kernel del SO, il quale dà la possibilità ad ogni processo creato all'intero dell'ambiente, di avere una visione diversa, come se avesse una CPU completamente dedicata. Ciò permette ai vari processi eseguiti sulla macchina astratta M di condividere l'uso delle unità reali di elaborazione (tale virtualizzazione si basa sulle politiche di *scheduling*) tramite l'allocazione in modo esclusivo ad ogni processo di un'unità virtuale di elaborazione. Di fatto la macchina astratta M offre l'illusione che il sistema sia composto da tante unità di elaborazione, quanti siano i processi in esecuzione;
- meccanismo di **sincronizzazione** e **comunicazione**, estende le potenzialità delle unità reali di elaborazione, rendendo disponibile alle unità virtuali strumenti mediante i quali sincronizzarsi e comunicare.
Oltre ai meccanismi di multiprogrammazione e interazione, è presente anche il meccanismo di **protezione** (controllo degli accessi alle risorse): importante per rilevare eventuali interferenze tra i processi; può essere realizzato in hardware o software nel supporto a tempo di esecuzione; comprende capabilities e ACL.

#### Architettura della Macchina M
In base all'organizzazione logica di M vengono definiti due modelli di interazione tra i processi:
1. Modello a **memoria comune**, ovvero le macchine astratte M sono collegate ad un'unica memoria principale. La visione proposta è aderente al modello del *multiprocessore*. Se queste sono le caratteristiche della macchina astratta, le unità di elaborazione astratte/virtuali prevedono l'interazione dei processi tramite oggetti contenuti in memoria comune (modello ad ambiente globale).
2. Modello a **scambio di messaggi**, ovvero gli elaboratori astratti realizzati dalla macchina M non condividono memoria. Sono posti in collegamento da una rete di comunicazione, ma non hanno possibilità di accedere alle stesse aree di memoria (tipico dei sistemi *multicomputer*). Ciascuna di queste aree virtuali viene fornita ad un certo processo, e sarà compito della macchina M fornire dei meccanismi opportuni che consentano la comunicazione fra i processi che eseguono (modello ad ambiente locale).

### Costrutti Linguistici per la Specifica della Concorrenza
Qualunque siano le caratteristiche della macchina astratta, il linguaggio di programmazione (concorrente) deve fornire costrutti che consentano di gestire i processi.\
Esistono due modelli diversi:

#### Fork/Join
Questo modello comprende appunto due primitive fondamentali: *fork* e *join*.

**Fork**: permette di creare e attivare un processo che inizia la propria esecuzione in *parallelo* con quella del processo chiamante.
```
NB: non va confusa con la system call di UNIX: in questo caso riguarda un modello più
generale e, a differenza della primitiva UNIX, si passa una funzione, col codice da e-
seguire, alla fork.
```

La fork ha un comportamento simile ad una exec: mentre quest'ultima implica l'attivazione di un processo che esegue il programma chiamato e la sospensione del programma chiamante, la fork prevede che il programma chiamante prosegua contemporaneamente con l'esecuzione della funzione chiamata. Coincide infatti con una biforcazione del grafo.

**Join**: consente di sincronizzare un processo con la terminazione di un altro processo, precedentemente creato tramite una fork.

In un grafo di precedenza, il nodo che rappresenta l'evento join ha due predecessori.

```
NB: a differenza della wait UNIX, nella join è necessario specificare il processo da
attendere, mentre nella wait no, di conseguenza quest'ultima si mette in attesa della 
terminazione di uno qualunque dei processi figli.
```

<img width="70%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Fork%20Join.png" alt="Fork/Join"/>

#### Cobegin/Coend
Questo modello trae ispirazione dalla programmazione strutturata, permettendo di esprimere la concorrenza tramite opportuni blocchi da inserire nel codice di opportuni programmi concorrenti. Si basa su due primitive fondamentali: *cobegin* e *coend*.

**Cobegin**: specifica l'inizio di un blocco di codice che deve essere eseguito in parallelo. All'interno di questo blocco si possono specificare una serie di operazioni o processi: la caratteristica degli statement in questo blocco è che ognuno di essi verrà eseguito concorrentemente rispetto agli altri di tale blocco. Inoltre, è possibile innestare un blocco dentro l'altro. 

**Coend**: indica la fine di un blocco di istruzioni parallele.

```
NB: fork/join è un formalismo più generale di cobegin/coend: tutti i grafi di preceden-
za possono essere espressi tramite fork/join ma non tutti possono essere espressi con 
cobegin/coend.
```

<img width="50%" src="https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/03%20-%20Programmazione%20Concorrente/Cobegin%20Coend.png" alt="Cobegin/Coend"/>

### Proprietà dei Programmi
I seguenti concetti permettono di specificare cosa succede quando il programma viene eseguito, di conseguenza sono utili per verificare la correttezza dei programmi realizzati.

**Traccia dell'esecuzione**: sequenza degli stati attraversati dal sistema di elaborazione durante l'esecuzione del programma. L'esecuzione di un programma è descritta dalla sua traccia.

**Stato**: insieme dei valori delle variabili definite nel programma più le variabili "implicite" (ad esempio il valore del program counter, o di altri registri).

#### Verifica della Correttezza di un Programma
**Programma sequenziale**: nei programmi sequenziali ogni esecuzione di un certo programma P su un particolare insieme di dati D genera sempre la stessa traccia (la verifica può essere svolta facilmente tramite debugging).\
**Programma concorrente**: nei programmi concorrenti l'esito dell'esecuzione dipende da quale sia l'effettiva sequenza cronologica di esecuzione delle istruzioni contenute, dunque ogni esecuzione di un certo programma P su un particolare insieme di dati D può dare origine a una traccia diversa, in quanto lo scheduling dei processi non è deterministico (la verifica è molto più difficile).

#### Proprietà di Safety e Liveness
**Proprietà di un programma**: attributo che è sempre vero, in ogni possibile traccia generata dalla sua esecuzione. Oltre alle proprietà di correttezza di un programma definite in precedenza, esistono anche altre proprietà, che solitamente si classificano in due categorie: *safety properties* e *liveness properties*.

**Safety**: garantisce che durante l'esecuzione di un programma *non si entrerà mai in uno stato "errato"*, ovvero in cui le variabili assumono valori non desiderati.

**Liveness**: garantisce che durante l'esecuzione del programma, *prima o poi si entrerà in uno stato "corretto"*, ovvero in cui le variabili assumono valori desiderati.

##### Proprietà dei Programmi Sequenziali
Le proprietà fondamentali che ogni programma sequenziale deve avere sono:
- *la correttezza del risultato finale*, ovvero che per ogni esecuzione, al termine del programma, il risultato ottenuto sia giusto -> **Safety**;
- *la terminazione*, ovvero prima o poi l'esecuzione del programma deve terminare -> **Liveness**.

<!-- lezione 2021/10/12 -->
##### Proprietà dei Programmi Concorrenti
Le proprietà fondamentali che ogni programma concorrente deve avere sono:
- *correttezza del risultato finale* -> **Safety**;
- *terminazione*, -> **Liveness**;
- *mutua esclusione nell'accesso a risorse condivise*, ovvero per ogni esecuzione non accadrà mai che più di un processo acceda contemporaneamente alla stessa risorsa -> **Safety**;
- *assenza di deadlock*, ovvero per ogni esecuzione non si verificheranno mai situazioni di blocco critico -> **Safety**;
- *asseenza di starvation*, ovvero prima o poi ogni processo potrà accedere alle risorse richieste -> **Liveness**.

###### Verifica di Proprietà nei Programmi Concorrenti
Poiché lo scheduling dei processi non è deterministico, il semplice testing su vari set di dati, per diverse ripetizioni dell'esecuzione, non dimostra rigorosamente il soddisfacimento di proprietà. Per questo motivo, un possibile approccio è l'utilizzo di una specifica "formale": tramite un processo di dimostrazione matematica si possono verificare le proprietà di un programma concorrente.

### Modelli di Interazione tra Processi
L'interazione tra processi può avvenire sostanzialmente secondo due modelli:
- modello a *memoria comune* (ambiente globale, memoria condivisa). In questo caso, la macchina astratta aderisce al modello multiprocessore, cioé offre ai programmi (che sono gli utilizzatori di tale macchina) un "modello" basato su un insieme di unità virtuali di elaborazione, ciascuna per l'esecuzione di un diverso processo, che condividono la stessa memoria. I processi possono vedere e accedere alle stesse aree di memoria.
- modello a *scambio di messaggi* (ambiente locale, memoria distribuita). In questo caso, i processori non condivisono memoria gli uni con gli altri, ma ognuno fa riferimento alla propria "memoria privata".

## 04 - Modello a Memoria Comune [![Vai al Capitolo Singolo](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/gfx/icon-document.png)](https://github.com/mikyll/Sistemi-Operativi-M/blob/main/capitoli/04%20-%20Modello%20a%20Memoria%20Comune.md "Vai al Capitolo Singolo")

