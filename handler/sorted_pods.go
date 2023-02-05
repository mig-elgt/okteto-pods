package handler

import (
	"net/http"
	"strings"

	pods "github.com/mig-elgt/okteto-pods"
	"github.com/mig-elgt/sender"
	"github.com/mig-elgt/sender/codes"
	"github.com/sirupsen/logrus"
)

type pod struct {
	Name     string `json:"name"`
	Restarts int32  `json:"restarts"`
	Status   string `json:"status,omitempty"`
	Age      string `json:"age"`
}

type response struct {
	Status int   `json:"status"`
	Pods   []pod `json:"pods"`
}

// GetSortedPods represents a handle for the request GET /pods?sort=name:asc,restarts:desc,age:asc
func (h handler) GetSortedPods(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("sort")
	var sortFields []pods.FieldOrder
	if len(q) > 0 {
		res, err := h.validateSortParameter(r)
		if err != nil {
			sender.
				NewJSON(w, http.StatusBadRequest).
				WithFieldError(codes.InvalidArgument, "sort", err.Error()).
				Send()
			return
		}
		sortFields = res
	}
	pods, err := h.podsvc.List("mig-elgt")
	if err != nil {
		logrus.Errorf("could not get pods: %v", err)
		sender.
			NewJSON(w, http.StatusInternalServerError).
			WithError(codes.Internal, "Something went wrong...").
			Send()
		return
	}
	pods = h.sorter.Sort(pods, sortFields)
	var podList []pod
	for _, p := range pods {
		podList = append(podList, pod{
			Name:     p.Name,
			Status:   p.Status,
			Restarts: p.Restarts,
			Age:      p.Age.String(),
		})
	}
	sender.NewJSON(w, http.StatusOK).Send(&response{
		Status: http.StatusOK,
		Pods:   podList,
	})
}

func (h handler) validateSortParameter(r *http.Request) ([]pods.FieldOrder, error) {
	q := r.URL.Query().Get("sort")
	var sortFields []pods.FieldOrder
	fields := strings.Split(q, ",")
	for _, pair := range fields {
		field := strings.Split(pair, ":")
		if len(field) != 2 {
			logrus.Errorf("sort query bad format: %v", q)
			return nil, pods.ErrIncorrectSortValue
		}
		if field[0] != "name" && field[0] != "restarts" && field[0] != "age" {
			logrus.Errorf("sort field incorrect value: %v", q)
			return nil, pods.ErrIncorrectSortValue
		}
		if field[1] != "asc" && field[1] != "desc" {
			logrus.Errorf("sort field order incorrect value: %v", q)
			return nil, pods.ErrIncorrectSortValue
		}
		sortFields = append(sortFields, pods.FieldOrder{Field: field[0], Order: field[1]})
	}
	return sortFields, nil
}
