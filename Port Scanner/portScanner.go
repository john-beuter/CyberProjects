package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)


func scanPort(host string){
		connect, err := net.Dial("tcp", host) //Trying to make connection
		port := host[len(host)-2:]
	
		if err != nil {
			//This would print each port that we failed to connect to 
		}
		
		if connect != nil{ // IF we have a connection, then record that
			fmt.Println("Port " + port + " is open")
		}
} 



// Function should take an IP address as a parameter and enumerate that IP
func scanIP(address string) {
	var wait sync.WaitGroup
	for i := 0; i < 24; i++ {
		wait.Add(1)
		port := strconv.Itoa(i)
		s := net.JoinHostPort(address, port)
		go scanPort(s)
	}	
	wait.Wait() //Got it to somewhat work with goroutines but having deadlock!
}

func main() {
	scanIP("192.168.4.76")
}
