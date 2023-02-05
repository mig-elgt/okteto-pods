package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	pods "github.com/mig-elgt/okteto-pods"
	"github.com/mig-elgt/okteto-pods/mocks"
)

func TestHandlerGetSortedPods(t *testing.T) {
	type args struct {
		listPodsFnMock func(namespace string) ([]*pods.Pod, error)
		query          string
	}
	testCases := map[string]struct {
		args           args
		wantStatusCode int
		wantResponse   []byte
	}{
		"sort parameter order value bad format, should return a bad request error": {
			args: args{
				query: "sort=foo,bar:asc",
			},
			wantStatusCode: http.StatusBadRequest,
			wantResponse:   []byte("{\"error\":{\"status\":400,\"error\":\"INVALID_ARGUMENT\",\"description\":\"One or more fields raised validation errors.\",\"fields\":{\"sort\":\"Incorrect field order value, should be asc or desc.\"}}}\n"),
		},
		"sort field name incorrect value, should return a bad request error": {
			args: args{
				query: "sort=foo:foo,bar:bar",
			},
			wantStatusCode: http.StatusBadRequest,
			wantResponse:   []byte("{\"error\":{\"status\":400,\"error\":\"INVALID_ARGUMENT\",\"description\":\"One or more fields raised validation errors.\",\"fields\":{\"sort\":\"Incorrect field order value, should be asc or desc.\"}}}\n"),
		},
		"sort field' order incorrect value, should return a bad request error": {
			args: args{
				query: "sort=name:asc,age:bar",
			},
			wantStatusCode: http.StatusBadRequest,
			wantResponse:   []byte("{\"error\":{\"status\":400,\"error\":\"INVALID_ARGUMENT\",\"description\":\"One or more fields raised validation errors.\",\"fields\":{\"sort\":\"Incorrect field order value, should be asc or desc.\"}}}\n"),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest("GET", fmt.Sprintf("/pods?%s", tc.args.query), nil)
			w := httptest.NewRecorder()
			h := handler{
				podsvc: &mocks.PodListerMock{
					ListFn: tc.args.listPodsFnMock,
				},
			}
			h.GetSortedPods(w, req)
			if got, want := w.Code, tc.wantStatusCode; got != want {
				t.Fatalf("%v: GetSortedPods(w,r) got %v; want %v", name, got, want)
			}
			if got, want := w.Body.Bytes(), tc.wantResponse; !reflect.DeepEqual(got, want) {
				t.Fatalf("%v; GetSortedPods(w,r) \ngot %v; \ngot %v", name, string(got), string(want))
			}
		})
	}

}
