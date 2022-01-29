// Svolto da: Michele Righi
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

const PI = 0 // Ponte Inferiore
const PS = 1 // Ponte Superiore

const SALITA = 0
const DISCESA = 1

const SALITA_VIAGGIATORE = 0
const DISCESA_VIAGGIATORE = 1
const SALITA_EQUIPAGGIO = 2
const DISCESA_EQUIPAGGIO = 3

const TS = 20 // capacità massima del ponte Superiore
const TI = 20 // capacità massima del ponte Inferiore

// ==================================================================
// STRUTTURE DATI ===================================================
type Richiesta struct {
	id  int
	ack chan int
}

// ==================================================================
// CANALI ===========================================================
var entrataScala [4]chan Richiesta
var uscitaScala = make(chan int)

// ==================================================================
// CANALI DI JOIN ===================================================
var done = make(chan bool)
var termina = make(chan bool)
var terminaCabina = make(chan bool)

// ==================================================================
// FUNZIONI =========================================================
// Guardia logica
func when(b bool, c chan Richiesta) chan Richiesta {
	if !b {
		return nil
	}
	return c
}

// Sleep per un tot di secondi randomico nell'intervallo [x, y)
func sleepRandTimeRange(x, y int) {
	if x >= 0 && y > 0 && x < y {
		time.Sleep(time.Duration(rand.Intn(y)+x) * time.Second)
	}
}

func getTipo(t int) string {
	switch t {
	case SALITA_VIAGGIATORE:
		return "viaggiatore che saliva"
	case DISCESA_VIAGGIATORE:
		return "viaggiatore che scendeva"
	case SALITA_EQUIPAGGIO:
		return "membro dell'equipaggio che saliva"
	case DISCESA_EQUIPAGGIO:
		return "membro dell'equipaggio che scendeva"
	default:
		return ""
	}
}
func getTipoScale(t int) string {
	if t == SALITA {
		return "salita"
	} else if t == DISCESA {
		return "discesa"
	} else {
		return ""
	}
}
func getPonte(t int) string {
	if t == PI {
		return "PI"
	} else if t == PS {
		return "PS"
	} else {
		return ""
	}
}

