/*Per i voli di lunga percorrenza la compagnia aerea AIRBOLO usa aerei del tipo “jumbo jet”. In questi aerei la cabina è strutturata in due zone:
• un ponte superiore, nel quale sono collocati i posti di prima classe ed un salottino-bar;
• un ponte inferiore, dove sono collocati i posti di classe economy e le toilette.
Dal ponte superiore si può accedere a quello inferiore, e viceversa, tramite una scala a chiocciola, che viene utilizzata a senso unico alternato: ad ogni istante la scala può essere occupata da persone che la percorrono nella stessa direzione.
Si assuma che:
• ilpontesuperioreabbiacapacitàmassimapariaTSpersone;
• ilponteinferioreabbiacapacitàmassimapariaTIpersone;
• l’aereo non parta mai completamente pieno (cioè, il numero di
persone imbarcate sia sempre minore del valore TI+TS). Le persone imbarcate sull’aereo si classificano in due categorie:
• Viaggiatori,
• Equipaggio, cioè le persone di servizio sull’aereo (ad esempio,
hostess, steward, ecc.).
Entrambi i tipi di persone possono spostarsi liberamente da un ponte all’altro.
Realizzare un’applicazione ADA nella quale Scala, Viaggiatori e membri dell’Equipaggio siano rappresentati da TASK distinti.
La politica di controllo degli accessi alla scala dovrà favorire le persone provenienti dalla zona più affollata (cioè quella in cui il numero di posti liberi è minore); inoltre, nell’ambito di una stessa direzione, i membri dell’equipaggio dovranno avere la priorità sui viaggiatori.*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const TI = 20
const TS = 20
const MAXBUFF = 100
const S = 1
const I = 0
const L = 2
const MAXVIAGGSUP = 12
const MAXEQUIPSUP = 8
const MAXEQUIPINF = 8
const MAXVIAGGINF = 12

var done = make(chan bool)
var termina = make(chan bool)

var viaggiatore_superiore_sale_scala = make(chan int, MAXBUFF)
var viaggiatore_superiore_scende_scala = make(chan int, MAXBUFF)
var viaggiatore_inferiore_scende_scala = make(chan int, MAXBUFF)
var viaggiatore_inferiore_sale_scala = make(chan int, MAXBUFF)
var viaggiatore_superiore_scala_salita_ok = make(chan int, MAXBUFF)
var viaggiatore_superiore_scala_discesa_ok = make(chan int, MAXBUFF)
var viaggiatore_inferiore_scala_salita_ok = make(chan int, MAXBUFF)
var viaggiatore_inferiore_scala_discesa_ok = make(chan int, MAXBUFF)

var equipaggio_superiore_sale_scala = make(chan int, MAXBUFF)
var equipaggio_superiore_scende_scala = make(chan int, MAXBUFF)
var equipaggio_inferiore_scende_scala = make(chan int, MAXBUFF)
var equipaggio_inferiore_sale_scala = make(chan int, MAXBUFF)
var equipaggio_superiore_scala_salita_ok = make(chan int, MAXBUFF)
var equipaggio_superiore_scala_discesa_ok = make(chan int, MAXBUFF)
var equipaggio_inferiore_scala_salita_ok = make(chan int, MAXBUFF)
var equipaggio_inferiore_scala_discesa_ok = make(chan int, MAXBUFF)

var ACK_SALITA_VIAGGIATORE_SUP [TS]chan int
var ACK_DISCESA_VIAGGIATORE_SUP [TS]chan int
var ACK_SALITA_VIAGGIATORE_INF [TI]chan int
var ACK_DISCESA_VIAGGIATORE_INF [TI]chan int
var ACK_SALITA_EQUIP_SUP [TS]chan int
var ACK_DISCESA_EQUIP_SUP [TS]chan int
var ACK_SALITA_EQUIP_INF [TI]chan int
var ACK_DISCESA_EQUIP_INF [TI]chan int

func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}

func viaggiatore(myid int, partenza int) {

	var tt int

	switch partenza {

	case S:

		tt = rand.Intn(5) + 1
		time.Sleep(time.Duration(tt) * time.Second)
		fmt.Printf("\n Il viaggiatore S%d vorrebbe scendere al piano inferiore \n", myid)
		viaggiatore_superiore_scende_scala <- myid
		<-ACK_DISCESA_VIAGGIATORE_SUP[myid]
		tt = rand.Intn(5) + 2
		fmt.Printf("\n Il viaggiatore S%d sta scendendo la scala, ci impiegherà %d secondi \n", myid, tt)
		time.Sleep(time.Duration(tt) * time.Second)
		viaggiatore_superiore_scala_discesa_ok <- myid
		fmt.Printf("\n Il viaggiatore S%d è ora nel piano inferiore \n", myid)
		tt = rand.Intn(5) + 3
		time.Sleep(time.Duration(tt) * time.Second)
		fmt.Printf("\n Il viaggiatore S%d ha finito la visita al piano inferiore e vorrebbe ora risalire \n", myid)
		viaggiatore_superiore_sale_scala <- myid
		<-ACK_SALITA_VIAGGIATORE_SUP[myid]
		tt = rand.Intn(5) + 2
		fmt.Printf("\n Il viaggiatore S%d sta salendo la scala, ci impiegherà %d secondi \n", myid, tt)
		time.Sleep(time.Duration(tt) * time.Second)
		viaggiatore_superiore_scala_salita_ok <- myid
		fmt.Printf("\n Il viaggiatore S%d è salito nuovamente al piano superiore, ora può terminare\n", myid)

	case I:

		tt = rand.Intn(5) + 1
		time.Sleep(time.Duration(tt) * time.Second)
		fmt.Printf("\n Il viaggiatore I%d vorrebbe salire al piano superiore \n", myid)
		viaggiatore_inferiore_sale_scala <- myid
		<-ACK_SALITA_VIAGGIATORE_INF[myid]
		tt = rand.Intn(5) + 2
		fmt.Printf("\n Il viaggiatore I%d sta salendo la scala, ci impiegherà %d secondi \n", myid, tt)
		time.Sleep(time.Duration(tt) * time.Second)
		viaggiatore_inferiore_scala_salita_ok <- myid
		fmt.Printf("\n Il viaggiatore I%d è ora nel piano superiore \n", myid)
		tt = rand.Intn(5) + 3
		time.Sleep(time.Duration(tt) * time.Second)
		fmt.Printf("\n Il viaggiatore I%d ha finito la visita al piano superiore e vorrebbe ora riscendere \n", myid)
		viaggiatore_inferiore_scende_scala <- myid
		<-ACK_DISCESA_VIAGGIATORE_INF[myid]
		tt = rand.Intn(5) + 2
		fmt.Printf("\n Il viaggiatore I%d sta scendendo la scala, ci impiegherà %d secondi \n", myid, tt)
		time.Sleep(time.Duration(tt) * time.Second)
		viaggiatore_inferiore_scala_discesa_ok <- myid
		fmt.Printf("\n Il viaggiatore I%d è sceso nuovamente al piano inferiore, ora può terminare\n", myid)

	}

	done <- true
}

func equipaggio(myid int, partenza int) {

	var tt int

	switch partenza {

	case S:

		tt = rand.Intn(5) + 1
		time.Sleep(time.Duration(tt) * time.Second)
		fmt.Printf("\n Il membro dell'equipaggio S%d vorrebbe scendere al piano inferiore \n", myid)
		equipaggio_superiore_scende_scala <- myid
		<-ACK_DISCESA_EQUIP_SUP[myid]
		tt = rand.Intn(5) + 2
		fmt.Printf("\n Il membro dell'equipaggio S%d sta scendendo la scala, ci impiegherà %d secondi \n", myid, tt)
		time.Sleep(time.Duration(tt) * time.Second)
		equipaggio_superiore_scala_discesa_ok <- myid
		fmt.Printf("\n Il membro dell'equipaggio S%d è ora nel piano inferiore \n", myid)
		tt = rand.Intn(5) + 2
		time.Sleep(time.Duration(tt) * time.Second)
		fmt.Printf("\n Il membro dell'equipaggio S%d ha finito la visita al piano inferiore e vorrebbe ora risalire \n", myid)
		equipaggio_superiore_sale_scala <- myid
		<-ACK_SALITA_EQUIP_SUP[myid]
		tt = rand.Intn(5) + 3
		fmt.Printf("\n Il membro dell'equipaggio S%d sta salendo la scala, ci impiegherà %d secondi \n", myid, tt)
		time.Sleep(time.Duration(tt) * time.Second)
		equipaggio_superiore_scala_salita_ok <- myid
		fmt.Printf("\n Il membro dell'equipaggio S%d è salito nuovamente al piano superiore, ora può terminare\n", myid)

	case I:

		tt = rand.Intn(5) + 1
		time.Sleep(time.Duration(tt) * time.Second)
		fmt.Printf("\n Il membro dell'equipaggio I%d vorrebbe salire al piano superiore \n", myid)
		equipaggio_inferiore_sale_scala <- myid
		<-ACK_SALITA_EQUIP_INF[myid]
		tt = rand.Intn(5) + 2
		fmt.Printf("\n Il membro dell'equipaggio I%d sta salendo la scala, ci impiegherà %d secondi \n", myid, tt)
		time.Sleep(time.Duration(tt) * time.Second)
		equipaggio_inferiore_scala_salita_ok <- myid
		fmt.Printf("\n Il membro dell'equipaggio I%d è ora nel piano superiore \n", myid)
		tt = rand.Intn(5) + 3
		time.Sleep(time.Duration(tt) * time.Second)
		fmt.Printf("\n Il membro dell'equipaggio I%d ha finito la visita al piano superiore e vorrebbe ora riscendere \n", myid)
		equipaggio_inferiore_scende_scala <- myid
		<-ACK_DISCESA_EQUIP_INF[myid]
		tt = rand.Intn(5) + 2
		fmt.Printf("\n Il membro dell'equipaggio I%d sta scendendo la scala, ci impiegherà %d secondi \n", myid, tt)
		time.Sleep(time.Duration(tt) * time.Second)
		equipaggio_inferiore_scala_discesa_ok <- myid
		fmt.Printf("\n Il membro dell'equipaggio I%d è sceso nuovamente al piano inferiore, ora può terminare\n", myid)

	}

	done <- true
}

func scala(sup int, inf int) {

	var num_scala_sup int   //numero di persone che percorrono la scala dal piano SUP verso INF
	var num_scala_inf int   //numero di persone che percorrono la scala dal piano INF verso SUP
	var num_persone_sup int //numero di persone che sono sul piano SUP
	var num_persone_inf int //numero di persone che sono sul piano INF
	var num_viagg_sup_term int
	var num_viagg_inf_term int
	var num_equip_sup_term int
	var num_equip_inf_term int

	num_scala_sup = 0
	num_scala_inf = 0
	num_persone_sup = sup
	num_persone_inf = inf
	num_viagg_sup_term = 0
	num_viagg_inf_term = 0
	num_equip_sup_term = 0
	num_equip_inf_term = 0

	for {
		fmt.Printf("\n Ci sono %d persone sulla scala da SUP a INF, %d persone sulla scala da INF a SUP, %d persone al piano sup e %d persone al piano inf\n", num_scala_sup, num_scala_inf, num_persone_sup, num_persone_inf)
		select {

		case x := <-when(((num_persone_inf+num_scala_sup <= num_persone_sup+num_scala_inf) && (num_scala_inf == 0) && (len(equipaggio_superiore_scende_scala) == 0) && (len(equipaggio_inferiore_scende_scala) == 0)), viaggiatore_superiore_scende_scala):
			// se il numero di persone al piano inferiore più quelle che stanno percorrendo la scala verso il piano inferiore è minore del numero di persone che sono al piano superiore aumentato del
			// numero di persone che stanno percorrendo la scala in salita significa che il numero di persone al piano superiore è maggiore del numero di persone del piano inferiroe
			// per cui la politica è rispettata. Se non c'è nessuno nella direzione opposta e nessun membro dell'equipaggio che opera nella stessa direzione allora può passare
			num_scala_sup++
			num_persone_sup--
			ACK_DISCESA_VIAGGIATORE_SUP[x] <- 1

		case x := <-when((((num_persone_sup+num_scala_inf <= num_persone_inf+num_scala_sup) && (num_scala_sup == 0) && (len(equipaggio_superiore_sale_scala) == 0) && (len(equipaggio_inferiore_sale_scala) == 0)) || (num_viagg_inf_term+num_equip_inf_term == inf)), viaggiatore_superiore_sale_scala):
			num_scala_inf++
			num_persone_inf--
			ACK_SALITA_VIAGGIATORE_SUP[x] <- 1

		case x := <-when((((num_persone_inf+num_scala_sup <= num_persone_sup+num_scala_inf) && (num_scala_inf == 0) && (len(equipaggio_superiore_scende_scala) == 0) && (len(equipaggio_inferiore_scende_scala) == 0)) || (num_viagg_sup_term+num_equip_sup_term == sup)), viaggiatore_inferiore_scende_scala):
			num_scala_sup++
			num_persone_sup--
			ACK_DISCESA_VIAGGIATORE_INF[x] <- 1

		case x := <-when(((num_persone_sup+num_scala_inf <= num_persone_inf+num_scala_sup) && (num_scala_sup == 0) && (len(equipaggio_superiore_sale_scala) == 0) && (len(equipaggio_inferiore_sale_scala) == 0)), viaggiatore_inferiore_sale_scala):
			num_scala_inf++
			num_persone_inf--
			ACK_SALITA_VIAGGIATORE_INF[x] <- 1

		case x := <-when(((num_persone_inf+num_scala_sup <= num_persone_sup+num_scala_inf) && (num_scala_inf == 0)), equipaggio_superiore_scende_scala):
			num_scala_sup++
			num_persone_sup--
			ACK_DISCESA_EQUIP_SUP[x] <- 1

		case x := <-when(((num_persone_sup+num_scala_inf <= num_persone_inf+num_scala_sup) && ((num_scala_sup == 0) || (num_viagg_inf_term+num_equip_inf_term == inf))), equipaggio_superiore_sale_scala):
			num_scala_inf++
			num_persone_inf--
			ACK_SALITA_EQUIP_SUP[x] <- 1

		case x := <-when(((num_persone_inf+num_scala_sup <= num_persone_sup+num_scala_inf) && ((num_scala_inf == 0) || (num_viagg_sup_term+num_equip_sup_term == sup))), equipaggio_inferiore_scende_scala):
			num_scala_sup++
			num_persone_sup--
			ACK_DISCESA_EQUIP_INF[x] <- 1

		case x := <-when(((num_persone_sup+num_scala_inf <= num_persone_inf+num_scala_sup) && (num_scala_sup == 0)), equipaggio_inferiore_sale_scala):
			num_scala_inf++
			num_persone_inf--
			ACK_SALITA_EQUIP_INF[x] <- 1

		case <-viaggiatore_superiore_scala_discesa_ok:
			num_scala_sup--
			num_persone_inf++

		case <-viaggiatore_superiore_scala_salita_ok:
			num_scala_inf--
			num_persone_sup++
			num_viagg_sup_term++

		case <-viaggiatore_inferiore_scala_discesa_ok:
			num_scala_sup--
			num_persone_inf++
			num_viagg_inf_term++

		case <-viaggiatore_inferiore_scala_salita_ok:
			num_scala_inf--
			num_persone_sup++

		case <-equipaggio_superiore_scala_discesa_ok:
			num_scala_sup--
			num_persone_inf++

		case <-equipaggio_superiore_scala_salita_ok:
			num_scala_inf--
			num_persone_sup++
			num_equip_sup_term++

		case <-equipaggio_inferiore_scala_discesa_ok:
			num_scala_sup--
			num_persone_inf++
			num_equip_inf_term++
		case <-equipaggio_inferiore_scala_salita_ok:
			num_scala_inf--
			num_persone_sup++

		case <-termina:

			fmt.Println("\n Termino\n")
			done <- true
			return

		}

	}

}

func main() {

	var n_viag_piano_sup int
	var n_viag_piano_inf int
	var n_equip_piano_inf int
	var n_equip_piano_sup int

	fmt.Printf("\n Quanti viaggiatori al piano superiore (max %d)? ", TS)
	fmt.Scanf("%d", &n_viag_piano_sup)
	fmt.Printf("\n Quanti viaggiatori al piano inferiore (max %d)? ", TI)
	fmt.Scanf("%d", &n_viag_piano_inf)
	fmt.Printf("\n Quanti membri dell'equipaggio al piano superiore (max %d)? ", TS-n_viag_piano_sup)
	fmt.Scanf("%d", &n_equip_piano_sup)
	fmt.Printf("\n Quanti membri dell'equipaggio al piano inferiore (max %d)? ", TI-n_viag_piano_inf)
	fmt.Scanf("%d", &n_equip_piano_inf)

	for i := 0; i < n_viag_piano_sup; i++ {
		ACK_SALITA_VIAGGIATORE_SUP[i] = make(chan int, MAXBUFF)
		ACK_DISCESA_VIAGGIATORE_SUP[i] = make(chan int, MAXBUFF)
	}

	for i := 0; i < n_viag_piano_inf; i++ {
		ACK_SALITA_VIAGGIATORE_INF[i] = make(chan int, MAXBUFF)
		ACK_DISCESA_VIAGGIATORE_INF[i] = make(chan int, MAXBUFF)
	}

	for i := 0; i < n_viag_piano_sup; i++ {
		ACK_SALITA_EQUIP_SUP[i] = make(chan int, MAXBUFF)
		ACK_DISCESA_EQUIP_SUP[i] = make(chan int, MAXBUFF)
	}

	for i := 0; i < n_viag_piano_inf; i++ {
		ACK_SALITA_EQUIP_INF[i] = make(chan int, MAXBUFF)
		ACK_DISCESA_EQUIP_INF[i] = make(chan int, MAXBUFF)
	}

	rand.Seed(time.Now().Unix())

	go scala(n_viag_piano_sup+n_equip_piano_sup, n_viag_piano_inf+n_equip_piano_inf)

	for i := 0; i < n_viag_piano_sup; i++ {
		go viaggiatore(i, S)
	}

	for i := 0; i < n_viag_piano_inf; i++ {
		go viaggiatore(i, I)
	}

	for i := 0; i < n_equip_piano_sup; i++ {
		go equipaggio(i, S)
	}

	for i := 0; i < n_equip_piano_inf; i++ {
		go equipaggio(i, I)
	}

	for i := 0; i < n_viag_piano_inf+n_viag_piano_sup+n_equip_piano_sup+n_equip_piano_inf; i++ {
		<-done
	}

	termina <- true
	<-done

}
