// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package errors

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

// HumanError 格式化GRPC相关错误
func HumanError(err error, endpoints []string, configFile string) (resultErr error, isConnError bool) {
	if err == nil {
		return err, false
	}
	errStatus, ok := status.FromError(err)
	if !ok {
		return err, false
	}
	switch errStatus.Code() {
	case codes.InvalidArgument:
		return fmt.Errorf("错误的RPC参数：%w", err), false
	case codes.DeadlineExceeded:
		return fmt.Errorf("RPC操作超时，请重试：%w", err), false
	case codes.Unimplemented:
		return fmt.Errorf("请求的RPC服务或方法不存在，可能是没有升级API节点或者当前节点没有升级：%w", err), false
	case codes.Unavailable:
		return fmt.Errorf("RPC当前不可用：<br/>1、请确认当前节点的api.yaml（<em>%s</em>）配置中的地址（<em>%s</em>）是否已填写正确；<br/>2、请确保API节点已启动，并检查当前节点和API节点之间的网络连接是正常的。<hr/>错误信息：%w", configFile, strings.Join(endpoints, ", "), err), true
	}

	return err, false
}
