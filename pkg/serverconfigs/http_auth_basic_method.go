// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"encoding/base64"
	"encoding/json"
	stringutil "github.com/iwind/TeaGo/utils/string"
	"net/http"
)

// HTTPAuthBasicMethodUser BasicAuth中的用户
type HTTPAuthBasicMethodUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Encoder  string `json:"encoding"`
}

func (this *HTTPAuthBasicMethodUser) Validate(password string) (bool, error) {
	switch this.Encoder {
	case "md5":
		return this.Password == stringutil.Md5(password), nil
	case "base64":
		return this.Password == base64.StdEncoding.EncodeToString([]byte(password)), nil
	default:
		return this.Password == password, nil
	}
}

// HTTPAuthBasicMethod BasicAuth方法定义
type HTTPAuthBasicMethod struct {
	Users []*HTTPAuthBasicMethodUser `json:"users"`

	userMap map[string]*HTTPAuthBasicMethodUser // username => *User
}

func NewHTTPAuthBasicMethod() *HTTPAuthBasicMethod {
	return &HTTPAuthBasicMethod{}
}

func (this *HTTPAuthBasicMethod) Init(params map[string]interface{}) error {
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

func (this *HTTPAuthBasicMethod) Filter(req *http.Request, doSubReq func(subReq *http.Request) (status int, err error), formatter func(string) string) (bool, error) {
	username, password, ok := req.BasicAuth()
	if !ok {
		return false, nil
	}
	user, ok := this.userMap[username]
	if !ok {
		return false, nil
	}
	return user.Validate(password)
}
