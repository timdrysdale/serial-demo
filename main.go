package main

import (
	"fmt"
	"log"
	"time"

	"go.bug.st/serial"
)

func main() {

	mode := &serial.Mode{
		BaudRate: 57600,
	}

	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1)

	n, err := port.Write([]byte(`{"set":"port","to":"open"}\n\r`))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

	buff := make([]byte, 640)
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		fmt.Printf("%v", string(buff[:n]))
	}

}
