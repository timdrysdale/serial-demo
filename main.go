package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/google/gousb"
)

func main() {
	// Initialize a new Context.
	ctx := gousb.NewContext()

	ctx.Debug(4)

	defer ctx.Close()

	// Open any device with a given VID/PID using a convenience function.
	dev, err := ctx.OpenDeviceWithVIDPID(0x0403, 0x6001)
	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
	}
	defer dev.Close()

	err = dev.SetAutoDetach(true)
	if err != nil {
		log.Fatalf("Could not set AutoDetach true: %v", err)
	}
	/*
		err = dev.Reset()
		if err != nil {
			log.Fatalf("Could not set reset connection: %v", err)
		}*/

	// Claim the default interface using a convenience function.
	// The default interface is always #0 alt #0 in the currently active
	// config.
	intf, done, err := dev.DefaultInterface()
	if err != nil {
		log.Fatalf("%s.DefaultInterface(): %v", dev, err)
	}
	defer done()

	//[0x02(2,OUT) 0x81(1,IN)]

	// Note handle_events: error: libusb: interrupted [code -10]
	// https://github.com/google/gousb/issues/87

	// Open an OUT endpoint.
	ep, err := intf.OutEndpoint(2)
	if err != nil {
		log.Fatalf("%s.OutEndpoint(2): %v", intf, err)
	}

	// Generate some data to write.
	data := []byte(`{"set":"port","to":"open"}\n`)
	var numBytes int

	// Write data to the USB device.
	numBytes, err = ep.Write(data)
	if numBytes != len(data) {
		log.Fatalf("%s.Write([%d]): only %d bytes written, returned error is %v", ep, len(data), numBytes, err)
	}
	fmt.Printf("%d bytes successfully sent to the endpoint\n", len(data))

	// In this interface open endpoint #6 for reading.
	epIn, err := intf.InEndpoint(1)
	if err != nil {
		log.Fatalf("%s.InEndpoint(1): %v", intf, err)
	}

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(1 * time.Second)
			numBytes, err := ep.Write(data)
			if numBytes != len(data) {
				log.Fatalf("%s.Write([%d]): only %d bytes written, returned error is %v", ep, len(data), numBytes, err)
			}
			fmt.Printf("%d bytes successfully sent to the endpoint\n", len(data))
		}
	}()

	//go func() {
	log.Infof("%s.InEndpoint(1) MaxPacketSize: %d", intf, epIn.Desc.MaxPacketSize)
	resp := make([]byte, epIn.Desc.MaxPacketSize*3)

	for i := 0; i < 100; i++ {
		numBytes, err = epIn.Read(resp)
		if err != nil {
			log.Fatalf("%s.InEndpoint(1): %v", intf, err)
			break
		}
		log.Infof("read %d: %s\n", numBytes, string(resp[0:numBytes-1]))
		time.Sleep(100 * time.Millisecond)
		//func (e *InEndpoint) Read(buf []byte) (int, error)
	}
	//}()

	//time.Sleep(5)

}
