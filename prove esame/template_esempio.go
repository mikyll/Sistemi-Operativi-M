/*
CONCETTO:
C'è un gestore di risorse, e dei clienti di tipo diverso che vi accedono,
richiedendole, ottenendole e rilasciandole.
L'ottenimento delle risorse è vincolato da politiche di precedenza
fra i tipo diversi di clienti che accedono a queste.

Sostanzialmente, per ogni tipo di cliente che vuole accedere ad una
risorsa, per cui la traccia impone delle condizioni o regole di precedenza,
bisogna creare un canale dedicato; nel server (gestore), vi sarà un caso
con i controlli relativi a ciascun canale, utilizzando le guardie logiche.

In Go per realizzare una guardia logica utilizziamo una funzione
when(b bool, c chan) chan.
b è la condizione da verificare;
c è il canale da cui viene letto un valore se la condizione è verificata

il valore di ritorno è il canale c da cui è possibile leggere un valore,
tramite l'operatore <-

Se viene inserito in un caso della select, all'interno di un ciclo for,
la select leggerà un valore dal canale, solo se:
1. la condizione di when (guardia logica) è verificata, e
2. c'è un valore da leggere nel canale.
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
const MAXPROC = 10
const MAX = 5 // capacità

const TIPO1 = 0 // Risorsa di tipo1 (stampante)
const TIPO2 = 1 // Risorsa di tipo2 (computer)
const TIPO3 = 2 // Risorsa di tipo3 (telefono)

const NM = 5         // Numero di mascherine rifornite dall'operatore
const DIM_STANZA = 5 // massimo numero di utenti che possono essere contemporaneamente presenti nella stanza

// ==================================================================
// STRUTTURE DATI ===================================================
type Richiesta struct {
	id  int
	ack chan int
}

// ==================================================================
// CANALI ===========================================================
var richiestaRisorsa [3]chan Richiesta
var entrataOperatore = make(chan Richiesta, MAXBUFF)
var rilascioRisorsa = make(chan int, MAXBUFF)
var uscitaOperatore = make(chan Richiesta)

// ==================================================================
// CANALI DI JOIN ===================================================
var done = make(chan bool)
var termina = make(chan bool)
var terminaServer = make(chan bool)

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

// Funzione di debug
func debug(params ...int) {
	res := "[[[DEBUG]]] "
	for _, p := range params {
		res += (string(p) + ", ")
	}
	fmt.Println(res)
}

func getTipo(t int) string {
	if t == TIPO1 {
		return "stampante"
	} else if t == TIPO2 {
		return "computer"
	} else if t == TIPO3 {
		return "telefono"
	} else {
		return ""
	}
}

// ==================================================================
// GOROUTINE ========================================================
func client(id int) {
	fmt.Printf("[CLIENTE %d] Inizio\n", id)
	tipo := rand.Intn(3)
	r := Richiesta{id, make(chan int)}

	// 1. Richiede risorsa
	fmt.Printf("[CLIENTE %d] Voglio la risorsa %s\n", id, strings.ToUpper(getTipo(tipo)))
	richiestaRisorsa[tipo] <- r // inserisce la richiesta
	<-r.ack                     // attende che la risorsa gli venga assegnata

	// 2. Fa cose
	fmt.Printf("[CLIENTE %d] Uso la risorsa...\n", id)
	sleepRandTime(5) // simula l'utilizzo

	// 3. Rilascia risorsa
	fmt.Printf("[CLIENTE %d] Rilascio la risorsa %s\n", id, strings.ToUpper(getTipo(tipo)))
	rilascioRisorsa <- tipo

	fmt.Printf("[CLIENTE %d] Termino\n", id)
	done <- true
}
func operatore(id int) {
	fmt.Printf("[OPERATORE %d] Inizio\n", id)
	r := Richiesta{id, make(chan int)}

	// do things
	for {

		// 1. Richiede risorsa
		fmt.Printf("[OPERATORE %d] Voglio entrare\n", id)
		entrataOperatore <- r
		<-r.ack

		// 2. Fa cose
		fmt.Printf("[OPERATORE %d] Svolgo la mia attività...\n", id)
		sleepRandTime(5)

		// 3. Rilascia risorsa
		fmt.Printf("[OPERATORE %d] Esco\n", id)
		uscitaOperatore <- r
		<-r.ack

		select {
		case <-termina:
			fmt.Printf("[OPERATORE %d] Termino\n", id)
			done <- true
			return
		default:
			sleepRandTime(5)
		}
	}
}

/*
ESEMPIO DI SERVER:
I clienti accedono alle risorse e per farlo devono prendere una mascherina. L'operatore ogni tanto rifornisce le mascherine.

VINCOLI:
- Nella stanza ci possono essere contemporaneamente al massimo 5 persone, operatore compreso.
- Una stampante (TIPO1) può essere usata da massimo 2 utenti contemporaneamente
- Un computer (TIPO2) può essere usato da massimo 3 utenti contemporaneamente
- Un telefono (TIPO3) può essere usato da massimo 1 utente contemporaneamente

POLITICA DI PRECEDENZA:
Stampanti > Computer > Telefono

CONDIZIONI:
- Operatore - l'operatore può entrare solo se:
	1. c'è spazio nella stanza
- Stampante (TIPO1) - la stampante può essere utilizzata solo se:
	1. la stanza non è piena (< 5)
	2. la risorsa non ha raggiunto il numero massimo di utenti (< 2)
	3. c'è almeno una mascherina
	4. Precedenza: non c'è l'operatore in attesa di rifornire le mascherine
- Computer (TIPO2) - il computer può essere utilizzato solo se:
	1. la stanza non è piena
	2. la risorsa non ha raggiunto il numero massimo di utenti (< 3)
	3. c'è almeno una mascherina
	4. Precedenza: non c'è l'operatore in attesa e la coda per accedere alla stampante è vuota
- Telefono (TIPO3) - il telefono può essere utilizzato solo se:
	1. la stanza non è piena
	2. la risorsa non ha raggiunto il numero massimo di utenti (< 1)
	3. c'è almeno una mascherina
	4. Precedenza: non c'è l'operatore in attesa e le code per accedere a stampante e computer sono vuote
*/
func server() {
	// Variabili varie
	stanza := 0
	risorse := make([]int, 3) // tiene traccia di quante persone stanno usando una risorsa
	mascherine := 3

	fmt.Printf("[SERVER] Inizio\n")
	for {
		select {
		case r := <-when(stanza < DIM_STANZA && risorse[TIPO1] < 2 && mascherine > 0 && len(entrataOperatore) == 0, richiestaRisorsa[TIPO1]): // condizione 1 (stampante)
			// NB: Quando la condizione di when(<condizione>, <canale>) è verificata, viene letto il primo valore di canale ed eseguito il codice seguente:
			stanza++
			mascherine--
			risorse[TIPO1]++
			fmt.Printf("[SERVER] Il cliente %d è entrato ed inizia a usare la stampante.\n", r.id)
			fmt.Printf("\t[SERVER] Utilizzo risorse - S: %d/%d, C: %d/%d, T: %d/%d. Mascherine: %d. Attesa - S: %d, C: %d, T: %d, O: %d\n", risorse[0], 2, risorse[1], 3, risorse[2], 1, mascherine, len(richiestaRisorsa[0]), len(richiestaRisorsa[1]), len(richiestaRisorsa[2]), len(entrataOperatore))
			r.ack <- 1
		case r := <-when(stanza < DIM_STANZA && risorse[TIPO2] < 3 && mascherine > 0 && len(entrataOperatore) == 0 && len(richiestaRisorsa[TIPO1]) == 0, richiestaRisorsa[TIPO2]): // condizione 2 (computer)
			stanza++
			mascherine--
			risorse[TIPO2]++
			fmt.Printf("[SERVER] Il cliente %d è entrato ed inizia a usare il computer.\n", r.id)
			fmt.Printf("\t[SERVER] Utilizzo risorse - S: %d/%d, C: %d/%d, T: %d/%d. Mascherine: %d. Attesa - S: %d, C: %d, T: %d, O: %d\n", risorse[0], 2, risorse[1], 3, risorse[2], 1, mascherine, len(richiestaRisorsa[0]), len(richiestaRisorsa[1]), len(richiestaRisorsa[2]), len(entrataOperatore))
			r.ack <- 1
		case r := <-when(stanza < DIM_STANZA && risorse[TIPO3] < 3 && mascherine > 0 && len(entrataOperatore) == 0 && len(richiestaRisorsa[TIPO1]) == 0 && len(richiestaRisorsa[TIPO2]) == 0, richiestaRisorsa[TIPO3]): // condizione 3
			stanza++
			mascherine--
			risorse[TIPO3]++
			fmt.Printf("[SERVER] Il cliente %d è entrato ed inizia a usare il telefono.\n", r.id)
			fmt.Printf("\t[SERVER] Utilizzo risorse - S: %d/%d, C: %d/%d, T: %d/%d. Mascherine: %d. Attesa - S: %d, C: %d, T: %d, O: %d\n", risorse[0], 2, risorse[1], 3, risorse[2], 1, mascherine, len(richiestaRisorsa[0]), len(richiestaRisorsa[1]), len(richiestaRisorsa[2]), len(entrataOperatore))
			r.ack <- 1

		case r := <-when(stanza < DIM_STANZA, entrataOperatore):
			stanza++ // entrata operatore
			fmt.Printf("[SERVER] È entrato l'operatore %d per rifornire le mascherine.\n", r.id)
			fmt.Printf("\t[SERVER] Utilizzo risorse - S: %d/%d, C: %d/%d, T: %d/%d. Mascherine: %d. Attesa - S: %d, C: %d, T: %d, O: %d\n", risorse[0], 2, risorse[1], 3, risorse[2], 1, mascherine, len(richiestaRisorsa[0]), len(richiestaRisorsa[1]), len(richiestaRisorsa[2]), len(entrataOperatore))
			r.ack <- 1

		case tipo := <-rilascioRisorsa:
			risorse[tipo]--
			stanza--
			fmt.Printf("[SERVER] Risorsa %s rilasciata da un cliente\n", strings.ToUpper(getTipo(tipo)))
			fmt.Printf("\t[SERVER] Utilizzo risorse - S: %d/%d, C: %d/%d, T: %d/%d. Mascherine: %d. Attesa - S: %d, C: %d, T: %d, O: %d\n", risorse[0], 2, risorse[1], 3, risorse[2], 1, mascherine, len(richiestaRisorsa[0]), len(richiestaRisorsa[1]), len(richiestaRisorsa[2]), len(entrataOperatore))

		case r := <-uscitaOperatore:
			mascherine += 5 // rifornimento mascherine
			stanza--
			fmt.Printf("[SERVER] Mascherine rifornite. Disponibili: %d\n", mascherine)
			r.ack <- 1

		case <-terminaServer:
			fmt.Printf("[SERVER] Termino\n")
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
	var nOperatori int

	fmt.Printf("\nQuanti Client (max %d)? ", 10)
	fmt.Scanf("%d\n", &nClient)
	/*fmt.Printf("\nQuanti Operatori (max %d)? ", 2)
	fmt.Scanf("%d\n", &nOperatori)*/

	nClient = 10
	nOperatori = 1

	// INIZIALIZZAZIONE CANALI
	// Essendo richiestaRisorsa un array di canali, devo inizializzare ciascuno.
	for i := 0; i < len(richiestaRisorsa); i++ {
		richiestaRisorsa[i] = make(chan Richiesta)
	}

	// ESECUZIONE GOROUTINE
	// Server (Gestore risorse)
	go server()

	// Client
	for i := 0; i < nClient; i++ {
		go client(i)
	}
	for i := 0; i < nOperatori; i++ {
		go operatore(i)
	}

	// TERMINAZIONE (JOIN) GOROUTINE
	// Client (terminano da soli)
	for i := 0; i < nClient; i++ {
		<-done
	}
	// Operatori (terminano dopo i client)
	for i := 0; i < nOperatori; i++ {
		termina <- true
		<-done
	}
	// Server
	terminaServer <- true
	<-done

	fmt.Printf("\n\n[MAIN] Fine\n")
}
