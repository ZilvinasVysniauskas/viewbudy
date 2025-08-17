package main

/*
#cgo CFLAGS: -x objective-c -fmodules -fobjc-arc
#cgo LDFLAGS: -L${SRCDIR}/bin -lplayer
#include <stdlib.h>

// Add the typedef here so Go knows about it
typedef void (*CLogCallback)(const char*);

#include "libplayer.h"
*/
import "C"
import (
	"fmt"
	"log"
	"net"
	"time"
)

var (
	sendToPeerChan       chan string
	receivedFromPeerChan chan string
)

//export SendPauseRequest
func SendPauseRequest() {
	sendToPeerChan <- "pause"
}

//export StartGoLogic
func StartGoLogic() {
	sendToPeerChan = make(chan string)
	receivedFromPeerChan = make(chan string)

	go func() {
		connectPeerToPeer()
	}()

	go func() {
		for msg := range receivedFromPeerChan {
			if msg == "pause" {
				C.pauseCurrentVideo()
			}
			fmt.Println("Received from receivedFromPeerChan:", msg)
		}
	}()
}

const SERVER_ADDR = "70.34.248.18:9000"

func connectPeerToPeer() {
	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		log.Fatalf("Failed to create UDP socket: %v", err)
	}
	defer conn.Close()
	log.Printf("Listening on local UDP address: %s", conn.LocalAddr().String())

	serverAddr, err := net.ResolveUDPAddr("udp", SERVER_ADDR)
	if err != nil {
		log.Fatalf("Failed to resolve server address: %v", err)
	}

	log.Println("Sending initial packet to rendezvous server...")
	if _, err := conn.WriteTo([]byte("register"), serverAddr); err != nil {
		log.Fatalf("Failed to send packet to server: %v", err)
	}

	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFrom(buffer)
	if err != nil {
		log.Fatalf("Failed to read from connection: %v", err)
	}
	otherPeerAddrStr := string(buffer[:n])
	otherPeerAddr, err := net.ResolveUDPAddr("udp", otherPeerAddrStr)
	if err != nil {
		log.Fatalf("Failed to resolve other peer's address: %v", err)
	}
	log.Printf("Received other peer's address: %s", otherPeerAddr.String())

	punchAHole(conn, otherPeerAddr)

	setupDataListener(conn, buffer, otherPeerAddrStr)
	setupDataSender(conn, otherPeerAddr)

}

func setupDataListener(conn net.PacketConn, buffer []byte, otherPeerAddrStr string) {
	for {
		n, remoteAddr, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Printf("Error reading from connection: %v", err)
			continue
		}
		// Make sure the message is from the peer we expect
		if remoteAddr.String() == otherPeerAddrStr {
			receivedFromPeerChan <- string(buffer[:n])
		}
	}
}

func setupDataSender(conn net.PacketConn, otherPeerAddr *net.UDPAddr) {
	for msg := range sendToPeerChan {
		if _, err := conn.WriteTo([]byte(msg), otherPeerAddr); err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}

func punchAHole(conn net.PacketConn, otherPeerAddr *net.UDPAddr) {
	log.Println("Starting UDP hole punching...")
	for i := 0; i < 5; i++ {
		if _, err := conn.WriteTo([]byte("ping"), otherPeerAddr); err != nil {
			log.Printf("Error sending punch packet: %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {}
