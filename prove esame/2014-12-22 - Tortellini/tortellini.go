package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ==================================================================
// COSTANTI =========================================================
const MAXBUFF = 100
const MAXPROC = 50
const MAX = 25 // capacità
const N = 20   // capacità frigorifero (< MAX)

// stato prenotazione
const PRENOTAZIONE_NON_VALIDA = -1
const RITIRO_NON_VALIDO = -1

// valore ticket
const NON_ASSEGNATO = 0
const ASSEGNATO = 1
const USATO = 2

// ==================================================================
// STRUTTURE DATI ===================================================
type Richiesta struct {
	id  int
	n   int
	ack chan int
}

// ==================================================================
// CANALI ===========================================================
var prenotazione = make(chan Richiesta, MAXBUFF)
var ritiro = make(chan Richiesta, MAXBUFF)
var deposito = make(chan Richiesta, MAXBUFF)

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
func whenBool(b bool, c chan bool) chan bool {
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

// ==================================================================
// GOROUTINE ========================================================
func cliente(id int) {
	n_conf := rand.Intn(MAX * 2)
	fmt.Printf("[CLIENTE %d] Voglio acquistare %d confezioni!\n", id, n_conf)

	r := Richiesta{id: id, n: n_conf, ack: make(chan int)}

	// PRENOTAZIONE
	// 1. Prenota num di confezioni (se > MAX => prenotazione rifiutata)
	prenotazione <- r
	// 2. Ottiene ticket UNIVOCO
	ticket := <-r.ack

	// Prenotazione non valida
	if ticket == PRENOTAZIONE_NON_VALIDA {
		fmt.Printf("[CLIENTE %d] La mia prenotazione è stata rifiutata\n", id)
	} else { // Prenotazione andata a buon fine
		fmt.Printf("[CLIENTE %d] Ho ottenuto il ticket #%d\n", id, ticket)
		r.n = ticket

		sleepRandTime(5)
		// RITIRO CONFEZIONI
		ottenute := 0
		for ottenute < n_conf { // Finché non ottiene tutte le confezioni non termina
			// 3. Invia il ticket
			ritiro <- r

			// 4. Attende le confezioni (potrebbero non essere tutte)
			res := <-r.ack

			if res < 0 { // Ritiro non andato a buon fine
				fmt.Printf("[CLIENTE %d] Il mio ticket è stato rifiutato\n", id)
				break
			} else {
				fmt.Printf("[CLIENTE %d] Ho ottenuto %d confezioni. Prenotazione: %d/%d\n", id, res, ottenute, n_conf)
			}
			ottenute += res
			sleepRandTime(5)
		}
	}
	done <- true
	fmt.Printf("[CLIENTE %d] Termino\n", id)
}

func sfoglina(id int) {
	r := Richiesta{id: id, n: 1, ack: make(chan int)}

	fmt.Printf("[SFOGLINA %d] Inizio\n", id)
	for {
		// Produce una confezione
		sleepRandTime(3)
		fmt.Printf("[SFOGLINA %d] Ho prodotto una confezione\n", id)

		// Deposita la confezione nel frigo
		deposito <- r

		// Attende l'esito
		res := <-r.ack
		if res < 0 {
			// terminazione
			done <- true
			fmt.Printf("[SFOGLINA %d] Termino\n", id)
			return
		}
	}
}

// gestore
func laboratorio() {
	frigo := 0                           // posti occupati nel frigo
	ticket := 0                          // contatore ticket corrente
	ticket_usati := make([]int, MAXPROC) // ticket usati
	prenotati := make([]int, MAXPROC)    // numero di confezioni prenotate da ciascun cliente
	fine := false

	fmt.Printf("[LABORATORIO] Inizio\n")
	for {
		select {
		case r := <-when(true, prenotazione):
			if r.n <= MAX { // Prenotazione OK
				ticket_usati[ticket] = ASSEGNATO // Segno che il ticket è stato ASSEGNATO
				prenotati[ticket] = r.n          // Segno QUANTE confezioni il cliente ha prenotato
				r.ack <- ticket                  // Invio al cliente il ticket
				fmt.Printf("[LABORATORIO] Assegno il ticket #%d al cliente %d\n", ticket, r.id)
				ticket++
			} else { // Il cliente ha richiesto troppe confezioni
				r.ack <- PRENOTAZIONE_NON_VALIDA
				fmt.Printf("[LABORATORIO] Prenotazione del cliente %d non valida! (%d > MAX)\n", r.id, r.n)
			}
		case r := <-when(len(prenotazione) == 0, ritiro):
			switch {
			// Ticket assegnato e prenotazione ancora da soddisfare
			case ticket_usati[r.n] == ASSEGNATO:
				if prenotati[r.n] > frigo { // Nel frigo non ci sono abbastanza confezioni per soddisfare la richiesta del cliente => invio quelle attualmente disponibili
					fmt.Printf("[LABORATORIO] Invio al cliente %d (ticket #%d) una parte delle confezioni prenotate (%d/%d)\n", r.id, r.n, frigo, prenotati[r.n])
					prenotati[r.n] -= frigo
					r.ack <- frigo
					frigo = 0
				} else {
					fmt.Printf("[LABORATORIO] Invio al cliente %d (ticket #%d) tutte le confezioni della sua prenotazione (%d)\n", r.id, r.n, prenotati[r.n])
					frigo -= prenotati[r.n]   // Sottraggo dal frigo il numero delle confezioni prenotate
					ticket_usati[r.n] = USATO // Segno che il ticket è stato USATO
					r.ack <- prenotati[r.n]   // Invio al cliente le confezioni
					prenotati[r.n] = 0        // Segno che la prenotazione è stata soddisfatta
				}
			// Ticket non ancora assegnato (il cliente vuole fare il furbetto)
			case ticket_usati[r.n] == NON_ASSEGNATO:
				r.ack <- RITIRO_NON_VALIDO
				fmt.Printf("[LABORATORIO] Ritiro non valido! Il ticket #%d non è ancora stato assegnato!!!\n", r.n)
			// Ticket già usato (il cliente vuole fare il furbetto)
			case ticket_usati[r.n] == USATO:
				r.ack <- RITIRO_NON_VALIDO
				fmt.Printf("[LABORATORIO] Ritiro non valido! Il ticket #%d è già stato usato!!!\n", r.n)
			// Altro
			default:
				r.ack <- RITIRO_NON_VALIDO
				fmt.Printf("[LABORATORIO] Ritiro non valido!!!\n")
			}
		case r := <-when(!fine && frigo < N && (len(prenotazione) == 0 && len(ritiro) == 0), deposito):
			frigo++
			r.ack <- 1
			fmt.Printf("[LABORATORIO] Confezione depositata correttamente. Stato deposito: %d/%d\n", frigo, N)
		case r := <-when(fine, deposito):
			r.ack <- -1
		case <-whenBool(!fine, termina):
			fine = true
			fmt.Printf("\n[LABORATORIO] Avviso le sfogline\n")
		case <-whenBool(fine, termina):
			done <- true
			fmt.Printf("\n[LABORATORIO] Termino\n")
			return
		}
	}
}

// ==================================================================
// MAIN =============================================================
func main() {
	fmt.Printf("[MAIN] Inizio\n\n")
	rand.Seed(time.Now().UnixNano())

	nClienti := 10
	nSfogline := 5

	// Esecuzione goroutine
	go laboratorio()

	for i := 0; i < nClienti; i++ {
		go cliente(i)
	}
	for i := 0; i < nSfogline; i++ {
		go sfoglina(i)
	}

	// Join goroutine
	for i := 0; i < nClienti; i++ {
		<-done
	}

	termina <- true
	for i := 0; i < nSfogline; i++ {
		<-done
	}
	termina <- true
	<-done
	fmt.Printf("\n\n[MAIN] Fine\n")
}
