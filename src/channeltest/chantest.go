package main

import (
	"fmt"
	"time"
)

func doSomething(s chan string) {
	time.Sleep(time.Second)
	fmt.Print(" --- 1 --- ")
	s <- "done"
}

func doSomethingElse(s chan string) {
	time.Sleep(time.Second + time.Second)
	fmt.Print(" --- 2 --- ")
	s <- "done"
}

//simple test
func main() {
	chan1 := make(chan string, 2)
	chan2 := make(chan string)

	go doSomething(chan1)
	go doSomethingElse(chan2)

	<-chan1
	fmt.Print("ok...")
	<-chan2
	chan1 <- "d"
	<-chan1

	fmt.Print("Done ... ")
}
