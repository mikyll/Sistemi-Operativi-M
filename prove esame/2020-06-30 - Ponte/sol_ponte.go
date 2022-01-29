package main

import (
	"fmt"
	"math/rand"
	"time"
)

/////////////////////////////////////////////////////////////////////
// Strutture Dati
/////////////////////////////////////////////////////////////////////
type Richiesta struct {
	id  int
	ack chan int
}

/////////////////////////////////////////////////////////////////////
// Costanti
/////////////////////////////////////////////////////////////////////
const MAXBUFF = 100
const MAXPROC = 60
const MAXBAR = 6
const MAX_AUTO = 5 // capacità

const sTA, sTB int = 0, 1 // stati ponte

const tNS, tSN int = 0, 1 // stati transito

/////////////////////////////////////////////////////////////////////
// Canali
/////////////////////////////////////////////////////////////////////
var ponteB [2]chan Richiesta
var ponteAin [4]chan Richiesta
var ponteAout = make(chan Richiesta, MAXBUFF)

const IND_B_IN, IND_B_OUT int = 0, 1                  // indici richieste canale ponte barche
const IND_N, IND_S, IND_P_N, IND_P_S int = 0, 1, 2, 3 // indici richieste canale ponte auto (normale e pubblico P)

/////////////////////////////////////////////////////////////////////
// Join Goroutine
/////////////////////////////////////////////////////////////////////
var done = make(chan bool)
var termina = make(chan bool)

/////////////////////////////////////////////////////////////////////
// Funzioni Ausiliarie
/////////////////////////////////////////////////////////////////////
// se si usa struttura dati modificare la when
func when(b bool, c chan Richiesta) chan Richiesta {
	if !b {
		return nil
	}
	return c
}

func whenInt(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}

func sleepTimeMill(t int) {
	if t > 0 {
		time.Sleep(time.Duration(t) * time.Millisecond)
	}
}

func sleepTime(t int) {
	if t > 0 {
		time.Sleep(time.Duration(t) * time.Second)
	}
}

func sleepRandTime(timeLimit int) {
	if timeLimit > 0 {
		time.Sleep(time.Duration(rand.Intn(timeLimit)+1) * time.Second)
	}
}

func sleepRandTimeBetw(min int, max int) {
	if min > 0 && min < max {
		time.Sleep(time.Duration(rand.Intn(max-min+1)+min) * time.Second)
	}
}

func makeChannels() {
	for i := 0; i < 2; i++ {
		ponteB[i] = make(chan Richiesta, MAXBUFF)
		ponteAin[i] = make(chan Richiesta, MAXBUFF)
	}
	ponteAin[2] = make(chan Richiesta, MAXBUFF)
	ponteAin[3] = make(chan Richiesta, MAXBUFF)
}

/////////////////////////////////////////////////////////////////////
// Goroutine
/////////////////////////////////////////////////////////////////////
func veicolo(id int, tipo int) {
	sleepRandTime(15)

	richiesta := Richiesta{id, make(chan int)}

	fmt.Printf("\n[veicolo %d]: tipo %d. Richiedo utilizzo ponte.", id, tipo)
	ponteAin[tipo] <- richiesta
	<-richiesta.ack
	fmt.Printf("\n[veicolo %d]: tipo %d. Percorro il ponte...", id, tipo)

	sleepTimeMill(600)

	ponteAout <- richiesta
	<-richiesta.ack
	fmt.Printf("\n[veicolo %d]: tipo %d. Ponte percorso.", id, tipo)

	done <- true
}

func barca(id int) {
	sleepRandTime(15)

	richiesta := Richiesta{id, make(chan int)}

	fmt.Printf("\n[barca %d]: richiedo utilizzo ponte", id)
	ponteB[IND_B_IN] <- richiesta
	<-richiesta.ack
	fmt.Printf("\n[barca %d]: percorro...", id)

	sleepTime(2)

	ponteB[IND_B_OUT] <- richiesta
	<-richiesta.ack
	fmt.Printf("\n[barca %d]: tratta percorsa.", id)

	done <- true
}

