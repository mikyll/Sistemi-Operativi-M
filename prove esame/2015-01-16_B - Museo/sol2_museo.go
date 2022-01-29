/* Si consideri un grande complesso museale, suddiviso in due aree distinte:
•Il museo, che consiste in ’esposizione permanente
•una sala mostre, dove vengono organizzate mostre temporanee dedicate a temi vari.
I visitatori del complesso museale possono essere di tre tipi:
• visitatori che accedono solo al museo,
• visitatori che accedono solo alla sala mostre,
• visitatori che accedono a entrambe le aree (museo e sala mostre).
Le due aree hanno una capacità limitata fissata rispettivamente a Nm (massimo numero di clienti ammessi al museo) e Ns (massimo numero di clienti ammessi alla sala mostre).
Ognuna delle 2 aree è supervisionata da operatori specializzati, dedicati all’assistenza dei visitatori presenti nell’area; in particolare vi sono 2 tipi di operatori:
• operatori museo;
• operatori sala mostre.
Gli operatori possono arbitrariamente entrare e uscire (ad esempio, per una pausa) dal complesso museale tenendo conto che:
• per garantire un sufficiente grado di sicurezza, è necessario che in presenza di uno o più visitatori ognuna delle 2 aree sia supervisionata da almeno un operatore.
Pertanto ingressi e uscite verso/da il complesso museale di visitatori e operatori devono essere opportunamente regolamentati tenendo conto dei vincoli sopra menzionati.
Realizzare un’applicazione a scelta nel linguaggio go oppure nel linguaggio Ada, nella quale visitatori, operatori e gestore del complesso museale siano rappresentati da processi (goroutine o task) concorrenti.
La sincronizzazione tra i processi dovrà tenere conto dei vincoli dati.
Inoltre, per quanto riguarda l’ingresso dei visitatori, il gestore del complesso servirà le richieste applicando il seguente criterio di priorità:
1.Richieste di ingresso di visitatori che accedono solo al MUSEO; 2.Richieste di ingresso di visitatori che accedono solo alla SALA MOSTRE; 3.Richieste di ingresso di visitatori che accedono a entrambe le aree. */
package main

import (
	"fmt"
	"math/rand"
	"time"
)


const NM = 3
const NS = 3
const MAXVISM = 20
const MAXVISS = 20
const MAXVISSM = 20
const MAXOPM = 10
const MAXOPS = 10
const MAXBUFF = 100
const M = 0
const S = 1
const SM = 2

var done = make(chan bool)
var termina = make(chan bool)
var entrata_visitatore_museo = make(chan int, MAXBUFF)
var entrata_visitatore_sala = make(chan int, MAXBUFF)
var entrata_visitatore_sm = make(chan int, MAXBUFF)
var uscita_visitatore_museo = make(chan int)
var uscita_visitatore_sala = make(chan int)
var uscita_visitatore_sm = make(chan int)

var entrata_operatore_museo = make(chan int, MAXBUFF)
var entrata_operatore_sala = make(chan int, MAXBUFF)
var uscita_operatore_museo = make(chan int, MAXBUFF)
var uscita_operatore_sala = make(chan int, MAXBUFF)

var ACK_E_V_M [MAXVISM]chan int
var ACK_E_V_S [MAXVISS]chan int
var ACK_E_V_SM [MAXVISSM]chan int
var ACK_E_O_M [MAXOPM]chan int
var ACK_E_O_S [MAXOPM]chan int
var ACK_U_O_M [MAXOPM]chan int
var ACK_U_O_S [MAXOPM]chan int

func when(b bool, c chan int) chan int{
	if(!b){
		return nil
	}
	return c
}


