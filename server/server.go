package server

import (
	"codesearch/server/api"
	"fmt"
	"net/http"
)

const PORT = "10001"
const ADDR = "localhost"

func NewServer() error {
	topMux := http.NewServeMux()

	apiMux := api.NewApiMux()
	topMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	fmt.Println("Listening on http://" + ADDR + ":" + PORT)
	err := http.ListenAndServe(ADDR+":"+PORT, topMux)
	return err
}
