package handler

import (
	"net/http"
	"strings"

	pods "github.com/mig-elgt/okteto-pods"
	"github.com/mig-elgt/sender"
	"github.com/mig-elgt/sender/codes"
	"github.com/sirupsen/logrus"
)

// GET /pods?sort=foo:asc,bar:desc
func (h handler) GetSortedPods(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("sort")
	fields := strings.Split(q, ",")
	var sortFields []pods.FieldOrder
	for _, pair := range fields {
		field := strings.Split(pair, ":")
		if len(field) != 2 {
			logrus.Errorf("sort query bad format: %v", q)
			sender.
				NewJSON(w, http.StatusBadRequest).
				WithFieldError(codes.InvalidArgument, "sort", pods.ErrIncorrectSortValue.Error()).
				Send()
			return
		}
		if field[0] != "name" && field[0] != "restarts" && field[0] != "age" {
			logrus.Errorf("sort field incorrect value: %v", q)
			sender.
				NewJSON(w, http.StatusBadRequest).
				WithFieldError(codes.InvalidArgument, "sort", pods.ErrIncorrectSortValue.Error()).
				Send()
			return
		}
		if field[1] != "asc" && field[1] != "desc" {
			logrus.Errorf("sort field order incorrect value: %v", q)
			sender.
				NewJSON(w, http.StatusBadRequest).
				WithFieldError(codes.InvalidArgument, "sort", pods.ErrIncorrectSortValue.Error()).
				Send()
			return
		}
		sortFields = append(sortFields, pods.FieldOrder{Field: field[0], Order: field[1]})
	}
}
