package tree

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"path/filepath"
)

const ZIP = ".zip"

func HandlerZIP(h *ExtHandler) HandlerFunc {
	if h == nil {
		h = DefaultExtHandler
	}

	return func(b []byte, path string) (interface{}, error) {
		data := make(map[string]interface{})

		r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
		if err != nil {
			return nil, err
		}

		for _, f := range r.File {
			if f.FileInfo().IsDir() {
				Insert(data, f.Name, make(map[string]interface{}))
				continue
			}

			rc, err := f.Open()
			if err != nil {
				SkipFile(filepath.Join(path, f.Name)+":", err)
				continue
			}

			b, err := ioutil.ReadAll(rc)
			if err != nil {
				SkipFile(filepath.Join(path, f.Name)+":", err)
				continue
			}

			v, err := h.HandleFile(b, filepath.Join(path, f.Name))
			if err != nil {
				SkipFile(filepath.Join(path, f.Name)+":", err)
				continue
			}

			Insert(
				data,
				f.Name[:len(f.Name)-len(filepath.Ext(f.Name))],
				v,
			)
		}

		return data, nil
	}
}

func HandleZIP(h *ExtHandler) {
	if h == nil {
		h = DefaultExtHandler
	}

	h.HandleFunc(ZIP, HandlerZIP(h))
}
