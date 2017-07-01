package tree

const TXT = ".txt"

func HandlerTXT(b []byte, path string) (interface{}, error) {
	return string(b), nil
}

func HandleYAML(h *ExtHandler) {
	if h == nil {
		h = DefaultExtHandler
	}

	h.HandleFunc(TXT, HandlerTXT)
}
