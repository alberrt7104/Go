
package main

import "fmt"
import "time"


func numbs(c chan<- int) {

         c<-2

   for i := 3; i<=10000; i=i+2 {

      if( i %2!=0){

         c <- i  

      }

   }

}



func newprime(ch1<-chan int,ch2 chan int, num int) {


   for next := range ch1 {

      if next%num != 0 {

         ch2 <- next

      }
   }



}

func worker(done chan bool) {

fmt.Print("working...")

time.Sleep(time.Second)

fmt.Println("done")


done <- true

}

func main() {

   ch1 := make(chan int)  
// done := make(chan bool, 1)

   go numbs(ch1)      

   start := time.Now()


   for i := 1; i<=1229 ; i++{

      prime := <-ch1   

      

      ch2 := make(chan int) 

      go newprime(ch1, ch2, prime)

      ch1 = ch2
      fmt.Println("the ", i, "prime number is",prime)

   }

   end := time.Now().Sub(start)

   fmt.Println("running time :",end)

}







