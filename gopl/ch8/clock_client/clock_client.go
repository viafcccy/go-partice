package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	for index, val := range os.Args {
		if index == 0 {
			continue
		}

		go dial(val)
	}

	func() {
		forever := make(chan bool)
		<-forever
	}()

}

func dial(port string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
