package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	pods "github.com/mig-elgt/okteto-pods"
)

// handler defines a HTTP Handler to perform POD API Requests
type handler struct {
	podsvc pods.PodLister
}

// New creates a new HTTP Handler
func New() http.Handler {
	r := mux.NewRouter()
	h := handler{}
	r.HandleFunc("/hello", h.HelloWorld).Methods("GET")
	return r
}

func (handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
