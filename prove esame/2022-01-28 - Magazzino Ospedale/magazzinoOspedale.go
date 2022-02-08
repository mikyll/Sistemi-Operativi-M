// Svolto da: Michele Righi (Esame 2022-01-28)
// ==================================================================
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
const MAXAR = 20 // Massimo numero di Addetti Reparto

const F = 0 // mascherina Ffp2
const C = 1 // mascherina Chirurgica
const M = 2 // mascherine Miste

const NF = 4000 // max Numero di mascherine Ffp2 che lo scaffale può contenere
const NC = 3000 // max Numero di mascherine Chirurgiche che lo scaffale può contenere

const LF = 700 // Lotto Ffp2
const LC = 300 // Lotto Chirurgighe
const LM = 500 // Lotto Misto

// ==================================================================
// STRUTTURE DATI ===================================================
type Richiesta struct {
	id   int
	tipo int
	ack  chan int
}

// ==================================================================
// CANALI ===========================================================
var prelievo [3]chan Richiesta
var rifornimento [2]chan Richiesta
var finePrelievo = make(chan Richiesta, MAXBUFF)
var fineRifornimento = make(chan Richiesta)

// ==================================================================
// CANALI DI JOIN ===================================================
var done = make(chan bool)
var terminaMagazzino = make(chan bool)
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
	if t == C {
		return "chirurgiche"
	} else if t == F {
		return "ffp2"
	} else if t == M {
		return "miste"
	} else {
		return ""
	}
}

func debug(pf, pc, lenf, lenc, lenm int) {
	fmt.Printf("Prelievi - PF: %d, PC: %d | LEN_F: %d, LEN_C: %d, LEN_M: %d\n", pf, pc, lenf, lenc, lenm)
}

// ==================================================================
// GOROUTINE ========================================================
// Addetto Reparto AR
func addetto(id int) {
	tipo := -1
	r := Richiesta{id: id, tipo: tipo, ack: make(chan int)}
	// ciclicamente accede al magazzino per prelevare un lotto di uno dei 3 particolari formati (LF, LC, LC)
	// decide in modo randomico che tipo di prelievo fare

	fmt.Printf("[ADDETTO %d] Inizio\n", id)
	for i := 0; i < 5; i++ {
		// Scelta tipo di prelievo da fare (LF, LC, LM)
		tipo := rand.Intn(100)
		if tipo >= 80 { // la probabilità che l'addetto voglia mascherine miste è inferiore
			r.tipo = M
		} else {
			r.tipo = tipo % 2
		}

		// 1. Richiede mascherine (Richiesta)
		fmt.Printf("[ADDETTO %d] Voglio prelevare le mascherine %s\n", id, strings.ToUpper(getTipo(tipo)))
		prelievo[r.tipo] <- r
		<-r.ack

		// 2. Prelievo tempo non trascurabile
		fmt.Printf("[ADDETTO %d] Sto prelevando le mascherine %s...\n", id, strings.ToUpper(getTipo(tipo)))
		sleepRandTime(3)

		// 3. Fine prelievo (Rilascio risorsa)
		finePrelievo <- r // ERRORE (-2 punti): serviva una send sincrona (<-r.ack)
		/*
			Spiegazione:
			L'addetto fa più cicli: se la richiesta del fine rifornimento non è stata ancora gestita, potrebbe essere presa in carico una nuova richiesta di prelievo,
			senza aver decrementato il contatore di quelle disponibili.
			Di conseguenza, verrà decrementato dopo che viene accettato un nuovo prelievo e si avrà un numero di mascherine negativo.
		*/
	}
	done <- true
	fmt.Printf("[ADDETTO %d] Termino\n", id)
}
func fornitore(t int) {
	r := Richiesta{id: 0, tipo: t, ack: make(chan int)}
	// accede con frequenza variabile e casuale al magazzino

	fmt.Printf("[FORNITORE %s] Inizio\n", strings.ToUpper(getTipo(t)))
	for {
		sleepRandTimeRange(5, 10)

		// 1. Rifornisce mascherine (Richiesta)
		// Ottenimento risorsa (attende che non ci siano addetti che prelevano da quel reparto)
		fmt.Printf("[FORNITORE %s] Voglio rifornire il reparto\n", strings.ToUpper(getTipo(t)))
		rifornimento[t] <- r
		<-r.ack

		// 2. Rifornimento (tempo trascurabile) - nel mentre nessuno può prelevare le scatole
		fmt.Printf("[FORNITORE %s] Sto rifornendo gli scaffali...\n", strings.ToUpper(getTipo(t)))
		sleepRandTimeRange(3, 5)

		// 3. Rifornimento completato (Rilascio)
		fineRifornimento <- r
		<-r.ack
		fmt.Printf("[FORNITORE %s] Ho concluso il rifornimento\n", strings.ToUpper(getTipo(t)))

		select {
		case <-terminaFornitore:
			fmt.Printf("[FORNITORE %s] Termino\n", strings.ToUpper(getTipo(t)))
			done <- true
			return
		default:
			continue
		}
	}
}

