package tree

import "encoding/json"

const JSON = ".json"

func HandlerJSON(b []byte, path string) (interface{}, error) {
	var filedata map[string]interface{}

	err := json.Unmarshal(b, &filedata)
	if err != nil {
		return b, err
	}

	return filedata, nil
}

func HandleJSON(h *ExtHandler) {
	if h == nil {
		h = DefaultExtHandler
	}

	h.HandleFunc(JSON, HandlerJSON)
}
