package imutil

import (
	"../model"
)


func BuildAnwser(to string) []byte {
	m := model.Message{
		From:"server",
		To:to,
		Text:"",
		Type:ANWSER,
		MessageId:Guid(),
	}
	return EncodeToByteArray(&m)
}

func BuildOffline(to, text string) []byte {
	m := model.Message{
		From:"server",
		To:to,
		Text:text,
		Type:OFFLINE,
		MessageId:Guid(),
	}
	return EncodeToByteArray(&m)
}
