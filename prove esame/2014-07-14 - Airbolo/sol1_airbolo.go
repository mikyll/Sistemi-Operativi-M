package main

import(
	"fmt"
	"time"
	"math/rand"
)

//definizione tipo
type Biglietto struct{
	id int
	piano int
}

//aereo mai completamente pieno

//costanti
const CAPACITA_PONTE_SUPERIORE int = 20
const CAPACITA_PONTE_INFERIORE int = 30
const N_PASSEGGERI int = 20
const N_EQUIPAGGIO int = 15
const MAX_ATTESA int = 10
const MAX_BUFF int = 20

const PRIMA_CLASSE int = 0
const SECONDA_CLASSE int = 1

//canali
var done = make(chan bool)
var termina = make(chan bool)
var discesa_viaggiatore = make(chan int, MAX_BUFF)
var salita_viaggiatore = make(chan int, MAX_BUFF)
var discesa_equipaggio = make(chan int, MAX_BUFF)
var salita_equipaggio = make(chan int, MAX_BUFF)
var salita_aereo = make(chan Biglietto, MAX_BUFF)

//ack
var ack_discesa_viaggiatore[N_PASSEGGERI] chan bool
var ack_salita_viaggiatore[N_PASSEGGERI] chan bool
var ack_discesa_equipaggio[N_EQUIPAGGIO] chan bool
var ack_salita_equipaggio[N_EQUIPAGGIO] chan bool

var ack_salita_aereo[N_PASSEGGERI+N_EQUIPAGGIO] chan bool

//funzioni utility
func when(b bool, c chan int) chan int{
	if(!b){
		return nil
	}
	return c
}


func viaggiatore(id int){
	
	var tempo_di_attesa int
	var classe int = rand.Intn(2)
	var biglietto Biglietto = Biglietto{id, classe}
	
	salita_aereo <- biglietto
	
	<-ack_salita_aereo[id]
	
	//ciclo di vita di un passeggero
	for{
		//aspetto un tempo random
		tempo_di_attesa = rand.Intn(MAX_ATTESA)+1 //+1 perchè randomizza da 0 a MAX_ATTESA
		time.Sleep(time.Duration(tempo_di_attesa) * time.Second)
		
		switch classe{
		
			case PRIMA_CLASSE:
				fmt.Printf("[Viaggiatore %d] Scendo a farmi un giro\n", id)
				discesa_viaggiatore <- id
				<- ack_discesa_viaggiatore[id]
				
				fmt.Printf("[Viaggiatore %d] Sono di sopra\n", id)
				//passo del tempo giù
				tempo_di_attesa = rand.Intn(MAX_ATTESA)+1 //+1 perchè randomizza da 0 a MAX_ATTESA
				time.Sleep(time.Duration(tempo_di_attesa) * time.Second)
				
				salita_viaggiatore <- id
				<- ack_salita_viaggiatore[id]		
				fmt.Printf("[Viaggiatore %d] Sono tornato giù\n", id)		
			
			case SECONDA_CLASSE:
			
				fmt.Printf("[Viaggiatore %d] Vado a farmi un giro su\n", id)
				salita_viaggiatore <- id
				<- ack_salita_viaggiatore[id]
				
				fmt.Printf("[Viaggiatore %d] Sono di sopra\n", id)
				//passo del tempo giù
				tempo_di_attesa = rand.Intn(MAX_ATTESA)+1 //+1 perchè randomizza da 0 a MAX_ATTESA
				time.Sleep(time.Duration(tempo_di_attesa) * time.Second)
				
				discesa_viaggiatore <- id
				<- ack_discesa_viaggiatore[id]		
				fmt.Printf("[Viaggiatore %d] Sono tornato giù\n", id)
		}
	
	}
		
	//ho finito il mio lavoro
	done <- true
}

func equipaggio(id int){

	var tempo_di_attesa int
	var classe int = rand.Intn(2)
	
	var biglietto Biglietto = Biglietto{id, classe}
	
	salita_aereo <- biglietto
	<-ack_salita_aereo[id]

	//ciclo di vita di un equipaggio
	for{
		//aspetto un tempo random
		tempo_di_attesa = rand.Intn(MAX_ATTESA)+1 //+1 perchè randomizza da 0 a MAX_ATTESA
		time.Sleep(time.Duration(tempo_di_attesa) * time.Second)
		
		switch classe{
		
			case PRIMA_CLASSE:
				fmt.Printf("[Equipaggio %d] Scendo a farmi un giro\n", id)
				discesa_equipaggio <- id
				<- ack_discesa_equipaggio[id]
				
				fmt.Printf("[Equipaggio %d] Sono di sopra\n", id)
				//passo del tempo giù
				tempo_di_attesa = rand.Intn(MAX_ATTESA)+1 //+1 perchè randomizza da 0 a MAX_ATTESA
				time.Sleep(time.Duration(tempo_di_attesa) * time.Second)
				
				salita_equipaggio <- id
				<- ack_salita_equipaggio[id]		
				fmt.Printf("[Equipaggio %d] Sono tornato giù\n", id)		
			
			case SECONDA_CLASSE:
			
				fmt.Printf("[Equipaggio %d] Vado a farmi un giro su\n", id)
				salita_equipaggio <- id
				<- ack_salita_equipaggio[id]
				
				fmt.Printf("[Equipaggio %d] Sono di sopra\n", id)
				
				//passo del tempo giù
				tempo_di_attesa = rand.Intn(MAX_ATTESA)+1 //+1 perchè randomizza da 0 a MAX_ATTESA
				time.Sleep(time.Duration(tempo_di_attesa) * time.Second)
				
				discesa_equipaggio <- id
				<- ack_discesa_equipaggio[id]		
				fmt.Printf("[Viaggiatore %d] Sono tornato giù\n", id)
		}
	}
	done <- true

}

