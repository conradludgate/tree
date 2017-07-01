package tree

import (
	"gopkg.in/yaml.v2"
)

const YAML = ".yaml"

func HandlerYAML(b []byte, path string) (interface{}, error) {
	var filedata map[string]interface{}

	err := yaml.Unmarshal(b, &filedata)
	if err != nil {
		return b, err
	}

	return filedata, nil
}

func HandleYAML(h *ExtHandler) {
	if h == nil {
		h = DefaultExtHandler
	}

	h.HandleFunc(YAML, HandlerYAML)
}
