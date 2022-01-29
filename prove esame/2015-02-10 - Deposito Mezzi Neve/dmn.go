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
const MAX_SPAZZANEVE = 20
const MAX_SPARGISALE = 20
const MAX_CAMION = 5

const K = 10 // capacità massima del silos (m³ di sale)
const N = 2  // capacità dell'area del DMN
const S = 4  // capacità serbatoio

const SPAZZANEVE = 0
const SPARGISALE = 1
const CAMION = 2

// ==================================================================
// STRUTTURE DATI ===================================================
type Richiesta struct {
	id  int
	ack chan int
}

// ==================================================================
// CANALI ===========================================================
var entrata [3]chan Richiesta
var uscita = make(chan Richiesta)

// ==================================================================
// CANALI DI JOIN ===================================================
var done = make(chan bool)
var termina = make(chan bool)
var terminaDMM = make(chan bool)

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

// ==================================================================
// GOROUTINE ========================================================
func spazzaneve(id int) {
	r := Richiesta{id: id, ack: make(chan int)}

	fmt.Printf("[SPAZZANEVE %d] Inizio\n", id)
	for i := 0; i < 5; i++ {
		// Spazza la neve
		fmt.Printf("[SPAZZANEVE %d] Spazzo la neve...\n", id)
		sleepRandTime(4)

		if rand.Intn(4) == 2 { // 25% di probabilità di entrare nel DMN
			// 1. Entra nel DMN
			fmt.Printf("[SPAZZANEVE %d] Voglio entrare nel DMN per sostare\n", id)
			entrata[SPAZZANEVE] <- r
			<-r.ack

			// 2. Sosta per un tempo arbitrario
			fmt.Printf("[SPAZZANEVE %d] Sosto...\n", id)
			sleepRandTime(5)

			// 3. Esce dal DMN
			fmt.Printf("[SPAZZANEVE %d] Esco dal DMN...\n", id)
			uscita <- r
			<-r.ack
		}
	}
	fmt.Printf("[SPAZZANEVE %d] Termino\n", id)
	done <- true
}
func spargisale(id int) {
	serbatoio := rand.Intn(S)
	r := Richiesta{id: id, ack: make(chan int)}

	fmt.Printf("[SPARGISALE %d] Inizio\n", id)
	for i := 0; i < 5; i++ {
		if serbatoio == 0 {
			// 1. Entra nel DMN
			fmt.Printf("[SPARGISALE %d] Devo entrare nel DMN per riempire il serbatoio (0%%)\n", id)
			entrata[SPARGISALE] <- r
			serbatoio = (<-r.ack) * S

			// 2. Si rifornisce di sale (sottraggono 1 m³ al silos)
			fmt.Printf("[SPARGISALE %d] Rifornimento sale in corso...\n", id)
			sleepRandTime(5)

			// 3. Esce dal DMN
			fmt.Printf("[SPARGISALE %d] Esco dal DMN\n", id)
			uscita <- r
			<-r.ack
		}
		// Sparge il sale
		fmt.Printf("[SPARGISALE %d] Spargo il sale... (serbatoio: %d%%)\n", id, serbatoio*100/4)
		sleepRandTime(4)
		serbatoio--
	}
	fmt.Printf("[SPARGISALE %d] Termino\n", id)
	done <- true
}
func camion(id int) {
	r := Richiesta{id: id, ack: make(chan int)}

	fmt.Printf("[CAMION %d] Inizio\n", id)
	for {
		sleepRandTime(20)

		// 1. Entrano nel DMN
		fmt.Printf("[CAMION %d] Voglio entrare nel DMN\n", id)
		entrata[CAMION] <- r
		<-r.ack

		// 2. Riempiono COMPLETAMENTE il silos
		fmt.Printf("[CAMION %d] Riempimento silos...\n", id)
		sleepRandTime(10)

		// 3. Escono dal DMN
		fmt.Printf("[CAMION %d] Esco dal DMN\n", id)
		uscita <- r
		<-r.ack

		select {
		case <-termina:
			fmt.Printf("[CAMION %d] Termino\n", id)
			done <- true
			return
		default:
			continue
		}
	}
}

// gestore
func dmn(k int) {
	silos := k // m³ di sale attualmente presenti nel silos
	n := 0     // posti occupati nel DMN

	fmt.Printf("[DMN] Inizio\n")
	for {
		select {
		case r := <-when(n < N, entrata[SPAZZANEVE]):
			n++
			fmt.Printf("[DMN] È entrato uno SPAZZANEVE. Stato DMN: %d/%d m³, %d/%d posti occupati\n", silos, K, n, N)
			r.ack <- 1
		case r := <-when(n < N && len(entrata[SPAZZANEVE]) == 0, entrata[CAMION]):
			n++
			fmt.Printf("[DMN] È entrato un CAMION. Stato DMN: %d/%d m³, %d/%d posti occupati\n", silos, K, n, N)
			silos = K
			r.ack <- 1
		case r := <-when(n < N && silos > 0 && len(entrata[SPAZZANEVE]) == 0 && len(entrata[CAMION]) == 0, entrata[SPARGISALE]):
			n++
			fmt.Printf("[DMN] È entrato uno SPARGISALE. Stato DMN: %d/%d m³, %d/%d posti occupati\n", silos, K, n, N)
			silos--
			r.ack <- 1
		case r := <-uscita:
			n--
			fmt.Printf("[DMN] È uscito un veicolo. Stato DMN: %d/%d m³, %d/%d posti occupati\n", silos, K, n, N)
			r.ack <- 1
		case <-terminaDMM:
			fmt.Printf("[DMN] Termino\n")
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

	var kIniziali int
	var nSpazzaneve int
	var nSpargisale int
	var nCamion int

	fmt.Printf("\nQuanti spazzaneve (max %d)? ", MAX_SPAZZANEVE)
	fmt.Scanf("%d\n", &nSpazzaneve)
	fmt.Printf("\nQuanti spargisale (max %d)? ", MAX_SPARGISALE)
	fmt.Scanf("%d\n", &nSpargisale)
	fmt.Printf("\nQuanti camion (max %d)? ", MAX_CAMION)
	fmt.Scanf("%d\n", &nCamion)

	kIniziali = K
	/*nSpazzaneve = 10
	nSpargisale = 10
	nCamion = 2*/

	// Inizializzazione canali
	for i := 0; i < len(entrata); i++ {
		entrata[i] = make(chan Richiesta, MAXBUFF)
	}

	// Esecuzione goroutine
	go dmn(kIniziali)

	for i := 0; i < nSpazzaneve; i++ {
		go spazzaneve(i)
	}
	for i := 0; i < nSpargisale; i++ {
		go spargisale(i)
	}
	for i := 0; i < nCamion; i++ {
		go camion(i)
	}

	// Join goroutine
	for i := 0; i < nSpazzaneve+nSpargisale; i++ {
		<-done
	}

	for i := 0; i < nCamion; i++ {
		termina <- true
		<-done
	}

	terminaDMM <- true
	<-done

	fmt.Printf("\n\n[MAIN] Fine\n")
}
