package configutils

import (
	"gopkg.in/yaml.v3"
	"os"
)

func UnmarshalYamlFile(file string, ptr interface{}) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, ptr)
}
