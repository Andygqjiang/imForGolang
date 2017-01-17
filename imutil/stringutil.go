package imutil

import (
	json "encoding/json"
	"fmt"

	m "../model"
)

/*
json转对象
*/
func Parse(text []byte) *m.Message {
	message := &m.Message{}
	json.Unmarshal(text, message)
	return message
}

/*
对象转json
*/
func Encode(msg *m.Message) string {
	r, err := json.Marshal(msg)
	if err != nil {
		println(err)
	}
	return string(r)
}

/*
调试输出
*/
func Ilog(v ...interface{}) {
	fmt.Println(v...)
}
