package main

import (
	"fmt"
	"os"
	"time"
)

//Have a function called time delay
//Call that function as a thread from the main program
//The time spins down and then runs a thread to trigger the attack at certain time.

func main() {

	timer2 := time.NewTimer(time.Duration(1) * time.Minute)
	go func() {
		<-timer2.C
		fmt.Println("Finished the timer")
		//Here I could just call the seperate thread to run the program.
		os.Exit(0)
	}()

	fmt.Println("Printing the output of the example")
}
