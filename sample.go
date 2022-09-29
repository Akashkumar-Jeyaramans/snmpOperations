package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	conn, err := net.DialTimeout("udp", "10.6.50.8:161", 2*time.Second)
	if err != nil {
		fmt.Printf(`error connecting to ("udp", "10.6.50.8") : %s `, err)
	} else {
		fmt.Print("connected", conn)
	}
}
