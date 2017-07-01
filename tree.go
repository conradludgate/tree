// The tree package will recursively read from a specified file or directory, forming the data into a tree
package tree

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var LogOutput = false

func SkipFile(path string, err error) {
	if LogOutput {
		log.Println("Skipping", path+":", err.Error())
	}
}

func GenerateTreeFromFile(f *os.File, handler Handler) (interface{}, error) {
	if handler == nil {
		handler = DefaultExtHandler
	}

	data := make(map[string]interface{})

	fi, err := f.Stat()
	if err != nil {
		return data, err
	}

	if !fi.IsDir() {
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return data, err
		}

		v, err := handler.HandleFile(b, fi.Name())

		if err != nil {
			return data, err
		}

		return v, nil
	}

	names, err := f.Readdirnames(0)
	if err != nil {
		return data, err
	}

	for _, name := range names {
		fullpath := filepath.Join(f.Name(), name)

		file, err := os.Open(fullpath)
		if err != nil {
			SkipFile(fullpath, err)
			continue
		}

		v, err := GenerateTreeFromFile(file, handler)
		if err != nil {
			SkipFile(fullpath, err)
			continue
		}

		name = name[:len(name)-len(filepath.Ext(name))]
		data[name] = v
	}

	return data, nil
}

func GenerateTree(path string, handler Handler) (interface{}, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return GenerateTreeFromFile(f, handler)
}
