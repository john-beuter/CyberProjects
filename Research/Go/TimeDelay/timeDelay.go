package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	timer2 := time.NewTimer(time.Duration(1) * time.Minute)
	go func() {
		<-timer2.C
		fmt.Println("Finished the timer")
		os.Exit(0)
	}()

	fmt.Println("Printing the output of the example")
}
