package main

import (
	"code.google.com/p/go.net/ipv4"
	"log"
	"net"
)

func main() {

        c, err := net.ListenPacket("ip", "0.0.0.0")

	if err != nil {
		log.Fatal(err)
	}

	conn := ipv4.NewPacketConn(c)

	if err != nil {
		log.Fatal(err)
	}

    var b []byte

    n, _, src, err := conn.ReadFrom(b)

	log.Println("%+v", b)
	log.Println("%+v", n)
	log.Println("%+v", src)
	log.Println("%+v", err)
}
