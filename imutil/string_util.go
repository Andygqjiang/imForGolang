package imutil

import (
	"encoding/json"
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
对象转json字串
*/
func Encode(msg *m.Message) string {
	r, err := json.Marshal(msg)
	if err != nil {
		println(err)
	}
	return string(r)
}

/*
对象转json byte数组
*/
func EncodeToByteArray(v interface{}) []byte {
	r, err := json.Marshal(v)
	if err != nil {
		println(err)
	}
	return r
}

/*
调试输出
*/
func Ilog(v ...interface{}) {
	fmt.Println(v...)
}


