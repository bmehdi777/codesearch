package server

import (
	"hsh/server/api"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"
)

const PORT = "10001"
const ADDR = "localhost"

//go:embed all:webapp
var assetsFolder embed.FS

func NewServer() error {
	assets, err := fs.Sub(assetsFolder, "webapp/dist")
	if err != nil {
		panic(err)
	}

	topMux := http.NewServeMux()

	apiMux := api.NewApiMux()
	topMux.Handle("/api/", http.StripPrefix("/api", apiMux))
	topMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleWebappRequest(w, r, assets, "/")
	})

	fmt.Println("Listening on http://" + ADDR + ":" + PORT)
	err = http.ListenAndServe(ADDR+":"+PORT, topMux)
	return err
}

func handleWebappRequest(w http.ResponseWriter, r *http.Request, fs fs.FS, prefixUrl string) {
	filePath := path.Clean(r.URL.Path)
	prefixClean := path.Clean(prefixUrl)

	if filePath == prefixClean {
		filePath = "index.html"
	} else {
		filePath = strings.TrimPrefix(filePath, prefixUrl)
	}

	file, err := fs.Open(filePath)
	if os.IsNotExist(err) || filePath == "index.html" {
		http.ServeFileFS(w, r, fs, "index.html")
		return
	} else if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Change url before giving corresponding file : we don't have prefix in
	// the FS
	r.URL.Path = filePath

	http.FileServer(http.FS(fs)).ServeHTTP(w, r)
}
