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

				//推送离线消息至客户端
				offlineMsg := imutil.HgetAll(message.From)
				if offlineMsg !=nil && len(offlineMsg) > 0 {
					offlineMsgs := make([]string, len(offlineMsg))
					index := 0
					for _,v := range offlineMsg {
						offlineMsgs[index] = v
						index++
					}
					offline := imutil.BuildOffline(message.From, string(imutil.EncodeToByteArray(offlineMsgs)))
					conn.Write(offline)
				}
			} else if message.Type == imutil.SINGLE_CHAT_REQ {
				//单聊发信息
				component.DispatchSingle(message)
			}
		}
	}
}
