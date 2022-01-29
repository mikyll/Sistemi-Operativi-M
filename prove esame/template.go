// Nome Cognome - Matricola
// ==================================================================
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

// ==================================================================
// STRUTTURE DATI ===================================================
type Richiesta struct {
	id  int
	ack chan int
}

// ==================================================================
// CANALI ===========================================================
var entrata = make(chan int, MAXBUFF)
var uscita = make(chan int)
var ACK [MAXPROC]chan int // alternativa alla richiesta

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
func whenInt(b bool, c chan int) chan int {
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

// Sleep per un tot di secondi randomico nell'intervallo [x, y)
func sleepRandTimeRange(x, y int) {
	if x >= 0 && y > 0 && x < y {
		time.Sleep(time.Duration(rand.Intn(y)+x) * time.Second)
	}
}

// Rimuove un un elemento da uno slice
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

// Restituisce uno slice di n float64 randomici
func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

// ==================================================================
// GOROUTINE ========================================================
func client(id int) {
	fmt.Printf("[CLIENTE %d] Inizio\n", id)
	r := Richiesta{id, make(chan int)} // se si usa struttura

	// do things

	fmt.Printf("[CLIENTE %d] Termino\n", id)
	done <- true
}

// gestore risorse
func server() {
	// vari counter

	fmt.Printf("[SERVER] Inizio\n")
	for {
		select {
		case r := <-when(true, make(chan Richiesta)): // condizione 1 da verificare, canale da cui leggere
			r.ack <- 1
		case r := <-when(true, make(chan Richiesta)): // condizione 2 da verificare, canale da cui leggere
			r.ack <- 1
		case x := <-uscita: // nessuna condizione da verificare prima di leggere un messaggio da uscita
			if x == 0 {
				// do smth
			} else {
				// do smth else
			}
		case <-termina:
			fmt.Println("[SERVER] Termino")
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

	var nClient int

	fmt.Printf("\nQuanti Client (max %d)? ", MAXPROC)
	fmt.Scanf("%d\n", &nClient)

	// Inizializzazione canali
	for i := 0; i < len(ACK); i++ {
		ACK[i] = make(chan int, MAXBUFF)
	}

	// Esecuzione goroutine
	go server()

	for i := 0; i < nClient; i++ {
		go client(i)
	}

	// Join goroutine
	for i := 0; i < nClient; i++ {
		<-done
	}

	termina <- true
	<-done
	fmt.Printf("\n\n[MAIN] Fine\n")
}
