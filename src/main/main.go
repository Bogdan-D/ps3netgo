package main

import (
	"log"
	"net"

	"os"

	"github.com/alecthomas/kingpin"
)

const (
	MIN_PORT = 1024
	MAX_PORT = 65535

	BUFFER_SIZE = 3 * 1048576
	MAX_CLIENTS = 5
)

const (
	VISO_NONE = iota
	VISO_PS3
	VISO_ISO
)

var (
	root = kingpin.Arg("dir", "root directory").Required().ExistingDir()
	port = kingpin.Flag("port", "listen port, default: 38008").Short('p').Default("38008").Int()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	os.Chdir(*root)

	if *port < MIN_PORT || *port > MAX_PORT {
		log.Fatalf("Port should be in %d-%d range", MIN_PORT, MAX_PORT)
	}

	listen := &net.TCPAddr{IP: nil, Port: *port}
	socket, err := net.ListenTCP("tcp", listen)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Waiting for client...")

	for {
		conn, err := socket.AcceptTCP()

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Connection from %s", conn.RemoteAddr().String())

		c := &client{
			socket: conn,
		}
		go func() {
			c.Run()
			defer c.socket.Close()
		}()
	}
}
