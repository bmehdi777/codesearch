package api

import "net/http"

func NewApiMux() *http.ServeMux {
	apiMux := http.NewServeMux()

	health := Health{}
	search := Search{}
	apiMux.HandleFunc("/health", health.router)
	apiMux.HandleFunc("/search", search.router)

	return apiMux
}

