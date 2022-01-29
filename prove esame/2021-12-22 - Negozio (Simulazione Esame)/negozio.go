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
const MAXBUFF = 100 // buffer canale
const MAXCL = 50    // numero massimo di clienti
const MINCL = 10    // numero minimo di clienti
const MAXCO = 10    // numero massimo di commessi
const MINCO = 2     // numero minimo di commessi
const MAX = 8       // dimensione massima del negozio

const CLIENTE_A = 0 // Cliente Abituale
const CLIENTE_O = 1 // Cliente Occasionale
const COMMESSO = 2  // Commesso

const FUORI = -1 // Il commesso è fuori dal negozio

var NM = 10   // lotto di mascherine consegnate dal fornitore
var MAXM = 30 // numero massimo di mascherine nel distributore

// ==================================================================
// STRUTTURE DATI ===================================================
type Richiesta struct {
	id  int
	ack chan int
}

// ==================================================================
// CANALI ===========================================================
var entrata [3]chan Richiesta
var uscitaCliente = make(chan int)
var uscitaCommesso = make(chan Richiesta, MAXBUFF)
var rifornimentoMascherine = make(chan Richiesta)

// ==================================================================
// CANALI DI JOIN ===================================================
var done = make(chan bool)
var terminaCommesso = make(chan bool, MAXBUFF)
var terminaNegozio = make(chan bool)
var terminaFornitore = make(chan bool)

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
	if t == CLIENTE_A {
		return "abituale"
	} else if t == CLIENTE_O {
		return "occasionale"
	} else {
		return ""
	}
}

// Inizializza l'array di commessi
func initCommessi(nc int) []int {
	res := make([]int, nc)
	for i := 0; i < len(res); i++ {
		res[i] = FUORI
	}
	return res
}

// Restituisce il primo commesso che ha un posto disponibile
func getPrimoDisponibile(c []int) int {
	for i := 0; i < len(c); i++ {
		if c[i] >= 0 && c[i] < 3 {
			return i
		}
	}
	return -1
}

// restituisce un commesso disponibile in modo "più equo" (quello con meno clienti)
func getPrimoDisponibileEquo(c []int) int {
	for i := 0; i < len(c); i++ {
		if c[i] == 0 {
			return i
		}
	}
	for i := 0; i < len(c); i++ {
		if c[i] == 1 {
			return i
		}
	}
	for i := 0; i < len(c); i++ {
		if c[i] == 2 {
			return i
		}
	}
	return -1
}

// Soluzione equa As Fuck: restituisce il primo commesso disponibile, con meno clienti, iniziando la ricerca da un indice casuale
func getPrimoDisponibileEquoAF(c []int) int {
	si := rand.Intn(len(c))
	for i := si; i < len(c); i++ {
		if c[i] == 0 {
			return i
		}
		if i == len(c)-1 {
			i = 0
			continue
		}
		if i == si-1 {
			break
		}
	}
	for i := si; i < len(c); i++ {
		if c[i] == 1 {
			return i
		}
		if i == len(c)-1 {
			i = 0
			continue
		}
		if i == si-1 {
			break
		}
	}
	for i := si; i < len(c); i++ {
		if c[i] == 2 {
			return i
		}
		if i == len(c)-1 {
			i = 0
			continue
		}
		if i == si-1 {
			break
		}
	}
	return -1
}

func printStatoCommessi(c []int, d int) {
	res := "\nSTATO COMMESSI - disp: " + fmt.Sprintf("%d", d) + " | "
	for i := 0; i < len(c); i++ {
		res += fmt.Sprintf("%d", c[i]) + "/3, "
	}
	res += "\n"
	fmt.Printf("%s", res)
}

