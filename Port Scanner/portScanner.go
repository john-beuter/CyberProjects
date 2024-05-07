package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

var wait sync.WaitGroup

func scanPort(port int, c chan string, address string) {

	for i := range port {
		portValue := strconv.Itoa(i)
		host := net.JoinHostPort(address, portValue)
		connect, err := net.Dial("tcp", host) //Trying to make connection
		portValue = host[len(host)-2:]
		var output string
		if err != nil {
			//This would print each port that we failed to connect to
			//fmt.Println("Port " + port + " is closed")
			output = portValue + " is closed"
		}

		if connect != nil { // IF we have a connection, then record that
			//fmt.Println("Port " + port + " is open")
			output = portValue + " is open"
		}

		c <- output
	}
}

// Function should take an IP address as a parameter and enumerate that IP
func scanIP(address string) {
	c := make(chan string, 24)
	go scanPort(cap(c), c, address) //use the first in as a parameter for the loop.
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	scanIP("localhost")
}
