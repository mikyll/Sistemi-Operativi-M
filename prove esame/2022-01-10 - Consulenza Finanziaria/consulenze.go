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

const AMMINISTRATORE = 0    // Amministratore
const PROPRIETARIO_SOLO = 1 // Proprietario Solo
const PROPRIETARIO_ACC = 2  // Proprietario Accompagnato

const SUPERBONUS = 0
const ALTRO = 1

const NU = 5   // Numero Uffici della filiale
const MAXS = 3 // Dimensione Sala d'Aspetto

// ==================================================================
// STRUTTURE DATI ===================================================
type Richiesta struct {
	t   int       // tipo utente
	ack chan bool // canale di ACK
}

/*
NB: la scelta del bool invece di un int è stata fatta in quanto il
boolean in Go ha dimensione di un solo byte, mentre un int32 può
avere da 4 a 8 byte di dimensione (in base al tipo dell'architettura,
32bit o 64); e in questo caso, non era necessario restituire più di 2
valori differenti all'utente.
*/

// ==================================================================
// CANALI ===========================================================
var entrataSA [3]chan Richiesta      // Ingresso Sala d'Aspetto
var entrataUfficio [2]chan Richiesta // Ingresso Ufficio Consulente
var uscita = make(chan int)          // Uscita dalla filiale

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
func getTipoUtente(t int) string {
	if t == AMMINISTRATORE {
		return "amministratore di condominio"
	} else if t == PROPRIETARIO_SOLO {
		return "proprietario solo"
	} else if t == PROPRIETARIO_ACC {
		return "proprietario accompagnato"
	} else {
		return ""
	}
}
func getTipoFinanziamento(t int) string {
	if t == SUPERBONUS {
		return "superbonus 110%"
	} else if t == ALTRO {
		return "altro"
	} else {
		return ""
	}
}

// ==================================================================
// GOROUTINE ========================================================
func utente(id int) {
	fmt.Printf("[UTENTE %d] Inizio\n", id)

	tipo_u := rand.Intn(3) // tipo utente (Amministratore, Proprietario Solo o Proprietario Accompagnato)
	tipo_f := rand.Intn(2) // tipo finanziamento (Superbonus o Altro)

	r := Richiesta{t: tipo_u, ack: make(chan bool)}

	// accede al servizio senza prenotazione
	fmt.Printf("[UTENTE %d] Sono un %s, voglio entrare in sala d'aspetto\n", id, getTipoUtente(tipo_u))
	// 1. Entra nella sala d'aspetto (se c'è spazio)
	entrataSA[tipo_u] <- r
	<-r.ack

	fmt.Printf("[UTENTE %d] Sono attendo il mio turno\n", id)
	sleepRandTime(2)

	// 2. Entra nell'ufficio di un consulente
	fmt.Printf("[UTENTE %d] Voglio una consulenza finanziaria di tipo %s\n", id, strings.ToTitle(getTipoFinanziamento(tipo_f)))
	entrataUfficio[tipo_f] <- r
	<-r.ack

	// 3. Rimane nell'ufficio per tutto il tempo necessario alla consulenza
	sleepRandTime(4) // consulenza in corso

	// 4. Esce dalla filiale
	fmt.Printf("[UTENTE %d] Ho ottenuto il finanziamento di tipo %s\n", id, strings.ToTitle(getTipoFinanziamento(tipo_f)))
	uscita <- tipo_u

	fmt.Printf("[UTENTE %d] Termino\n", id)
	done <- true
}

