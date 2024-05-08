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
	}else{ // IF we have a connection record that then close
		output = port + " is open"
		connect.Close()
	}

	fmt.Println(output)
	defer wait.Done()
}

// Function should take an IP address as a parameter and enumerate that IP
func scanIP(address string, start int, end int) {
	for i := start; i < end; i++ {
		wait.Add(1)
		port := strconv.Itoa(i)
		s := net.JoinHostPort(address, port)
		go scanPort(s)
	}
	wait.Wait()
}

func main() {
	scanIP("localhost", 1, 35)
}
