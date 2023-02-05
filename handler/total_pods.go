package handler

import (
	"net/http"

	"github.com/mig-elgt/sender"
	"github.com/mig-elgt/sender/codes"
	"github.com/sirupsen/logrus"
)

// GetTotalPods performs requests for the endpoint GET /pods/total
func (h handler) GetTotalPods(w http.ResponseWriter, r *http.Request) {
	pods, err := h.podsvc.Total("mig-elgt")
	if err != nil {
		logrus.Errorf("could not get total pods: %v", err)
		sender.
			NewJSON(w, http.StatusInternalServerError).
			WithError(codes.Internal, "Something went wrong...").
			Send()
		return
	}
	type response struct {
		Total int `json:"total"`
	}
	sender.NewJSON(w, http.StatusOK).Send(&response{Total: pods})
}