// ==================================================================
// GOROUTINE ========================================================
func cliente(id int) {
	tipo := rand.Intn(2) // cliente Abituale o Occasionale
	r := Richiesta{id: id, ack: make(chan int)}

	sleepRandTimeRange(id, id+1)
	fmt.Printf("[CLIENTE %s %d] Inizio\n", strings.ToUpper(getTipo(tipo)), id)

	// 1. Entra prendendo una mascherina e gli viene assegnato un commesso
	fmt.Printf("[CLIENTE %s %d] Voglio entrare nel negozio\n", strings.ToUpper(getTipo(tipo)), id)
	entrata[tipo] <- r
	commesso := <-r.ack // riceve la "risorsa", ovvero un commesso (ID)

	// 2. Permane
	fmt.Printf("[CLIENTE %s %d] Visito il negozio...\n", strings.ToUpper(getTipo(tipo)), id)
	sleepRandTimeRange(2, 5)

	// 3. Esce
	fmt.Printf("[CLIENTE %s %d] Esco dal negozio\n", strings.ToUpper(getTipo(tipo)), id)
	uscitaCliente <- commesso // rilascia la "risorsa"

	// 4. Termina
	fmt.Printf("[CLIENTE %s %d] Termino\n", strings.ToUpper(getTipo(tipo)), id)
	done <- true
}
func commesso(id int) {
	r := Richiesta{id: id, ack: make(chan int)}

	fmt.Printf("[COMMESSO %d] Inizio\n", id)
	for {
		// 1. Entra
		fmt.Printf("[COMMESSO %d] Entro nel negozio\n", id)
		entrata[COMMESSO] <- r
		<-r.ack

		// 2. Lavora
		sleepRandTimeRange(5, 10)

		// 3. Esce (può farlo solo quando non ha più clienti assegnati)
		fmt.Printf("[COMMESSO %d] Voglio andare in pausa\n", id)
		uscitaCommesso <- r
		<-r.ack

		// 4. Va in pausa
		fmt.Printf("[COMMESSO %d] Esco dal negozio...\n", id)

		select {
		case <-terminaCommesso:
			fmt.Printf("[COMMESSO %d] Termino\n", id)
			done <- true
			return
		default:
			sleepRandTimeRange(5, 10)
		}
	}

}
func fornitore() {
	r := Richiesta{id: NM, ack: make(chan int)}
	fmt.Printf("[FORNITORE] Inizio\n")
	for {
		sleepRandTimeRange(10, 15)

		// Rifornisce le mascherine
		fmt.Printf("[FORNITORE] Rifornisco un lotto di %d mascherine\n", NM)
		rifornimentoMascherine <- r
		<-r.ack

		select {
		case <-terminaFornitore:
			fmt.Printf("[FORNITORE] Termino\n")
			done <- true
			return
		default:
			continue
		}
	}
}

