// Svolto da: Michele Righi
/*
NB: nella traccia c'è scritto di controllare solo che il numero di
clienti del museo non superi una certa soglia, quindi ho pensato
intendesse solo i visitatori. In caso contrario, bisognerebbe
aggiungere un controllo all'ingresso dei visitatori e degli operatori
per evitare che la dimensione massima venga sforata.
*/
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ==================================================================
// COSTANTI =========================================================
const MAXBUFF = 100
const MAXV = 50
const MINV = 5
const MAXO = 10
const MINO = 2

const Nm = 10
const Ns = 10
const MUSEO = 0
const MOSTRE = 1
const ENTRAMBI_MUSEO = 2
const ENTRAMBI_MOSTRE = 3
const USCITA_OP_MUSEO = 2
const USCITA_OP_MOSTRE = 3

// ==================================================================
// STRUTTURE DATI ===================================================
type Richiesta struct {
	id   int
	sala int
	ack  chan int
}

// ==================================================================
// CANALI ===========================================================
var entrataVisitatori [4]chan Richiesta
var entrataOperatori = make(chan Richiesta, MAXBUFF)
var uscitaVisitatori = make(chan int)
var uscitaOperatori [2]chan Richiesta

// ==================================================================
// CANALI DI JOIN ===================================================
var done = make(chan bool)
var termina = make(chan bool)
var terminaComplesso = make(chan bool)

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

// Sleep per un tot di secondi randomico nell'intervallo [x, y)
func sleepRandTimeRange(x, y int) {
	if x >= 0 && y > 0 && x < y {
		time.Sleep(time.Duration(rand.Intn(y)+x) * time.Second)
	}
}

func getTipo(t int) string {
	switch t {
	case MUSEO:
		return "museo"
	case MOSTRE:
		return "mostre"
	case ENTRAMBI_MUSEO:
		return "entrambi"
	case ENTRAMBI_MOSTRE:
		return "entrambi"
	default:
		return ""
	}
}

func debug(nvm, nvs, nom, nos, attvm, attvs, attevm, attevs int) {
	fmt.Printf("[[[DEBUG]]] NV[MUSEO]: %d, NV[MOSTRA]: %d, NO[MUSEO]: %d, NO[MOSTRA]: %d\nIn attesa VM: %d, VS: %d, VEM: %d, VES: %d\n\n", nvm, nvs, nom, nos, attvm, attvs, attevm, attevs)
}

// ==================================================================
// GOROUTINE ========================================================
func visitatore(id int) {
	tipo := rand.Intn(4) // tipo visitatore (museo, sala mostre, entrambi-prima-museo-poi-mostre, entrambi-prima-mostre-poi-museo)
	r := Richiesta{id: id, sala: tipo % 2, ack: make(chan int)}

	fmt.Printf("[VISITATORE %s %d] Inizio\n", strings.ToUpper(getTipo(tipo)), id)

	// NB: in r.sala salvo la sala che vuole visitare quando manda la richiesta di ingresso

	sleepRandTimeRange(id, id+1)

	// 1. Entra Visitatore
	fmt.Printf("[VISITATORE %s %d] Entro in %s\n", strings.ToUpper(getTipo(tipo)), id, getTipo(r.sala))
	entrataVisitatori[tipo] <- r
	<-r.ack

	// 2. Visita
	fmt.Printf("[VISITATORE %s %d] Visito %s...\n", strings.ToUpper(getTipo(tipo)), id, getTipo(r.sala))
	sleepRandTimeRange(2, 5)

	// 3. Uscita Visitatore
	fmt.Printf("[VISITATORE %s %d] Esco da %s\n", strings.ToUpper(getTipo(tipo)), id, getTipo(r.sala))
	uscitaVisitatori <- r.sala

	if tipo > 1 { // Se visitatore ENTRAMBI
		if tipo == ENTRAMBI_MUSEO {
			r.sala = MOSTRE
			tipo = ENTRAMBI_MOSTRE
		} else if tipo == ENTRAMBI_MOSTRE {
			r.sala = MUSEO
			tipo = ENTRAMBI_MUSEO
		}

		// 1. Entra Visitatore (seconda sala)
		fmt.Printf("[VISITATORE %s %d] Entro in %s\n", strings.ToUpper(getTipo(tipo)), id, getTipo(r.sala))
		entrataVisitatori[tipo] <- r
		<-r.ack

		// 2. Visita
		fmt.Printf("[VISITATORE %s %d] Visito %s...\n", strings.ToUpper(getTipo(tipo)), id, getTipo(r.sala))
		sleepRandTimeRange(2, 5)

		// 3. Uscita Visitatore
		fmt.Printf("[VISITATORE %s %d] Esco da %s\n", strings.ToUpper(getTipo(tipo)), id, getTipo(r.sala))
		uscitaVisitatori <- r.sala
	}
	fmt.Printf("[VISITATORE %s %d] Termino\n", strings.ToUpper(getTipo(tipo)), id)
	done <- true
}

