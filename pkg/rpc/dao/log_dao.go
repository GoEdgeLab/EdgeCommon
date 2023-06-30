package dao

import (
	"context"
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/langs"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
)

var SharedLogDAO = NewLogDAO()

type LogDAO struct {
	BaseDAO
}

func NewLogDAO() *LogDAO {
	return &LogDAO{}
}

func (this *LogDAO) CreateUserLog(ctx context.Context, level string, action string, description string, ip string) error {
	_, err := this.RPC().LogRPC().CreateLog(ctx, &pb.CreateLogRequest{
		Level:       level,
		Description: description,
		Action:      action,
		Ip:          ip,
	})
	return err
}

func (this *LogDAO) CreateAdminLog(ctx context.Context, level string, action string, description string, ip string, langMessageCode langs.MessageCode, langMessageArgs []any) error {
	var langMessageArgsJSON []byte
	var err error
	if len(langMessageArgs) > 0 {
		langMessageArgsJSON, err = json.Marshal(langMessageArgs)
		if err != nil {
			return err
		}
	}

	_, err = this.RPC().LogRPC().CreateLog(ctx, &pb.CreateLogRequest{
		Level:                level,
		Description:          description,
		Action:               action,
		Ip:                   ip,
		LangMessageCode:      langMessageCode.String(),
		LangMessagesArgsJSON: langMessageArgsJSON,
	})
	return err
}
