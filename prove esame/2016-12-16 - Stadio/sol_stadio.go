//Corni Gabriele 0000796972
//TURNO 1 TEMA 3

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAXBUFF = 100 //massimo buffer per i canali bufferizzati
const MAXPROC = 50  //massimo numero di goroutine per ogni tipo
const NUM = 3       //numero di operatori al varco
const LOCALI = 0
const OSPITI = 1

//definizione canali per terminazioni
var done = make(chan bool)
var terminaStadio = make(chan bool)
var terminaBiglietteria = make(chan bool)

//definizioni canali per richieste al server
var acquisto = make(chan int, MAXBUFF)

var ingressoLocali = make(chan int, MAXBUFF)
var ingressoOspiti = make(chan int, MAXBUFF)
var uscitaLocali = make(chan int)
var uscitaOspiti = make(chan int)

//definizioni canali per 3-way handshake
var ACK [MAXPROC]chan int

//guardia logica
func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}

//stato del server
func risorseStadio(l int, o int, oo int) {
	fmt.Printf("\n\n//////\n IN LOCALI: %d  IN OSPITI: %d   OPERATORI OCCUPATI: %d/%d\n/////\n\n\n", l, o, oo, NUM)
}

func spettatore(pid int) {

	//inizializzazione
	var tt, biglietto, soldi int
	var entrata, uscita chan int
	var nomeTribuna string

	fmt.Printf("[ spettatore %d ] -> partito\n", pid)

	tt = rand.Intn(5) + 1
	time.Sleep(time.Duration(tt) * time.Second)

	//acquisto biglietto versando offerta
	soldi = rand.Intn(10) + 1
	fmt.Printf("[ spettatore %d ] -> acquisto biglietto in biglietteria con offerta di euro %d\n", pid, soldi)

	acquisto <- pid
	biglietto = <-ACK[pid] //per semplicità, il biglietto non è numerato ma indica solo LOCALI/OSPITI

	if biglietto == LOCALI {
		entrata = ingressoLocali
		uscita = uscitaLocali
		nomeTribuna = "LOCALI"
	} else { //OSPITI
		entrata = ingressoOspiti
		uscita = uscitaOspiti
		nomeTribuna = "OSPITI"
	}

	//ingresso controlli di sicurezza
	tt = rand.Intn(5) + 1
	time.Sleep(time.Duration(tt) * time.Second)
	fmt.Printf("[ spettatore %d ] -> richiedo accesso a tribuna %s\n", pid, nomeTribuna)

	entrata <- pid
	<-ACK[pid]

	//tempo non trascurabile di controllo
	fmt.Printf("[ spettatore %d ] -> controllo in corso...\n", pid)
	tt = rand.Intn(5) + 1
	time.Sleep(time.Duration(tt) * time.Second)

	//uscita controlli di sicurezza
	fmt.Printf("[ spettatore %d ] -> ok, entro alla mia tribuna (%s)\n", pid, nomeTribuna)
	uscita <- pid

	//terminazione
	done <- true
	return

}

func biglietteria() {

	var tribuna int

	fmt.Printf("[ biglietteria ] -> partito\n")
	for {

		select {

		case x := <-acquisto:
			tribuna = rand.Intn(2) //0 o 1, ovvero locali o ospiti
			fmt.Printf("[ biglietteria ] -> spettatore %d va in tribuna %d\n", x, tribuna)
			ACK[x] <- tribuna // termine "call"

		case <-terminaBiglietteria:
			fmt.Printf("[ tribuna ] -> ho finito\n")
			done <- true
			return

		} //select
	} //for
} //server

func stadio() {

	var numLoc, numOsp int    //numero di spettatori per ogni tribuna
	var operatoriOccupati int //per tener traccia di quanti op liberi ci sono al varco

	fmt.Printf("[ stadio ] -> partito\n")
	for {
		//print dello stato
		risorseStadio(numLoc, numOsp, operatoriOccupati)
		select {

		

		//--------------richieste prioritarie---------------------------------
		case x := <-when(operatoriOccupati < NUM && ((numLoc >= numOsp) || (numLoc < numOsp && len(ingressoOspiti) == 0)), ingressoLocali):
			operatoriOccupati++
			fmt.Printf("[ stadio ] -> spettatore %d inizia controlli per LOCALI\n", x)
			ACK[x] <- 1 // termine "call"

		case x := <-when(operatoriOccupati < NUM && ((numOsp >= numLoc) || (numOsp < numLoc && len(ingressoLocali) == 0)), ingressoOspiti):
			operatoriOccupati++
			fmt.Printf("[ stadio ] -> spettatore %d inizia controlli per OSPITI\n", x)
			ACK[x] <- 1 // termine "call"

			//-------------richieste non prioritarie------------------------------
		case x := <-uscitaLocali: //completato controllo di sicurezza per uno spettatore della tribuna LOCALI
			numLoc++ //la tribuna si popola quando uno spettatore esce dal varco
			operatoriOccupati--
			fmt.Printf("[ stadio ] -> spettatore %d ha superato controlli, entra in LOCALI\n", x)

		case x := <-uscitaOspiti: //completato controllo di sicurezza per uno spettatore della tribuna OSPITI
			numOsp++ //la tribuna si popola quando uno spettatore esce dal varco
			operatoriOccupati--
			fmt.Printf("[ stadio ] -> spettatore %d ha superato controlli, entra in OSPITI\n", x)

		case <-terminaStadio:
			fmt.Printf("[ stadio ] -> ho finito\n")
			done <- true
			return

		} //select
	} //for
} //server

func main() {
	var S int

	fmt.Printf("\nQuanti spettatori (max %d)? ", MAXPROC)
	fmt.Scanf("%d", &S)

	//inizializzazione canali per handshake
	for i := 0; i < S; i++ {
		ACK[i] = make(chan int, MAXBUFF)
	}

	//avvio goroutines
	rand.Seed(time.Now().Unix())
	go stadio()
	go biglietteria()

	for i := 0; i < S; i++ {
		go spettatore(i)
	}

	//attesa terminazione spettatori
	for i := 0; i < S; i++ {
		<-done
	}

	//terminazione stadio e biglietteria
	terminaStadio <- true
	terminaBiglietteria <- true
	<-done
	<-done

	fmt.Printf("\nHO FINITO\n")
}
