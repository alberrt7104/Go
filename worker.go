
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



func worker1(ch1<-chan int,ch2 chan int, num int) {



   for next := range ch1 {

      if next%num != 0 {

         ch2 <- next

      }



   }





}

func worker2(ch1<-chan int,ch3 chan int, num int) {



   for next := range ch1 {

      if next%num != 0 {

         ch3 <- next

      }



   }





}

func worker3(ch1<-chan int,ch4 chan int, num int) {



   for next := range ch1 {

      if next%num != 0 {

         ch4 <- next

      }



   }





}

func worker4(ch1<-chan int,ch5 chan int, num int) {



   for next := range ch1 {

      if next%num != 0 {

         ch5 <- next

      }



   }





}
func main() {

   ch1 := make(chan int)  

   go numbs(ch1)      

   start := time.Now()

	//var count int=0

	

   for i := 1; i<=1229; i++{
	
           

      if i%4==0{
	prime := <-ch1  
      ch2 := make(chan int) 

      go worker1(ch1, ch2, prime)

      ch1 = ch2
      fmt.Println("the worker1", i, "prime number is",prime)


}
      if i%4==1{
prime := <-ch1  
      ch3 := make(chan int) 

      go worker2(ch1, ch3, prime)

      ch1 = ch3
      fmt.Println("the worker2", i, "prime number is",prime)


}
      if i%4==2{
prime := <-ch1  
      ch4 := make(chan int) 

      go worker4(ch1, ch4, prime)

      ch1 = ch4
      fmt.Println("the worker3", i, "prime number is",prime)


}
     
 
      if i%4==3{
prime := <-ch1  
      ch5 := make(chan int) 

      go worker4(ch1, ch5, prime)

      ch1 = ch5
      fmt.Println("the worker4", i, "prime number is",prime)


}

  
}
   end := time.Now().Sub(start)

   fmt.Println("running time :",end)



}





