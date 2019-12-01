package main

import (
	"httpsniffer/network"
	"httpsniffer/tcp_handle"
	"net"
)

func main(){
	c := make(chan *net.UDPAddr, 1)
	network.BoardCastClient(c)
	remoteAddr := <- c

	network.TcpController( remoteAddr.IP.String() , tcp_handle.ControllerHandle)
}
