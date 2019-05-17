package main

import "fmt"
import "math/rand"

var size = 20 //size of array
var ran = 100 //range of number

func main() {
	// initialize the unsorted array
	var v []int = make([]int, size)
	for i := 0; i < size; i++{
		v[i] = rand.Intn(ran)
		fmt.Printf("%d ", v[i])
	}
	fmt.Printf(" init\n")

	//merge channels
	var in1 chan int = make (chan int, size)
	var in2 chan int = make (chan int, size)
	var out chan int = make (chan int, size)
	
	//channels used to synchronize routines
	var check1 chan bool = make (chan bool)
	var check2 chan bool = make (chan bool)
	var check3 chan bool = make (chan bool)

	//size of in1 and in2 in each round
	var subsize = 1

	//index of in1 and in2
	var index1 int
	var index2 int
	var block int
	
	//each round merges size / (subsize * 2) blocks
	for subsize <= size {
		block = 1
		
		//check1, check2 and check3 are used to synchronize
		for i := block * subsize; i <= size; i = block * subsize {
			
				for index1 = i - subsize; index1< i && index1<size ; index1++ {
					in1 <- v[index1]
				}
				for index2 = i; index2 < i + subsize && index2 < size; index2++ {
					in2 <- v[index2]
				}
				check1 <- true
			
			go Merge(in1, in2, out, check1, check2)
			go func() {
				<- check2
				for j := i - subsize; j < i + subsize && j < size; j++ {
					v[j] =<- out
				}
				check3 <- true
			}()
			<- check3
			block += 2
		}
		fmt.Printf("\n")
		for k := 0; k < size; k++ {
			fmt.Printf("%d ", v[k])
		}
		fmt.Printf("\n")
		subsize *= 2
	}
}

func Merge(in1 chan int, in2 chan int, out chan int, check1 chan bool, check2 chan bool)  {
	var v1, v2, v3, tag int
	<- check1
	tag = 0

	//read data only if in1 and in2 are not empty
	if (len(in1) > 0) {
		v1 =<- in1
	} else {
		v1 = ran + 1
	}
	if (len(in2) > 0) {
		v2 =<- in2
	} else {
		v2 = ran + 1
	}

	//read and compare and out<-
	for len(in1) > 0 && len(in2) > 0 {
		if (v1 < v2) {
			v3 = v1
			if (len(in1) > 0) {
				v1 = <- in1
			}
		} else {
			v3 = v2
			if (len(in2) > 0) {
				v2 = <- in2
			}
		}
		out <- v3
	}

	//deal with the remaining
	if (v1 < v2 && v1 < ran) {
		out <- v1
		if (v2 < ran) {
			tag = 2
		}
	} else if (v2 < v1 && v2 < ran) {
		out <- v2
		if (v1 < ran) {
			tag = 1
		}
	}			
	for len(in1) > 0 {
		v3 = <- in1
		if (tag == 1) {
			if (v1 < v3) {
				tag = 0
				out <- v1
			}
		} else if (tag == 2) {
			if (v2 < v3) {
				tag = 0
				out <- v2
			}
		}
		out <- v3
	}
	for len(in2) > 0 {
		v3 = <- in2
		if (tag == 1) {
			if (v1 < v3) {
				tag = 0
				out <- v1
			}
		} else if (tag == 2) {
			if (v2 < v3) {
				tag = 0
				out <- v2
			}
		}
		out <- v3
	}
	if (tag == 1) {
		out <- v1
	} else if (tag == 2) {
		out <- v2
	}
	check2 <- true
}
