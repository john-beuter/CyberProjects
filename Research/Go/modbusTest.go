package main

import (
	"fmt"
	"time"

	"github.com/simonvetter/modbus"
)

func main() {
	var client *modbus.ModbusClient
	var err error

	// for a TCP endpoint
	// (see examples/tls_client.go for TLS usage and options)
	client, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://192.168.14.236:502",
		Timeout: 1 * time.Second,
	})

	err = client.WriteCoil(0, false)
	fmt.Println("Done")

	if err != nil {
		fmt.Println(err)
	}

}
