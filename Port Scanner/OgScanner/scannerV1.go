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
	port := host[len(host)-3:] //I could string strip the : 
	var output string
	if err != nil {
		output = port + " is closed"
	}else{ // IF we have a connection, then record that
		//fmt.Println("Port " + port + " is open")
		output = port + " is open"
		connect.Close()
		//fmt.Sprintf("%s is open", port)
	}

	fmt.Println(output)
	defer wait.Done()
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
