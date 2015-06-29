package main

import (
	"fmt"
	"time"
)

//func main() {
//	ci := make(chan int)
//
//	go write(ci)
//
//	//ci <- 4
//
//	value := <-ci
//
//	fmt.Println(value)
//}
//
//func write(c chan int) {
//	c <- 4
//}

//func main() {
//	c := make(chan int)
//	//<-c
//	//fmt.Println("should never reach here")
//
//	go func() {
//		for {
//			c <- 2
//		}
//	}()
//
//	fmt.Println("should never reach here 1")
//	for {
//		select {
//		case <-c:
//			//close(c)
//			//return
//			fmt.Println("data come for chan c")
//			fmt.Println("---------------0--------------")
//			//return
//			//case <-time.After(10 * time.Second):
//			//	fmt.Println("timer of 10 second arrived!")
//			//	fmt.Println("---------------1--------------")
//			//default:
//			//	fmt.Println("---------------2--------------")
//			//	c <- 2
//			//	return
//		}
//	}
//
//	fmt.Println("should never reach here 2")
//}

func main() {
	c := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		c <- 10
	}()

	for {
		select {
		case <-c:
			fmt.Println("trigger from chan C")
			return
		}
	}
}
