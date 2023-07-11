// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

import (
	"bytes"
	"github.com/iwind/TeaGo/types"
	"github.com/tdewolff/minify/v2"
	"io"
	"net/http"
	"strings"
)

var httpPageOptimizationLimiter = make(chan bool, 64)

type HTTPPageOptimizationMimeType = string

const (
	HTTPPageOptimizationMimeTypeHTML       HTTPPageOptimizationMimeType = "text/html"
	HTTPPageOptimizationMimeTypeJavascript HTTPPageOptimizationMimeType = "text/javascript"
	HTTPPageOptimizationMimeTypeCSS        HTTPPageOptimizationMimeType = "text/css"
)

type HTTPPageOptimizationConfig struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`

	HTML       *HTTPHTMLOptimizationConfig       `yaml:"html" json:"html"`
	Javascript *HTTPJavascriptOptimizationConfig `yaml:"javascript" json:"javascript"`
	CSS        *HTTPCSSOptimizationConfig        `yaml:"css" json:"css"`

	isOn           bool
	minifyInstance *minify.M
}

func NewHTTPPageOptimizationConfig() *HTTPPageOptimizationConfig {
	return &HTTPPageOptimizationConfig{
		IsPrior:    false,
		HTML:       NewHTTPHTMLOptimizationConfig(),
		Javascript: NewHTTPJavascriptOptimizationConfig(),
		CSS:        NewHTTPCSSOptimizationConfig(),
	}
}

func (this *HTTPPageOptimizationConfig) Init() error {
	this.isOn = this.CheckIsOn()

	if this.isOn {
		// MUST NOT create instance for every config
		this.minifyInstance = minify.New()
	}

	if this.HTML != nil {
		err := this.HTML.Init()
		if err != nil {
			return err
		}
		if this.HTML.IsOn {
			this.isOn = true
			this.minifyInstance.Add(HTTPPageOptimizationMimeTypeHTML, this.HTML.AsMinifier())

		}
	}
	if this.Javascript != nil {
		err := this.Javascript.Init()
		if err != nil {
			return err
		}
		if this.Javascript.IsOn {
			this.isOn = true
			this.minifyInstance.Add(HTTPPageOptimizationMimeTypeJavascript, this.Javascript.AsMinifier())
		}
	}
	if this.CSS != nil {
		err := this.CSS.Init()
		if err != nil {
			return err
		}
		if this.CSS.IsOn {
			this.isOn = true
			this.minifyInstance.Add(HTTPPageOptimizationMimeTypeCSS, this.CSS.AsMinifier())
		}
	}

	return nil
}

func (this *HTTPPageOptimizationConfig) IsOn() bool {
	return this.isOn
}

func (this *HTTPPageOptimizationConfig) CheckIsOn() bool {
	return (this.HTML != nil && this.HTML.IsOn) ||
		(this.Javascript != nil && this.Javascript.IsOn) ||
		(this.CSS != nil && this.CSS.IsOn)
}

func (this *HTTPPageOptimizationConfig) FilterResponse(resp *http.Response) error {
	if !this.isOn || this.minifyInstance == nil {
		return nil
	}

	var contentType = resp.Header.Get("Content-Type")
	if len(contentType) == 0 {
		return nil
	}

	// validate content length
	if resp.ContentLength <= 0 || resp.ContentLength > (1<<20) {
		return nil
	}

	contentType, _, _ = strings.Cut(contentType, ";")
	var mimeType = ""
	switch contentType {
	case "text/html":
		if this.HTML != nil && this.HTML.IsOn {
			mimeType = HTTPPageOptimizationMimeTypeHTML
		}
	case "text/javascript", "application/javascript":
		if this.Javascript != nil && this.Javascript.IsOn {
			mimeType = HTTPPageOptimizationMimeTypeJavascript
		}
	case "text/css":
		if this.CSS != nil && this.CSS.IsOn {
			mimeType = HTTPPageOptimizationMimeTypeCSS
		}
	default:
		return nil
	}

	if len(mimeType) == 0 {
		return nil
	}

	// concurrent limiter, to prevent memory overflow
	select {
	case httpPageOptimizationLimiter <- true:
		defer func() {
			<-httpPageOptimizationLimiter
		}()

		var contentLength int64
		var err error
		resp.Body, contentLength, err = this.minify(mimeType, resp.Body)
		if err != nil {
			return err
		}

		// fix resp.ContentLength and Content-Length header
		resp.ContentLength = contentLength
		resp.Header.Set("Content-Length", types.String(contentLength))
	default:
	}
	return nil
}

func (this *HTTPPageOptimizationConfig) minify(mimeType HTTPPageOptimizationMimeType, rawReader io.ReadCloser) (newReader io.ReadCloser, newContentLength int64, err error) {

	var rawData []byte
	rawData, err = io.ReadAll(rawReader)
	if err != nil {
		return
	}

	resultData, err := this.minifyInstance.Bytes(mimeType, rawData)
	if err != nil {
		return io.NopCloser(bytes.NewReader(rawData)), int64(len(rawData)), nil // return rawData, and ignore error
	}

	return io.NopCloser(bytes.NewReader(resultData)), int64(len(resultData)), nil
}
