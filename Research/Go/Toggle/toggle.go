package main

import (
	"fmt"
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
	var ip_address string
	fmt.Println("Enter an ip address")
	fmt.Scanln(&ip_address)

	var port string
	fmt.Println("Enter a port (default Modbus port is 502)")
	fmt.Scanln(&port)
	if port == "" {
		port = "502"
	}

	var coil uint16
	fmt.Println("Enter a coil:")
	fmt.Scanln(&coil)

	var coil_value int
	fmt.Println("Enter a coil value:")
	fmt.Scanln(&coil_value)

	toggle(ip_address, port, coil, coil_value)

}
