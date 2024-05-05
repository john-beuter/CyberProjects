package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

var wait sync.WaitGroup

func scanPort(host string) {
	connect, err := net.Dial("tcp", host) //Trying to make connection
	port := host[len(host)-3:]
	if err != nil {
		//This would print each port that we failed to connect to
		fmt.Println("Port " + port + " is closed")
	}

	if connect != nil { // IF we have a connection, then record that
		fmt.Println("Port " + port + " is open")
	}

	defer wait.Done() //When the function is done then decrement 
}

// Function should take an IP address as a parameter and enumerate that IP
func scanIP(address string) {
	for i := 200; i < 224; i++ {
		wait.Add(1)
		port := strconv.Itoa(i)
		s := net.JoinHostPort(address, port)
		go scanPort(s)
	}
	wait.Wait() 
}

func main() {
	scanIP("localhost")
}
