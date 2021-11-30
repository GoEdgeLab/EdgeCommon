// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package errors

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HumanError 格式化GRPC相关错误
func HumanError(err error) error {
	if err == nil {
		return err
	}
	errStatus, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch errStatus.Code() {
	case codes.InvalidArgument:
		return errors.New("错误的RPC参数：" + err.Error())
	case codes.DeadlineExceeded:
		return errors.New("RPC操作超时，请重试：" + err.Error())
	case codes.Unimplemented:
		return errors.New("请求的RPC服务或方法不存在，可能是没有升级API节点或者当前节点没有升级：" + err.Error())
	case codes.Unavailable:
		return errors.New("RPC当前不可用，1、请确保当前节点的api.yaml配置中的地址填写正确；2、请确保API节点已启动，并检查当前节点和API节点之间的网络连接是正常的。错误信息：" + err.Error())
	}

	return err
}
