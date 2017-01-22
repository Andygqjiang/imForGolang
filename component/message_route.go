package component

import (
	cm "../connection"
	"../imutil"
	m "../model"
)

func DispatchSingle(msg *m.Message) {

	//写入redis成离线消息
	OffileHandlerSingle(msg)

	//写回执给发送方， 告知发送方，服务器已接收到消息
	srcConn := cm.Get(msg.From)
	srcConn.Write(imutil.Packet(imutil.BuildAnwser(msg.From)))
	//路由至接收方
	targetConn := cm.Get(msg.To)

	if targetConn != nil {
		json := imutil.Encode(msg)
		targetConn.Write(imutil.PacketMsg(json))
	}
}
