package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	pods "github.com/mig-elgt/okteto-pods"
)

// handler defines a HTTP Handler to perform POD API Requests
type handler struct {
	podsvc pods.PodLister
	sorter pods.Sorter
}

// New creates a new HTTP Handler
func New(podsvc pods.PodLister, s pods.Sorter) http.Handler {
	r := mux.NewRouter()

	h := handler{
		podsvc: podsvc,
		sorter: s,
	}
	r.HandleFunc("/pods/total", h.GetTotalPods).Methods("GET")

	return r
}
