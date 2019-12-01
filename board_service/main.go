package main

import (
	"fmt"
	"httpsniffer/network"
	"net"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Pls, input argument which is 'server' or 'connect'")
		return
	}

	switch os.Args[1] {
	case "server":
		network.BoardCastServer()

		//client
	case "connect":
		c := make(chan *net.UDPAddr, 1)
		network.BoardCastClient(c)
		fmt.Println(<-c)
	default:
		fmt.Println(os.Args[0] + " arguments is 'server' or 'connect'")
	}
}
