package main

import (
	"fmt"
	"github.com/LiamHaworth/go-tproxy"
	"github.com/spf13/viper"
	"io"
	"log"
	"net"
	"net/http"
	"transparent-proxifier/vagrant/conf"
)

// main will initialize the TProxy
// handling application
func main() {
	conf.InitConfig()

	log.Println("Starting GoLang TProxy example")

	go startListener(viper.GetInt(conf.TcpTransparentProxyPort), connectDirectly)
	startListener(viper.GetInt(conf.TlsTransparentProxyPort), issueConnectRequest)

	log.Println("TProxy listener closing")
}

func startListener(port int, remoteDialer func(dest string) (net.Conn, error)) {
	log.Println(fmt.Sprintf("Binding TCP TProxy listener to 0.0.0.0:%d", port))

	tcpListener, err := tproxy.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: port})
	if err != nil {
		log.Panicln("Failed to establish TCP listener on port", port, err)
	}

	listen(tcpListener, remoteDialer)
}

// listenTCP runs in a routine to
// accept TCP connections and hand them
// off into their own routines for handling
func listen(listener net.Listener, remoteDialer func(dest string) (net.Conn, error)) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				log.Printf("Temporary error while accepting connection: %s", netErr)
			}

			log.Fatalf("Unrecoverable error while accepting connection: %s", err)
			return
		}

		go handleTCPConn(conn, remoteDialer)
	}
}

func handleTCPConn(conn net.Conn, remoteDialer func(dest string) (net.Conn, error)) {
	log.Printf("Accepting TCP connection from %s with destination of %s", conn.RemoteAddr().String(), conn.LocalAddr().String())
	defer conn.Close()

	remoteConn, err := remoteDialer(conn.LocalAddr().String())
	if err != nil {
		log.Fatalln(err)
	}
	defer remoteConn.Close()

	streamConn := func(dst io.Writer, src io.Reader) {
		io.Copy(dst, src)
	}
	go streamConn(remoteConn, conn)
	streamConn(conn, remoteConn)
}

func issueConnectRequest(dest string) (net.Conn, error) {

	conn, err := net.Dial("tcp", viper.GetString(conf.ProxyLocation))
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		Dial: func(network, address string) (net.Conn, error) {
			return conn, nil
		},
	}
	pr, _ := io.Pipe()
	req, err := http.NewRequest("CONNECT", "http://"+dest, pr)
	if err != nil {
		return nil, err
	}
	resp, err := tr.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	log.Printf("resp: %v", resp)
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return conn, nil
}

func connectDirectly(dest string) (net.Conn, error) {
	return net.Dial("tcp", dest)
}
