//客户端发送封包
package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"../imutil"
	iconn "../connection"

)

func sender(conn net.Conn) {

	text := `{"From":"1","To":"2","Text":"我是1", "Type": 0}`
	conn.Write(imutil.Packet([]byte(text)))

	for i := 0; i < 2; i++ {
		time.Sleep(time.Second * 5)
		text = `{"From":"1","To":"2","Text":"我是1", "Type": 1}`
		conn.Write(imutil.Packet([]byte(text)))
	}
}

func main() {
	server := "127.0.0.1:8090"
	tcpAddr, err := net.ResolveTCPAddr("tcp6", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	imutil.Ilog("connected")

	go iconn.HandleConnection(conn, reader)
	go sender(conn)

	for {
		time.Sleep(time.Hour * 1)
	}
}

func reader(readerChannel chan []byte, conn net.Conn) {
	for {
		select {
		case data := <-readerChannel:
			imutil.Ilog(string(data))
		}
	}
}
