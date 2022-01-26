package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Costanti generali
const NUFF = 5  // numero uffici (e consulenti)
const MAXS = 10 //capacità sala attesa
const N_UTENTI = 100

const MAXBUF = 50

// Costanti relative alla priorità nella gestione della sala d'attesa
const TIPI_UT = 3
const amm = 0
const priv_solo = 1
const priv_accomp = 2

// Costanti relative alla priorità nella gestione degli uffici
const TIPI_FIN = 2
const Superbonus = 0
const Altro = 1

// Tipi
type descrUtente struct {
	id    int
	tipoU int // tipo dell'utente richiedente(amministratore, propr, ecc)
	tipoF int // tipo del finanziamento
	reply chan int
}

// Canali di comunicazione generali
var termina chan bool
var done chan bool

// Canali :
var entraSala [TIPI_UT]chan descrUtente
var entraUfficio [TIPI_FIN]chan descrUtente
var esceUfficio chan int

func sleepRandom() {
	time.Sleep(time.Duration(1e9 * ((rand.Intn(30)) + 1)))
}

func when(condition bool, ch chan descrUtente) chan descrUtente {
	if !condition {
		return nil
	}
	return ch
}

func server() {
	var i int
	personeInSalaAttesa := 0
	ufficiOccupati := 0
	var ufficioOccupato [NUFF]bool
	for i = 0; i < NUFF; i++ {
		ufficioOccupato[i] = false
	}

	fmt.Printf("Il servizio di consulenza è aperto.\n\n")
	for {
		select {
		case richiesta := <-when(personeInSalaAttesa < MAXS, entraSala[amm]):
			personeInSalaAttesa += 1
			fmt.Printf("SERVER: entrato amministratore %d.\n", richiesta.id)
			richiesta.reply <- 1
		case richiesta := <-when(personeInSalaAttesa < MAXS && len(entraSala[amm]) == 0, entraSala[priv_solo]):
			personeInSalaAttesa += 1
			fmt.Printf("SERVER:entrato privato da solo %d.\n", richiesta.id)
			richiesta.reply <- 1
		case richiesta := <-when(personeInSalaAttesa+2 <= MAXS && len(entraSala[amm]) == 0 && len(entraSala[priv_solo]) == 0, entraSala[priv_accomp]):
			personeInSalaAttesa += 2
			fmt.Printf("SERVER: entrato privato con accompagnatore %d.\n", richiesta.id)
			richiesta.reply <- 1
		case richiesta := <-when(ufficiOccupati < NUFF, entraUfficio[Superbonus]):
			for i = 0; i < NUFF; i++ { //individuazione ufficio da occupare
				if !ufficioOccupato[i] {
					break
				}
			}
			ufficioOccupato[i] = true
			ufficiOccupati++
			if richiesta.tipoU == priv_accomp {
				personeInSalaAttesa -= 2 //libero 2 posti in sala d'attesa
				fmt.Printf("SERVER: il privato con accompagnatore per superbonus   %d è entrato nell'ufficio %d.\n", richiesta.id, i)
			} else {
				personeInSalaAttesa -= 1 //libero 1 posto in sala d'attesa
				if richiesta.tipoU == amm {
					fmt.Printf("SERVER: l'amministratore per superbonus  %d è entrato nell'ufficio %d.\n", richiesta.id, i)
				} else {
					fmt.Printf("SERVER: privato singolo per superbonus  %d è entrato nell'ufficio %d.\n", richiesta.id, i)
				}
			}
			richiesta.reply <- i
		case richiesta := <-when(ufficiOccupati < NUFF && len(entraUfficio[Superbonus]) == 0, entraUfficio[Altro]):
			for i = 0; i < NUFF; i++ { //individuazione ufficio da occupare
				if !ufficioOccupato[i] {
					break
				}
			}
			ufficioOccupato[i] = true
			ufficiOccupati++
			if richiesta.tipoU == priv_accomp {
				personeInSalaAttesa -= 2 //libero 2 posti in sala d'attesa
				fmt.Printf("SERVER: il privato con accompagnatore per Altro   %d è entrato nell'ufficio %d.\n", richiesta.id, i)
			} else {
				personeInSalaAttesa -= 1 //libero 1 posto in sala d'attesa
				if richiesta.tipoU == amm {
					fmt.Printf("SERVER: l'amministratore per Altro  %d è entrato nell'ufficio %d.\n", richiesta.id, i)
				} else {
					fmt.Printf("SERVER: privato singolo per Altro  %d è entrato nell'ufficio %d.\n", richiesta.id, i)
				}
			}
			richiesta.reply <- i
		case rilascio := <-esceUfficio:
			ufficioOccupato[rilascio] = false
			ufficiOccupati--

		case <-termina:
			fmt.Printf("Il servizio di consulenza chiude.\n")
			done <- true
			return
		}
	}
}

func utente(id int) {

	tipo_ut := rand.Intn(TIPI_UT)   // amministratore, singolo o accompagnato
	tipo_fin := rand.Intn(TIPI_FIN) //tipo di finanziamento
	var ack = make(chan int)
	var richiesta descrUtente

	richiesta.id = id
	richiesta.tipoU = tipo_ut
	richiesta.tipoF = tipo_fin
	richiesta.reply = ack

	sleepRandom()

	entraSala[tipo_ut] <- richiesta
	<-richiesta.reply

	entraUfficio[tipo_fin] <- richiesta
	ottenuto := <-richiesta.reply

	sleepRandom()

	esceUfficio <- ottenuto

	fmt.Printf("Utente [%d]: sono uscito dall'ufficio %d. Termino.\n", id, ottenuto)
	done <- true
	return
}

func main() {

	rand.Seed(time.Now().UnixNano())
	// Generali
	termina = make(chan bool)
	done = make(chan bool)

	// inizializzazione canali
	for i := 0; i < 3; i++ {
		entraSala[i] = make(chan descrUtente, MAXBUF)
	}

	// Ambulatorio
	for i := 0; i < 2; i++ {
		entraUfficio[i] = make(chan descrUtente, MAXBUF)
	}
	esceUfficio = make(chan int, MAXBUF)

	go server()

	for id := 0; id < N_UTENTI; id++ {
		go utente(id)
	}
	for id := 0; id < N_UTENTI; id++ {
		<-done
	}
	termina <- true
	<-done
}