func filiale() {
	sala := 0           // posti occupati in sala d'aspetto
	ufficiOccupati := 0 // numero uffici occupati

	// EXTRA
	superbonus := 0 // consulenze di tipo superbonus effettuate
	altro := 0      // consulenze di tipo altro effettuate

	fmt.Printf("[FILIALE] Inizio\n")
	for {
		select {
		case r := <-when(sala < MAXS, entrataSA[AMMINISTRATORE]):
			// occupa il/i posto/i in sala
			switch {
			case r.t == AMMINISTRATORE || r.t == PROPRIETARIO_SOLO:
				sala++
				r.ack <- true
				fmt.Printf("[FILIALE] Un %s è entrato in sala d'aspetto. Stato sala: %d/%d\n", getTipoUtente(r.t), sala, MAXS)
			case r.t == PROPRIETARIO_ACC:
				sala += 2
				r.ack <- true
				fmt.Printf("[FILIALE] Un %s è entrato in sala d'aspetto. Stato sala: %d/%d\n", getTipoUtente(r.t), sala, MAXS)
			default:
				r.ack <- false
				fmt.Printf("[FILIALE] Tipo utente %d non valido!\n", r.t)
			}
		case r := <-when(sala < MAXS && len(entrataSA[AMMINISTRATORE]) == 0, entrataSA[PROPRIETARIO_SOLO]):
			// occupa il/i posto/i in sala
			switch {
			case r.t == AMMINISTRATORE || r.t == PROPRIETARIO_SOLO:
				sala++
				r.ack <- true
				fmt.Printf("[FILIALE] Un %s è entrato in sala d'aspetto. Stato sala: %d/%d\n", getTipoUtente(r.t), sala, MAXS)
			case r.t == PROPRIETARIO_ACC:
				sala += 2
				r.ack <- true
				fmt.Printf("[FILIALE] Un %s è entrato in sala d'aspetto. Stato sala: %d/%d\n", getTipoUtente(r.t), sala, MAXS)
			default:
				r.ack <- false
				fmt.Printf("[FILIALE] Tipo utente %d non valido!\n", r.t)
			}
		case r := <-when(sala+1 < MAXS && (len(entrataSA[AMMINISTRATORE]) == 0 && len(entrataSA[PROPRIETARIO_SOLO]) == 0), entrataSA[PROPRIETARIO_ACC]):
			// occupa il/i posto/i in sala
			switch {
			case r.t == AMMINISTRATORE || r.t == PROPRIETARIO_SOLO:
				sala++
				r.ack <- true
				fmt.Printf("[FILIALE] Un %s è entrato in sala d'aspetto. Stato sala: %d/%d\n", getTipoUtente(r.t), sala, MAXS)
			case r.t == PROPRIETARIO_ACC:
				sala += 2
				r.ack <- true
				fmt.Printf("[FILIALE] Un %s è entrato in sala d'aspetto. Stato sala: %d/%d\n", getTipoUtente(r.t), sala, MAXS)
			default:
				r.ack <- false
				fmt.Printf("[FILIALE] Tipo utente %d non valido!\n", r.t)
			}
		case r := <-when(ufficiOccupati < NU, entrataUfficio[SUPERBONUS]):
			// libera il/i posto/i in sala d'aspetto e occupa un ufficio
			switch {
			case r.t == AMMINISTRATORE || r.t == PROPRIETARIO_SOLO:
				sala--
				ufficiOccupati++
				r.ack <- true
				fmt.Printf("[FILIALE] Inizio consulenza di tipo \"Superbonus %%\". Stato uffici: %d/%d\n", ufficiOccupati, NU)
				superbonus++
			case r.t == PROPRIETARIO_ACC:
				sala -= 2
				ufficiOccupati++
				r.ack <- true
				fmt.Printf("[FILIALE] Inizio consulenza di tipo \"Superbonus %%\". Stato uffici: %d/%d\n", ufficiOccupati, NU)
				superbonus++
			default:
				r.ack <- false
				fmt.Printf("[FILIALE] Tipo utente %d non valido!\n", r.t)
			}
		case r := <-when(ufficiOccupati < NU && len(entrataUfficio[SUPERBONUS]) == 0, entrataUfficio[ALTRO]):
			// libera il/i posto/i in sala d'aspetto e occupa un ufficio
			switch {
			case r.t == AMMINISTRATORE || r.t == PROPRIETARIO_SOLO:
				sala--
				ufficiOccupati++
				r.ack <- true
				fmt.Printf("[FILIALE] Inizio consulenza di tipo altro...\n")
				altro++
			case r.t == PROPRIETARIO_ACC:
				sala -= 2
				ufficiOccupati++
				r.ack <- true
				fmt.Printf("[FILIALE] Inizio consulenza di tipo altro...\n")
				altro++
			default:
				r.ack <- false
				fmt.Printf("[FILIALE] Tipo utente %d non valido!\n", r.t)
			}
		case <-uscita:
			// libera l'ufficio
			ufficiOccupati--
			fmt.Printf("[FILIALE] Utente uscito dalla filiale. Stato uffici: %d/%d\n", ufficiOccupati, NU)
		case <-termina:
			fmt.Printf("[FILIALE] Termino. Consulenze effettuate: %d (%d di tipo \"Superbonus 110%%\", %d di altro tipo)\n", superbonus+altro, superbonus, altro)
			done <- true
			return
		}
	}
}

// ==================================================================
// MAIN =============================================================
func main() {
	/*var boolean bool
	var integer int
	fmt.Println(unsafe.Sizeof(boolean))
	fmt.Println(unsafe.Sizeof(integer))
	return*/
	fmt.Printf("[MAIN] Inizio\n\n")
	rand.Seed(time.Now().UnixNano())

	/*var V1 int
	var V2 int

	fmt.Printf("\nQuanti Thread tipo1 (max %d)? ", MAXPROC)
	fmt.Scanf("%d\n", &V1)
	fmt.Printf("\nQuanti Thread tipo2 (max %d)? ", MAXPROC)
	fmt.Scanf("%d\n", &V2)*/

	nUtenti := 10

	// Inizializzazione canali
	for i := 0; i < len(entrataSA); i++ {
		entrataSA[i] = make(chan Richiesta, MAXBUFF)
	}
	for i := 0; i < len(entrataUfficio); i++ {
		entrataUfficio[i] = make(chan Richiesta, MAXBUFF)
	}

	// Esecuzione goroutine
	// Esecuzione Utenti
	for i := 0; i < nUtenti; i++ {
		go utente(i)
	}
	go filiale()

	// Join goroutine
	for i := 0; i < nUtenti; i++ {
		<-done
	}

	// Terminazione Filiale
	termina <- true
	<-done
	fmt.Printf("\n\n[MAIN] Fine\n")
}
