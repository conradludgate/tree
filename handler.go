package tree

// Handlers will format file bytes so that they are easier to work with.
// b is the source data for the file.
// path is the 'path' to the file. Do not use to find the resource as it may not be a valid file path
// For example, path may be example/foo.zip/bar, even though foo.zip is a file and not a directory
//
// If an error occurs when handling the data, the tree will still be processed, but the file will be ignored and the error logged.
type Handler interface {
	HandleFile(b []byte, path string) (interface{}, error)
}

type HandlerFunc func(b []byte, path string) (interface{}, error)

func (f HandlerFunc) HandleFile(b []byte, path string) (interface{}, error) {
	return f(b, path)
}

func Handle(ext string, handler Handler) {
	DefaultExtHandler.Handle(ext, handler)
}

func HandleFunc(ext string, handler func(b []byte, path string) (interface{}, error)) {
	DefaultExtHandler.Handle(ext, HandlerFunc(handler))
}
