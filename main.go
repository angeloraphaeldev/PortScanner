package main

import (
	"portScanner/port"
)

func main() {
	port.GetOpenPorts("www.freecodecamp.com", port.PortRange{Start: 80, End: 85})

	port.GetOpenPorts("104.26.10.78", port.PortRange{Start: 8079, End: 8090})
	// Simple open port
	port.GetOpenPorts("104.26.10.78", port.PortRange{Start: 440, End: 450})
	// Simple open port
	port.GetOpenPorts("137.74.187.104", port.PortRange{Start: 440, End: 450})
	// Multiples Ports Returned
	port.GetOpenPorts("scanme.nmap.org", port.PortRange{Start: 20, End: 80})
	// Invalid Ip
	port.GetOpenPorts("www.dsaldkasdlsakdlas.com", port.PortRange{Start: 0, End: 100})
}
