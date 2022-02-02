package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAXPROC = 25
const MAXCICLI = 3
const NC = 10 //capacità scaffale chirurgiche
const NF = 10 //capacità scaffale ffp2
const LF = 4  //numerosità lotto FFP2
const LC = 3  //numerosità lotto Chir
const LM = 2  //numerosità lotto Misto

const T_MIX = 0  //tipo formato misto
const T_FFP2 = 1 //tipo formato FFP2
const T_CHIR = 2 //tipo formato CHIRURGICHE

const F_FFP2 = 0 // fornitore FFP2
const F_CHIR = 1 // fornitore CHIR

var done = make(chan bool)
var termina = make(chan bool)
var chiudi = make(chan bool)

//canali dedicati agli addetti
var inizio_prelievo [3]chan richiesta
var fine_prelievo = make(chan richiesta, 100)

//canali dedicati ai fornitori
var inizio_consegna [2]chan richiesta
var fine_consegna = make(chan richiesta, 100)

type richiesta struct {
	id   int
	tipo int
	ack  chan int
}

func sleep(t int) {
	time.Sleep(time.Duration(rand.Intn(t)+1) * time.Second)
}
func when(b bool, c chan richiesta) chan richiesta {
	if !b {
		return nil
	}
	return c

}

func AR(id int) { // addetto di reparto
	tipo := rand.Intn(3) //0 per Misto, 1 per ffp2 , 2 per chirurgiche
	r := richiesta{id, tipo, make(chan int)}
	sleep(10)
	cicli := rand.Intn(MAXCICLI) + 1
	for i := 0; i < cicli; i++ { // ciclo prelievi
		tipo := rand.Intn(3) //0 per Misto, 1 per ffp2 , 2 per chirurgiche
		r.tipo = tipo
		if tipo == T_MIX {
			fmt.Printf("[AR %d] richiedo prelievo di un lotto  misto\n", id)
		} else if tipo == T_FFP2 {
			fmt.Printf("[AR %d] richiedo prelievo di un lotto  FFP2\n", id)
		} else {
			fmt.Printf("[AR %d] richiedo prelievo di un lotto  Chirurgiche\n", id)
		}
		inizio_prelievo[tipo] <- r
		<-r.ack
		sleep(10) //tempo impiegato per il prelievo
		//richiesta fine erogazione
		fine_prelievo <- r
		<-r.ack
	}
	//terminazione
	fmt.Printf("[AR %d] termino \n", id)
	done <- true

}
func fornitore(tipo int) { //tipo è FFP2 o chirurgica
	r := richiesta{tipo, tipo, make(chan int)}

	for { // ciclo consegne
		sleep(5)
		fmt.Printf("[Fornitore %d] richiesta riempimento mascherine  di tipo %d\n", tipo, tipo)
		inizio_consegna[tipo] <- r
		flag := <-r.ack
		sleep(20) //tempo impiegato per riempire lo scaffale
		//fine rifornimento:
		if flag == 0 { //il fornitore deve terminare
			done <- true
			fmt.Printf("[fornitore% d] finito!\n", tipo)
			return
		}
		fine_consegna <- r
		<-r.ack
		fmt.Printf("[Fornitore %d] ho finito di riempire lo scaffale di mascherine di tipo %d\n", tipo, tipo)
		sleep(3)
	}
}
func magazzino() {
	var Disp_FFP2 = NF   //numero scatole disponibili nello scaffale FFP2
	var Disp_CHIR = NC   //numero scatole disponibili nello scaffale CHIR
	var Forn_in_FFP2 = 0 //numero fornitori in consegna sullo scaffale FFP2
	var Forn_in_CHIR = 0 //numero fornitori in consegna sullo scaffale CHIR
	var AR_in_FFP2 = 0   //numero addetti in prelievo dallo scaffale FFP2
	var AR_in_CHIR = 0   //numero addetti in prelievo dallo scaffale CHIR
	var END = false      //diventa true quando tutti gli AR sono terminati
	for {
		select {
		// PRELIEVI:
		case x := <-when(Disp_CHIR >= LM && Disp_FFP2 >= LM && Forn_in_CHIR == 0 && Forn_in_FFP2 == 0, inizio_prelievo[T_MIX]):
			AR_in_CHIR++
			AR_in_FFP2++
			Disp_CHIR -= LM
			Disp_FFP2 -= LM
			fmt.Printf("[Magazzino] l'addetto %d inizia a prelevare un lotto di tipo %d \n", x.id, x.tipo)
			x.ack <- 1
		case x := <-when(Disp_FFP2 >= LF && Forn_in_FFP2 == 0 && len(inizio_prelievo[T_MIX]) == 0, inizio_prelievo[T_FFP2]):
			AR_in_FFP2++
			Disp_FFP2 -= LF
			fmt.Printf("[Magazzino] l'addetto %d inizia a prelevare un lotto di tipo %d \n", x.id, x.tipo)
			x.ack <- 1
		case x := <-when(Disp_CHIR >= LC && Forn_in_CHIR == 0 && len(inizio_prelievo[T_MIX]) == 0 && len(inizio_prelievo[T_FFP2]) == 0, inizio_prelievo[T_CHIR]):
			AR_in_CHIR++
			Disp_CHIR -= LC
			fmt.Printf("[Magazzino] l'addetto %d inizia a prelevare un lotto di tipo %d \n", x.id, x.tipo)
			x.ack <- 1
		case x := <-fine_prelievo:
			if x.tipo == T_MIX {
				AR_in_CHIR--
				AR_in_FFP2--
			} else if x.tipo == T_FFP2 {
				AR_in_FFP2--
			} else {
				AR_in_CHIR--
			}
			fmt.Printf("[Magazzino] l'addetto %d ha terminato il prelievo\n", x.id)
			x.ack <- 1
		//CONSEGNE:
		case x := <-when(END == false && Disp_CHIR < NC && Forn_in_CHIR == 0 && AR_in_CHIR == 0 && ((Disp_CHIR >= Disp_FFP2 && len(inizio_consegna[F_FFP2]) == 0) || Disp_CHIR < Disp_FFP2), inizio_consegna[F_CHIR]):
			Disp_CHIR = NC
			Forn_in_CHIR++
			x.ack <- 1
		case x := <-when(END == false && Disp_FFP2 < NF && Forn_in_FFP2 == 0 && AR_in_FFP2 == 0 && ((Disp_CHIR < Disp_FFP2 && len(inizio_consegna[F_CHIR]) == 0) || Disp_CHIR >= Disp_FFP2), inizio_consegna[F_FFP2]):
			Disp_FFP2 = NF
			Forn_in_FFP2++
			fmt.Printf("[Magazzino] il fornitore %d ha iniziato a riempire lo scaffale %d\n", x.id, x.tipo)
			x.ack <- 1

		case x := <-fine_consegna:
			if x.tipo == F_FFP2 {
				Forn_in_FFP2--
			} else {
				Forn_in_CHIR--
			}
			fmt.Printf("[Magazzino] il fornitore %d ha terminato di riempire lo scaffale %d\n", x.id, x.tipo)
			x.ack <- 1
		case <-termina:
			END = true
			fmt.Printf("[Magazzino] il magazzino sta per chiudere ...\n")

		case x := <-when(END == true, inizio_consegna[0]):
			x.ack <- 0
		case x := <-when(END == true, inizio_consegna[1]):
			x.ack <- 0
		case <-chiudi:
			END = true
			fmt.Printf("[Magazzino] il magazzino CHIUDE!\n")
			done <- true
			return

		}
		// per debug:
		//fmt.Printf("\n[Magazzino] stato attuale: \nFFP2=%d, \nCHIR=%d, \nARin prelievo C= %d\nARin prelievoF=%d\nFornitori_inConsegnaF=%d\nFornitori_inconsegnaC=%d\n",
		//	Disp_FFP2, Disp_CHIR, AR_in_CHIR, AR_in_FFP2, Forn_in_FFP2, Forn_in_CHIR)
	}
}
func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2; i++ {
		inizio_consegna[i] = make(chan richiesta, 100)
	}
	for i := 0; i < 3; i++ {
		inizio_prelievo[i] = make(chan richiesta, 100)
	}
	numeroAR := rand.Intn(MAXPROC) + 2 //per avere almeno 2 addetti
	fmt.Printf("Il numero di AR è: %d\n", numeroAR)
	go magazzino()
	go fornitore(F_CHIR)
	go fornitore(F_FFP2)
	for i := 0; i < numeroAR; i++ {
		go AR(i)
	}
	for i := 0; i < numeroAR; i++ { // attesa AR
		<-done
	}

	fmt.Printf("[MAIN] tutti gli AR sono terminati!\n")
	termina <- true //avverto il server che tutti gli addetti sono terminati

	for i := 0; i < 2; i++ { // attendo i due fornitori
		<-done
	}
	chiudi <- true // comando al server di terminare
	<-done         // attesa server
}
