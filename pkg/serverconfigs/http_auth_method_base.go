// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/iwind/TeaGo/lists"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
)

var httpAuthTimestampRegexp = regexp.MustCompile(`^\d{10}$`)

type HTTPAuthBaseMethod struct {
	Exts    []string `json:"exts"`
	Domains []string `json:"domains"`
}

func (this *HTTPAuthBaseMethod) SetExts(exts []string) {
	this.Exts = exts
}

func (this *HTTPAuthBaseMethod) SetDomains(domains []string) {
	this.Domains = domains
}

func (this *HTTPAuthBaseMethod) removeQueryArgs(query string, args []string) string {
	var pieces = strings.Split(query, "&")
	var result = []string{}
Loop:
	for _, piece := range pieces {
		for _, arg := range args {
			if strings.HasPrefix(piece, arg+"=") {
				continue Loop
			}
		}
		result = append(result, piece)
	}
	return strings.Join(result, "&")
}

func (this *HTTPAuthBaseMethod) matchTimestamp(timestamp string) bool {
	return httpAuthTimestampRegexp.MatchString(timestamp)
}

func (this *HTTPAuthBaseMethod) MatchRequest(req *http.Request) bool {
	if len(this.Exts) > 0 {
		var ext = filepath.Ext(req.URL.Path)
		if len(ext) == 0 {
			return false
		}

		// ext中包含点符号
		ext = strings.ToLower(ext)
		if !lists.ContainsString(this.Exts, ext) {
			return false
		}
	}

	if len(this.Domains) > 0 {
		var domain = req.Host
		if len(domain) == 0 {
			return false
		}
		if !configutils.MatchDomains(this.Domains, domain) {
			return false
		}
	}

	return true
}
