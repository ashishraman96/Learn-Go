package main 
import (
	"fmt"
	"time"
	"math"
	"math/rand"
	"strconv"
	"os"
	)
type data struct{
	id int
	prime []int
}
func main() {
	m,_:= strconv.Atoi(os.Args[1])
	//p := make(data)
	// p2 := make(data)
	// p3 := make(data)
	c1 := make(chan data)
	c2 := make(chan data)
	c3 := make(chan data)
	// 
	// go randPrime(m,1,c1,p)
	// go randPrime(m,2,c2,p)
	// go randPrime(m,3,c3,p)
	go randPrime(m,1,c1)
	time.Sleep(time.Second)
	go randPrime(m,2,c2)
	time.Sleep(time.Second)
	go randPrime(m,3,c3)
	time.Sleep(time.Second)
	r1:= <-c1
	r2:= <-c2
	r3:= <-c3
	duplicates(r1)
	duplicates(r2)
	duplicates(r3)
	// for i:=0; i<3; i++ {
	// 	select{
	// 		case r1:= <-c1: 
	// 			fmt.Println("Go routine", r1.id, "ended")
	// 		case r2:= <-c2: 
	// 			fmt.Println("Go routine", r2.id, "ended")
	// 		case r3:= <-c3: 
	// 			fmt.Println("Go routine", r3.id, "ended")
	// 		}
	// }
	
}

func randPrime(m,k int, c chan data){
	var p data 
	fmt.Println("Go routine", k, "started")
	var primes []int 
	rand.Seed(time.Now().Unix())
	count:=0
	for {
		r:=rand.Int()%1000
		l:= checkPrime(r)
		if l ==true {
			//fmt.Println(r)
			primes = append(primes, r)
			count++
			if(count==m){
				break
			}
		}	
	}
	p.id = k
	p.prime = primes
	c<-p 
}

func checkPrime(n int) bool{
	flag:=true
	sq:= math.Floor(math.Sqrt(float64(n)))
	//fmt.Println("number:", n, "root:",sq)
	
	for i:=2; i <= int(sq) ; i++{
		if n%i==0{
			flag=false
			break
		}
	
	}
	return flag
}

func duplicates(r data){
	count := make(map[int]float64, len(r.prime))
	for _, num :=range r.prime{
		count[num]++
	}
	// var s float64
	// for k,n := range count{
	// 	if n>1{
	// 		s=s+n
	// 	}
	// }
	// fmt.Printf("\n %f \t",s)
	// fmt.Printf("%f \n",float64(len(r.prime)))
	per := (float64(len(count))/float64(len(r.prime)))*100
	fmt.Printf("The id %d has %.2f %% of duplicacy \n", r.id, per)
}