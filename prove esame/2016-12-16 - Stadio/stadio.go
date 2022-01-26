package main

import (
	"fmt"
	"math/rand"
	"time"
)

const LOCALI = 0
const OSPITI = 1
const NUM = 5 // Operatori controllo sicurezza

const MAXPROC = 50 // numero spettatori

type Richiesta struct {
	ID    int
	soldi float64
	ack   chan int
}

var richiesta_acquisto = make(chan Richiesta)
var richiesta_operatore [2]chan Richiesta
var rilascio_operatore [2]chan int

// goroutine join
var terminateServer = make(chan bool)
var done = make(chan bool)

// aux
func when(b bool, c chan Richiesta) chan Richiesta {
	if !b {
		return nil
	}
	return c
}
func sleepRandTime(timeLimit int) {
	if timeLimit > 0 {
		time.Sleep(time.Duration(rand.Intn(timeLimit)+1) * time.Second)
	}
}
func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}
func getTribuna(t int) string {
	switch t {
	case LOCALI:
		return "LOCALI"
	case OSPITI:
		return "OSPITI"
	default:
		return ""
	}
}

func spettatore(ID int) {
	r := Richiesta{ID: ID, soldi: randFloats(1.0, 100.0, 1)[0], ack: make(chan int)}

	fmt.Printf("[Spettatore %d] Voglio acquistare un biglietto per %.2f€\n", r.ID, r.soldi)

	// 1. Acquisto biglietto
	// Richiesta acquisto biglietto
	richiesta_acquisto <- r

	// 2. ACK (# tribuna)
	tribuna := <-r.ack

	// 3. Recarsi al varco
	sleepRandTime(3)

	// 4. Ingresso controllo di sicurezza
	// Richiesta risorsa operatore
	richiesta_operatore[tribuna] <- r

	<-r.ack

	// 5. Controllo di Sicurezza
	fmt.Printf("[Spettatore %d] Inizio il controllo di sicurezza per la tribuna %s\n", r.ID, getTribuna(tribuna))
	sleepRandTime(5)

	// 6. Uscita controllo di sicurezza
	// Rilascio risorsa operatore
	rilascio_operatore[tribuna] <- ID

	done <- true
}

func biglietteria() {
	cassa := 0.0
	var tribuna int

	fmt.Println("[Biglietteria] Inizio")
	for {
		select {
		// 1. Riceve richiesta biglietto
		case r := <-richiesta_acquisto:
			cassa += float64(r.soldi)
			// 2. Assegna biglietto per tribuna a caso
			tribuna = rand.Intn(2)
			fmt.Printf("[Biglietteria] Allo spettatore %d è stata assegnata la tribuna %s\n", r.ID, getTribuna(tribuna))
			r.ack <- tribuna
		case <-terminateServer:
			fmt.Printf("[Biglietteria] Ho finito. Abbiamo raccolto %.2f€\n", cassa)
			done <- true
			return
		}
	}
}

func stadio() {
	tL, tO := 0, 0 // counter spettatori in tribuna Locali, tribuna Ospiti
	oo := 0        // operatori occupati

	fmt.Println("[Stadio] Inizio")
	for {
		select {
		case r := <-when(oo < NUM &&
			(tL <= tO || len(richiesta_operatore[OSPITI]) == 0), richiesta_operatore[LOCALI]):
			fmt.Printf("[Stadio] Lo spettatore %d ha terminato il controllo (ospiti)\n", r.ID)
			r.ack <- 1
		case r := <-when(oo < NUM &&
			(tL >= tO || len(richiesta_operatore[LOCALI]) == 0), richiesta_operatore[OSPITI]):
			fmt.Printf("[Stadio] Lo spettatore %d ha terminato il controllo (ospiti)\n", r.ID)
			r.ack <- 1
		case id := <-rilascio_operatore[LOCALI]:
			fmt.Printf("[Stadio] Lo spettatore %d ha terminato il controllo (locali)\n", id)
			tL++
			oo--
		case id := <-rilascio_operatore[OSPITI]:
			fmt.Printf("[Stadio] Lo spettatore %d ha terminato il controllo (ospiti)\n", id)
			tO++
			oo--
		// when operatori liberi == 0 => attesa
		case <-terminateServer:
			fmt.Println("[Stadio] Termino")
			done <- true
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var n int
	richiesta_operatore[LOCALI] = make(chan Richiesta)
	richiesta_operatore[OSPITI] = make(chan Richiesta)
	rilascio_operatore[LOCALI] = make(chan int)
	rilascio_operatore[OSPITI] = make(chan int)

	fmt.Printf("\nQuanti spettatori (max %d)? ", MAXPROC)
	fmt.Scanf("%d", &n)

	// avvio goroutines
	for i := 0; i < n; i++ {
		go spettatore(i)
	}
	go biglietteria()
	go stadio()

	// terminazione spettatori
	for i := 0; i < n; i++ {
		<-done
	}

	terminateServer <- true
	<-done
	terminateServer <- true
	<-done
}
