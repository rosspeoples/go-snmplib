package main

import (
	"log"
	"net"

	snmplib "github.com/deejross/go-snmplib"
)

type snmpHandler struct{}

func (h snmpHandler) OnError(addr net.Addr, err error) {
	log.Println(addr.String(), err)
}

func (h snmpHandler) OnTrap(addr net.Addr, trap snmplib.Trap) {
	log.Println(addr.String(), trap)
}

func main() {
	server, err := snmplib.NewTrapServer("0.0.0.0", 162)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Listening for traps on port 162")
	server.ListenAndServe(snmpHandler{})
}
