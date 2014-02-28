package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/akrennmair/gopcap"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	TYPE_IP  = 0x0800
	TYPE_ARP = 0x0806
	TYPE_IP6 = 0x86DD

	IP_ICMP = 1
	IP_INIP = 4
	IP_TCP  = 6
	IP_UDP  = 17
)

func NewHandle(device string, port int, snaplen int) (*pcap.Pcap, error) {
	handle, err := pcap.Openlive(device, int32(snaplen), true, 0)

	if handle == nil {
		return nil, fmt.Errorf("Couldn't Open live connection to interface: %s", err)
	}

	err = handle.Setfilter("port " + strconv.Itoa(port))

	if err != nil {
		return nil, fmt.Errorf("Failed to add filter to handle: %s", err)
	}

	return handle, nil
}

// Capture packets on the given device using libpcap. Only packets sent to or
// from the given port are captured. The packets are reassembled back into the
// HTTP traffic.
//
func sniff(handle *pcap.Pcap, sink Sink) {
	defer handle.Close()

	times := map[uint32]time.Time{}
	reqs := map[uint32]*http.Request{}

	for pkt := handle.Next(); pkt != nil; pkt = handle.Next() {
		pkt.Decode()
		if len(pkt.Payload) > 0 {
			//Parse the packet twice? Not good
			buf := bufio.NewReader(bytes.NewReader(pkt.Payload))
			req, _ := http.ReadRequest(buf)

			if req != nil {
				times[pkt.TCP.Ack] = time.Now()
				reqs[pkt.TCP.Ack] = req
				continue
			}

			buf = bufio.NewReader(bytes.NewReader(pkt.Payload))
			resp, _ := http.ReadResponse(buf, req)

			if resp != nil {
				start, _ := times[pkt.TCP.Seq]
				req, foundRequest := reqs[pkt.TCP.Seq]

				if foundRequest {
					sink.Put(req, resp, time.Now().Sub(start))
				}

				delete(times, pkt.TCP.Seq)
				delete(reqs, pkt.TCP.Seq)
			}
		}
	}
}

func main() {
	var device, esurl string
	var port, snaplen int

	flag.StringVar(&device, "i", "", "interface")
    flag.StringVar(&esurl, "e", "http://localhost:9200/coiltap", "elastic-url")
	flag.IntVar(&port, "p", 80, "port")
	flag.IntVar(&snaplen, "s", 1600, "snaplen")

	flag.Usage = func() {
		log.Fatalf("usage: %s [ -i interface ] [ -s snaplen ] [ -p port ] [-e elastic-url]", os.Args[0])
		os.Exit(1)
	}

	flag.Parse()

	if device == "" {
		devs, err := pcap.Findalldevs()
		if err != nil {
			log.Fatalf("Couldn't find any devices: %s", err)
		}
		if 0 == len(devs) {
			flag.Usage()
		}
		device = devs[0].Name
	}

	sink, err := NewSink(esurl)
    sink.Run()

	if err != nil {
		log.Fatal(err)
	}

	for {
		handle, err := NewHandle(device, port, snaplen)

		if err != nil {
			log.Fatal(err)
		}

		sniff(handle, sink)

        time.sleep(time.Second * 1)
	}
}
