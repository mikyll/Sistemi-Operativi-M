// Svolto da: Michele Righi
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ==================================================================
// COSTANTI =========================================================
const MAXBUFF = 100
const MAXSTUD = 100

const TRIENNALE_NON_LAUREANDO = 0
const TRIENNALE_LAUREANDO = 1
const MAGISTRALE_NON_LAUREANDO = 2
const MAGISTRALE_LAUREANDO = 3

const N = 5

// ==================================================================
// STRUTTURE DATI ===================================================
type Richiesta struct {
	id  int
	ack chan int
}

// ==================================================================
// CANALI ===========================================================
var entrataPortineria = make(chan Richiesta)
var entrataBiblioteca [4]chan Richiesta
var uscitaBiblioteca = make(chan int)
var uscitaPortineria = make(chan Richiesta) // forse richiesta?

// ==================================================================
// CANALI DI JOIN ===================================================
var done = make(chan bool)
var termina = make(chan bool)

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

func getTipoStudente(t int) string {
	if t == TRIENNALE_NON_LAUREANDO {
		return "triennale non laureando"
	} else if t == TRIENNALE_LAUREANDO {
		return "triennale laureando"
	} else if t == MAGISTRALE_NON_LAUREANDO {
		return "magistrale non laureando"
	} else if t == MAGISTRALE_LAUREANDO {
		return "magistrale laureando"
	} else {
		return ""
	}
}

