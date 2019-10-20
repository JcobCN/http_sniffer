package main

import (
	"fmt"
	"github.com/google/gopacket"
	_ "github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"strings"
	"time"
)

func main(){
	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	var netCardName string

	// Print device information
	fmt.Println("Devices found:")
	for _, d := range devices {
		fmt.Println("\nName: ", d.Name)
		fmt.Println("Description: ", d.Description)
		fmt.Println("Devices addresses: ", d.Description)

		if strings.Contains( strings.ToLower(d.Description), "pcie") {
			netCardName = d.Name
			break
		}

		for _, address := range d.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}

	//netCardName1 := "\\Device\\NPF_{A03165A0-9781-4C35-8298-FEC0E040754A}"
	handle, err := pcap.OpenLive(netCardName, 1024, false, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		log.Println(packet)
	}
}