func operatore(id int) {
	sala := id
	if id > 1 { // gli assegno una sala in modo randomico (solo se id > 1 così ho almeno un operatore per ciascuna sala)
		sala = rand.Intn(id) % 2
	}

	r := Richiesta{id: id, sala: sala, ack: make(chan int)}

	fmt.Printf("[OPERATORE %s %d] Inizio\n", strings.ToUpper(getTipo(sala)), id)
	for {
		// 1. Entra
		fmt.Printf("[OPERATORE %s %d] Entro\n", strings.ToUpper(getTipo(sala)), id)
		entrataOperatori <- r
		<-r.ack

		// 2. Esegue attività operatore
		sleepRandTimeRange(5, 10)

		// 3. Esce
		fmt.Printf("[OPERATORE %s %d] Voglio andare in pausa\n", strings.ToUpper(getTipo(sala)), id)
		uscitaOperatori[sala] <- r
		<-r.ack

		// 4. Pausa
		fmt.Printf("[OPERATORE %s %d] Sono in pausa...\n", strings.ToUpper(getTipo(sala)), id)
		sleepRandTimeRange(5, 10)

		select {
		case <-termina:
			fmt.Printf("[OPERATORE %s %d] Termino\n", strings.ToUpper(getTipo(sala)), id)
			done <- true
			return
		default:
			continue
		}
	}
}

