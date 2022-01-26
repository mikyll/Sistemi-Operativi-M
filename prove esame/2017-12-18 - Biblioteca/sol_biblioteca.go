package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Costanti
*/
const MAXBUFF = 100
const TRIENNALE = 0
const TRIENNALE_LAUREANDO = 1
const MAGISTRALE = 2
const MAGISTRALE_LAUREANDO = 3

const N = 3

const N_STUD_TRI = 10
const N_STUD_MAG = 19
const N_STUD_TRI_L = 15
const N_STUD_MAG_L = 7
const N_STUD_TOT = N_STUD_TRI + N_STUD_TRI_L + N_STUD_MAG + N_STUD_MAG_L

/*
	Strutture
*/
type Richiesta struct {
	ID  int
	ack chan int
}

/*
	Canali
*/
var richiesta_consegna = make(chan Richiesta, MAXBUFF)
var richiesta_ingresso [4]chan Richiesta
var richiesta_uscita [4]chan Richiesta
var richiesta_ritiro = make(chan Richiesta, MAXBUFF)

/*
	Goroutine join
*/
var terminateServer = make(chan bool)
var done chan bool

/*
	Funzioni aux
*/
func when(b bool, c chan Richiesta) chan Richiesta {
	if !b {
		return nil
	}
	return c
}

func printTipo(typ int) string {
	switch typ {
	case TRIENNALE:
		return "triennale"
	case MAGISTRALE:
		return "magistrale"
	case TRIENNALE_LAUREANDO:
		return "triennale laureando"
	case MAGISTRALE_LAUREANDO:
		return "magistrale laureando"
	default:
		return ""
	}
}

/*
	Goroutines
*/
func studente(ID int, typ int) {
	richiesta := Richiesta{ID: ID, ack: make(chan int)}

	richiesta_consegna <- richiesta
	<-richiesta.ack
	fmt.Println(printTipo(typ), " ID ", ID, " consegna documento")
	richiesta_ingresso[typ] <- richiesta
	//fmt.Println(printTipo(typ), " ID ", ID, " richiesta ingresso in bibilioteca")
	<-richiesta.ack
	fmt.Println(printTipo(typ), " ID ", ID, " sto studiando.. ")
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	richiesta_uscita[typ] <- richiesta
	<-richiesta.ack
	//	fmt.Println(printTipo(typ), " ID ", ID, " ho ritirato il documento")
	richiesta_ritiro <- richiesta
	<-richiesta.ack
	//	fmt.Println(printTipo(typ), " ID ", ID, " sono uscito")
	done <- true
}

func portiniere() {
	fmt.Println("Portiniere start!")
	var documenti int = 0
	for {
		select {
		case richiesta := <-richiesta_consegna:
			richiesta.ack <- 1
			documenti++
			fmt.Println("Ricevuto documento da studente matricola ", richiesta.ID, " documenti dentro: ", documenti)
		case richiesta := <-richiesta_ritiro:
			richiesta.ack <- 1
			documenti--
			fmt.Println("Restituito documento studente matricola ", richiesta.ID, " documenti dentro: ", documenti)
		case <-terminateServer:
			done <- true
			return
		}
	}
}

