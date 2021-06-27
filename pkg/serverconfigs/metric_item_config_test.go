// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"testing"
)

func TestMetricItemConfig_ProcessRequest(t *testing.T) {
	var metric = &MetricItemConfig{
		Keys:  []string{"${remoteAddr}", "${status}", "${requestPath}"},
		Value: "${trafficIn}",
	}
	key, hash, value := metric.ParseRequest(func(s string) string {
		return configutils.ParseVariables(s, func(varName string) (value string) {
			switch varName {
			case "trafficIn":
				return "1000"
			}
			return "[" + varName + "]"
		})
	})
	t.Log("key:", key, "hash:", hash)
	t.Logf("value: %f", value)
}

func BenchmarkMetricItemConfig_ProcessRequest(b *testing.B) {
	var metric = &MetricItemConfig{
		Keys: []string{"${remoteAddr}", "${status}"},
	}
	for i := 0; i < b.N; i++ {
		metric.ParseRequest(func(s string) string {
			return configutils.ParseVariables(s, func(varName string) (value string) {
				return "[" + varName + "]"
			})
		})
	}
}
