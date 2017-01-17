package main

import (
	"net"

	"../component"
	iconn "../connection"
	"../imutil"
)

func main() {
	ln, err := net.Listen("tcp", ":8090")
	if err != nil {
		imutil.Ilog(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			imutil.Ilog(err)
			return
		}
		go iconn.HandleConnection(conn, reader)
	}

}

func reader(readerChannel chan []byte, conn net.Conn) {
	for {
		select {
		case data := <-readerChannel:
			message := imutil.Parse(data)
			if message.Type == imutil.LOGIN {
				iconn.Put(message.From, conn)
			} else if message.Type == imutil.SINGLE_CHAT_REQ {
				//单聊发信息
				component.Dispatch(message)
			}
		}
	}
}
