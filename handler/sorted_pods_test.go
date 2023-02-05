package handler

import (
	"errors"
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
		sortFnMock     func(pods []*pods.Pod, fields []pods.FieldOrder) []*pods.Pod
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
		"kubernetes API server is not available": {
			args: args{
				query: "sort=name:asc,age:desc",
				listPodsFnMock: func(_ string) ([]*pods.Pod, error) {
					return nil, errors.New("failed to get pods from etcd cluster")
				},
			},
			wantStatusCode: http.StatusInternalServerError,
			wantResponse:   []byte("{\"error\":{\"status\":500,\"error\":\"INTERNAL\",\"description\":\"Something went wrong...\"}}\n"),
		},
		"base case without any order": {
			args: args{
				query: "sort=",
				listPodsFnMock: func(_ string) ([]*pods.Pod, error) {
					return []*pods.Pod{
						{Name: "foo", Restarts: 1, Age: 1000000000},
						{Name: "bar", Restarts: 1, Age: 5000000000},
					}, nil
				},
				sortFnMock: func(_ []*pods.Pod, fields []pods.FieldOrder) []*pods.Pod {
					return []*pods.Pod{
						{Name: "foo", Restarts: 1, Age: 1000000000},
						{Name: "bar", Restarts: 1, Age: 5000000000},
					}
				},
			},
			wantStatusCode: http.StatusOK,
			wantResponse:   []byte("{\"status\":200,\"pods\":[{\"name\":\"foo\",\"restarts\":1,\"age\":\"1s\"},{\"name\":\"bar\",\"restarts\":1,\"age\":\"5s\"}]}\n"),
		},
		"base case sort by name": {
			args: args{
				query: "sort=name:asc",
				listPodsFnMock: func(_ string) ([]*pods.Pod, error) {
					return []*pods.Pod{
						{Name: "foo", Restarts: 1, Age: 1000000000},
						{Name: "bar", Restarts: 1, Age: 5000000000},
					}, nil
				},
				sortFnMock: func(_ []*pods.Pod, fields []pods.FieldOrder) []*pods.Pod {
					return []*pods.Pod{
						{Name: "bar", Restarts: 1, Age: 5000000000},
						{Name: "foo", Restarts: 1, Age: 1000000000},
					}
				},
			},
			wantStatusCode: http.StatusOK,
			wantResponse:   []byte("{\"status\":200,\"pods\":[{\"name\":\"bar\",\"restarts\":1,\"age\":\"5s\"},{\"name\":\"foo\",\"restarts\":1,\"age\":\"1s\"}]}\n"),
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
				sorter: &mocks.SorterMock{
					SortFn: tc.args.sortFnMock,
				},
			}
			h.GetSortedPods(w, req)
			if got, want := w.Code, tc.wantStatusCode; got != want {
				t.Fatalf("%v: GetSortedPods(w,r) got %v; want %v", name, got, want)
			}
			if got, want := w.Body.Bytes(), tc.wantResponse; !reflect.DeepEqual(got, want) {
				t.Fatalf("%v; GetSortedPods(w,r) \ngot %v; \nwant %v", name, string(got), string(want))
			}
		})
	}

}
