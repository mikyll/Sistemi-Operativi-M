/*
Nicola Alessi (nicola.alessi@studio.unibo.it)
NOTE: mi ha tolto 1 punto perchè non le piaceva la terminazione. Il resto andava bene.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/////////////////////////////////////////////////////////////////////
//Constants
/////////////////////////////////////////////////////////////////////
const MAXBUFF = 100
const PA = 2
const PV = 1
const Y = 3
const Z = 2
const X = 10
const K = 20
const MAX_P = 100
const MAX_V = 150

const CONTANTI = 0
const BANCOMAT = 1
const BONIFICO = 1

/////////////////////////////////////////////////////////////////////
//Structures
/////////////////////////////////////////////////////////////////////
type Richiesta struct {
	ID  int
	ack chan int
}

/////////////////////////////////////////////////////////////////////
//Channels
/////////////////////////////////////////////////////////////////////
var chAcquirenteContanti = make(chan Richiesta, MAXBUFF)
var chAcquirenteBancomat = make(chan Richiesta, MAXBUFF)
var chFornitoreContanti = make(chan Richiesta, MAXBUFF)
var chFornitoreBonifico = make(chan Richiesta, MAXBUFF)

/////////////////////////////////////////////////////////////////////
//GORoutine join
/////////////////////////////////////////////////////////////////////
var terminateServer = make(chan bool)
var done chan bool

/////////////////////////////////////////////////////////////////////
//GORoutines
/////////////////////////////////////////////////////////////////////
func acquirente(ID int) {

	richiesta := Richiesta{ID: ID, ack: make(chan int)}
	pagamento := rand.Intn(2)

	switch pagamento {
	case BANCOMAT:
		chAcquirenteBancomat <- richiesta
		fmt.Println("acquirenteID ", ID, " richiesta acquisto ", Y, "bottiglie con BANCOMAT")

	case CONTANTI:
		chAcquirenteContanti <- richiesta
		fmt.Println("acquirenteID ", ID, " richiesta acquisto ", Y, "bottiglie in CONTANTI")

	default:
		fmt.Println("error")
		done <- true
		return
	}

	spesa := <-richiesta.ack
	if spesa >= 0 {
		fmt.Println("acquirenteID ", ID, " acquisto riuscito! ho speso: ", spesa, " ho consegnato ", Z, " bottiglie vuote")
	} else {
		fmt.Println("acquirenteID ", ID, " termina...")
	}

	done <- true
}
func fornitore(ID int) {

	for {
		rifornimento := Richiesta{ID: ID, ack: make(chan int)}
		pagamento := rand.Intn(2)

		switch pagamento {
		case BONIFICO:
			chFornitoreBonifico <- rifornimento
			fmt.Println("fornitoreID ", ID, " rifornimento ", X, " bottiglie piene, pagamento BONIFICO")

		case CONTANTI:
			chFornitoreContanti <- rifornimento
			fmt.Println("acquirenteID ", ID, "rifornimento ", X, " bottiglie piene, pagamento CONTANTI")

		default:
			fmt.Println("error")
			done <- true
			return
		}

		incasso := <-rifornimento.ack
		if incasso == -1 {
			fmt.Println("rifornitoreID ", ID, "termina...")
			done <- true
			return
		}
		fmt.Println("rifornitoreID ", ID, "ho incassato ", incasso)
		time.Sleep(1e9)
	}
}

func ditta() {

	fmt.Println("Ditta GoRoutine start!")
	running := true
	cassa := X * PV
	contoCorrente := X * PV

	numPiene := 0
	numVuote := 0
	for {
		select {
		//acquirente
		case richiesta := <-when(running && numPiene >= Y && numVuote+Z <= MAX_V && (cassa < K || len(chAcquirenteBancomat) == 0), chAcquirenteContanti):
			spesaAcquirente := Y * PA
			cassa = cassa + spesaAcquirente
			numPiene = numPiene - Y
			numVuote = numVuote + Z
			richiesta.ack <- spesaAcquirente
		case richiesta := <-when(running && numPiene >= Y && numVuote+Z <= MAX_V && (cassa >= K || len(chAcquirenteContanti) == 0), chAcquirenteBancomat):
			spesaAcquirente := Y * PA
			contoCorrente = contoCorrente + spesaAcquirente
			numPiene = numPiene - Y
			numVuote = numVuote + Z
			richiesta.ack <- spesaAcquirente

		//fornitore
		case richiesta := <-when(running && numPiene+X <= MAX_P && cassa >= X*PV && (cassa >= K || len(chFornitoreBonifico) == 0), chFornitoreContanti):
			numVuote = 0
			numPiene = numPiene + X
			incassoFornitore := X * PV
			cassa = cassa - incassoFornitore
			richiesta.ack <- incassoFornitore
		case richiesta := <-when(running && numPiene+X <= MAX_P && cassa >= X*PV && (cassa < K || len(chFornitoreContanti) == 0), chFornitoreBonifico):
			numVuote = 0
			numPiene = numPiene + X
			incassoFornitore := X * PV
			contoCorrente = contoCorrente - incassoFornitore
			richiesta.ack <- incassoFornitore

		//terminazione
		case richiesta := <-when(!running, chAcquirenteContanti):
			richiesta.ack <- -1
		case richiesta := <-when(!running, chAcquirenteBancomat):
			richiesta.ack <- -1
		case richiesta := <-when(!running, chFornitoreContanti):
			richiesta.ack <- -1
		case richiesta := <-when(!running, chFornitoreBonifico):
			richiesta.ack <- -1

		case <-terminateServer:
			if running {
				fmt.Println("ditta: start termination")
				running = false
			} else {
				fmt.Println("ditta: terminated..")
				done <- true
				return
			}
		}
	}
}

/////////////////////////////////////////////////////////////////////
//Test main
/////////////////////////////////////////////////////////////////////
func main() {
	var numAcquirenti int
	var numFornitori int
	rand.Seed(time.Now().UTC().UnixNano()) //rand.Intn(2)

	fmt.Printf("\n[main] Quanti acquirenti? \n")
	fmt.Scanf("%d", &numAcquirenti)

	fmt.Printf("\n[main] Quanti fornitori? \n")
	fmt.Scanf("%d", &numFornitori)

	done = make(chan bool, numAcquirenti+numFornitori)

	go ditta()

	for i := 0; i < numAcquirenti; i++ {
		go acquirente(i)

	}

	for i := 0; i < numFornitori; i++ {
		go fornitore(i)

	}

	//GoRoutine join
	for i := 0; i < numAcquirenti; i++ {
		<-done
	}

	terminateServer <- true

	for i := 0; i < numFornitori; i++ {
		<-done
	}

	fmt.Println("main: start termination")
	terminateServer <- true
	<-done
	fmt.Println("main: terminated")
}

/////////////////////////////////////////////////////////////////////
// Auxiliary functions
/////////////////////////////////////////////////////////////////////
func when(b bool, c chan Richiesta) chan Richiesta {
	if !b {
		return nil
	}
	return c
}
