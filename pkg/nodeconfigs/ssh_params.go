// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package nodeconfigs

func DefaultSSHParams() *SSHParams {
	return &SSHParams{Port: 22}
}

type SSHParams struct {
	Port int `json:"port"`
}
