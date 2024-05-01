package main

import (
	"net"
	//"fmt"
	"strconv"
)

// Function should take an IP address as a parameter and enumerate that IP
// Blast it with packets and listen for the connection?
func scanIP(address string) {

	for i := 0; i < 1000; i++ {
		s := net.JoinHostPort(address, strconv.Itoa(i))
		net.Dial("tcp", s) //Trying to make connection
	}

	//Only listen back on a select IP

}

func main() {
	println("test")
}
