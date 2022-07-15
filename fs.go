package http2

import (
	"net/http"
	"path"
	"strings"
)

type fileHandler struct {
	root http.FileSystem
}

func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	ServeFile(w, r, f.root, path.Clean(upath))
}

func ServeFile(w http.ResponseWriter, r *http.Request, fs http.FileSystem, name string) {
	f, err := fs.Open(name)
	if err != nil {
		Error(w, r)
		return
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		Error(w, r)
		return
	}

	if d.IsDir() {
		Error(w, r)
		return
	}

	http.ServeContent(w, r, d.Name(), d.ModTime(), f)
}

func FileServer(root http.FileSystem) http.Handler {
	return &fileHandler{root}
}
