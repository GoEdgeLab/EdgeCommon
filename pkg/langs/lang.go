// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package langs

import (
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"strings"
)

const varPrefix = "lang."

type LangCode = string

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

// Compile variable to literal strings
func (this *Lang) Compile() error {
	for code, oldMessage := range this.messageMap {
		message, err := this.get(code, 0)
		if err != nil {
			return errors.New("compile '" + code + "': '" + oldMessage + "' failed: " + err.Error())
		}
		this.messageMap[code] = message
	}
	return nil
}

func (this *Lang) get(messageCode string, loopIndex int) (string, error) {
	if loopIndex >= 8 /** max recurse **/ {
		return "", errors.New("too many recurse")
	}
	loopIndex++

	message, ok := this.messageMap[messageCode]
	if len(message) == 0 {
		if !ok && loopIndex > 1 {
			// recover as variable
			return "${" + varPrefix + messageCode + "}", errors.New("can not find message for code '" + messageCode + "'")
		}
		return "", nil
	}

	return configutils.ParseVariablesError(message, func(varName string) (value string, err error) {
		if !strings.HasPrefix(varName, varPrefix) {
			return "${" + varName + "}", nil
		}

		return this.get(varName[len(varPrefix):], loopIndex)
	})
}
