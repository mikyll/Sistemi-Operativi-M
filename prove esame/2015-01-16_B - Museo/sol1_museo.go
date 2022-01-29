//150116

package main

import(
	"fmt"
	"math/rand"
	"time"
)

//COSTANTI:
const MAX_S int = 10
const MAX_M int = 10

const N_VIS_M int = 2
const N_VIS_S int = 2
const N_VIS_SM int = 3

const N_OP_M int = 2
const N_OP_S int = 2

const MAX_BUFF int = 100

//CANALI:
//mi servono SEMPRE i due canali per lo stato di fine processo
var done = make(chan bool)
var termina_museo = make(chan bool)

//mi serve un canale per ogni richiesta, di ogni tipo, da ogni tipo di utente
var richiesta_ingresso_visitatore_m = make(chan int, MAX_BUFF)
var richiesta_ingresso_visitatore_s = make(chan int, MAX_BUFF)
var richiesta_ingresso_visitatore_sm = make(chan int, MAX_BUFF)

var comunica_uscita_visitatore_m = make(chan int, MAX_BUFF)
var comunica_uscita_visitatore_s = make(chan int, MAX_BUFF)
var comunica_uscita_visitatore_ms = make(chan int, MAX_BUFF)

var richiesta_ingresso_operatore_m = make(chan int, MAX_BUFF)
var richiesta_ingresso_operatore_s = make(chan int, MAX_BUFF)

var richiesta_uscita_operatore_m = make(chan int, MAX_BUFF)
var richiesta_uscita_operatore_s = make(chan int, MAX_BUFF)

//ACK:
var ack_ingresso_visitatore_m[N_VIS_M] chan bool
var ack_ingresso_visitatore_s[N_VIS_S] chan bool
var ack_ingresso_visitatore_sm[N_VIS_SM] chan bool

var ack_uscita_visitatore_m[N_VIS_M] chan bool
var ack_uscita_visitatore_s[N_VIS_S] chan bool
var ack_uscita_visitatore_sm[N_VIS_SM] chan bool

var ack_ingresso_operatore_m[N_OP_M] chan bool
var ack_ingresso_operatore_s[N_OP_S]  chan bool

var ack_uscita_operatore_m[N_OP_M] chan bool
var ack_uscita_operatore_s[N_OP_S] chan bool

//FUNZIONI UTILITY:
//per ogni tipo definito bisogna (eventualmente) ridefinire una when
func when (b bool, c chan int) chan int{
	if(!b){
		return nil
	}
	return c
}

// funzione client per il visitatore M
func visitatore_m(id int){
	//richiesta ingresso
//	time.Sleep(time.Duration(rand.Intn(3))*time.Second)
	richiesta_ingresso_visitatore_m<-id
	fmt.Printf("[Visitatore M %d]: richiesta ingresso.\n", id)
	<-ack_ingresso_visitatore_m[id]
	fmt.Printf("[Visitatore M %d]: entrato.\n", id)
	
	//comunicazione uscita
//	time.Sleep(time.Duration(rand.Intn(3))*time.Second)
	comunica_uscita_visitatore_m<-id
	fmt.Printf("[Visitatore M %d]: uscita.\n", id)
	<-ack_uscita_visitatore_m[id]
	
	fmt.Printf("[Visitatore M %d]: torno a casa.\n", id)
	done<-true
}

// funzione client per il visitatore S
func visitatore_s(id int){
	//richiesta ingresso
//	time.Sleep(time.Duration(rand.Intn(3))*time.Second)
	richiesta_ingresso_visitatore_s<-id
	fmt.Printf("[Visitatore S %d]: richiesta ingresso.\n", id)
	<-ack_ingresso_visitatore_s[id]
	fmt.Printf("[Visitatore S %d]: entrato.\n", id)
	
	//comunicazione uscita
//	time.Sleep(time.Duration(rand.Intn(3))*time.Second)
	comunica_uscita_visitatore_s<-id
	fmt.Printf("[Visitatore S %d]: uscita.\n", id)
	<-ack_uscita_visitatore_s[id]
	
	fmt.Printf("[Visitatore S %d]: torno a casa.\n", id)
	done<-true
}

// funzione client per il visitatore SM
func visitatore_sm(id int){
	//richiesta ingresso
//	time.Sleep(time.Duration(rand.Intn(3))*time.Second)
	richiesta_ingresso_visitatore_sm<-id
	fmt.Printf("[Visitatore SM %d]: richiesta ingresso.\n", id)
	<-ack_ingresso_visitatore_sm[id]
	fmt.Printf("[Visitatore SM %d]: entrato.\n", id)
	
	//comunicazione uscita
//	time.Sleep(time.Duration(rand.Intn(3))*time.Second)
	comunica_uscita_visitatore_ms<-id
	fmt.Printf("[Visitatore SM %d]: uscita.\n", id)
	<-ack_uscita_visitatore_sm[id]
	
	fmt.Printf("[Visitatore SM %d]: torno a casa.\n", id)
	done<-true
}

