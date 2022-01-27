package main
import (
	"fmt"
	"math/rand"
	"time"
)

const MAXSPAZZANEVE = 20
const MAXSPARGISALE = 20
const MAXCAMION = 20
const N = 5 // numero mezzi che possono operare nell'area
const K = 5 // capacità silo
const MAXBUFF = 100

var done = make(chan bool)
var termina = make(chan bool)
var entrata_spazzaneve = make(chan int, MAXBUFF)
var entrata_spargisale = make(chan int, MAXBUFF)
var entrata_camion = make(chan int, MAXBUFF)
var uscita_spazzaneve = make(chan int, MAXBUFF)
var uscita_spargisale = make(chan int, MAXBUFF)
var uscita_camion = make(chan int, MAXBUFF)
var camion_rifornisce = make(chan int, MAXBUFF)
var ACK_E_SPAZZANEVE [MAXSPAZZANEVE] chan int
var ACK_E_SPARGISALE [MAXSPARGISALE] chan int
var ACK_E_CAMION [MAXCAMION] chan int
var ACK_U_CAMION [MAXCAMION] chan int

func when(b bool, c chan int) chan int {
	if !b {
		return nil
	}
	return c
}


func spazzaneve(myid int){

  var tt int
  tt = rand.Intn(4)+1
  time.Sleep(time.Duration(tt) * time.Second)
  fmt.Printf("\n Sono il mezzo spazzaneve SPAZZA%d e sono in strada a pulire la neve\n", myid);
  tt = rand.Intn(10)+10
  time.Sleep(time.Duration(tt) * time.Second)
  fmt.Printf("\n Sono il mezzo spazzaneve SPAZZA%d e ho terminato il turno, mi dirigo al deposito per una sosta\n", myid);
  entrata_spazzaneve <- myid
  <- ACK_E_SPAZZANEVE[myid]
  tt = rand.Intn(10)+10
  fmt.Printf("\n Sono il mezzo spazzaneve SPAZZA%d e sono nel deposito per una sosta, ci rimango per %d secondi\n", myid, tt);
  time.Sleep(time.Duration(tt) * time.Second)
  fmt.Printf("\n Sono il mezzo spazzaneve SPAZZA%d e ho terminato la sosta, esco dal deposito\n", myid);
  uscita_spazzaneve <- 1
  done <- true

}

func spargisale(myid int){

  var tt int
  tt = rand.Intn(4)+1
  time.Sleep(time.Duration(tt) * time.Second)
  fmt.Printf("\n Sono il mezzo spargisale SPARGI%d e sono in strada per spargere il sale\n", myid);
  tt = rand.Intn(5)+5
  time.Sleep(time.Duration(tt) * time.Second)
  fmt.Printf("\n Sono il mezzo spargisale SPARGI%d e ho terminato il sale, mi dirigo al deposito per il rifornimento\n", myid);
  entrata_spargisale <- myid
  <- ACK_E_SPARGISALE[myid]
  tt = rand.Intn(10)+10
  fmt.Printf("\n Sono il mezzo spargisale SPARGI%d e sono nel deposito per il rifornimento, ci sosto per %d secondi\n", myid, tt);
  time.Sleep(time.Duration(tt) * time.Second)
  fmt.Printf("\n Sono il mezzo spargisale SPARGI%d e ho terminato il rifornimento, esco dall'area\n", myid);
  uscita_spargisale <- 1
  done <- true
}

func camion(myid int){

  var tt int
  var risultato int
  risultato = 0
  for risultato == 0{
    tt = rand.Intn(15)+15
    time.Sleep(time.Duration(tt) * time.Second)
    fmt.Printf("\n Sono il camion CAMION%d e mi dirigo al deposito per rifornire il silo\n", myid)
    entrata_camion <- myid
    <- ACK_E_CAMION[myid]
    tt = rand.Intn(10)+10
    fmt.Printf("\n Sono il camion CAMION%d e sono nel deposito. Rifornirò il silo in %d secondi\n", myid, tt)
    time.Sleep(time.Duration(tt) * time.Second)
    camion_rifornisce <- 1
    uscita_camion <- myid
    risultato = <- ACK_U_CAMION[myid]
  }
  fmt.Printf("\n CAMION%d termina!\n", myid)
  done <- true
}

func gestore(num_spargi int){

  var num_mezzi_deposito = 0
  var capacita = K
  var spargisale_terminati = 0

  for{

    select{

    case x := <- when((num_mezzi_deposito < N), entrata_spazzaneve):
      num_mezzi_deposito++
      ACK_E_SPAZZANEVE[x] <- 1

    case x := <- when((num_mezzi_deposito < N && len(entrata_spazzaneve) == 0), entrata_camion):
      num_mezzi_deposito++
      ACK_E_CAMION[x] <- 1

    case x := <- when((num_mezzi_deposito < N && len(entrata_spazzaneve) == 0 && len(entrata_camion) == 0 && capacita > 0), entrata_spargisale):
      num_mezzi_deposito++
      ACK_E_SPARGISALE[x] <- 1

    case <- camion_rifornisce:
      capacita = K

    case x := <- uscita_camion:
      num_mezzi_deposito--
      if spargisale_terminati == num_spargi {
        ACK_U_CAMION[x] <- 1
      }else{
        ACK_U_CAMION[x] <- 0
      }

    case <- uscita_spazzaneve:
      num_mezzi_deposito--

    case <- uscita_spargisale:
      num_mezzi_deposito--
      spargisale_terminati++

    case <-termina:

      fmt.Println("\n Termino\n")
      done <- true
      return
    }

  }
}

func main(){

  var num_camion int
  var num_spazzaneve int
  var num_spargisale int

  fmt.Printf("\n Quanti camion (max %d)? ", MAXCAMION)
  fmt.Scanf("%d", &num_camion)
  fmt.Printf("\n Quanti spazzaneve (max %d)? ", MAXSPAZZANEVE)
  fmt.Scanf("%d", &num_spazzaneve)
  fmt.Printf("\n Quanti spargisale (max %d)? ", MAXSPARGISALE)
  fmt.Scanf("%d", &num_spargisale)

  for i := 0; i < num_camion; i++{
    ACK_E_CAMION[i] = make(chan int, MAXBUFF)
    ACK_U_CAMION[i] = make(chan int, MAXBUFF)
  }

  for i := 0; i < num_spazzaneve; i++{
    ACK_E_SPAZZANEVE[i] = make(chan int, MAXBUFF)
  }

  for i := 0; i < num_spargisale; i++{
    ACK_E_SPARGISALE[i] = make(chan int, MAXBUFF)
  }

  rand.Seed(time.Now().Unix())

  go gestore(num_spargisale)

  for i := 0; i < num_camion; i++{
    go camion(i)
  }

  for i := 0; i < num_spazzaneve; i++{
    go spazzaneve(i)
  }

  for i := 0; i < num_spargisale; i++{
    go spargisale(i)
  }

  for i := 0; i < num_spazzaneve+num_spargisale; i++ {
		<-done
	}

  for i := 0; i < num_camion; i++ {
		<-done
	}
  termina <- true
  <- done



}