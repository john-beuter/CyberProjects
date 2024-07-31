package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/goburrow/modbus" //Import the necessary modbus library
)

func modbusConnection(ip_address string, port string) modbus.Client {
	target := ip_address + ":" + port //Send Modbus traffic the specified IP on the Modbus port
	handler := modbus.NewTCPClientHandler(target)
	handler.Timeout = 10 * time.Second

	err := handler.Connect()
	if err != nil {
		fmt.Print(err)
	}

	defer handler.Close()
	client := modbus.NewClient(handler)
	return client
}

func toggle(ip_address string, port string, coil uint16, coil_value int) {
	if coil_value == 1 {
		results, err := modbusConnection(ip_address, port).WriteSingleCoil(coil, 0xFF00) //Writes a one to coil at position 0. 0xFF00 for 1 and 0x0000 for 0
		if err != nil {
			fmt.Println(err)
			fmt.Println(results)
		}
	} else {
		results, err := modbusConnection(ip_address, port).WriteSingleCoil(coil, 0x0000)
		if err != nil {
			fmt.Println(err)
			fmt.Println(results)
		}
	}
}

func main() {

	argsWithoutProg := os.Args[1:]
	var ip_address string
	var port string
	var coil uint16
	var coil_value int

	if len(argsWithoutProg) < 3 {
		fmt.Println("Enter an ip address")
		fmt.Scanln(&ip_address)

		fmt.Println("Enter a port (default Modbus port is 502)")
		fmt.Scanln(&port)
		if port == "" {
			port = "502"
		}

		fmt.Println("Enter a coil:")
		fmt.Scanln(&coil)

		fmt.Println("Enter a coil value:")
		fmt.Scanln(&coil_value)
	} else {

		var arrIPs = [3]string{"192.168.12.151", "192.168.13.64", "192.168.13.86"}
		//ip_selected, err1 := strconv.Atoi(argsWithoutProg[0])
		var ip_selected int

		if argsWithoutProg[0] == "PDU" {
			ip_selected = 0
		} else if argsWithoutProg[0] == "Master" {
			ip_selected = 1
		} else if argsWithoutProg[0] == "John" {
			ip_selected = 2
		}

		coil_selected, err1 := strconv.Atoi(argsWithoutProg[1])
		coil_value_selected, err2 := strconv.Atoi(argsWithoutProg[2])

		ip_address = arrIPs[ip_selected]

		if err1 != nil {
			fmt.Println("Could not convert arg to int")
		}

		if err2 != nil {
			fmt.Println("Could not convert arg to int")
		}

		port = "502"
		coil = uint16(coil_selected)
		coil_value = coil_value_selected
	}

	toggle(ip_address, port, coil, coil_value)

}
