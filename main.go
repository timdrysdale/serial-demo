package main

import (
	"fmt"
	"log"
	"os"
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

	ready := make(chan interface{})

	go func() {
		for {
			select {
			case <-time.After(5 * time.Second):
				fmt.Println("not ready after 5sec, exiting")
				os.Exit(1) //timeout opening channel
			default:
				buff := make([]byte, 640)
				n, err := port.Read(buff)
				if err != nil {
					log.Fatal(err)
				}
				if n == 0 {
					fmt.Println("\nEOF")

				}
				fmt.Printf("ready: %v", string(buff[:n]))
				close(ready)
				return
			}
		}
	}()

	<-ready

	done := make(chan interface{})

	go func() {
		i := 0
		for {
			select {
			case <-done:
				fmt.Println("done reading")
				return
			default:
				buff := make([]byte, 640)
				n, err := port.Read(buff)
				if err != nil {
					log.Fatal(err)
				}
				if n == 0 {
					fmt.Println("\nEOF")

				}
				i = i + 1
				fmt.Printf("read %d: %v", i, string(buff[:n]))
			}
		}
	}()

	time.Sleep(2 * time.Second)

	n, err := port.Write([]byte(`{"set":"port","to":"open"}\n\r`))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

	time.Sleep(time.Second)

	n, err = port.Write([]byte(`{"set":"port","to":"short"}\n\r`))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

	time.Sleep(time.Second)
	n, err = port.Write([]byte(`{"set":"port","to":"open"}\n\r`))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

	time.Sleep(time.Second)

	time.Sleep(5 * time.Second)

	close(done)

}
