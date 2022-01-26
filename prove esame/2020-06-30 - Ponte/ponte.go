package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ==================================================================
// COSTANTI =========================================================
const MAXBUFF = 100
const MAXA = 50
const MAXB = 20
const MAX = 5 // capacità massima di autoveicoli che possono transitare contemporaneamente su di esso

const TA = true      // stato A: transito Autoveicoli (ponte abbassato)
const TB = false     // stato B: transito Barche (ponte alzato)
const PUBBLICO_A = 0 // veicolo pubblico che percorre in senso A
const PUBBLICO_B = 1 // veicolo pubblico che percorre in senso B
const PRIVATO_A = 2  // veicolo privato che percorre in senso A
const PRIVATO_B = 3  // veicolo privato che percorre in senso B

// ==================================================================
// STRUTTURE DATI ===================================================
type Richiesta struct {
	id  int
	ack chan int
}

// ==================================================================
// CANALI ===========================================================
var entrataBarca = make(chan Richiesta, MAXBUFF)
var entrataVeicolo [4]chan Richiesta
var uscitaBarca = make(chan int) // NB: si poteva usare anche un unico canale di uscita
var uscitaVeicolo = make(chan int)

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
func getTipoVeicolo(t int) string {
	switch t {
	case PUBBLICO_A:
		return "mezzo pubblico in senso A"
	case PUBBLICO_B:
		return "mezzo pubblico in senso B"
	case PRIVATO_A:
		return "mezzo privato in senso A"
	case PRIVATO_B:
		return "mezzo privato in senso B"
	default:
		return ""
	}
}
func getStatoPonte(t bool) string {
	if t {
		return "TA"
	} else {
		return "TB"
	}
}

// ==================================================================
// GOROUTINE ========================================================
// cliente di tipo 1
func imbarcazione(id int) {
	r := Richiesta{id: id, ack: make(chan int)}
	fmt.Printf("[BARCA %d] Inizio\n", id)

	// 1. Richiesta risorsa PONTE
	entrataBarca <- r
	// 2. Ottenimento risorsa PONTE
	<-r.ack

	// 3. Attraversamento PONTE
	fmt.Printf("[BARCA %d] Attraverso il ponte...\n", id)
	sleepRandTime(10)

	// 4. Rilascio risorsa PONTE
	uscitaBarca <- 1

	fmt.Printf("[BARCA %d] Termino\n", id)
	done <- true
}

// cliente di tipo 2
func autoveicolo(id int) {
	tipo := rand.Intn(4) // tipo e senso del veicolo
	r := Richiesta{id: id, ack: make(chan int)}
	fmt.Printf("[VEICOLO %d] Inizio. Tipo: %s\n", id, getTipoVeicolo(tipo))

	// 1. Richiesta risorsa PONTE
	entrataVeicolo[tipo] <- r
	// 2. Ottenimento risorsa PONTE
	<-r.ack

	// 3. Attraversamento PONTE
	fmt.Printf("[VEICOLO %d] Attraverso il ponte...\n", id)
	sleepRandTime(4)

	// 4. Rilascio risorsa PONTE
	uscitaVeicolo <- tipo

	fmt.Printf("[VEICOLO %d] Termino\n", id)
	done <- true
}

