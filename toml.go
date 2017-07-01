package tree

import "github.com/BurntSushi/toml"

const TOML = ".toml"

func HandlerTOML(b []byte, path string) (interface{}, error) {
	var filedata map[string]interface{}

	_, err := toml.Decode(string(b), &filedata)
	if err != nil {
		return b, err
	}

	return filedata, nil
}

func HandleTOML(h *ExtHandler) {
	if h == nil {
		h = DefaultExtHandler
	}

	h.HandleFunc(TOML, HandlerTOML)
}
