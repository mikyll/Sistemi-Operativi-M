// compito 22 dicembre 2014

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAXCLI = 100
const MAXSFO = 50
const MAXPREN = 60
const MAXFRIGO = 20
const MAXBUFF = 100

type ReqR struct { // struttura messaggio per il ritiro
	id     int
	ticket int
}

var done = make(chan bool)
var termina = make(chan bool)
var prenota = make(chan int, MAXBUFF) // necessità di accodamento per priorità
var ritira = make(chan ReqR, MAXBUFF)
var deposita = make(chan int, MAXBUFF)

var ACK_PRENOTA [MAXCLI]chan int   //risposte prenotazioni: ticket oppure -1
var ACK_RITIRO [MAXCLI]chan bool   //risposte ritiro: false se ticket non valido
var ACK_DEPOSITO [MAXSFO]chan bool // risposte deposito: false se le sfogline devono finire
var r int

func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}
func whenR(b bool, c chan ReqR) chan ReqR { // guardia logica per il ritiro(in canale è di tipo ReqR)
	if !b {
		return nil
	}
	return c
}
func cliente(myid int) {
	var risP int
	var answ bool
	var tt int
	var msgRitiro ReqR

	tt = rand.Intn(5) + 1

	fmt.Printf("inizializzazione cliente  %d ! \n", myid)

	time.Sleep(time.Duration(tt) * time.Second)
	prenota <- myid            // send asincrona
	risP = <-ACK_PRENOTA[myid] // attesa x sincronizzazione
	fmt.Printf("[cliente %d]  richiesta prenotazione - ottenuto ticket %d\n ", myid, risP)
	if risP == -1 { // terminazione
		done <- true
		return
	}

	tt = rand.Intn(5)
	time.Sleep(time.Duration(tt) * time.Second) // attesa prima del ritiro
	msgRitiro = ReqR{myid, risP}
	ritira <- msgRitiro
	answ = <-ACK_RITIRO[myid]
	if answ {
		fmt.Printf("[cliente %d]  ritirati tortellini!\n ", myid)
	} else {
		fmt.Printf("[cliente %d]  ritiro negato..\n ", myid)
	}
	done <- true
}

func sfoglina(myid int) {

	var tt int
	var esito bool

	for {
		tt = rand.Intn(5) + 1
		time.Sleep(time.Duration(tt) * time.Second)
		deposita <- myid // send asincrona
		esito = <-ACK_DEPOSITO[myid]
		if esito == true {
			fmt.Printf("[sfoglina %d]  ho depositato una nuova confezione.\n ", myid)
		} else {
			fmt.Printf("[sfoglina %d]  termino\n ", myid)
			done <- true
			return
		}
	}

}

func server(cli int, sfo int) {
	var contatore int // valore ticket corrente
	var daritirare int
	var fine bool
	var ticket [MAXCLI]int //ticket assegnati
	var ris int            //messaggio di risposta prenotazione
	var rr bool            //messaggio di risposta ritiro
	var infrigo int
	var i int
	var pre int
	contatore = 1
	fine = false
	infrigo = 0
	pre = -1
	fmt.Printf("\n*** TEST con %d CLIENTI e %d SFOGLINE ***\n ", cli, sfo)
	for {

		select {
		case x := <-when(((infrigo < MAXFRIGO) && (fine == false)), deposita): //deposito accettato
			infrigo++
			fmt.Printf("[server]  depositata nuova confezione da sfoglina %d !", x)
			fmt.Printf("(infrigo=%d, daritirare=%d)!  \n", infrigo, daritirare)
			ACK_DEPOSITO[x] <- true

		case x := <-when((fine == true), deposita): //deposito negato
			fmt.Printf("[server]  termino la sfoglina %d !", x)
			fmt.Printf("(infrigo=%d, daritirare=%d)!  \n", infrigo, daritirare)
			ACK_DEPOSITO[x] <- false
		case x := <-prenota: //prenotazione
			if contatore < MAXPREN { // accetta prenotazione
				ticket[x] = contatore
				ris = contatore
				contatore++
				daritirare++
				fmt.Printf("[server] prenotata confezione per cliente %d: assegnato ticket %d ", x, ris)
				fmt.Printf("(infrigo=%d, daritirare=%d)!  \n", infrigo, daritirare)
			} else {
				ris = -1
				fmt.Printf("[server] rifiutata prenotazione per cliente %d: assegnato ticket %d  \n", x, ris)
			}
			ACK_PRENOTA[x] <- ris // termine "call"
		case x := <-whenR((infrigo > 0) && (len(prenota) == 0), ritira): //ritiro
			for i = 0; i < cli; i++ {
				if x.ticket == ticket[i] { // il ticket è valido
					pre = i
				}
			}
			if pre >= 0 { // ritiro accettato
				infrigo--
				daritirare--
				ticket[pre] = 0 //ticket consumato
				rr = true
				fmt.Printf("[server]  cliente %d ha ritirato!  ", i)
				fmt.Printf("(infrigo=%d, daritirare=%d)!  \n", infrigo, daritirare)
				if daritirare == 0 {
					fine = true
					fmt.Printf("[server]  completati ritiri!!!! ")
					fmt.Printf("(infrigo=%d, daritirare=%d)!  \n", infrigo, daritirare)
				}
			} else {
				rr = false
				fmt.Printf("[server]  cliente %d - RITIRO NEGATO!  \n", i)
			}
			ACK_RITIRO[x.id] <- rr // termine "call"
		case <-termina: // quando tutti i processi hanno finito
			fmt.Printf("[server] FINE (sono rimaste %d confezioni in frigo..)!!!!!!\n", infrigo)
			done <- true
			return
		}

	}
}

func main() {
	var CLI int
	var SFO int

	fmt.Printf("\n quanti clienti (max %d)? ", MAXCLI)
	fmt.Scanf("%d", &CLI)
	fmt.Printf("\n quante sfogline (max %d)? ", MAXSFO)
	fmt.Scanf("%d", &SFO)

	//inizializzazione canali
	for i := 0; i < CLI; i++ {
		ACK_PRENOTA[i] = make(chan int, MAXBUFF)
		ACK_RITIRO[i] = make(chan bool, MAXBUFF)
	}
	for i := 0; i < SFO; i++ {
		ACK_DEPOSITO[i] = make(chan bool, MAXBUFF)
	}

	rand.Seed(time.Now().Unix())
	go server(CLI, SFO)

	for i := 0; i < CLI; i++ {
		go cliente(i)
	}
	for i := 0; i < SFO; i++ {
		go sfoglina(i)
	}

	for i := 0; i < CLI+SFO; i++ {
		<-done
	}
	termina <- true
	<-done
	fmt.Printf("\n HO FINITO.\n ")
}
