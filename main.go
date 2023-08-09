package main

import (
	"net"
	"time"
)

func main() {
	// Set up a UDP broadcast address and port
	broadcastAddr := &net.UDPAddr{IP: net.IPv4(255, 255, 255, 255), Port: 8888}

	// Create a UDP connection to the broadcast address
	conn, err := net.DialUDP("udp", nil, broadcastAddr)
	if err != nil {
		panic(err)
	}

	// Loop forever, sending a message every second
	for {
		msg := []byte("Hello, world!")
		//fmt.Println("hello world")
		_, err := conn.Write(msg)
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}
