// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"context"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"testing"
)

func TestReverseProxyConfig_Init(t *testing.T) {
	var config = &ReverseProxyConfig{}
	config.Scheduling = &SchedulingConfig{
		Code:    "random",
		Options: nil,
	}
	config.AddPrimaryOrigin(&OriginConfig{
		Addr: &NetworkAddressConfig{Host: "127.0.0.1"},
		IsOn: true,
	})
	config.AddPrimaryOrigin(&OriginConfig{
		Addr: &NetworkAddressConfig{Host: "127.0.0.2"},
		IsOn: true,
	})
	config.AddPrimaryOrigin(&OriginConfig{
		Addr:    &NetworkAddressConfig{Host: "127.0.0.3"},
		Domains: []string{"*.www.example.com", ".example.com"},
		IsOn:    true,
	})
	config.AddBackupOrigin(&OriginConfig{
		Addr: &NetworkAddressConfig{Host: "127.0.0.4"},
		IsOn: true,
	})
	err := config.Init(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	for domain, group := range config.schedulingGroupMap {
		for _, origin := range group.PrimaryOrigins {
			t.Log(domain, "primary", origin.Addr.Host)
		}
		for _, origin := range group.BackupOrigins {
			t.Log(domain, "backup", origin.Addr.Host)
		}
	}

	//config.ResetScheduling()

	nextOrigin := config.NextOrigin(&shared.RequestCall{
		Formatter:         nil,
		Request:           nil,
		Domain:            "a.example.com",
		ResponseCallbacks: nil,
		Options:           nil,
	})
	if nextOrigin == nil {
		t.Log("not found")
	} else {
		t.Log("result:", nextOrigin.Addr.Host)
	}
}
