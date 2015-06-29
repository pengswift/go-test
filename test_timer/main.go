package main

import (
	"fmt"
	"time"
)

func main() {

	//time.AfterFunc(5*time.Second, func() {
	//	fmt.Println("expired1")
	//})

	//timer := time.NewTimer(5 * time.Second)
	//<-timer.C
	//fmt.Println("expired2")

	//<-time.After(5 * time.Second)
	//fmt.Println("expired3")

	time.Sleep(10 * time.Second)
	fmt.Println("...........end............")

}
