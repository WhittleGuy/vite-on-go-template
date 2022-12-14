package hiddenFs

import (
	"net/http"
	"os"
	"strings"
)

type FileSystem struct {
	fs http.FileSystem
}

// Create unindexed fs for security [https://github.com/jordan-wright/unindexed/blob/master/unindexed.go]
func (hiddenFs FileSystem) Open(name string) (http.File, error) {
	f, err := hiddenFs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(name, "/") + "/index.html"
		_, err := hiddenFs.fs.Open(index)
		if err != nil {
			return nil, os.ErrPermission
		}
	}
	return f, nil
}

// Create drop-in replacement for http.Dir [https://github.com/jordan-wright/unindexed/blob/master/unindexed.go]
func Dir(filepath string) http.FileSystem {
	return FileSystem{
		fs: http.Dir(filepath),
	}
}