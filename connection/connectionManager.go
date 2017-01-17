package connection

import (
	"net"
	"../imutil"
)

var connections map[string]net.Conn

func init() {
	connections = make(map[string]net.Conn)
}

func Put(name string, conn net.Conn) {
	connections[name] = conn
}

func Get(name string) net.Conn {
	return connections[name]
}


func HandleConnection(conn net.Conn, reader func(readerChannel chan []byte, conn net.Conn)) {
	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0)

	//声明一个管道用于接收解包的数据
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel, conn)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			imutil.Ilog(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		tmpBuffer = imutil.Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
}



