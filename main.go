package main

import (
	"fmt"
	"github.com/google/gopacket"
	_ "github.com/google/gopacket"
	"github.com/google/gopacket/layers"
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
	handle, err := pcap.OpenLive(netCardName, 1600, true, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	//设置过滤
	if err := handle.SetBPFFilter("tcp and (port 80)"); err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	//for packet := range packetSource.Packets() {
	//	// Process packet here
	//	log.Println(packet)
	//}
	packets := packetSource.Packets()

	for {
		select{
		case packet := <-packets:
			if packet == nil {
				return
			}

			//log.Println("Packet Detail:")
			//log.Println(packet)
			log.Println("TCP:")
			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				log.Println("Unusable packet")
				continue
			}
			tcp := packet.TransportLayer().(*layers.TCP)
			//log.Println(tcp)
			log.Println(tcp.TransportFlow())
			log.Printf("payload:%v\n", string(tcp.BaseLayer.Payload))

		}
	}
}

