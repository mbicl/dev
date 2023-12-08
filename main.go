package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strconv"
)

var (
	Message = []byte("Pong")
)

func handler(conn net.Conn) {
	defer conn.Close()

	var (
		buf = make([]byte, 1<<10)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)

	for {
		n, err := r.Read(buf)

		switch err {
		case io.EOF:
			continue
		case nil:
			data := buf[:n]
			log.Println("Received: ", string(data))
			w.Write(Message)
			w.Flush()
			log.Println("Sent: ", Message)
		default:
			log.Panic("Failed: ", err.Error())
		}
	}

	
}

func socket(port int) {
	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer listen.Close()

	log.Printf("Listening in port: %d", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("Cannot accept connection: %s", err.Error())
			continue
		}
		go handler(conn)
	}
}

func main() {
	port := 3000
	socket(port)
}
