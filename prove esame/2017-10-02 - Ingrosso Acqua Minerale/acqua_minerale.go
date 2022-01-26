package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ==================================================================
// COSTANTI =========================================================
const MAXBUFF = 100
const MAXPROC = 10
const MAX = 5 // capacità

const Y = 5     // Numero Bottiglie ACQUISTATE dall'acquirente
const Z = 5     // Numero VUOTI di Bottiglie precedentemente acquistate
const X = 10    // Numero Bottiglie CONSEGNATE dai fornitori
const PA = 2.00 // Prezzo Acquisto 1 Bottiglia (da acquirenti)
const PV = 1.00 // Prezzo Vendita 1 Bottiglia (da fornitori)

const MAX_P = 20 // Massimo numero di BOTTIGLIE PIENE che il magazzino può contenere
const MAX_V = 30 // Massimo numero di BOTTIGLIE VUOTE che il magazzino può contenere

const K = 5.00 // Valore di soglia per la cassa

const CONTANTI = 0
const BONIFICO = 1
const BANCOMAT = 1

// ==================================================================
// STRUTTURE DATI ===================================================
type Richiesta struct {
	id  int
	ack chan float64
}

// ==================================================================
// CANALI ===========================================================
var acquisto [2]chan Richiesta
var consegna [2]chan Richiesta

// ==================================================================
// CANALI DI JOIN ===================================================
var done = make(chan bool)
var terminaMagazzino = make(chan bool)

// ==================================================================
// FUNZIONI =========================================================
// Guardia logica
func when(b bool, c chan Richiesta) chan Richiesta {
	if !b {
		return nil
	}
	return c
}

// Sleep per un tot di secondi randomico nell'intervallo [0, timeLimit)
func sleepRandTime(timeLimit int) {
	if timeLimit > 0 {
		time.Sleep(time.Duration(rand.Intn(timeLimit)+1) * time.Second)
	}
}

func getModalita(m int) string {
	if m == CONTANTI {
		return "contanti"
	} else if m == BANCOMAT {
		return "bancomat"
	} else if m == BONIFICO {
		return "bonifico"
	} else {
		return ""
	}
}

// ==================================================================
// GOROUTINE ========================================================
func acquirente(id int) {
	fmt.Printf("[ACQUIRENTE %d] Inizio\n", id)
	r := Richiesta{id: id, ack: make(chan float64)}

	ma := rand.Intn(2) // Modalità pagamentoAcquisto

	// Richiede di acquistare Y bottiglie e consegnare
	fmt.Printf("[ACQUIRENTE %d] Voglio acquistare %d bottiglie pagando con %s, e consegnare %d vuoti\n", id, Y, getModalita(ma), Z)
	sleepRandTime((id / 2) + 1) // tempo impiegato per l'acquisto dopo che questo è stato accettato
	acquisto[ma] <- r

	sleepRandTime(2) // tempo impiegato per l'acquisto dopo che questo è stato accettato
	s := <-r.ack
	fmt.Printf("[ACQUIRENTE %d] Ho speso %.2f€\n", id, s)

	done <- true
	fmt.Printf("[ACQUIRENTE %d] Termino.\n", id)
}

func fornitore(id int) {
	fmt.Printf("[FORNITORE %d] Inizio\n", id)
	r := Richiesta{id: id, ack: make(chan float64)}
	mc := -1

	for {
		// Scelgo modalità di pagamento della consegna
		mc = rand.Intn(2)

		// ogni tanto consegna X bottiglie alla ditta.
		// quando consegna: 1. Ritira TUTTI i vuoti di bottiglia presenti in magazzino; 2. Incassa la somma delle X bottiglie acquistate
		fmt.Printf("[FORNITORE %d] Voglio consegnare %d bottiglie e pagare con %s\n", id, X, getModalita(mc))
		consegna[mc] <- r

		sleepRandTime(3) // tempo impiegato per la consegna dopo che questa è stata accettata
		res := <-r.ack
		if res < 0 {
			done <- true
			fmt.Printf("[FORNITORE %d] Termino\n", id)
			return
		} else {
			fmt.Printf("[FORNITORE %d] Ho ottenuto %.2f€\n", id, res)
		}
	}
}

