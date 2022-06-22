// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/iwind/TeaGo/Tea"
	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/iwind/TeaGo/types"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

type ServiceInfo struct {
	Name     string        `json:"name"`
	Methods  []*MethodInfo `json:"methods"`
	Filename string        `json:"filename"`
	Doc      string        `json:"doc"`
}

type MethodInfo struct {
	Name                string `json:"name"`
	RequestMessageName  string `json:"requestMessageName"`
	ResponseMessageName string `json:"responseMessageName"`
	Code                string `json:"code"`
	Doc                 string `json:"doc"`
}

type MessageInfo struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Doc  string `json:"doc"`
}

type RPCList struct {
	Services []*ServiceInfo `json:"services"`
	Messages []*MessageInfo `json:"messages"`
}

func readComments(data []byte) string {
	var lines = bytes.Split(data, []byte{'\n'})
	var comments = [][]byte{}
	for i := len(lines) - 1; i >= 0; i-- {
		var line = bytes.TrimLeft(lines[i], " \t")
		if len(line) == 0 {
			comments = append([][]byte{{' '}}, comments...)
			continue
		}

		if bytes.HasPrefix(line, []byte("//")) {
			line = bytes.TrimSpace(bytes.TrimLeft(line, "/"))
			comments = append([][]byte{line}, comments...)
		} else {
			break
		}
	}
	return string(bytes.TrimSpace(bytes.Join(comments, []byte{'\n'})))
}

// 生成JSON格式API列表
func main() {
	var quiet = false
	flag.BoolVar(&quiet, "quiet", false, "")
	flag.Parse()

	var dirs = []string{Tea.Root + "/../pkg/rpc/protos/", Tea.Root + "/../pkg/rpc/protos/models"}

	var services = []*ServiceInfo{}
	var messages = []*MessageInfo{}

	for _, dir := range dirs {
		func(dir string) {
			dir = filepath.Clean(dir)

			files, err := filepath.Glob(dir + "/*.proto")
			if err != nil {
				fmt.Println("[ERROR]list proto files failed: " + err.Error())
				return
			}

			for _, path := range files {
				func(path string) {
					data, err := ioutil.ReadFile(path)
					if err != nil {
						fmt.Println("[ERROR]" + err.Error())
						return
					}

					// 先将rpc代码替换成临时代码
					var methodCodeMap = map[string][]byte{} // code => method
					var methodIndex = 0
					var methodReg = regexp.MustCompile(`rpc\s+(\w+)\s*\(\s*(\w+)\s*\)\s*returns\s*\(\s*(\w+)\s*\)\s*;`)
					data = methodReg.ReplaceAllFunc(data, func(methodData []byte) []byte {
						methodIndex++
						var code = "METHOD" + types.String(methodIndex)
						methodCodeMap[code] = methodData
						return []byte("\n" + code)
					})

					// 服务列表
					// TODO 这里需要改进一下，当前实现方法如果方法注释里有括号（}），就会导致部分方法解析不到
					var serviceNameReg = regexp.MustCompile(`(?sU)\n\s*service\s+(\w+)\s*\{(.+)}`)
					var serviceMatches = serviceNameReg.FindAllSubmatch(data, -1)
					var serviceNamePositions = serviceNameReg.FindAllIndex(data, -1)
					for serviceMatchIndex, serviceMatch := range serviceMatches {
						var serviceName = string(serviceMatch[1])
						var serviceNamePosition = serviceNamePositions[serviceMatchIndex][0]
						var comment = readComments(data[:serviceNamePosition])

						// 方法列表
						var methods = []*MethodInfo{}
						var serviceData = serviceMatch[2]
						var methodCodeReg = regexp.MustCompile(`\b(METHOD\d+)\b`)
						var methodCodeMatches = methodCodeReg.FindAllSubmatch(serviceData, -1)
						var methodCodePositions = methodCodeReg.FindAllIndex(serviceData, -1)
						for methodMatchIndex, methodMatch := range methodCodeMatches {
							var methodCode = string(methodMatch[1])
							var methodData = methodCodeMap[methodCode]
							var methodPieces = methodReg.FindSubmatch(methodData)
							var methodCodePosition = methodCodePositions[methodMatchIndex]

							methods = append(methods, &MethodInfo{
								Name:                string(methodPieces[1]),
								RequestMessageName:  string(methodPieces[2]),
								ResponseMessageName: string(methodPieces[3]),
								Code:                string(methodData),
								Doc:                 readComments(serviceData[:methodCodePosition[0]]),
							})
						}

						services = append(services, &ServiceInfo{
							Name:     serviceName,
							Methods:  methods,
							Filename: filepath.Base(path),
							Doc:      comment,
						})
					}

					// 消息列表
					var topMessageCodeMap = map[string][]byte{} // code => message
					var allMessageCodeMap = map[string][]byte{}
					var messageCodeIndex = 0
					var messagesReg = regexp.MustCompile(`(?sU)\n\s*message\s+(\w+)\s*\{([^{}]+)\n\s*}`)
					var firstMessagesReg = regexp.MustCompile(`message\s+(\w+)`)
					var messageCodeREG = regexp.MustCompile(`MESSAGE\d+`)
					for {
						var hasMessage = false

						data = messagesReg.ReplaceAllFunc(data, func(messageData []byte) []byte {
							messageCodeIndex++
							hasMessage = true

							// 是否包含子Message
							var subMatches = messageCodeREG.FindAllSubmatch(messageData, -1)
							for _, subMatch := range subMatches {
								var subMatchCode = string(subMatch[0])
								delete(topMessageCodeMap, subMatchCode)
							}

							var code = "MESSAGE" + types.String(messageCodeIndex)
							topMessageCodeMap[code] = messageData
							allMessageCodeMap[code] = messageData
							return []byte("\n" + code)
						})
						if !hasMessage {
							break
						}
					}

					for messageCode, messageData := range topMessageCodeMap {
						// 替换其中的子Message
						for {
							if messageCodeREG.Match(messageData) {
								messageData = messageCodeREG.ReplaceAllFunc(messageData, func(messageCodeData []byte) []byte {
									return allMessageCodeMap[string(messageCodeData)]
								})
							} else {
								break
							}
						}

						// 注释
						var index = bytes.Index(data, []byte(messageCode))
						var messageName = string(firstMessagesReg.FindSubmatch(messageData)[1])
						messages = append(messages, &MessageInfo{
							Name: messageName,
							Code: string(bytes.TrimSpace(messageData)),
							Doc:  readComments(data[:index]),
						})
					}
				}(path)
			}
		}(dir)
	}

	var countServices = len(services)
	var countMethods = 0
	var countMessages = len(messages)
	for _, service := range services {
		countMethods += len(service.Methods)
	}

	var rpcList = &RPCList{
		Services: services,
		Messages: messages,
	}
	jsonData, err := json.MarshalIndent(rpcList, "", "  ")
	if err != nil {
		fmt.Println("[ERROR]marshal to json failed: " + err.Error())
		return
	}

	var jsonFile = Tea.Root + "/rpc.json"
	err = ioutil.WriteFile(jsonFile, jsonData, 0666)
	if err != nil {
		fmt.Println("[ERROR]write json to file failed: " + err.Error())
		return
	}

	if !quiet {
		fmt.Println("services:", countServices, "methods:", countMethods, "messages:", countMessages)
		fmt.Println("===")
		fmt.Println("generated " + filepath.Base(jsonFile) + " successfully")
	}
}
