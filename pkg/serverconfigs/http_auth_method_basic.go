// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"encoding/json"
	"net/http"
)

// HTTPAuthBasicMethodUser BasicAuth中的用户
type HTTPAuthBasicMethodUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (this *HTTPAuthBasicMethodUser) Validate(password string) (bool, error) {
	return this.Password == password, nil
}

// HTTPAuthBasicMethod BasicAuth方法定义
type HTTPAuthBasicMethod struct {
	HTTPAuthBaseMethod

	Users   []*HTTPAuthBasicMethodUser `json:"users"`
	Realm   string                     `json:"realm"`
	Charset string                     `json:"charset"`

	userMap map[string]*HTTPAuthBasicMethodUser // username => *User
}

func NewHTTPAuthBasicMethod() *HTTPAuthBasicMethod {
	return &HTTPAuthBasicMethod{}
}

func (this *HTTPAuthBasicMethod) Init(params map[string]any) error {
	this.userMap = map[string]*HTTPAuthBasicMethodUser{}

	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return err
	}
	err = json.Unmarshal(paramsJSON, this)
	if err != nil {
		return err
	}

	for _, user := range this.Users {
		this.userMap[user.Username] = user
	}

	return nil
}

func (this *HTTPAuthBasicMethod) Filter(req *http.Request, doSubReq func(subReq *http.Request) (status int, err error), formatter func(string) string) (ok bool, newURI string, uriChanged bool, err error) {
	username, password, ok := req.BasicAuth()
	if !ok {
		return false, "", false, nil
	}
	user, ok := this.userMap[username]
	if !ok {
		return false, "", false, nil
	}
	ok, err = user.Validate(password)
	return ok, "", false, err
}