// gestore risorsa (server)
func ponte() {
	sp := true // Stato Ponte (TA == true o TB == false)
	nb := 0    // Numero Barche in transito sul ponte
	nvA := 0   // Numero Veicoli in transito sul ponte (< MAX) in senso A
	nvB := 0   // Numero Veicoli in transito sul ponte (< MAX) in senso B

	fmt.Printf("[PONTE] Inizio\n")
	for {
		select {
		case r := <-when(!sp, entrataBarca):
			nb++
			fmt.Printf("[PONTE %s] È entrata una barca\n", getStatoPonte(sp))
			r.ack <- 1

		case r := <-when(sp &&
			len(entrataBarca) == 0 &&
			nvA < MAX &&
			nvB == 0, entrataVeicolo[PUBBLICO_A]):
			nvA++
			fmt.Printf("[PONTE %s] È entrato un mezzo pubblico in senso A\n", getStatoPonte(sp))
			r.ack <- 1

		case r := <-when(sp &&
			len(entrataBarca) == 0 &&
			nvB < MAX &&
			nvA == 0, entrataVeicolo[PUBBLICO_B]):
			nvB++
			fmt.Printf("[PONTE %s] È entrato un mezzo pubblico in senso B\n", getStatoPonte(sp))
			r.ack <- 1

		case r := <-when(sp &&
			len(entrataBarca) == 0 &&
			nvA < MAX &&
			nvB == 0 &&
			len(entrataVeicolo[PUBBLICO_A]) == 0, entrataVeicolo[PRIVATO_A]):
			nvA++
			fmt.Printf("[PONTE %s] È entrato un mezzo privato in senso A\n", getStatoPonte(sp))
			r.ack <- 1

		case r := <-when(sp &&
			len(entrataBarca) == 0 &&
			nvB < MAX &&
			nvA == 0 &&
			len(entrataVeicolo[PUBBLICO_B]) == 0, entrataVeicolo[PRIVATO_B]):
			nvB++
			fmt.Printf("[PONTE %s] È entrato un mezzo privato in senso B\n", getStatoPonte(sp))
			r.ack <- 1

		case <-uscitaBarca:
			fmt.Printf("[PONTE %s] È uscito dal ponte una barca\n", getStatoPonte(sp))
			nb--
			// se il ponte è vuoto e ci sono auto in attesa, inizio ad abbassare il ponte
			if nb == 0 &&
				(len(entrataVeicolo[PUBBLICO_A]) > 0 || len(entrataVeicolo[PUBBLICO_B]) > 0 ||
					len(entrataVeicolo[PRIVATO_A]) > 0 || len(entrataVeicolo[PRIVATO_B]) > 0) {
				fmt.Printf("[PONTE TB->TA] Chiusura del ponte in corso...\n")
				time.Sleep(time.Duration(2 * time.Second)) // Simulo la chiusura del ponte
				sp = TA
			}

		case x := <-uscitaVeicolo:
			if x == PUBBLICO_A || x == PRIVATO_A {
				nvA--
			} else if x == PUBBLICO_B || x == PRIVATO_B {
				nvB--
			}
			fmt.Printf("[PONTE %s] È uscito dal ponte un %s\n", getStatoPonte(sp), getTipoVeicolo(x))

			// se il ponte è vuoto e ci sono barche in attesa, inizio ad alzare il ponte
			if (nvA+nvB) == 0 && len(entrataBarca) > 0 {
				fmt.Printf("[PONTE TA->TB] Apertura del ponte in corso...\n")
				time.Sleep(time.Duration(2 * time.Second)) // Simulo l'apertura del ponte
				sp = TB
			}

		case <-termina:
			fmt.Printf("[PONTE] Termino\n")
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

	var nImbarcazioni int
	var nAutoveicoli int

	fmt.Printf("\nQuanti veicoli (max %d)? ", MAXA)
	fmt.Scanf("%d\n", &nImbarcazioni)
	fmt.Printf("\nQuante imbarcazioni (max %d)? ", MAXB)
	fmt.Scanf("%d\n", &nAutoveicoli)

	// Inizializzazione canali
	for i := 0; i < len(entrataVeicolo); i++ {
		entrataVeicolo[i] = make(chan Richiesta, MAXBUFF)
	}

	/*nImbarcazioni = 10
	nAutoveicoli = 50*/

	// Esecuzione goroutine
	go ponte()

	for i := 0; i < nAutoveicoli; i++ {
		go autoveicolo(i)
	}
	time.Sleep(time.Duration(2 * time.Second))
	for i := 0; i < nImbarcazioni; i++ {
		sleepRandTime(2)
		go imbarcazione(i)
	}

	// Join goroutine
	for i := 0; i < nAutoveicoli+nImbarcazioni; i++ {
		<-done
	}

	termina <- true
	<-done
	fmt.Printf("\n\n[MAIN] Fine\n")
}
