package tree

import (
	"errors"
	"path/filepath"
	"strings"
)

func isMap(tree interface{}) bool {
	_, ok := tree.(map[string]interface{})
	return ok
}

// Risky
func getMap(tree interface{}) map[string]interface{} {
	return tree.(map[string]interface{})
}

var ErrValueNotFound = errors.New("Value not found")
var ErrCannotUsePath = errors.New("Cannot use path")

// Get takes the path and splits it on forward slashes.
// It then recursively travels through the tree until it either errors or returns the data
// If an error occurs, the last Data value in the tree found will be returned
func Get(tree interface{}, path string) (interface{}, error) {
	filepath.Clean(path)
	return GetWithSlice(tree, strings.Split(path, "/"))
}

func GetWithSlice(tree interface{}, path []string) (interface{}, error) {
	if len(path) == 0 {
		return tree, nil
	}

	if len(path) == 1 {
		if isMap(tree) {
			v, ok := getMap(tree)[path[0]]
			if !ok {
				return tree, ErrValueNotFound
			}
			return v, nil
		}

		return tree, ErrValueNotFound
	}

	if isMap(tree) {
		v, ok := getMap(tree)[path[0]]
		if !ok {
			return nil, ErrValueNotFound
		}
		return GetWithSlice(v, path[1:])
	}

	return tree, ErrValueNotFound
}

func Insert(tree interface{}, path string, v interface{}) error {
	if path == "" {
		return InsertWithSlice(tree, []string{}, v)
	}

	filepath.Clean(path)
	return InsertWithSlice(tree, strings.Split(path, "/"), v)
}

func InsertWithSlice(tree interface{}, path []string, v interface{}) error {
	if len(path) == 0 {
		return nil
	}

	if len(path) == 1 {
		if isMap(tree) {
			getMap(tree)[path[0]] = v
			return nil
		}

		return ErrCannotUsePath
	}

	if isMap(tree) {
		m, ok := getMap(tree)[path[0]]
		if !ok {
			m = make(map[string]interface{})
		}

		return InsertWithSlice(m, path[1:], v)
	}

	return ErrCannotUsePath
}