func ponte() {
	var stato int = sTA
	var trans int = tNS

	var vTrans int = 0

	for {
		select {
		case x := <-when(stato == sTB, ponteB[IND_B_IN]):
			vTrans++
			fmt.Printf("\n[ponte]: barca %d in transito.\tSt: %d\tVtrans: %d", x.id, stato, vTrans)
			x.ack <- 1

		case x := <-ponteB[IND_B_OUT]:
			vTrans--
			fmt.Printf("\n[ponte]: barca %d passata.\tSt: %d\tVtrans: %d", x.id, stato, vTrans)
			x.ack <- 1
			if len(ponteB[IND_B_IN]) == 0 && vTrans == 0 {
				fmt.Printf("\n[ponte]: barche passate. Abbasso il ponte!")
				stato = sTA
			}

		case x := <-when(stato == sTA && ((vTrans > 0 && vTrans < MAX_AUTO && trans == tNS) || (vTrans == 0 && trans == tSN)) && len(ponteB[IND_B_IN]) == 0, ponteAin[IND_P_N]):
			if trans == tSN {
				trans = tNS
			}
			vTrans++
			fmt.Printf("\n[ponte]: veicolo pubblico %d in transito da N a S.\tSt: %d\tVtrans: %d", x.id, stato, vTrans)
			x.ack <- 1

		case x := <-when(stato == sTA && ((vTrans > 0 && vTrans < MAX_AUTO && trans == tSN) || (vTrans == 0 && trans == tNS)) && len(ponteB[IND_B_IN]) == 0, ponteAin[IND_P_S]):
			if trans == tNS {
				trans = tSN
			}
			vTrans++
			fmt.Printf("\n[ponte]: veicolo pubblico %d in transito da S a N.\tSt: %d\tVtrans: %d", x.id, stato, vTrans)
			x.ack <- 1

		case x := <-when(stato == sTA && ((vTrans > 0 && vTrans < MAX_AUTO && trans == tNS) || (vTrans == 0 && trans == tSN)) &&
			len(ponteB[IND_B_IN]) == 0 && len(ponteAin[IND_P_S]) == 0 && len(ponteAin[IND_P_N]) == 0, ponteAin[IND_N]):
			if trans == tSN {
				trans = tNS
			}
			vTrans++
			fmt.Printf("\n[ponte]: veicolo privato %d in transito da N a S.\tSt: %d\tVtrans: %d", x.id, stato, vTrans)
			x.ack <- 1

		case x := <-when(stato == sTA && ((vTrans > 0 && vTrans < MAX_AUTO && trans == tSN) || (vTrans == 0 && trans == tNS)) &&
			len(ponteB[IND_B_IN]) == 0 && len(ponteAin[IND_P_S]) == 0 && len(ponteAin[IND_P_N]) == 0, ponteAin[IND_S]):
			if trans == tNS {
				trans = tSN
			}
			vTrans++
			fmt.Printf("\n[ponte]: veicolo privato %d in transito da S a N.\tSt: %d\tVtrans: %d", x.id, stato, vTrans)
			x.ack <- 1

		case x := <-ponteAout:
			vTrans--
			fmt.Printf("\n[ponte]: veicolo %d passato.\tSt: %d\tVtrans: %d", x.id, stato, vTrans)
			x.ack <- 1
			if vTrans == 0 && len(ponteB[IND_B_IN]) > 0 {
				fmt.Printf("\n[ponte]: barche in attesa. Alzo il ponte!")
				stato = sTB
			}

		case <-termina:
			fmt.Printf("\n\n[ponte]: Fine!")
			done <- true
			return
		}
	}
}

/////////////////////////////////////////////////////////////////////
// Main
/////////////////////////////////////////////////////////////////////
func main() {
	rand.Seed(time.Now().Unix())

	makeChannels()

	go ponte()

	for i := 0; i < MAXPROC; i++ {
		go veicolo(i, rand.Intn(4))
	}

	for i := 0; i < MAXBAR; i++ {
		go barca(i)
	}

	// join
	for i := 0; i < MAXPROC+MAXBAR; i++ {
		<-done
	}

	termina <- true
	<-done
	fmt.Printf("\n[Main]: fine.\n")
}