//funzione client per l'operatore M (worker)
func operatore_m(id int){
	//richiesta ingresso
//	time.Sleep(time.Duration(rand.Intn(3))*time.Second)
	richiesta_ingresso_operatore_m<-id
	fmt.Printf("[Operatore M %d]: richiesta ingresso.\n", id)
	<-ack_ingresso_operatore_m[id]
	fmt.Printf("[Operatore M %d]: entrato.\n", id)
	
	//comunicazione uscita
//	time.Sleep(time.Duration(rand.Intn(3))*time.Second)
	richiesta_uscita_operatore_m<-id
	fmt.Printf("[Operatore M %d]: uscita.\n", id)
	<-ack_uscita_operatore_m[id]
	
	fmt.Printf("[Operatore M %d]: torno a casa.\n", id)
	done<-true
}

//funzione client per l'operatore S (worker)
func operatore_s(id int){
	//richiesta ingresso
//	time.Sleep(time.Duration(rand.Intn(3))*time.Second)
	richiesta_ingresso_operatore_s<-id
	fmt.Printf("[Operatore S %d]: richiesta ingresso.\n", id)
	<-ack_ingresso_operatore_s[id]
	fmt.Printf("[Operatore S %d]: entrato.\n", id)
	
	//comunicazione uscita
//	time.Sleep(time.Duration(rand.Intn(3))*time.Second)
	richiesta_uscita_operatore_s<-id
	fmt.Printf("[Operatore S %d]: uscita.\n", id)
	<-ack_uscita_operatore_s[id]
	
	fmt.Printf("[Operatore S %d]: torno a casa.\n", id)
	done<-true
}

//funzione server per il museo
func museo(){
	//inizialmente nel museo
	var visitatori_s = 0
	var visitatori_m = 0
	var visitatori_sm = 0
	var operatori_m = 0
	var operatori_s = 0
	
	for{
		select{
			//gestisco richieste ingressi visitatori
			case x:=<-when(((visitatori_m+visitatori_sm+operatori_m)<MAX_M && operatori_m>0),richiesta_ingresso_visitatore_m):
				visitatori_m++
				ack_ingresso_visitatore_m[x]<-true
				fmt.Printf("Vis_m %d. Vis_s %d. Vis_sm %d. Op_m %d. Op_s %d.\n", visitatori_m, visitatori_s, visitatori_sm, operatori_m, operatori_s)
			case x:=<-when(((visitatori_s+visitatori_sm+operatori_s)<MAX_S && operatori_s>0),richiesta_ingresso_visitatore_s):
				visitatori_s++
				ack_ingresso_visitatore_s[x]<-true
				fmt.Printf("Vis_m %d. Vis_s %d. Vis_sm %d. Op_m %d. Op_s %d.\n", visitatori_m, visitatori_s, visitatori_sm, operatori_m, operatori_s)
			case x:=<-when(((visitatori_m+visitatori_sm+operatori_m)<MAX_M &&(visitatori_s+visitatori_sm+operatori_s)<MAX_S && operatori_s>0 && operatori_m>0),richiesta_ingresso_visitatore_sm):
				visitatori_sm++
				ack_ingresso_visitatore_sm[x]<-true
				fmt.Printf("Vis_m %d. Vis_s %d. Vis_sm %d. Op_m %d. Op_s %d.\n", visitatori_m, visitatori_s, visitatori_sm, operatori_m, operatori_s)
			
			//gestisco uscite visitatori (un visitatore può sempre uscire)
			case x:=<-when(1!=0,comunica_uscita_visitatore_m):
				visitatori_m--
				ack_uscita_visitatore_m[x]<-true
				fmt.Printf("Vis_m %d. Vis_s %d. Vis_sm %d. Op_m %d. Op_s %d.\n", visitatori_m, visitatori_s, visitatori_sm, operatori_m, operatori_s)
			case x:=<-when(1!=0,comunica_uscita_visitatore_s):
				visitatori_s--
				ack_uscita_visitatore_s[x]<-true
				fmt.Printf("Vis_m %d. Vis_s %d. Vis_sm %d. Op_m %d. Op_s %d.\n", visitatori_m, visitatori_s, visitatori_sm, operatori_m, operatori_s)
			case x:=<-when(1!=0,comunica_uscita_visitatore_ms):
				visitatori_sm--
				ack_uscita_visitatore_sm[x]<-true
				fmt.Printf("Vis_m %d. Vis_s %d. Vis_sm %d. Op_m %d. Op_s %d.\n", visitatori_m, visitatori_s, visitatori_sm, operatori_m, operatori_s)
			
			//gestisco ingressi operatori (un operatore può sempre entrare, se il MAX_M o il MAX_S sono rispettati)
			case x:=<-when((visitatori_m+visitatori_sm+operatori_m)<MAX_M,richiesta_ingresso_operatore_m):
				operatori_m++
				ack_ingresso_operatore_m[x]<-true
				fmt.Printf("Vis_m %d. Vis_s %d. Vis_sm %d. Op_m %d. Op_s %d.\n", visitatori_m, visitatori_s, visitatori_sm, operatori_m, operatori_s)
			case x:=<-when((visitatori_s+visitatori_sm+operatori_s)<MAX_S,richiesta_ingresso_operatore_s):
				operatori_s++
				ack_ingresso_operatore_s[x]<-true
				fmt.Printf("Vis_m %d. Vis_s %d. Vis_sm %d. Op_m %d. Op_s %d.\n", visitatori_m, visitatori_s, visitatori_sm, operatori_m, operatori_s)
			
			//gestisco uscite operatori (un operatore può uscire solo se in sala ci sono 0 visitatori oppure almeno un operatore)
			case x:=<-when(operatori_m>1||((len(richiesta_ingresso_visitatore_m)==0 && len(richiesta_ingresso_visitatore_sm)==0) && (operatori_m==1 && visitatori_m==0 && visitatori_sm==0)),richiesta_uscita_operatore_m):
				operatori_m--
				ack_uscita_operatore_m[x]<-true
				fmt.Printf("Vis_m %d. Vis_s %d. Vis_sm %d. Op_m %d. Op_s %d.\n", visitatori_m, visitatori_s, visitatori_sm, operatori_m, operatori_s)
			case x:=<-when(operatori_s>1||((len(richiesta_ingresso_visitatore_s)==0 && len(richiesta_ingresso_visitatore_sm)==0) && (operatori_s==1 && visitatori_s==0 && visitatori_sm==0)),richiesta_uscita_operatore_s):
				operatori_s--
				ack_uscita_operatore_s[x]<-true
				fmt.Printf("Vis_m %d. Vis_s %d. Vis_sm %d. Op_m %d. Op_s %d.\n", visitatori_m, visitatori_s, visitatori_sm, operatori_m, operatori_s)
			
			case<-termina_museo:
				<-done
				return	
		}
	}
}