func biblioteca() {
	fmt.Println("Biblioteca start!")
	var triennali int = 0
	var magistrali int = 0

	for {
		select {

		case richiesta := <-when(triennali+magistrali < N && (magistrali <= triennali || (len(richiesta_ingresso[TRIENNALE]) == 0 && len(richiesta_ingresso[TRIENNALE_LAUREANDO]) == 0)), richiesta_ingresso[MAGISTRALE_LAUREANDO]):
			richiesta.ack <- 1
			magistrali++
			fmt.Println("Ingresso in biblioteca magistrale laureando ", richiesta.ID, " magistrali dentro: ", magistrali, " triennali dentro: ", triennali)
		case richiesta := <-when(triennali+magistrali < N && ((magistrali <= triennali && len(richiesta_ingresso[MAGISTRALE_LAUREANDO]) == 0) || (len(richiesta_ingresso[TRIENNALE]) == 0 && len(richiesta_ingresso[TRIENNALE_LAUREANDO]) == 0) && len(richiesta_ingresso[MAGISTRALE_LAUREANDO]) == 0), richiesta_ingresso[MAGISTRALE]):
			richiesta.ack <- 1
			magistrali++
			fmt.Println("Ingresso in biblioteca magistrale ", richiesta.ID, " magistrali dentro: ", magistrali, " triennali dentro: ", triennali)
		case richiesta := <-when(triennali+magistrali < N && (triennali < magistrali || (len(richiesta_ingresso[MAGISTRALE]) == 0 && len(richiesta_ingresso[MAGISTRALE_LAUREANDO]) == 0)), richiesta_ingresso[TRIENNALE_LAUREANDO]):
			richiesta.ack <- 1
			triennali++
			fmt.Println("Ingresso in biblioteca triennale laureando ", richiesta.ID, " magistrali dentro: ", magistrali, " triennali dentro: ", triennali)
		case richiesta := <-when(triennali+magistrali < N && ((triennali < magistrali && len(richiesta_ingresso[TRIENNALE_LAUREANDO]) == 0) || (len(richiesta_ingresso[MAGISTRALE_LAUREANDO]) == 0 && len(richiesta_ingresso[MAGISTRALE]) == 0) && len(richiesta_ingresso[TRIENNALE_LAUREANDO]) == 0), richiesta_ingresso[TRIENNALE]):
			richiesta.ack <- 1
			triennali++
			fmt.Println("Ingresso in biblioteca triennale ", richiesta.ID, " magistrali dentro: ", magistrali, " triennali dentro: ", triennali)
		case richiesta := <-richiesta_uscita[TRIENNALE]:
			richiesta.ack <- 1
			triennali--
			fmt.Println("Uscita triennale ", richiesta.ID, " magistrali dentro: ", magistrali, " triennali dentro: ", triennali)
		case richiesta := <-richiesta_uscita[TRIENNALE_LAUREANDO]:
			richiesta.ack <- 1
			triennali--
			fmt.Println("Uscita magistrale ", richiesta.ID, " magistrali dentro: ", magistrali, " triennali dentro: ", triennali)
		case richiesta := <-richiesta_uscita[MAGISTRALE]:
			richiesta.ack <- 1
			magistrali--
			fmt.Println("Uscita magistrale ", richiesta.ID, " magistrali dentro: ", magistrali, " triennali dentro: ", triennali)
		case richiesta := <-richiesta_uscita[MAGISTRALE_LAUREANDO]:
			richiesta.ack <- 1
			magistrali--
			fmt.Println("Uscita magistrale laureando ", richiesta.ID, " magistrali dentro: ", magistrali, " triennali dentro: ", triennali)

		case <-terminateServer:
			done <- true
			return
		}
	}
}

/*
	Main di prova
*/
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	done = make(chan bool)

	for i := 0; i < 4; i++ {
		richiesta_ingresso[i] = make(chan Richiesta, MAXBUFF)
		richiesta_uscita[i] = make(chan Richiesta, MAXBUFF)
	}

	for i := 0; i < N_STUD_TRI; i++ {
		go studente(i, TRIENNALE)
		//time.Sleep(1e8)
	}
	for i := N_STUD_TRI; i < N_STUD_TRI+N_STUD_MAG; i++ {
		go studente(i, MAGISTRALE)
		//time.Sleep(1e8)
	}
	for i := N_STUD_TRI + N_STUD_MAG; i < N_STUD_TRI+N_STUD_MAG+N_STUD_TRI_L; i++ {
		go studente(i, TRIENNALE_LAUREANDO)
		//time.Sleep(1e8)
	}
	for i := N_STUD_TRI + N_STUD_MAG + N_STUD_TRI_L; i < N_STUD_TOT; i++ {
		go studente(i, MAGISTRALE_LAUREANDO)
		//time.Sleep(1e8)
	}
	go portiniere()
	//	time.Sleep(3e9)
	go biblioteca()

	//GoRoutine join
	for i := 0; i < N_STUD_TOT; i++ {
		<-done
	}

	fmt.Println("terminate")
	terminateServer <- true
	<-done
	terminateServer <- true
	<-done
	fmt.Println("FINITO")
}