// ==================================================================
// GOROUTINE ========================================================
func viaggiatore(id, stato_iniz int) {
	stato := stato_iniz
	r := Richiesta{id: id, ack: make(chan int)}

	fmt.Printf("[VIAGGIATORE %d %s] Inizio. Salgo sull'aereo\n", id, getPonte(stato))
	for {

		// 1. Vuole usare la scala
		fmt.Printf("[VIAGGIATORE %d %s] Voglio usare le scale in %s\n", id, getPonte(stato), strings.ToUpper(getTipoScale(stato)))
		entrataScala[stato] <- r
		<-r.ack

		// 2. Percorre la scala
		fmt.Printf("[VIAGGIATORE %d %s] Percorro le scale in %s...\n", id, getPonte(stato), strings.ToUpper(getTipoScale(stato)))
		sleepRandTimeRange(2, 5)

		// 3. Esce dalla scala
		fmt.Printf("[VIAGGIATORE %d %s] Esco dalle scale in %s\n", id, getPonte(stato), strings.ToUpper(getTipoScale(stato)))
		uscitaScala <- stato

		// 4. Cambia stato
		if stato == SALITA_VIAGGIATORE {
			stato = DISCESA_VIAGGIATORE
		} else {
			stato = SALITA_VIAGGIATORE
		}

		select {
		case <-termina:
			fmt.Printf("[VIAGGIATORE %d %s] Termino. Esco dall'aereo...\n", id, getPonte(stato))
			done <- true
			return
		default:
			sleepRandTimeRange(3, 5)
			continue
		}
	}
}
func equipaggio(id, stato_iniz int) {
	stato := stato_iniz
	r := Richiesta{id: id, ack: make(chan int)}

	fmt.Printf("[EQUIPAGGIO %d %s] Inizio. Salgo sull'aereo\n", id, getPonte(stato))
	for {
		// 1. Vuole usare la scala
		fmt.Printf("[EQUIPAGGIO %d %s] Voglio usare le scale in %s\n", id, getPonte(stato), strings.ToUpper(getTipoScale(stato)))
		entrataScala[stato+2] <- r // NB: SALITA_EQUIPAGGIO non è altro che SALITA + 2
		<-r.ack

		// 2. Percorre la scala
		fmt.Printf("[EQUIPAGGIO %d %s] Percorro le scale in %s...\n", id, getPonte(stato), strings.ToUpper(getTipoScale(stato)))
		sleepRandTimeRange(2, 5)

		// 3. Esce dalla scala
		fmt.Printf("[EQUIPAGGIO %d %s] Esco dalle scale in %s\n", id, getPonte(stato), strings.ToUpper(getTipoScale(stato)))
		uscitaScala <- stato

		// 4. Cambia stato
		if stato == SALITA_VIAGGIATORE {
			stato = DISCESA_VIAGGIATORE
		} else {
			stato = SALITA_VIAGGIATORE
		}

		select {
		case <-termina:
			fmt.Printf("[EQUIPAGGIO %d %s] Termino. Esco dall'aereo...\n", id, getPonte(stato))
			done <- true
			return
		default:
			continue
		}
	}
}
func cabina(p [2]int) {
	scala := make([]int, 2) // make inizializza allo zero value. Lo zero value degli interi è 0, quindi non serve altro.
	ponte := p

	fmt.Printf("[CABINA] Inizio. L'aereo inizia il decollo...\n")
	sleepRandTimeRange(2, 5) // decollo
	for {
		select {
		// ENTRATA SCALA - Controlli:
		// 0. L'aereo sia atterrato
		// 1. Dimensione massima Ponte destinatario
		// 2. Presenza di persone sulla scala (nel senso opposto)
		// 3. Precedenza verso la zona meno affollata
		// 4. Precedenza ai membri dell'equipaggio
		case r := <-when(ponte[PS] < TS &&
			scala[DISCESA] == 0 &&
			ponte[PS] < ponte[PI] &&
			len(entrataScala[SALITA_EQUIPAGGIO]) == 0, entrataScala[SALITA_VIAGGIATORE]):
			scala[SALITA]++
			fmt.Printf("[CABINA] Il VIAGGIATORE %d inizia a salire la scala\n", r.id)
			r.ack <- 1
		case r := <-when(ponte[PI] < TI &&
			scala[SALITA] == 0 &&
			ponte[PI] <= ponte[PS] &&
			len(entrataScala[DISCESA_EQUIPAGGIO]) == 0, entrataScala[DISCESA_VIAGGIATORE]):
			scala[DISCESA]++
			fmt.Printf("[CABINA] Il VIAGGIATORE %d inizia a scendere la scala\n", r.id)
			r.ack <- 1
		case r := <-when(ponte[PS] < TS &&
			scala[DISCESA] == 0 &&
			ponte[PS] < ponte[PI], entrataScala[SALITA_EQUIPAGGIO]):
			scala[SALITA]++
			fmt.Printf("[CABINA] Il membro dell'EQUIPAGGIO %d inizia a salire la scala\n", r.id)
			r.ack <- 1
		case r := <-when(ponte[PI] < TI &&
			scala[SALITA] == 0 &&
			ponte[PI] <= ponte[PS], entrataScala[DISCESA_EQUIPAGGIO]):
			scala[DISCESA]++
			fmt.Printf("[CABINA] Il membro dell'EQUIPAGGIO %d inizia a scendere la scala\n", r.id)
			r.ack <- 1

		// USCITA SCALA
		case x := <-uscitaScala:
			scala[x]--
			fmt.Printf("[CABINA] Uscito dalla scala un %s. Stato: PI: %d, PS: %d\n", getTipo(x), ponte[PI], ponte[PS])
			if x == SALITA { // PI -> PS
				ponte[PI]--
				ponte[PS]++
			} else { // PS -> PI
				ponte[PS]--
				ponte[PI]++
			}

		// ATTERRAGGIO
		case <-terminaCabina:
			fmt.Printf("\n[CABINA] L'aereo è atterrato, apertura porte...\n")
			sleepRandTimeRange(2, 5) // atterraggio
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

	nViaggiatori := 10
	nEquipaggio := 2

	fmt.Printf("\nQuanti Viaggiatori (max %d)? ", TS+TI-2)
	fmt.Scanf("%d\n", &nViaggiatori)
	fmt.Printf("\nQuanti membri dell'Equipaggio (max %d)? ", TS+TI-1-nViaggiatori)
	fmt.Scanf("%d\n", &nEquipaggio)

	// Inizializzazione canali
	for i := 0; i < len(entrataScala); i++ {
		entrataScala[i] = make(chan Richiesta, MAXBUFF)
	}

	// Esecuzione goroutine
	var ponte [2]int
	for i := 0; i < nViaggiatori; i++ {
		s := rand.Intn(2)
		go equipaggio(i, s)
		ponte[s]++
	}
	for i := 0; i < nViaggiatori; i++ {
		s := rand.Intn(2)
		go viaggiatore(i, s)
		ponte[s]++
	}
	go cabina(ponte) // Passo alla cabina il numero di posti occupati per ciascun ponte

	sleepRandTimeRange(10, 20) // Simulo un volo
	// Join goroutine
	for i := 0; i < nViaggiatori+nEquipaggio; i++ {
		termina <- true
		<-done
	}

	terminaCabina <- true
	<-done

	fmt.Printf("\n\n[MAIN] Fine\n")
}
