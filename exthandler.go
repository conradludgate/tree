package tree

import (
	"errors"
	"path/filepath"
	"strings"
	"sync"
)

var DefaultExtHandler = new(ExtHandler)

type ExtHandler struct {
	mu sync.RWMutex
	h  map[string]Handler
}

var ErrIgnoreFile = errors.New("Files with '.' prefix is ignored")

func (eh *ExtHandler) HandleFile(b []byte, path string) (interface{}, error) {
	name := filepath.Base(path)
	if name[0] == '.' {
		return nil, ErrIgnoreFile
	}

	if eh.h == nil {
		eh.h = make(map[string]Handler)
	}

	ext := strings.ToLower(filepath.Ext(name))

	h, ok := eh.h[ext]
	if !ok {
		return b, nil
	}

	return h.HandleFile(b, path)
}

func (eh *ExtHandler) Handle(ext string, handler Handler) {
	eh.mu.Lock()
	defer eh.mu.Unlock()

	if eh.h == nil {
		eh.h = make(map[string]Handler)
	}

	ext = strings.ToLower(ext)

	if handler == nil {
		delete(eh.h, ext)
		return
	}

	eh.h[ext] = handler
}

func (eh *ExtHandler) HandleFunc(ext string, handler func(b []byte, path string) (interface{}, error)) {
	eh.Handle(ext, HandlerFunc(handler))
}
