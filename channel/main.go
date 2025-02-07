package main

import (
	"fmt"
)

func printer(c chan bool) {
	fmt.Println("1")
	c <- true
}

func main() {
	d := make(chan bool)
	go printer(d)
	<-d
	fmt.Println("2")
}
