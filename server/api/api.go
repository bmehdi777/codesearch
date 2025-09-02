package api

import "net/http"

func NewApiMux() *http.ServeMux {
	apiMux := http.NewServeMux()

	health := Health{}
	apiMux.HandleFunc("/health", health.router)

	return apiMux
}

