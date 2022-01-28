package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAXBUFF int = 100
const MAX int = 18
const N_COMMESSI int = 8
const N_CLIENTI int = 70
const NM = 10

const ABITUALE int = 0
const OCCASIONALE int = 1

var tipoClienteStr [2]string = [2]string{"ABITUALE", "OCCASIONALE"}

type Richiesta struct {
	id  int
	ack chan bool
}

type Commesso struct {
	dentro                 bool
	vuoleUscire            bool
	clientiAssegnati       [3]int
	numeroClientiAssegnati int
	ackUscita              chan bool
}

func whenRichiesta(b bool, c chan Richiesta) chan Richiesta {
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

func sleepRandTime(timeLimit int) {
	if timeLimit > 0 {
		time.Sleep(time.Duration(rand.Intn(timeLimit)+1) * time.Second)
	}
}

func cliente(id int, tipo int, entra chan Richiesta, esci chan int, termina chan bool) {

	var ric Richiesta
	ric.id = id
	ric.ack = make(chan bool, MAXBUFF)

	sleepRandTime(5)

	fmt.Printf("[CLIENTE %s %d] Voglio entrare nel negozio...\n", tipoClienteStr[tipo], id)

	entra <- ric
	<-ric.ack

	fmt.Printf("[CLIENTE %s %d] Sono entrato nel negozio...\n", tipoClienteStr[tipo], id)

	sleepRandTime(7)

	esci <- id
	fmt.Printf("[CLIENTE %s %d] Sono uscito dal negozio...\n", tipoClienteStr[tipo], id)

	fmt.Printf("[CLIENTE %s %d] Termino...\n", tipoClienteStr[tipo], id)

	termina <- true

	return

}

func commesso(id int, entra chan Richiesta, esci chan Richiesta, termina chan bool, done chan bool) {
	var ric Richiesta
	ric.id = id
	ric.ack = make(chan bool, MAXBUFF)

	for true {

		sleepRandTime(5)

		fmt.Printf("[COMMESSO %d] Voglio entrare nel negozio...\n", id)

		entra <- ric
		<-ric.ack

		fmt.Printf("[COMMESSO %d] Sono entrato nel negozio...\n", id)

		sleepRandTime(9)

		esci <- ric
		<-ric.ack

		fmt.Printf("[COMMESSO %d] Sono uscito dal negozio...\n", id)

		select {
		case <-termina:
			{
				fmt.Printf("[COMMESSO %d] Termino...\n", id)
				done <- true
				return
			}
		default:
			{
				// non terminare
				sleepRandTime(2)
			}
		}
	}

}

func fornitore(deposita chan bool, termina chan bool) {

	for true {
		sleepRandTime(5)

		fmt.Printf("[FORNITORE] Voglio consegnare il lotto di mascherine...\n")
		deposita <- true
		<-deposita
		fmt.Printf("[FORNITORE] Consegna effettuata...\n")

		select {

		case <-termina:
			{
				fmt.Printf("[FORNITORE] Termino...\n")
				termina <- true
				return
			}
		default:
			{
				// non terminare
				sleepRandTime(2)
			}
		}
	}

}

func negozio(entraClienteAbituale chan Richiesta, entraClienteOccasionale chan Richiesta, entraCommesso chan Richiesta, esciCliente chan int, esciCommesso chan Richiesta, deposita chan bool, termina chan bool) {

	clientiDentro := 0
	commessiDentro := 0
	commessiLiberi := 0 // commessi che sono dentro al negozio e stanno supervisionando da 0 a 2 clienti [0,2]
	commessi := make([]Commesso, N_COMMESSI)

	for i := 0; i < N_COMMESSI; i++ {
		commessi[i].dentro = false
		commessi[i].vuoleUscire = false
		commessi[i].numeroClientiAssegnati = 0
		commessi[i].ackUscita = nil
		for j := 0; j < 3; j++ {
			commessi[i].clientiAssegnati[j] = -1
		}
	}

	mascherine := 0

	fmt.Printf("MAX: %d, NM: %d, N_CLIENTI: %d, N_COMMESSI: %d...\n", MAX, NM, N_CLIENTI, N_COMMESSI)

	for true {
		fmt.Printf("[NEGOZIO] ClientiDentro: %d, CommessiDentro: %d, CommessiLiberi: %d, Mascherine: %d...\n", clientiDentro, commessiDentro, commessiLiberi, mascherine)
		select {
		case <-deposita:
			{
				mascherine += NM
				fmt.Printf("[NEGOZIO] Il fornitore ha depositato %d mascherine...\n", NM)
				deposita <- true
			}
		case ric := <-whenRichiesta(clientiDentro+commessiDentro < MAX, entraCommesso):
			{
				commessiDentro++
				commessiLiberi++
				commessi[ric.id].dentro = true
				commessi[ric.id].vuoleUscire = false
				commessi[ric.id].numeroClientiAssegnati = 0
				for i := 0; i < 3; i++ {
					commessi[ric.id].clientiAssegnati[i] = -1
				}
				fmt.Printf("[NEGOZIO] Il commesso %d entra nel negozio...\n", ric.id)
				ric.ack <- true
			}
		case ric := <-esciCommesso:
			{

				if commessi[ric.id].numeroClientiAssegnati == 0 {

					// il commesso puo' uscire
					fmt.Printf("[NEGOZIO] Il commesso %d esce dal negozio...\n", ric.id)
					commessi[ric.id].dentro = false
					commessi[ric.id].vuoleUscire = false
					commessi[ric.id].ackUscita = nil
					ric.ack <- true
					commessiLiberi--
					commessiDentro--
				} else {
					// il commesso attende
					fmt.Printf("[NEGOZIO] Il commesso %d è in attesa di uscire dal negozio (%d clienti assegnati)...\n", ric.id, commessi[ric.id].numeroClientiAssegnati)
					commessi[ric.id].vuoleUscire = true
					commessi[ric.id].ackUscita = ric.ack
				}

			}
		case ric := <-whenRichiesta(commessiDentro > 0 && commessiLiberi > 0 && mascherine >= 1 && len(entraCommesso) == 0 && clientiDentro+commessiDentro < MAX, entraClienteAbituale):
			{
				found := false
				for i := 0; i < N_COMMESSI && !found; i++ {
					if commessi[i].dentro && commessi[i].numeroClientiAssegnati < 3 {
						for j := 0; j < 3 && !found; j++ {
							if commessi[i].clientiAssegnati[j] < 0 {
								commessi[i].clientiAssegnati[j] = ric.id // assegno l'id del cliente
								commessi[i].numeroClientiAssegnati++
								if commessi[i].numeroClientiAssegnati == 3 {
									commessiLiberi--
								}
								clientiDentro++
								mascherine--
								found = true
								ric.ack <- true
								fmt.Printf("[NEGOZIO] Il cliente ABITUALE %d entra nel negozio...\n", ric.id)
								fmt.Printf("[NEGOZIO] Assegno il commesso %d al cliente ABITUALE %d...\n", i, ric.id)
							}
						}
					}
				}

				if !found { // solo per debug
					fmt.Printf("[DEBUG NEGOZIO] Non riesco a trovare un commesso disponibile...\n")
				}
			}
		case ric := <-whenRichiesta(len(entraClienteAbituale) == 0 && commessiDentro > 0 && commessiLiberi > 0 && mascherine >= 1 && len(entraCommesso) == 0 && clientiDentro+commessiDentro < MAX, entraClienteOccasionale):
			{
				found := false
				for i := 0; i < N_COMMESSI && !found; i++ {
					if commessi[i].dentro && commessi[i].numeroClientiAssegnati < 3 {
						for j := 0; j < 3 && !found; j++ {
							if commessi[i].clientiAssegnati[j] < 0 {
								commessi[i].clientiAssegnati[j] = ric.id // assegno l'id del cliente
								commessi[i].numeroClientiAssegnati++
								if commessi[i].numeroClientiAssegnati == 3 {
									// il commesso ha raggiunto il limite di 3 cliente da supervisionare
									commessiLiberi--
								}
								clientiDentro++
								mascherine--
								found = true
								ric.ack <- true
								fmt.Printf("[NEGOZIO] Il cliente OCCASIONALE %d entra nel negozio...\n", ric.id)
								fmt.Printf("[NEGOZIO] Assegno il commesso %d al cliente OCCASIONALE %d...\n", i, ric.id)
							}
						}
					}
				}

				if !found { // solo per debug
					fmt.Printf("[DEBUG NEGOZIO] Non riesco a trovare un commesso disponibile...\n")
				}
			}
		case id := <-esciCliente:
			{
				found := false
				for i := 0; i < N_COMMESSI && !found; i++ {
					if commessi[i].dentro {
						for j := 0; j < 3 && !found; j++ {
							if commessi[i].clientiAssegnati[j] == id {

								commessi[i].clientiAssegnati[j] = -1 // reset id

								if commessi[i].numeroClientiAssegnati == 3 {
									// il commesso stava supervisionando 3 clienti, adesso non piu'
									commessiLiberi++
								}
								commessi[i].numeroClientiAssegnati--
								clientiDentro--
								found = true
								fmt.Printf("[NEGOZIO] Il cliente %d esce dal negozio...\n", id)
								fmt.Printf("[NEGOZIO] Libero il commesso %d dalla supervisione del cliente %d...\n", i, id)
								//verifico se il commesso è in attesa di uscire:
								if commessi[i].dentro && commessi[i].vuoleUscire && commessi[i].numeroClientiAssegnati == 0 {
									// il commesso puo' uscire
									fmt.Printf("[NEGOZIO] Il commesso %d esce dal negozio...\n", i)
									commessi[i].dentro = false
									commessi[i].vuoleUscire = false
									commessi[i].ackUscita <- true
									commessi[i].ackUscita = nil
									for j := 0; j < 3; j++ {
										commessi[i].clientiAssegnati[j] = -1
									}
									commessiLiberi--
									commessiDentro--
								}

							}
						}
					}
				}

			}
		case <-termina:
			{
				fmt.Printf("[NEGOZIO] Termino...\n")
				termina <- true
				return
			}
		}

	}

}

func main() {
	entraClienteAbituale := make(chan Richiesta, MAXBUFF)
	entraClienteOccasionale := make(chan Richiesta, MAXBUFF)
	entraCommesso := make(chan Richiesta, MAXBUFF)
	esciCliente := make(chan int)
	esciCommesso := make(chan Richiesta)

	deposita := make(chan bool)

	terminaCliente := make(chan bool)
	terminaCommesso := make([]chan bool, N_COMMESSI)
	done := make(chan bool)
	terminaFornitore := make(chan bool)
	terminaNegozio := make(chan bool)

	rand.Seed(time.Now().Unix())

	for i := 0; i < N_CLIENTI; i++ {

		var tipo int
		var entraCliente chan Richiesta

		if rand.Intn(100) > 70 {
			tipo = ABITUALE
			entraCliente = entraClienteAbituale
		} else {
			tipo = OCCASIONALE
			entraCliente = entraClienteOccasionale
		}
		go cliente(i, tipo, entraCliente, esciCliente, terminaCliente)

	}

	for i := 0; i < N_COMMESSI; i++ {
		terminaCommesso[i] = make(chan bool, MAXBUFF)
		go commesso(i, entraCommesso, esciCommesso, terminaCommesso[i], done)
	}

	go fornitore(deposita, terminaFornitore)

	go negozio(entraClienteAbituale, entraClienteOccasionale, entraCommesso, esciCliente, esciCommesso, deposita, terminaNegozio)

	for i := 0; i < N_CLIENTI; i++ {
		<-terminaCliente
	}

	terminaFornitore <- true
	<-terminaFornitore

	for i := 0; i < N_COMMESSI; i++ {
		// prima invio il messaggio di terminazione a tutti i commessi sul canale bufferizzato
		terminaCommesso[i] <- true
	}

	for i := 0; i < N_COMMESSI; i++ {
		// dopo aspetto le N terminazioni sul canale sincrono
		<-done
	}

	terminaNegozio <- true
	<-terminaNegozio

}
