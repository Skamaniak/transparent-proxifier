package main

import (
	"fmt"
	"github.com/LiamHaworth/go-tproxy"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	// tcpListener represents the TCP
	// listening socket that will receive
	// TCP connections from TProxy
	tcpListener net.Listener
)

// main will initialize the TProxy
// handling application
func main() {
	log.Println("Starting GoLang TProxy example")
	var err error

	log.Println("Binding TCP TProxy listener to 0.0.0.0:8080")
	tcpListener, err = tproxy.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 8080})
	if err != nil {
		log.Fatalf("Encountered error while binding listener: %s", err)
		return
	}

	defer tcpListener.Close()
	go listenTCP()

	interruptListener := make(chan os.Signal)
	signal.Notify(interruptListener, os.Interrupt)
	<-interruptListener

	log.Println("TProxy listener closing")
}

// listenTCP runs in a routine to
// accept TCP connections and hand them
// off into their own routines for handling
func listenTCP() {
	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				log.Printf("Temporary error while accepting connection: %s", netErr)
			}

			log.Fatalf("Unrecoverable error while accepting connection: %s", err)
			return
		}

		go handleTCPConn(conn)
	}
}

// handleTCPConn will open a connection
// to the original destination pretending
// to be the client. From there it will setup
// two routines to stream data between the
// connections
func handleTCPConn(conn net.Conn) {
	log.Printf("Accepting TCP connection from %s with destination of %s", conn.RemoteAddr().String(), conn.LocalAddr().String())
	defer conn.Close()

	remoteConn, err := net.Dial("tcp", conn.LocalAddr().String())
	if err != nil {
		log.Printf("Failed to connect to original destination [%s]: %s", conn.LocalAddr().String(), err)
		return
	}

	streamConn := func(dst io.Writer, src io.Reader) {
		io.Copy(scanningWriter{dst}, src)
	}
	go streamConn(remoteConn, conn)
	streamConn(conn, remoteConn)
}

type scanningWriter struct {
	inner io.Writer
}

func (c scanningWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return c.inner.Write(p)
}