func main(){
	fmt.Printf("AVVIO.\n")
	rand.Seed(time.Now().Unix())
	
	//inizializzo ACK
	for i:=0; i<N_VIS_M; i++{
		ack_ingresso_visitatore_m[i] = make(chan bool, MAX_BUFF)
	}
	for i:=0; i<N_VIS_S; i++{
		ack_ingresso_visitatore_s[i] = make(chan bool, MAX_BUFF)
	}
	for i:=0; i<N_VIS_SM; i++{
		ack_ingresso_visitatore_sm[i] = make(chan bool, MAX_BUFF)
	}
	for i:=0; i<N_OP_M; i++{
		ack_ingresso_operatore_m[i] = make(chan bool, MAX_BUFF)
	}
	for i:=0; i<N_OP_S; i++{
		ack_ingresso_operatore_s[i] = make(chan bool, MAX_BUFF)
	}
	for i:=0; i<N_VIS_M; i++{
		ack_uscita_visitatore_m[i] = make(chan bool, MAX_BUFF)
	}
	for i:=0; i<N_VIS_S; i++{
		ack_uscita_visitatore_s[i] = make(chan bool, MAX_BUFF)
	}
	for i:=0; i<N_VIS_SM; i++{
		ack_uscita_visitatore_sm[i] = make(chan bool, MAX_BUFF)
	}
	for i:=0; i<N_OP_M; i++{
		ack_uscita_operatore_m[i] = make(chan bool, MAX_BUFF)
	}
	for i:=0; i<N_OP_S; i++{
		ack_uscita_operatore_s[i] = make(chan bool, MAX_BUFF)
	}
	
	//lancio thread client e worker
	for i:=0; i<N_VIS_M; i++{
		go visitatore_m(i)
	}
	for i:=0; i<N_VIS_S; i++{
		go visitatore_s(i)
	}
	for i:=0; i<N_VIS_SM; i++{
		go visitatore_sm(i)
	}
	for i:=0; i<N_OP_M; i++{
		go operatore_m(i)
	}
	for i:=0; i<N_OP_S; i++{
		go operatore_s(i)
	}
	
	//lancio server
	go museo()
	
	//attendo terminazione dei client e dei worker
	for i:=1; i<N_VIS_M+N_VIS_S+N_VIS_SM+N_OP_M+N_OP_S; i++{
		<-done
	}
	
	//server termina
	termina_museo<-true
	<-done
	
	fmt.Printf("FINE.\n")
}