func magazzino(ci, cc float64) {
	bp := 0     // bottiglie piene
	bv := 0     // bottiglie vuote
	cassa := ci // fondo cassa
	conto := cc // conto corrente ditta

	fmt.Printf("[MAGAZZINO] Inizio\n")
	for {
		select {
		// ACQUISTO
		case r := <-when(bv+Z < MAX_V && bp >= Y && (cassa < K || len(acquisto[BANCOMAT]) == 0), acquisto[CONTANTI]):
			bp -= Y         // scalo le bottiglie acquistate
			bv += Z         // aggiungo i vuoti
			cassa += Y * PA // Aggiorno la cassa (pagamento in contanti)
			r.ack <- Y * PA
			fmt.Printf("[MAGAZZINO] Acquisto con CONTANTI di %d bottiglie e deposito %d vuoti. Stato: Bottiglie %d/%d, Vuoti %d/%d. Cassa corrente: %.2f€\n", Y, Z, bp, MAX_P, bv, MAX_V, cassa)
			fmt.Printf("[MAGAZZINO] PROCESSI IN ATTESA: AC: %d, AB: %d, CC: %d, CB: %d\n", len(acquisto[CONTANTI]), len(acquisto[BANCOMAT]), len(consegna[CONTANTI]), len(consegna[BONIFICO]))

		case r := <-when(bv+Z < MAX_V && bp >= Y && (cassa >= K || len(acquisto[CONTANTI]) == 0), acquisto[BANCOMAT]):
			bp -= Y         // scalo le bottiglie acquistate
			bv += Z         // aggiungo i vuoti
			conto += Y * PA // Aggiorno il conto corrente (pagamento con bancomat)
			r.ack <- Y * PA
			fmt.Printf("[MAGAZZINO] Acquisto con BANCOMAT di %d bottiglie e deposito %d vuoti. Stato: Bottiglie %d/%d, Vuoti %d/%d. Cassa corrente: %.2f€\n", Y, Z, bp, MAX_P, bv, MAX_V, cassa)
			fmt.Printf("[MAGAZZINO] PROCESSI IN ATTESA: AC: %d, AB: %d, CC: %d, CB: %d\n", len(acquisto[CONTANTI]), len(acquisto[BANCOMAT]), len(consegna[CONTANTI]), len(consegna[BONIFICO]))

		// CONSEGNA
		case r := <-when(bp+X < MAX_P && (cassa >= K || len(consegna[BONIFICO]) == 0), consegna[CONTANTI]):
			bp += X         // aggiungo le bottiglie piene consegnate
			bv = 0          // ritiro tutte le bottiglie vuote
			cassa -= X * PV // Aggiorno il conto corrente (pagamento con bonifico)
			r.ack <- X * PV
			fmt.Printf("[MAGAZZINO] Consegna con CONTANTI di %d bottiglie. Stato: Bottiglie %d/%d, Vuoti %d/%d. Cassa corrente: %.2f€\n", X, bp, MAX_P, bv, MAX_V, cassa)
			fmt.Printf("[MAGAZZINO] PROCESSI IN ATTESA: AC: %d, AB: %d, CC: %d, CB: %d\n", len(acquisto[CONTANTI]), len(acquisto[BANCOMAT]), len(consegna[CONTANTI]), len(consegna[BONIFICO]))

		case r := <-when(bp+X < MAX_P && (cassa < K || len(consegna[CONTANTI]) == 0), consegna[BONIFICO]):
			bp += X         // aggiungo le bottiglie piene consegnate
			bv = 0          // ritiro tutte le bottiglie vuote
			conto -= X * PV // Aggiorno il conto corrente (pagamento con bonifico)
			r.ack <- X * PV
			fmt.Printf("[MAGAZZINO] Consegna con BONIFICO di %d bottiglie. Stato: Bottiglie %d/%d, Vuoti %d/%d. Cassa corrente: %.2f€\n", X, bp, MAX_P, bv, MAX_V, cassa)
			fmt.Printf("[MAGAZZINO] PROCESSI IN ATTESA: AC: %d, AB: %d, CC: %d, CB: %d\n", len(acquisto[CONTANTI]), len(acquisto[BANCOMAT]), len(consegna[CONTANTI]), len(consegna[BONIFICO]))

		case r := <-when(len(acquisto[BANCOMAT]) == 0 && len(acquisto[CONTANTI]) == 0 && bp+X > MAX_P, consegna[BONIFICO]):
			r.ack <- -1
		case r := <-when(len(acquisto[BANCOMAT]) == 0 && len(acquisto[CONTANTI]) == 0 && bp+X > MAX_P, consegna[CONTANTI]):
			r.ack <- -1
		case <-terminaMagazzino:
			done <- true
			fmt.Printf("\n\n[MAGAZZINO] Termino. Fondo cassa: %.2f€, Conto corrente: %.2f€\n", cassa, conto)
			if conto+cassa < 0.0 {
				fmt.Println("[MAGAZZINO] Fallimento!")
			}
			return
		}
	}
}

// ==================================================================
// MAIN =============================================================
func main() {
	fmt.Printf("\n[MAIN] Inizio.\n")
	rand.Seed(time.Now().Unix())

	/*var V1 int
	var V2 int

	fmt.Printf("\nQuanti Thread tipo1 (max %d)? ", MAXPROC)
	fmt.Scanf("%d\n", &V1)
	fmt.Printf("\nQuanti Thread tipo2 (max %d)? ", MAXPROC)
	fmt.Scanf("%d\n", &V2)*/

	// inizializzazione canali
	acquisto[CONTANTI] = make(chan Richiesta, MAXBUFF)
	acquisto[BANCOMAT] = make(chan Richiesta, MAXBUFF)
	consegna[CONTANTI] = make(chan Richiesta, MAXBUFF)
	consegna[BONIFICO] = make(chan Richiesta, MAXBUFF)

	nAcquirenti := 10
	nFornitori := 2

	go magazzino(100.0, 1000.0)
	// esecuzione acquirenti
	for i := 0; i < nAcquirenti; i++ {
		go acquirente(i)
	}
	// esecuzione fornitori
	for i := 0; i < nFornitori; i++ {
		go fornitore(i)
	}

	// join
	for i := 0; i < nAcquirenti; i++ {
		<-done
	}

	// join magazzino
	terminaMagazzino <- true
	<-done

	fmt.Printf("\n[MAIN] Fine.\n")
}
