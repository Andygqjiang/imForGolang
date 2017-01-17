package component

import (
	cm "../connection"
	"../imutil"
	m "../model"
)

func Dispatch(msg *m.Message) {
	conn := cm.Get(msg.To)
	json := imutil.Encode(msg)
	conn.Write(imutil.Packet([]byte(json)))
}
