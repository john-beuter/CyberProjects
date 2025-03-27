package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
)

var wait sync.WaitGroup

func scanPort(c chan string, address string) {
	//The capacity of the channel is the port range. This meanss that each routine would scan 24 ports.
	connect, err := net.Dial("tcp", address) //Trying to make connection
	portValue := address[len(address)-2:]
	portValue = strings.ReplaceAll(portValue, ":", "")
	var output string

	if err != nil {
		output = portValue + " is closed"
	} else { // IF we have a connection, then record that
		//fmt.Println("Port " + port + " is open")
		output = portValue + " is open"
		connect.Close()
	}
	c <- output
}

func scanIP(address string) {
	c := make(chan string, 24) //Make 24 channnels to communicate through.
	for i := 0; i < cap(c); i++ {
		portValue := strconv.Itoa(i)
		host := net.JoinHostPort(address, portValue)
		go scanPort(c, host) //use the first in as a parameter for the loop.
	}

	for i := range c { //channel only returns one value from range
		fmt.Println(i)
	} //For each port
}

func main() {
	scanIP("localhost")
}

//This is a test change
