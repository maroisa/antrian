package internal

import (
	"antrian/web"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strings"
)

func (s *Server) Routes() {
	s.mux.HandleFunc("GET /loket/ws", s.WSHandler)

	distFS, err := fs.Sub(web.EmbeddedFiles, "dist")
	if err != nil {
		panic("Gagal memuat frontend")
	}

	fileServer := http.FileServer(http.FS(distFS))
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")

		_, err := distFS.Open(path)
		if os.IsNotExist(err) && path != "" {
			indexFile, _ := distFS.Open("index.html")
			defer indexFile.Close()

			seeker, ok := indexFile.(io.ReadSeeker)
			if !ok {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			stat, _ := indexFile.Stat()

			http.ServeContent(w, r, "index.html", stat.ModTime(), seeker)
			return
		}

		fileServer.ServeHTTP(w, r)
	})
}
