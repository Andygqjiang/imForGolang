//通讯协议处理，主要处理封包和解包的过程
package imutil

import (
	"bytes"
	"encoding/binary"
)

const (
	Header         = "cx"
	HeaderLength   = len(Header)
	LengthDataLength = 4
)

//封包
func Packet(message []byte) []byte {
	return append(append([]byte(Header), IntToBytes(len(message))...), message...)
}

//解包
func Unpack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)

	var i int
	for i = 0; i < length; i++ {
		if length < i+HeaderLength+LengthDataLength {
			break
		}
		if string(buffer[i:i+HeaderLength]) == Header {
			messageLength := BytesToInt(buffer[i+HeaderLength : i+HeaderLength+LengthDataLength])
			if length < i+HeaderLength+LengthDataLength+messageLength {
				break
			}
			data := buffer[i+HeaderLength+LengthDataLength : i+HeaderLength+LengthDataLength+messageLength]
			readerChannel <- data

			i += HeaderLength + LengthDataLength + messageLength - 1
		}
	}

	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func PacketMsg(text string) []byte {
	return Packet([]byte(text))
}