func scala(){
	
	var persone_prima_classe int = 0
	var persone_seconda_classe int = 0
	var scala_occupata bool = false
	
	//ciclo di vita del server
	for{
		select{
			
			case x:=<- salita_aereo:
				fmt.Printf("[Aereo] salito passeggero %d in classe %d\n", x.id, x.piano)
				switch x.piano{
					case PRIMA_CLASSE:
						persone_prima_classe++
					case SECONDA_CLASSE:
						persone_seconda_classe++
				}
				ack_salita_aereo[x.id] <- true
			
			//se non c'è nessuno in coda dall'altra parte
			case x:=<-when((!scala_occupata && ((CAPACITA_PONTE_SUPERIORE - persone_prima_classe) > (CAPACITA_PONTE_INFERIORE - persone_seconda_classe)) && (len(discesa_equipaggio) == 0) && (len(salita_equipaggio)==0)), salita_viaggiatore):
				fmt.Printf("[Viaggiatore %d] in transito (salita)\n", x)
				scala_occupata = true
				persone_seconda_classe--
				persone_prima_classe++
				scala_occupata = false
				ack_salita_viaggiatore[x] <- true
				
			case x:=<-when((!scala_occupata && ((CAPACITA_PONTE_INFERIORE - persone_seconda_classe) > (CAPACITA_PONTE_SUPERIORE - persone_prima_classe)) && (len(discesa_equipaggio) == 0) && (len(salita_equipaggio)==0)), discesa_viaggiatore):
				fmt.Printf("[Viaggiatore %d] in transito (discesa)\n", x)
				scala_occupata = true
				persone_prima_classe--
				persone_seconda_classe++
				scala_occupata = false
				ack_discesa_viaggiatore[x] <- true
			
			case x:=<-when((!scala_occupata && ((CAPACITA_PONTE_INFERIORE - persone_seconda_classe) > (CAPACITA_PONTE_SUPERIORE - persone_prima_classe))), discesa_equipaggio):
				fmt.Printf("[Equipaggio %d] in transito (discesa)\n", x)
				scala_occupata = true
				persone_prima_classe--
				persone_seconda_classe++
				scala_occupata = false
				ack_discesa_equipaggio[x] <- true
			
			case x:=<-when((!scala_occupata && ((CAPACITA_PONTE_SUPERIORE - persone_prima_classe) > (CAPACITA_PONTE_INFERIORE - persone_seconda_classe))), salita_equipaggio):
				fmt.Printf("[Equipaggio %d] in transito (salita)\n", x)
				scala_occupata = true
				persone_seconda_classe--
				persone_prima_classe++
				scala_occupata = false
				ack_salita_equipaggio[x] <- true
		
			case <-termina:
				done <- true
				return
		}
		fmt.Printf("[Scala] persone piano di sotto: %d, persone piano di sopra: %d\n", persone_seconda_classe, persone_prima_classe)
	}
}

func main(){
	
	fmt.Printf("Programma avviato\n")
	rand.Seed(time.Now().Unix())
	
	//inizializzo canali ack
	for i:=0; i<N_PASSEGGERI; i++{
		ack_salita_viaggiatore[i] = make(chan bool, MAX_BUFF)
		ack_discesa_viaggiatore[i] = make(chan bool, MAX_BUFF)
	}
	
	for i:=0; i<N_EQUIPAGGIO; i++{
	ack_salita_equipaggio[i] = make(chan bool, MAX_BUFF)
	ack_discesa_equipaggio[i] = make(chan bool, MAX_BUFF)
	}
	
	for i:=0; i<N_EQUIPAGGIO+N_PASSEGGERI;i++{
		ack_salita_aereo[i] = make(chan bool, MAX_BUFF)
	}
	
	//lancio threads
	for i:=0; i<N_PASSEGGERI;i++{
		go viaggiatore(i)
	}
	
	for i:=0; i<N_EQUIPAGGIO;i++{
		go equipaggio(i)
	}
	//lancio il server
	go scala()
	
	//attendo la terminazione dei clients
	for i:=0; i<N_PASSEGGERI + N_EQUIPAGGIO; i++{
		<-done
	}
	
	//avviso il server di terminare
	termina <- true
	
	//attendo la terminazione del server
	<-done
	
	fmt.Printf("Programma terminato\n")
}
