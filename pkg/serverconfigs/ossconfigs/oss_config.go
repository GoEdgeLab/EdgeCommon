// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package ossconfigs

import (
	"encoding/json"
	"reflect"
)

type OSSConfig struct {
	Type    OSSType `yaml:"oss" json:"type"`
	Options any     `yaml:"options" json:"options"`
}

func NewOSSConfig() *OSSConfig {
	return &OSSConfig{}
}

func (this *OSSConfig) Init() error {
	if this.Options != nil {
		// decode options
		if reflect.TypeOf(this.Options).Kind() == reflect.Map {
			optionsJSON, err := json.Marshal(this.Options)
			if err != nil {
				return err
			}

			newOptions, decodeErr := DecodeOSSOptions(this.Type, optionsJSON)
			if decodeErr != nil {
				return decodeErr
			}
			if newOptions != nil {
				this.Options = newOptions
			}
		}

		options, ok := this.Options.(OSSOptions)
		if ok {
			err := options.Init()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (this *OSSConfig) Summary() string {
	var name = ""
	var found = false
	for _, def := range FindAllOSSTypes() {
		if def.Code == this.Type {
			name = def.Name
			found = true
			break
		}
	}
	if !found {
		return ""
	}

	var summary = ""
	if this.Options != nil {
		// decode options
		if reflect.TypeOf(this.Options).Kind() == reflect.Map {
			optionsJSON, err := json.Marshal(this.Options)
			if err == nil { // ignore error
				newOptions, decodeErr := DecodeOSSOptions(this.Type, optionsJSON)
				if decodeErr == nil && newOptions != nil {
					this.Options = newOptions
				}
			}
		}
		options, ok := this.Options.(OSSOptions)
		if ok {
			summary = options.Summary()
		}
	}

	if len(summary) == 0 {
		return name
	}
	return name + " - " + summary
}
