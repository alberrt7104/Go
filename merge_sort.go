package main 

import "fmt"


func merge(ch1 chan int, ch2 chan int, out chan int, check1 chan bool,check2 chan bool,checko chan bool) {
var x, y, temp int

var checknull int=1000
var mode int=3

<- check1

if(len(ch1)>0){

x=<- ch1

}else{

x=checknull+100

}
if(len(ch2)>0){

y=<- ch2

}else{
y=checknull+100


}


for len(ch1)>0 && len(ch2)>0{



if(x<y){
temp=x
	if(len(ch1)>0){

	x=<- ch1
}
}else{
temp=y
	if(len(ch2)>0){
		y=<- ch2
}
}
fmt.Printf("%d \n",temp)
out <- temp
}

if(x>y && y<checknull){
fmt.Printf("%d \n",y)
out<- y
if(x<checknull){

mode=1
}
}else if(x<y && x<checknull){
fmt.Printf("%d \n",x)
out<- x
if(y<checknull){

mode=2

}

}

for len(ch1) >0{
temp =<- ch1
if(mode==1){
if (temp>x){
fmt.Printf("%d \n",x)
out<-x
mode=3

}
}else if(mode==2){
if(temp>y){
fmt.Printf("%d \n",y)
out<-y
mode=3

}
}
fmt.Printf("%d \n",temp)
out <- temp
}


for len(ch2) >0{
temp =<- ch2
if(mode==1){
if (temp>x){
fmt.Printf("%d \n",x)
out<-x
mode=3

}
}else if(mode==2){
if(temp>y){
fmt.Printf("%d \n",x)
out<-y
mode=3

}
}
fmt.Printf("%d \n",temp)
out <- temp
}

if(mode==1){
fmt.Printf("%d \n",x)
out<- x

}else if(mode==2){
fmt.Printf("%d \n",y)
out<-y

}
check2<- true

}


func main(){


//list := []int{ 8, 7, 4, 3, 2, 1, 10, 12, 6, 5}
list := []int{ 12,55,8, 7, 4, 3, 2, 1, 10, 12, 6, 5,11,33,22,44,56,31,24,909}
//var list1 int
//var list2 int

var j,k int
var count int=0
var size int=1 
var sizenum int=1
for i:=0;i<20;i++{
fmt.Printf("%d ",list[i])
}

ch1 :=make(chan int,20)
ch2 :=make(chan int,20)
out :=make(chan int,20)

check1 := make(chan bool)
check2 := make(chan bool)
checko := make(chan bool)

	

for size<=20 {
     sizenum=1
for i:=sizenum*size ;i<= 20; i=sizenum*size{
go func(){
	for j=i-size;j<i && j<20;j++{
	ch1<- list[j]
}
	for k=i;k<i+size && k<20;k++{
	ch2<- list[k]
}
check1<- true

}()

go merge(ch1,ch2,out,check1,check2,checko)

go func(){
<- check2
for m:=i-size;m<i+size && m<20;m++{
list[m]= <- out
}
checko <- true
}()
<- checko
//fmt.Printf("the %d time:\n",sizenum)
sizenum=sizenum+2

}

size=size+size
count++
//fmt.Printf("the %d time:\n",size)

fmt.Printf("\n")
fmt.Printf("the %d time:",count)
for n:=0;n<20;n++{
fmt.Printf("%d ",list[n])
}
fmt.Printf("\n")


}	


}




