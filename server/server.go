package server

import (
	"codesearch/server/api"
	"embed"
	"fmt"
	"net/http"
)

const PORT = "10001"
const ADDR = "localhost"

//go:embed all:webapp
var assetsFolder embed.FS

func NewServer() error {
	topMux := http.NewServeMux()

	apiMux := api.NewApiMux()
	topMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	fmt.Println("Listening on http://" + ADDR + ":" + PORT)
	err := http.ListenAndServe(ADDR+":"+PORT, topMux)
	return err
}
