package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// go run xxx.go -p 8181 -l Local

func main() {

	var port uint64
	var loc string

	flag.Uint64Var(&port, "p", 8181, "port, default 8181")
	flag.StringVar(&loc, "l", "Local", "location, default 'Local'")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, loc, port)
	}
}

func handleConn(c net.Conn, locString string, port uint64) {
	loc, err := time.LoadLocation(locString)
	if err != nil {
		fmt.Println("Error loading location:", err)
		loc, _ = time.LoadLocation("Local")
	}

	defer c.Close()
	for {
		_, err := io.WriteString(c, fmt.Sprintf("clock in port %d: %s %s",
			port, loc.String(), time.Now().In(loc).Format("15:04:05\n")))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
