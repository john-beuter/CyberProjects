package main

import (
	"fmt"
	"net"

	//"fmt"
	"strconv"
)

// Function should take an IP address as a parameter and enumerate that IP
// Blast it with packets and listen for the connection?
func scanIP(address string) {

	for i := 0; i < 24; i++ {
		port := strconv.Itoa(i)
		s := net.JoinHostPort(address, port)
		connect, err := net.Dial("tcp", s) //Trying to make connection

		if err != nil {
			//fmt.Println("failed to connect port: " + port) //If the connection fails then log that the we were unable to establish connection
		}
		
		if connect != nil{
			fmt.Println("Port " + port + " is open")

		}
	}

	//Only listen back on a select IP

}

func main() {
	scanIP("192.168.4.76")
}
