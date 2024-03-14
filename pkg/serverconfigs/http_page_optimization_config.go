// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package serverconfigs

type HTTPPageOptimizationMimeType = string

const (
	HTTPPageOptimizationMimeTypeHTML       HTTPPageOptimizationMimeType = "text/html"
	HTTPPageOptimizationMimeTypeJavascript HTTPPageOptimizationMimeType = "text/javascript"
	HTTPPageOptimizationMimeTypeCSS        HTTPPageOptimizationMimeType = "text/css"
)

type HTTPPageMinifier interface {
	Bytes(mediaType string, data []byte) ([]byte, error)
}

type HTTPPageOptimizationConfig struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`

	HTML       *HTTPHTMLOptimizationConfig       `yaml:"html" json:"html"`
	Javascript *HTTPJavascriptOptimizationConfig `yaml:"javascript" json:"javascript"`
	CSS        *HTTPCSSOptimizationConfig        `yaml:"css" json:"css"`

	isOn bool

	minifyInstance HTTPPageMinifier
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

	if this.HTML != nil {
		err := this.HTML.Init()
		if err != nil {
			return err
		}
		if this.HTML.IsOn {
			this.isOn = true
		}
	}
	if this.Javascript != nil {
		err := this.Javascript.Init()
		if err != nil {
			return err
		}
		if this.Javascript.IsOn {
			this.isOn = true
		}
	}
	if this.CSS != nil {
		err := this.CSS.Init()
		if err != nil {
			return err
		}
		if this.CSS.IsOn {
			this.isOn = true
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

func (this *HTTPPageOptimizationConfig) InternalInstance() HTTPPageMinifier {
	return this.minifyInstance
}

func (this *HTTPPageOptimizationConfig) SetInternalInstance(instance HTTPPageMinifier) {
	this.minifyInstance = instance
}
