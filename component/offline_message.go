package component

import (
	"../imutil"
	m "../model"
)

func OffileHandlerSingle(msg *m.Message) {
	msg.MessageId = imutil.Guid()
	imutil.Hset(msg.To, msg.MessageId, imutil.Encode(msg), 60*5)
}