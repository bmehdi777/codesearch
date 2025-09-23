package api

import (
	"home/search"
	"fmt"
	"net/http"
)

type Health struct{}

func (h *Health) router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Health) get(w http.ResponseWriter, _ *http.Request) {
	rg := search.Ripgrep{
		Path: "/home/bmehdi/Workspace/perso",
		Query: "test",
	}
	res, err := rg.Search()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("err", err)
		return
	}
	fmt.Println("res, ", res[0])
	w.WriteHeader(http.StatusOK)
}
 