/*
CONDIZIONI:
Prelievo
1. Sono disponibili abbastanza mascherine (anche con quelle che stanno venendo prelevate da altri AR)
2. Non c'è un rifornimento in corso
3. Precedenza (PM > PF > PC)

Rifornimento:
1. Il reparto è vuoto (nessuno stia prelevando)
2. Precedenza (Scaffale con meno scatole. Parità: FFP2, ovvero nessuno in coda in rifornimento chirurgiche)
*/
func magazzino() {
	mascherine := make([]int, 2) // numero di mascherine attualmente presenti
	mascherine[F] = NF
	mascherine[C] = NC

	prel_in_corso := make([]int, 2) // numero di addetti che stanno prelevando (inizializzati allo zero value - 0)
	rif_in_corso := make([]bool, 2) // numero di fornitori che stanno rifornendo (inizializzati allo zero value - false)

	fmt.Printf("[MAGAZZINO] Inizio. Stato magazzino: MC: %d/%d, MF: %d/%d\n", mascherine[C], NC, mascherine[F], NF)
	for {
		select {
		// PRELIEVO: Inizio (Richiesta)
		case r := <-when( // Condizioni
			((LF*prel_in_corso[F])+LF <= mascherine[F]) &&
				!rif_in_corso[F] &&
				len(prelievo[M]) == 0,
			prelievo[F]): // Canale
			// Occupazione risorsa
			prel_in_corso[F]++
			fmt.Printf("[MAGAZZINO] L'addetto %d inizia il prelievo di %d scatole di FFP2...\n", r.id, LF)
			//debug(prel_in_corso[F], prel_in_corso[C], len(prelievo[F]), len(prelievo[C]), len(prelievo[M]))
			r.ack <- 1
		case r := <-when( // Condizioni
			((LC*prel_in_corso[C])+LC <= mascherine[C]) &&
				!rif_in_corso[C] &&
				len(prelievo[M]) == 0 && len(prelievo[F]) == 0,
			prelievo[C]): // Canale
			// Occupazione risorsa
			prel_in_corso[C]++
			fmt.Printf("[MAGAZZINO] L'addetto %d inizia il prelievo di %d scatole di Chirurgiche...\n", r.id, LC)
			//debug(prel_in_corso[F], prel_in_corso[C], len(prelievo[F]), len(prelievo[C]), len(prelievo[M]))
			r.ack <- 1
		case r := <-when( // Condizioni
			((LF*prel_in_corso[F])+LM <= mascherine[F] &&
				(LC*prel_in_corso[C])+LM <= mascherine[C]) &&
				!rif_in_corso[F] && !rif_in_corso[C],
			prelievo[M]): // Canale
			// Occupazione risorsa
			prel_in_corso[C]++
			prel_in_corso[F]++
			fmt.Printf("[MAGAZZINO] L'addetto %d inizia il prelievo di %d scatole di FFP2 e %d di Chirurgiche...\n", r.id, LM, LM)
			//debug(prel_in_corso[F], prel_in_corso[C], len(prelievo[F]), len(prelievo[C]), len(prelievo[M]))
			r.ack <- 1

		// PRELIEVO: Fine (Rilascio)
		case r := <-finePrelievo:
			switch r.tipo { // sottraggo le mascherine prelevate
			case F:
				mascherine[F] -= LF
				prel_in_corso[F]-- // rilascio risorsa
			case C:
				mascherine[C] -= LC
				prel_in_corso[C]-- // rilascio risorsa
			case M:
				mascherine[F] -= LM
				mascherine[C] -= LM
				prel_in_corso[F]-- // rilascio risorsa
				prel_in_corso[C]-- // rilascio risorsa
			default:
				fmt.Printf("[MAGAZZINO] ERRORE: tipo non valido.\n\n")
			}
			fmt.Printf("[MAGAZZINO] L'addetto %d ha terminato il prelievo di %s. Stato magazzino: MC: %d/%d, MF: %d/%d\n", r.id, getTipo((r.tipo)), mascherine[C], NC, mascherine[F], NF)

		// RIFORNIMENTO: Inizio (Richiesta)
		case r := <-when( // Condizioni
			prel_in_corso[F] == 0 &&
				(mascherine[F] <= mascherine[C] || len(rifornimento[C]) == 0),
			rifornimento[F]): // Canale
			// Occupazione risorsa
			rif_in_corso[F] = true
			fmt.Printf("[MAGAZZINO] Rifornimento %s in corso...\n", strings.ToUpper(getTipo(r.tipo)))
			//debug(prel_in_corso[F], prel_in_corso[C], len(prelievo[F]), len(prelievo[C]), len(prelievo[M]))
			r.ack <- 1
		case r := <-when( // Condizioni
			prel_in_corso[C] == 0 &&
				(mascherine[C] < mascherine[F] || len(rifornimento[F]) == 0),
			rifornimento[C]): // Canale
			// Occupazione risorsa
			rif_in_corso[C] = true
			fmt.Printf("[MAGAZZINO] Rifornimento %s in corso...\n", strings.ToUpper(getTipo(r.tipo)))
			//debug(prel_in_corso[F], prel_in_corso[C], len(prelievo[F]), len(prelievo[C]), len(prelievo[M]))
			r.ack <- 1

		// RIFORNIMENTO: Fine (Rilascio)
		case r := <-fineRifornimento:
			fmt.Printf("[MAGAZZINO] Rifornimento %s concluso. Stato magazzino: MC: %d/%d, MF: %d/%d\n", strings.ToUpper(getTipo(r.tipo)), mascherine[C], NC, mascherine[F], NF)
			if r.tipo == C {
				mascherine[C] = NC
			} else if r.tipo == F {
				mascherine[F] = NF
			} else { // errore (tipo non valido)
				r.ack <- -1
			}
			rif_in_corso[r.tipo] = false // rilascio risorsa
			r.ack <- 1

		// TERMINAZIONE
		case <-terminaMagazzino:
			fmt.Printf("[MAGAZZINO] Termino\n")
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

	nAddetti := 10
	nFornitori := 2

	fmt.Printf("\n[MAIN] Quanti Addetti Reparto (max %d)? ", MAXAR)
	fmt.Scanf("%d\n", &nAddetti)

	if nAddetti < 2 {
		fmt.Printf("[MAIN] Numero addetti troppo basso, ne verranno usati 4.\n")
		nAddetti = 4
	}

	// Inizializzazione canali
	for i := 0; i < len(prelievo); i++ {
		prelievo[i] = make(chan Richiesta, MAXBUFF)
	}
	for i := 0; i < len(rifornimento); i++ {
		rifornimento[i] = make(chan Richiesta, MAXBUFF)
	}

	// ESECUZIONE GOROUTINE
	// Magazzino
	go magazzino()
	// Fornitori
	for i := 0; i < nFornitori; i++ {
		go fornitore(i)
	}
	// Addetti Reparto (AR)
	for i := 0; i < nAddetti; i++ {
		go addetto(i)
	}

	// JOIN GOROUTINE
	// Addetti Reparto
	for i := 0; i < nAddetti; i++ {
		<-done
	}
	// Fornitori
	terminaFornitore <- true
	terminaFornitore <- true
	<-done
	<-done

	terminaMagazzino <- true
	<-done

	fmt.Printf("\n\n[MAIN] Fine\n")
}
