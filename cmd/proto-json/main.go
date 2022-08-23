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
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type ServiceInfo struct {
	Name     string        `json:"name"`
	Methods  []*MethodInfo `json:"methods"`
	Filename string        `json:"filename"`
	Doc      string        `json:"doc"`
}

type MethodInfo struct {
	Name                string   `json:"name"`
	RequestMessageName  string   `json:"requestMessageName"`
	ResponseMessageName string   `json:"responseMessageName"`
	Code                string   `json:"code"`
	Doc                 string   `json:"doc"`
	Roles               []string `json:"roles"`
	IsDeprecated        bool     `json:"isDeprecated"`
}

type MessageInfo struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Doc  string `json:"doc"`
}

type LinkInfo struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type RPCList struct {
	Services []*ServiceInfo `json:"services"`
	Messages []*MessageInfo `json:"messages"`
	Links    []*LinkInfo    `json:"links"`
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

func removeDuplicates(s []string) []string {
	if len(s) == 0 {
		return s
	}
	var m = map[string]bool{}
	var result = []string{}
	for _, item := range s {
		_, ok := m[item]
		if ok {
			continue
		}
		result = append(result, item)
		m[item] = true
	}
	return result
}

// 生成JSON格式API列表
func main() {
	var quiet = false
	flag.BoolVar(&quiet, "quiet", false, "")
	flag.Parse()

	var methodRolesMap = map[string][]string{} // method => roles
	{
		var rootDir = filepath.Clean(Tea.Root + "/../../EdgeAPI/internal/rpc/services")
		files, err := filepath.Glob(rootDir + "/service_*.go")
		if err != nil {
			fmt.Println("[ERROR]list service implementation files failed: " + err.Error())
			return
		}

		var methodNameReg = regexp.MustCompile(`func\s*\(\w+\s+\*\s*(\w+Service)\)\s*(\w+)\s*\(`) // $1: serviceName, $2 methodName
		for _, file := range files {
			data, err := os.ReadFile(file)
			if err != nil {
				fmt.Println("[ERROR]read file '" + file + "' failed: " + err.Error())
				return
			}
			var sourceCode = string(data)

			var locList = methodNameReg.FindAllStringIndex(sourceCode, -1)
			for index, loc := range locList {
				var methodSource = ""
				if index == len(locList)-1 { // last one
					methodSource = sourceCode[loc[0]:]
				} else {
					methodSource = sourceCode[loc[0]:locList[index+1][0]]
				}

				// 方法名
				var submatch = methodNameReg.FindStringSubmatch(methodSource)
				if len(submatch) == 0 {
					continue
				}
				var serviceName = submatch[1]
				if serviceName == "BaseService" {
					continue
				}
				var methodName = submatch[2]
				if methodName[0] < 'A' || methodName[0] > 'Z' {
					continue
				}
				var roles = []string{}
				if strings.Contains(methodSource, ".ValidateNode(") {
					roles = append(roles, "node")
				}
				if strings.Contains(methodSource, ".ValidateUserNode(") {
					roles = append(roles, "user")
				}
				if strings.Contains(methodSource, ".ValidateAdmin(") {
					roles = append(roles, "admin")
				}
				if strings.Contains(methodSource, ".ValidateAdminAndUser(") {
					roles = append(roles, "admin", "user")
				}
				if strings.Contains(methodSource, ".ValidateNSNode(") {
					roles = append(roles, "dns")
				}
				if strings.Contains(methodSource, ".ValidateMonitorNode(") {
					roles = append(roles, "monitor")
				}
				if strings.Contains(methodSource, "rpcutils.UserTypeDNS") {
					roles = append(roles, "dns")
				}
				if strings.Contains(methodSource, "rpcutils.UserTypeUser") {
					roles = append(roles, "user")
				}
				if strings.Contains(methodSource, "rpcutils.UserTypeNode") {
					roles = append(roles, "node")
				}
				if strings.Contains(methodSource, "rpcutils.UserTypeMonitor") {
					roles = append(roles, "monitor")
				}
				if strings.Contains(methodSource, "rpcutils.UserTypeReport") {
					roles = append(roles, "report")
				}
				if strings.Contains(methodSource, "rpcutils.UserTypeCluster") {
					roles = append(roles, "cluster")
				}
				if strings.Contains(methodSource, "rpcutils.UserTypeAdmin") {
					roles = append(roles, "admin")
				}

				methodRolesMap[strings.ToLower(methodName)] = removeDuplicates(roles)
			}
		}
	}

	var services = []*ServiceInfo{}
	var messages = []*MessageInfo{}

	{
		var dirs = []string{Tea.Root + "/../pkg/rpc/protos/", Tea.Root + "/../pkg/rpc/protos/models"}
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
						var filename = filepath.Base(path)
						if filename == "service_authority_key.proto" || filename == "service_authority_node.proto" {
							return
						}

						data, err := os.ReadFile(path)
						if err != nil {
							fmt.Println("[ERROR]" + err.Error())
							return
						}

						// 先将rpc代码替换成临时代码
						var methodCodeMap = map[string][]byte{} // code => method
						var methodIndex = 0
						var methodReg = regexp.MustCompile(`(?s)rpc\s+(\w+)\s*\(\s*(\w+)\s*\)\s*returns\s*\(\s*(\w+)\s*\)\s*(\{.+})?\s*;`)
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

								var roles = methodRolesMap[strings.ToLower(string(methodPieces[1]))]
								if roles == nil {
									roles = []string{}
								}

								methods = append(methods, &MethodInfo{
									Name:                string(methodPieces[1]),
									RequestMessageName:  string(methodPieces[2]),
									ResponseMessageName: string(methodPieces[3]),
									IsDeprecated:        strings.Contains(string(methodPieces[4]), "deprecated"),
									Code:                string(methodData),
									Doc:                 readComments(serviceData[:methodCodePosition[0]]),
									Roles:               roles,
								})
							}

							services = append(services, &ServiceInfo{
								Name:     serviceName,
								Methods:  methods,
								Filename: filename,
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
	}

	var countServices = len(services)
	var countMethods = 0
	var countMessages = len(messages)
	for _, service := range services {
		countMethods += len(service.Methods)
	}

	// 链接
	var links = []*LinkInfo{}

	// json links
	{
		var dirs = []string{Tea.Root + "/../pkg/rpc/jsons"}
		for _, dir := range dirs {
			func(dir string) {
				dir = filepath.Clean(dir)

				files, err := filepath.Glob(dir + "/*.md")
				if err != nil {
					fmt.Println("[ERROR]list .md files failed: " + err.Error())
					return
				}

				for _, path := range files {
					func(path string) {
						var name = strings.TrimSuffix(filepath.Base(path), ".md")
						data, err := os.ReadFile(path)
						if err != nil {
							fmt.Println("[ERROR]read '" + path + "' failed: " + err.Error())
							return
						}

						links = append(links, &LinkInfo{
							Name:    "json:" + name,
							Content: string(data),
						})
					}(path)
				}
			}(dir)
		}
	}

	var rpcList = &RPCList{
		Services: services,
		Messages: messages,
		Links:    links,
	}
	jsonData, err := json.MarshalIndent(rpcList, "", "  ")
	if err != nil {
		fmt.Println("[ERROR]marshal to json failed: " + err.Error())
		return
	}

	var jsonFile = Tea.Root + "/rpc.json"
	err = os.WriteFile(jsonFile, jsonData, 0666)
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
