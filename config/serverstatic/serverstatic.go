package serverstatic

import (
	"errors"
	"gomodule/pkg/zlog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

func Master() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/assets", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, rq *http.Request) {
		})
		workDir, _ := os.Getwd()
		filesDir := http.Dir(filepath.Join(workDir, "assets"))
		fileServer(r, "/", filesDir)
	})
	return r
}
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		zlog.Fatal(errors.New("FileServer does not permit any URL parameters."))
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