/*
POLITICA DI PRECEDENZA
Commessi:
1. C'è spazio?

Clienti:
1. Ci sono sufficienti mascherine?
2. C'è spazio?
3. Ci sono commessi liberi?
4. Precedenza Commessi, Clienti Ab. e Clienti Occ.
*/
func negozio(nc int) {
	commessi := initCommessi(nc)    // numero di clienti a cui ciascun commesso è associato
	vuoleUscire := make([]bool, nc) // indica se un commesso vuole uscire. Nel caso in cui si voglia che non vengano più assegnati clienti ad un commesso che vuole uscire
	var ackUscita [MAXCO]chan int   // ack per uscita commesso (NB: serve salvarlo in quanto un commesso può uscire anche alla ricezione di un messaggio "esci cliente")
	disponibili := 0                // numero di commessi disponibili
	nm := NM                        // numero mascherine presenti nel distributore
	nco := 0                        // numero commessi nel negozio
	ncl := 0                        // numero clienti nel negozio

	fmt.Printf("[NEGOZIO] Apertura\n")
	for {
		select {
		// ENTRATA
		case r := <-when(nco+ncl < MAX, entrata[COMMESSO]):
			// IL COMMESSO PUÒ ENTRARE
			//printStatoCommessi(commessi, disponibili) // DEBUG
			nco++
			commessi[r.id] = 0 // inizializzo il contatore dei clienti assegnati a questo commesso
			vuoleUscire[r.id] = false
			disponibili++
			fmt.Printf("[NEGOZIO] È entrato il COMMESSO %d. Disponibili: %d/%d\n", r.id, disponibili, nc)
			r.ack <- 1
		case r := <-when(nm > 0 &&
			nco+ncl < MAX &&
			disponibili > 0 &&
			len(entrata[COMMESSO]) == 0, entrata[CLIENTE_A]):
			// IL CLIENTE ABITUALE PUÒ ENTRARE
			//printStatoCommessi(commessi, disponibili) // DEBUG
			nm--                                   // decremento il numero di mascherine nel distributore
			ncl++                                  // incremento il numero di clienti nel negozio
			i := getPrimoDisponibileEquo(commessi) // prendo il primo commesso disponibile (non ci preoccupiamo che non ci siano commessi disponibili perché lo ha già verificato la condizione del when)
			commessi[i]++                          // incremento il numero di clienti che sta seguendo
			// controllo se ha raggiunto il max numero di clienti che può seguire (in caso aggiorno disponibili)
			if commessi[i] == 3 {
				disponibili--
			}
			fmt.Printf("[NEGOZIO] È entrato il CLIENTE ABITUALE %d e gli è stato assegnato il commesso %d (%d/3).\n", r.id, i, commessi[i])
			// restituisco al cliente l'ID del commesso
			r.ack <- i
		case r := <-when(nm > 0 &&
			nco+ncl < MAX &&
			disponibili > 0 &&
			len(entrata[COMMESSO]) == 0 &&
			len(entrata[CLIENTE_A]) == 0, entrata[CLIENTE_O]):
			// IL CLIENTE OCCASIONALE PUÒ ENTRARE
			//printStatoCommessi(commessi, disponibili) // DEBUG
			nm--                                   // decremento il numero di mascherine nel distributore
			ncl++                                  // incremento il numero di clienti nel negozio
			i := getPrimoDisponibileEquo(commessi) // prendo il primo commesso disponibile (non ci preoccupiamo che non ci siano commessi disponibili perché lo ha già verificato la condizione del when)
			commessi[i]++                          // incremento il numero di clienti che sta seguendo
			// controllo se ha raggiunto il max numero di clienti che può seguire (in caso aggiorno disponibili)
			if commessi[i] == 3 {
				disponibili--
			}
			fmt.Printf("[NEGOZIO] È entrato il CLIENTE OCCASIONALE %d e gli è stato assegnato il commesso %d (%d/3).\n", r.id, i, commessi[i])
			// restituisco al cliente l'ID del commesso
			r.ack <- i

		// USCITA
		case r := <-uscitaCommesso:
			//printStatoCommessi(commessi, disponibili) // DEBUG
			if commessi[r.id] == 0 { // non ha clienti assegnati, quindi può uscire
				nco--
				disponibili--
				commessi[r.id] = FUORI
				ackUscita[r.id] = nil
				fmt.Printf("[NEGOZIO] È uscito il COMMESSO %d. Disponibili: %d/%d\n", r.id, disponibili, nc)
				r.ack <- 1
			} else { // ha clienti assegnati, quindi non può uscire e attende
				vuoleUscire[r.id] = true
				ackUscita[r.id] = r.ack
				fmt.Printf("[NEGOZIO] il COMMESSO %d non può uscire: Sta supervisionando ancora %d clienti!\n", r.id, commessi[r.id])
			}
		case i := <-uscitaCliente:
			//printStatoCommessi(commessi, disponibili) // DEBUG
			ncl--
			commessi[i]--
			fmt.Printf("[NEGOZIO] È uscito un CLIENTE associato al commesso %d (%d/3). Stato negozio: %d + %d / %d (Clienti+Commessi/Max)\n", i, commessi[i], ncl, nco, MAX)
			if commessi[i] == 2 {
				disponibili++ // il commesso i-esimo torna disponibile
			}
			if commessi[i] == 0 && vuoleUscire[i] { // il commesso esce
				disponibili--
				commessi[i] = -1
				fmt.Printf("[NEGOZIO] il COMMESSO %d può uscire\n", i)
				ackUscita[i] <- 1 // notifico al commesso che può uscire
				ackUscita[i] = nil
			}

		// RIFORNIMENTO
		case r := <-rifornimentoMascherine:
			if nm < MAXM-r.id {
				nm += r.id
			} else {
				nm = MAXM
			}
			fmt.Printf("[NEGOZIO] Distributore di mascherine rifornito: %d/%d\n", nm, MAXM)
			r.ack <- 1

		// TERMINAZIONE
		case <-terminaNegozio:
			fmt.Printf("[NEGOZIO] Chisura\n")
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

	nClienti := 10
	nCommessi := 2

	fmt.Printf("\nQuanti clienti (min %d, max %d)? ", MINCL, MAXCL)
	fmt.Scanf("%d\n", &nClienti)
	fmt.Printf("\nQuanti commessi (min %d, max %d)? ", MINCO, MAXCO)
	fmt.Scanf("%d\n", &nCommessi)

	// INIZIALIZZAZIONE CANALI
	for i := 0; i < len(entrata); i++ {
		entrata[i] = make(chan Richiesta, MAXBUFF)
	}

	// ESECUZIONE GOROUTINE
	go negozio(nCommessi)

	for i := 0; i < nCommessi; i++ {
		go commesso(i)
	}
	for i := 0; i < nClienti; i++ {
		go cliente(i)
	}
	go fornitore()

	// JOIN GOROUTINE
	// Clienti
	for i := 0; i < nClienti; i++ {
		<-done
	}
	// Fornitore
	terminaFornitore <- true
	<-done

	/* NB: quando voglio terminare un cliente con i segnali, posso fare in
	due modi:
	1) Uso 2 cicli spezzati. In questo caso terminaCommesso deve avere un
	buffer di capacità >= nCommessi, quindi >= MAXCO (ho usato MAXBUFF per
	comodità), come segue.
	*/
	// Commessi
	for i := 0; i < nCommessi; i++ {
		terminaCommesso <- true
	}
	for i := 0; i < nCommessi; i++ {
		<-done
	}

	/*2) Altrimenti si può usare un unico ciclo e fare le terminazioni entrambe
	in modo sincrono, come nel commento seguente. In questo caso, ovviamente,
	il buffer non serve.
	for i := 0; i < nCommessi; i++ {
		terminaCommesso <- true
		<-done
	}*/
	// Negozio
	terminaNegozio <- true
	<-done

	fmt.Printf("\n\n[MAIN] Fine\n")
}