func visitatore(myid int, tipo int){

  var tt int

  switch tipo{

  case M:

    //fmt.Printf("\n Inizializzazione visitatore M%d diretto al museo\n", myid)
    tt = rand.Intn(5)+1
    time.Sleep(time.Duration(tt) * time.Second)
    fmt.Printf("\n Visitatore M%d chiede accesso al museo \n", myid)
    entrata_visitatore_museo <- myid
    <- ACK_E_V_M[myid]
    tt = rand.Intn(5)+3
    fmt.Printf("\n Visitatore M%d visita il museo, uscirà fra %d secondi\n", myid, tt)
    time.Sleep(time.Duration(tt) * time.Second)
    uscita_visitatore_museo <- myid
    fmt.Printf("\n Visitatore M%d esce dal museo\n", myid)

  case S:

    //fmt.Printf("\n Inizializzazione visitatore S%d diretto alla sala mostre\n", myid)
    tt = rand.Intn(5)+1
    time.Sleep(time.Duration(tt) * time.Second)
    fmt.Printf("\n Visitatore S%d chiede accesso alla sala mostre\n", myid)
    entrata_visitatore_sala <- myid
    <- ACK_E_V_S[myid]
    tt = rand.Intn(5)+3
    fmt.Printf("\n Visitatore S%d visita la sala mostre, uscirà fra %d secondi\n", myid, tt)
    time.Sleep(time.Duration(tt) * time.Second)
    uscita_visitatore_sala <- myid
    fmt.Printf("\n Visitatore S%d esce dalla sala\n", myid)


  case SM:

    fmt.Printf("\n Inizializzazione visitatore SM%d diretto ad entrambe le sale\n", myid)
    tt = rand.Intn(5)+1
    time.Sleep(time.Duration(tt) * time.Second)
    fmt.Printf("\n Visitatore SM %d chiede accesso ad entrambe le sale\n", myid)
    entrata_visitatore_sm <- myid
    <- ACK_E_V_SM[myid]
    tt = rand.Intn(5)+3
    fmt.Printf("\n Visitatore SM %d visita entrambe le sale, uscirà fra %d secondi\n", myid, tt)
    time.Sleep(time.Duration(tt) * time.Second)
    uscita_visitatore_sm <- myid
    fmt.Printf("\n Visitatore SM %d esce da entrambe le sale\n", myid)

  }
  done <- true

}

func operatore(myid int, tipo int){

  var tt int

  for{

    switch tipo{

    case M:

      tt = rand.Intn(1)+1
      time.Sleep(time.Duration(tt) * time.Second)
      fmt.Printf("\n Operatore M%d prova ad entrare nel museo \n", myid)
      entrata_operatore_museo <- myid
      <- ACK_E_O_M[myid]
      fmt.Printf("\n Operatore M%d entrato nel museo \n", myid)
      tt = rand.Intn(5)+2
      time.Sleep(time.Duration(tt) * time.Second)
      tt = rand.Intn(5)+1
      fmt.Printf("\n Operatore M%d vorrebbe prendersi una pausa di %d secondi \n", myid, tt)
      uscita_operatore_museo <- myid
      x := <- ACK_U_O_M[myid]
      if x == 0 {
        fmt.Printf("\n Operatore M%d termina!\n", myid)
        done <- true
        return
      }
      fmt.Printf("\n Operatore M%d in pausa per %d secondi \n", myid, tt)
      time.Sleep(time.Duration(tt) * time.Second)

    case S:

      tt = rand.Intn(1)+1
      time.Sleep(time.Duration(tt) * time.Second)
      fmt.Printf("\n Operatore S%d prova ad entrare nella sala mostre \n", myid)
      entrata_operatore_sala <- myid
      <- ACK_E_O_S[myid]
      fmt.Printf("\n Operatore S%d entrato nella sala mostre \n", myid)
      tt = rand.Intn(5)+2
      time.Sleep(time.Duration(tt) * time.Second)
      tt = rand.Intn(5)+1
      fmt.Printf("\n Operatore S%d vorrebbe prendersi una pausa di %d secondi \n", myid, tt)
      uscita_operatore_sala <- myid
      x := <- ACK_U_O_S[myid]
      if x == 0 {
        fmt.Printf("\n Operatore S%d termina!\n", myid)
        done <- true
        return
      }
      fmt.Printf("\n Operatore S%d in pausa per %d secondi \n", myid, tt)
      time.Sleep(time.Duration(tt) * time.Second)

    }

  }

}

