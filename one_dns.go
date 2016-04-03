// One reply trivial DNS server
//
// Speed 112922 qps
// (C) Dmitry Podkorytov 2016
//
// Run it and test resolving by nslookup 1000.dip 127.0.0.1 
// and do benchmarks by dnsperf or something else
//

package main

import (
	"flag"
	"fmt"
	"net"
	"log"
)

var port = flag.Int("p", 53, "The listening port")
var udpPackageBufferSize = flag.Int("l", 1024, "The size of the udp package buffer")

func main() {
	flag.Parse()

	// open socket
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: *port,
	})

	if err != nil { fmt.Println("listen failed!", err)
		        return }

        log.Println("listen ", socket,port,err)

	defer socket.Close()

	for {
		// endless loop wait clients asks
		ask_data := make([]byte, *udpPackageBufferSize)
		readn, remoteAddr, _ := socket.ReadFromUDP(ask_data)

                //for fast working dont control returned error
		//if err != nil { fmt.Println("recvfrom error!", err)
		//	        continue }
		
		//go   
                process(socket, ask_data[:readn], remoteAddr)
	}
}


// here do reply on ask data
//
func process(conn *net.UDPConn, ask_data []byte, remoteAddr *net.UDPAddr) {

                id:=make([]byte,2)

                id[0]=ask_data[0]
                id[1]=ask_data[1]
                // compiled record with 1000.dip A DNS name and some IP
                rr_record := []byte{0,0,1,0,0,1,0,1,0,0,0,0,4,49,48,48,48,3,100,105,112,0,0,1,0,1,192,12,0,1,0,1,0,0,8,212,0,4,87,118,90,81}

                rr_record[0]=id[0]
                rr_record[1]=id[1]

 conn.WriteToUDP(rr_record, remoteAddr)

}