func complesso() {
	nv := make([]int, 2) // Numero di Visitatori attualmente nel Museo ([MUSEO] < Nm, [MOSTRE] < Ns)
	no := make([]int, 2) // Numero di Operatori attualmente nel Museo ([MUSEO], [MOSTRE])

	fmt.Printf("[COMPLESSO] Inizio\n")
	for {
		select {
		// INGRESSO VISITATORE
		case r := <-when(nv[MUSEO] < Nm && no[MUSEO] > 0, entrataVisitatori[MUSEO]):
			nv[MUSEO]++
			fmt.Printf("[COMPLESSO] È entrato il visitatore %d nella sala MUSEO\n", r.id)
			//debug(nv[MUSEO], nv[MOSTRE], no[MUSEO], no[MOSTRE], len(entrataVisitatori[MUSEO]), len(entrataVisitatori[MOSTRE]), len(entrataVisitatori[ENTRAMBI_MUSEO]), len(entrataVisitatori[ENTRAMBI_MOSTRE]))
			r.ack <- 1
		case r := <-when(nv[MOSTRE] < Nm && no[MOSTRE] > 0 && len(entrataVisitatori[MUSEO]) == 0, entrataVisitatori[MOSTRE]):
			nv[MOSTRE]++
			fmt.Printf("[COMPLESSO] È entrato il visitatore %d nella sala MOSTRE\n", r.id)
			//debug(nv[MUSEO], nv[MOSTRE], no[MUSEO], no[MOSTRE], len(entrataVisitatori[MUSEO]), len(entrataVisitatori[MOSTRE]), len(entrataVisitatori[ENTRAMBI_MUSEO]), len(entrataVisitatori[ENTRAMBI_MOSTRE]))
			r.ack <- 1
		case r := <-when(nv[MUSEO] < Nm && no[MUSEO] > 0 && len(entrataVisitatori[MUSEO]) == 0 && len(entrataVisitatori[MOSTRE]) == 0, entrataVisitatori[ENTRAMBI_MUSEO]):
			nv[MUSEO]++
			fmt.Printf("[COMPLESSO] È entrato il visitatore %d, che visita entrambe le sale, nella sala MUSEO\n", r.id)
			//debug(nv[MUSEO], nv[MOSTRE], no[MUSEO], no[MOSTRE], len(entrataVisitatori[MUSEO]), len(entrataVisitatori[MOSTRE]), len(entrataVisitatori[ENTRAMBI_MUSEO]), len(entrataVisitatori[ENTRAMBI_MOSTRE]))
			r.ack <- 1
		case r := <-when(nv[MOSTRE] < Nm && no[MOSTRE] > 0 && len(entrataVisitatori[MUSEO]) == 0 && len(entrataVisitatori[MOSTRE]) == 0, entrataVisitatori[ENTRAMBI_MOSTRE]):
			nv[MOSTRE]++
			fmt.Printf("[COMPLESSO] È entrato il visitatore %d, che visita entrambe le sale, nella sala MOSTRE\n", r.id)
			//debug(nv[MUSEO], nv[MOSTRE], no[MUSEO], no[MOSTRE], len(entrataVisitatori[MUSEO]), len(entrataVisitatori[MOSTRE]), len(entrataVisitatori[ENTRAMBI_MUSEO]), len(entrataVisitatori[ENTRAMBI_MOSTRE]))
			r.ack <- 1

		// INGRESSO OPERATORE
		case r := <-entrataOperatori:
			no[r.sala]++
			fmt.Printf("[COMPLESSO] È entrato l'operatore %d nella sala %s\n", r.id, getTipo(r.sala))
			//debug(nv[MUSEO], nv[MOSTRE], no[MUSEO], no[MOSTRE], len(entrataVisitatori[MUSEO]), len(entrataVisitatori[MOSTRE]), len(entrataVisitatori[ENTRAMBI_MUSEO]), len(entrataVisitatori[ENTRAMBI_MOSTRE]))
			r.ack <- 1

		// USCITA VISITATORI
		case x := <-uscitaVisitatori:
			nv[x]--
			fmt.Printf("[COMPLESSO] È uscito un visitatore dalla sala %s\n", getTipo(x))
			//debug(nv[MUSEO], nv[MOSTRE], no[MUSEO], no[MOSTRE], len(entrataVisitatori[MUSEO]), len(entrataVisitatori[MOSTRE]), len(entrataVisitatori[ENTRAMBI_MUSEO]), len(entrataVisitatori[ENTRAMBI_MOSTRE]))

		// USCITA OPERATORI
		case r := <-when((nv[MUSEO] == 0 && len(entrataVisitatori[MUSEO]) == 0 && len(entrataVisitatori[ENTRAMBI_MUSEO]) == 0) || no[MUSEO] > 1, uscitaOperatori[MUSEO]):
			no[MUSEO]--
			fmt.Printf("[COMPLESSO] È uscito l'operatore %d dalla sala MUSEO\n", r.id)
			//debug(nv[MUSEO], nv[MOSTRE], no[MUSEO], no[MOSTRE], len(entrataVisitatori[MUSEO]), len(entrataVisitatori[MOSTRE]), len(entrataVisitatori[ENTRAMBI_MUSEO]), len(entrataVisitatori[ENTRAMBI_MOSTRE]))
			r.ack <- 1
		case r := <-when((nv[MOSTRE] == 0 && len(entrataVisitatori[MOSTRE]) == 0 && len(entrataVisitatori[ENTRAMBI_MOSTRE]) == 0) || no[MOSTRE] > 1, uscitaOperatori[MOSTRE]):
			no[MOSTRE]--
			fmt.Printf("[COMPLESSO] È uscito l'operatore %d dalla sala MOSTRE\n", r.id)
			//debug(nv[MUSEO], nv[MOSTRE], no[MUSEO], no[MOSTRE], len(entrataVisitatori[MUSEO]), len(entrataVisitatori[MOSTRE]), len(entrataVisitatori[ENTRAMBI_MUSEO]), len(entrataVisitatori[ENTRAMBI_MOSTRE]))
			r.ack <- 1

		// TERMINAZIONE
		case <-terminaComplesso:
			fmt.Printf("[COMPLESSO] Termino\n")
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

	nVisitatori := 10
	nOperatori := 5

	/*fmt.Printf("\nQuanti Visitatori (min %d, max %d)? ", MINV, MAXV)
	fmt.Scanf("%d\n", &nVisitatori)
	fmt.Printf("\nQuanti Operatori (min %d, max %d)? ", MINO, MAXO)
	fmt.Scanf("%d\n", &nOperatori)
	if nVisitatori < MINV || nVisitatori > MAXV {
		nVisitatori = 10
	}
	if nOperatori < MINO || nOperatori > MAXO {
		nOperatori = 2
	}*/

	// Inizializzazione canali
	for i := 0; i < len(entrataVisitatori); i++ {
		entrataVisitatori[i] = make(chan Richiesta, MAXBUFF)
	}
	for i := 0; i < len(uscitaOperatori); i++ {
		uscitaOperatori[i] = make(chan Richiesta, MAXBUFF)
	}

	// Esecuzione goroutine
	go complesso()

	for i := 0; i < nVisitatori; i++ {
		go visitatore(i)
	}
	for i := 0; i < nOperatori; i++ {
		go operatore(i)
	}

	// Join goroutine
	for i := 0; i < nVisitatori; i++ {
		<-done
	}

	fmt.Printf("[MAIN] Tutti i visitatori sono usciti\n\n")

	for i := 0; i < nOperatori; i++ {
		termina <- true
		<-done
	}

	terminaComplesso <- true
	<-done

	fmt.Printf("\n\n[MAIN] Fine\n")
}