func gestore() {

  var num_op_m int
  var num_op_s int
  var num_vis_m int
  var num_vis_s int

  num_op_m = 0
  num_op_s = 0
  num_vis_m = 0
  num_vis_s = 0

  for{

    select{

    case x := <- when((num_op_m+num_vis_m < NM) && (num_op_m != 0), entrata_visitatore_museo):

      num_vis_m++
      fmt.Printf("\n Autorizzo l'ingresso del visitatore M%d al museo \n", x)
      ACK_E_V_M[x] <- 1

    case x := <- when((num_op_s+num_vis_s < NS) && (num_op_s != 0) && (len(entrata_visitatore_museo) == 0), entrata_visitatore_sala):

      num_vis_s++
      fmt.Printf("\n Autorizzo l'ingresso del visitatore S%d alla sala mostre \n", x)
      ACK_E_V_S[x] <- 1

    case x := <- when((num_op_s+num_vis_s < NS) && (num_op_m+num_vis_m < NM) && (num_op_s != 0) && (num_op_m != 0) && (len(entrata_visitatore_sala) == 0) && (len(entrata_visitatore_museo) == 0), entrata_visitatore_sm):

      num_vis_s++
      num_vis_m++
      fmt.Printf("\n Autorizzo l'ingresso del visitatore SM%d ad entrambe le sale \n", x)
      ACK_E_V_SM[x] <- 1

    case x := <- when((num_op_m+num_vis_m < NM), entrata_operatore_museo):

      num_op_m++
      fmt.Printf("\n Autorizzo l'ingresso dell'operatore M%d al museo \n", x)
      ACK_E_O_M[x] <- 1

    case x := <- when((num_op_s+num_vis_s < NS), entrata_operatore_sala):

      num_op_s++
      fmt.Printf("\n Autorizzo l'ingresso dell'operatore S%d alla sala \n", x)
      ACK_E_O_S[x] <- 1

    case x := <- when(( (num_op_m > 1) || ( (len(entrata_visitatore_museo) == 0) && (len(entrata_visitatore_sm) == 0) && (num_vis_m == 0))), uscita_operatore_museo):

      num_op_m--
      fmt.Printf("\n Autorizzo l'uscita dell'operatore M%d dalla sala \n", x)
      if num_vis_m == 0{
        ACK_U_O_M[x] <- 0
      }else{
        ACK_U_O_M[x] <- 1
      }
    case x := <- when(( (num_op_s > 1) || ( (len(entrata_visitatore_sala) == 0) && (len(entrata_visitatore_sm) == 0) && (num_vis_s == 0))), uscita_operatore_sala):

      num_op_s--
      fmt.Printf("\n Autorizzo l'uscita dell'operatore S%d dalla sala \n", x)
      if num_vis_s == 0{
        ACK_U_O_S[x] <- 0
      }else{
        ACK_U_O_S[x] <- 1
      }

    case x := <- uscita_visitatore_museo:

      num_vis_m--
      fmt.Printf("\n Visitatore M%d uscito dal museo \n", x)

    case x := <- uscita_visitatore_sala:

      num_vis_s--
      fmt.Printf("\n Visitatore S%d uscito dalla sala mostre \n", x)

    case x := <- uscita_visitatore_sm:

      num_vis_s--
      num_vis_m--
      fmt.Printf("\n Visitatore SM%d uscito dalle sale \n", x)

    case <-termina:

      fmt.Println("\n Termino\n")
			done <- true
			return

    }

  }

}

func main(){

  var num_museo int
  var num_sala int
  var num_sm int
  var op_m int
  var op_s int

  fmt.Printf("\n Quanti visitatori in ingresso al museo (max %d)? ", MAXVISM)
	fmt.Scanf("%d", &num_museo)
  fmt.Printf("\n Quanti visitatori in ingresso alla sala mostre (max %d)? ", MAXVISS)
  fmt.Scanf("%d", &num_sala)
  fmt.Printf("\n Quanti visitatori in ingresso ad entrambe le sale (max %d)? ", MAXVISSM)
	fmt.Scanf("%d", &num_sm)
  fmt.Printf("\n Quanti operatori per il museo (max %d)? ", MAXOPM)
	fmt.Scanf("%d", &op_m)
  fmt.Printf("\n Quanti operatori per la sala mostra (max %d)? ", MAXOPM)
	fmt.Scanf("%d", &op_s)

  for i := 0; i < num_museo; i++ {
		ACK_E_V_M[i] = make(chan int, MAXBUFF)
	}

  for i := 0; i < num_sala; i++ {
		ACK_E_V_S[i] = make(chan int, MAXBUFF)
	}

  for i := 0; i < num_sm; i++ {
		ACK_E_V_SM[i] = make(chan int, MAXBUFF)
	}

  for i := 0; i < op_m; i++ {
		ACK_E_O_M[i] = make(chan int, MAXBUFF)
    ACK_U_O_M[i] = make(chan int, MAXBUFF)
	}

  for i := 0; i < op_s; i++ {
		ACK_E_O_S[i] = make(chan int, MAXBUFF)
    ACK_U_O_S[i] = make(chan int, MAXBUFF)
	}

  rand.Seed(time.Now().Unix())

  go gestore()

  for i := 0; i < num_museo; i++ {
    go visitatore(i, M)
	}

  for i := 0; i < num_sala; i++ {
    go visitatore(i, S)
	}

  for i := 0; i < num_sm; i++ {
    go visitatore(i, SM)
	}

  for i := 0; i < op_m; i++ {
    go operatore(i, M)
	}

  for i := 0; i < op_s; i++ {
    go operatore(i, S)
	}

  for i := 0; i < num_museo+num_sala+num_sm; i++ {
		<-done
	}

  for i := 0; i < op_m+op_s; i++ {
		<-done
	}

	termina <- true
	<-done
	fmt.Printf("\n HO FINITO ")

}
