package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/mig-elgt/okteto-pods/mocks"
)

func TestHandlerGetTotalPods(t *testing.T) {
	type args struct {
		totalFnMock func(namespace string) (int, error)
	}
	testCases := map[string]struct {
		args           args
		wantStatusCode int
		wantResponse   []byte
	}{
		"Kubernetes API Server is not available": {
			args: args{
				totalFnMock: func(_ string) (int, error) {
					return 0, errors.New("Kubernetes API Internal Error")
				},
			},
			wantStatusCode: http.StatusInternalServerError,
			wantResponse:   []byte("{\"error\":{\"status\":500,\"error\":\"INTERNAL\",\"description\":\"Something went wrong...\"}}\n"),
		},
		"base case": {
			args: args{
				totalFnMock: func(_ string) (int, error) {
					return 100, nil
				},
			},
			wantStatusCode: http.StatusOK,
			wantResponse:   []byte("{\"status\":200,\"total\":100}\n"),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/pods/total", nil)
			w := httptest.NewRecorder()

			h := handler{
				podsvc: &mocks.PodListerMock{
					TotalFn: tc.args.totalFnMock,
				},
			}
			h.GetTotalPods(w, req)
			if got, want := w.Code, tc.wantStatusCode; got != want {
				t.Fatalf("%v: GetTotalPod(w,r) got %v; want %v", name, got, want)
			}
			if got, want := w.Body.Bytes(), tc.wantResponse; !reflect.DeepEqual(got, want) {
				t.Fatalf("%v; GetTotalPods(w,r) got %v; want %v", name, string(got), string(want))
			}
		})
	}
}
