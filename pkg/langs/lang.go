// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package langs

type Lang struct {
	code string

	messageMap map[string]string // message code => message text
}

func NewLang(code string) *Lang {
	return &Lang{
		code:       code,
		messageMap: map[string]string{},
	}
}

func (this *Lang) Set(messageCode string, messageText string) {
	this.messageMap[messageCode] = messageText
}

// Get 读取单条消息
// get single message with message code
func (this *Lang) Get(messageCode string) string {
	return this.messageMap[messageCode]
}

// GetAll 读取所有消息
// get all messages
func (this *Lang) GetAll() map[string]string {
	return this.messageMap
}