// ==================================================================
// GOROUTINE ========================================================
func studente(id int) {
	tipo := rand.Intn(4) // tipo di studente
	r := Richiesta{id: id, ack: make(chan int)}

	fmt.Printf("[STUDENTE %d - %s] Inizio\n", id, getTipoStudente(tipo))

	// 1. Consegno doc in portineria
	fmt.Printf("[STUDENTE %d - %s] Consegno il documento in portineria\n", id, getTipoStudente(tipo))
	entrataPortineria <- r
	<-r.ack

	// 2. Entro in biblioteca
	entrataBiblioteca[tipo] <- r
	<-r.ack

	// 3. Studio
	sleepRandTime(5)

	// 4. Esco dalla biblioteca
	uscitaBiblioteca <- tipo

	// 5. Ritiro il documento
	uscitaPortineria <- r
	<-r.ack

	fmt.Printf("[STUDENTE %d - %s] Termino\n", id, getTipoStudente(tipo))
	done <- true
}
func portineria() {
	var documenti = make([]bool, MAXSTUD) // inizializzati allo zero value (false per i boolean)
	fmt.Println(documenti)

	fmt.Printf("[PORTINERIA] Inizio\n")
	for {
		sleepRandTime(2)
		select {
		case r := <-entrataPortineria:
			// prendo il documento
			documenti[r.id] = true
			fmt.Printf("[PORTINERIA] Lo studente %d ha consegnato il documento\n", r.id)
			r.ack <- 1
		case r := <-uscitaPortineria:
			// restituisco il documento
			if documenti[r.id] {
				documenti[r.id] = false
				fmt.Printf("[PORTINERIA] Documento %d ritirato con successo\n", r.id)
				r.ack <- 1
			} else {
				fmt.Printf("[PORTINERIA] Il documento %d non è presente!\n", r.id)
				r.ack <- -1
			}
		case <-termina:
			fmt.Printf("[PORTINERIA] Termino\n")
			done <- true
			return
		default:
			continue
		}
	}
}
func biblioteca() {
	nt := 0
	nm := 0

	fmt.Printf("[BIBLIOTECA] Inizio\n")
	for {
		select {
		case r := <-when(nt+nm < N &&
			((nt < nm &&
				len(entrataBiblioteca[TRIENNALE_LAUREANDO]) == 0) ||
				(len(entrataBiblioteca[MAGISTRALE_NON_LAUREANDO]) == 0 && len(entrataBiblioteca[MAGISTRALE_LAUREANDO]) == 0)), entrataBiblioteca[TRIENNALE_NON_LAUREANDO]):
			nt++
			fmt.Printf("[BIBLIOTECA] Lo studente TRIENNALE NON LAUREANDO %d è entrato dalla biblioteca. Stato: %d/%d (T: %d, M: %d)\n", r.id, nt+nm, N, nt, nm)
			r.ack <- 1
		case r := <-when(nt+nm < N &&
			(nt < nm &&
				(len(entrataBiblioteca[MAGISTRALE_NON_LAUREANDO]) == 0 && len(entrataBiblioteca[MAGISTRALE_LAUREANDO]) == 0)), entrataBiblioteca[TRIENNALE_LAUREANDO]):
			nt++
			fmt.Printf("[BIBLIOTECA] Lo studente TRIENNALE LAUREANDO %d è entrato dalla biblioteca. Stato: %d/%d (T: %d, M: %d)\n", r.id, nt+nm, N, nt, nm)
			r.ack <- 1
		case r := <-when(nt+nm < N &&
			((nm <= nt &&
				len(entrataBiblioteca[MAGISTRALE_LAUREANDO]) == 0) ||
				(len(entrataBiblioteca[TRIENNALE_NON_LAUREANDO]) == 0 && len(entrataBiblioteca[TRIENNALE_LAUREANDO]) == 0)), entrataBiblioteca[MAGISTRALE_NON_LAUREANDO]):
			nm++
			fmt.Printf("[BIBLIOTECA] Lo studente MAGISTRALE NON LAUREANDO %d è entrato dalla biblioteca. Stato: %d/%d (T: %d, M: %d)\n", r.id, nt+nm, N, nt, nm)
			r.ack <- 1
		case r := <-when(nt+nm < N &&
			(nm <= nt ||
				(len(entrataBiblioteca[TRIENNALE_NON_LAUREANDO]) == 0 && len(entrataBiblioteca[TRIENNALE_LAUREANDO]) == 0)), entrataBiblioteca[MAGISTRALE_LAUREANDO]):
			nm++
			fmt.Printf("[BIBLIOTECA] Lo studente MAGISTRALE LAUREANDO %d è entrato dalla biblioteca. Stato: %d/%d (T: %d, M: %d)\n", r.id, nt+nm, N, nt, nm)
			r.ack <- 1
		case x := <-uscitaBiblioteca:
			switch {
			case x == TRIENNALE_NON_LAUREANDO || x == TRIENNALE_LAUREANDO:
				nt--
				fmt.Printf("[BIBLIOTECA] È uscito uno studente TRIENNALE. Stato: %d/%d (T: %d, M: %d)\n", nt+nm, N, nt, nm)
			case x == MAGISTRALE_NON_LAUREANDO || x == MAGISTRALE_LAUREANDO:
				nm--
				fmt.Printf("[BIBLIOTECA] È uscito uno studente MAGISTRALE. Stato: %d/%d (T: %d, M: %d)\n", nt+nm, N, nt, nm)
			}
		case <-termina:
			fmt.Printf("[BIBLIOTECA] Termino\n")
			done <- true
			return
		}
	}
}

// ==================================================================
// MAIN =============================================================
func main() {
	fmt.Printf("[MAIN] Inizio\n\n")
	rand.Seed(time.Now().UnixNano())

	var nStudenti int

	/*fmt.Printf("\nQuanti studenti (max %d)? ", MAXSTUD)
	fmt.Scanf("%d\n", &nStudenti)*/

	nStudenti = 10

	// Inizializzazione canali
	for i := 0; i < len(entrataBiblioteca); i++ {
		entrataBiblioteca[i] = make(chan Richiesta, MAXBUFF)
	}

	// Esecuzione goroutine
	for i := 0; i < nStudenti; i++ {
		go studente(i)
	}
	go portineria()
	go biblioteca()

	// Join goroutine
	for i := 0; i < nStudenti; i++ {
		<-done
	}

	// Biblioteca & Portineria
	termina <- true
	termina <- true
	<-done
	<-done
	fmt.Printf("\n\n[MAIN] Fine\n")
}
