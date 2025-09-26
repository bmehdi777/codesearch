package api

import (
	"encoding/json"
	"fmt"
	"hsh/search"
	"net/http"
)

type Search struct{}

func (s *Search) router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.post(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

type SearchBody struct {
	Query string `json:"query"`
}

func (s *Search) post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body SearchBody
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	rg := search.Ripgrep{
		Path:  "/home/bmehdi/Workspace/",
		Query: body.Query,
	}
	res, err := rg.Search()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("err", err)
		return
	}
	fmt.Println("res, ", res)
}
