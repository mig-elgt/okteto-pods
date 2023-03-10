package sort

import (
	"reflect"
	"testing"

	pods "github.com/mig-elgt/okteto-pods"
)

func TestSort(t *testing.T) {
	testCases := map[string]struct {
		pods   []*pods.Pod
		fields []pods.FieldOrder
		want   []*pods.Pod
	}{
		"without fields to sort": {
			pods: []*pods.Pod{
				{
					Name: "foo",
				},
				{
					Name: "bar",
				},
			},
			fields: []pods.FieldOrder{},
			want: []*pods.Pod{
				{
					Name: "foo",
				},
				{
					Name: "bar",
				},
			},
		},
		"sort fields name asc": {
			pods: []*pods.Pod{
				{
					Name: "foo",
				},
				{
					Name: "bar",
				},
			},
			fields: []pods.FieldOrder{{Field: "name", Order: "asc"}},
			want: []*pods.Pod{
				{
					Name: "bar",
				},
				{
					Name: "foo",
				},
			},
		},
		"sort fields name and restarts asc": {
			pods: []*pods.Pod{
				{
					Name:     "foo",
					Restarts: 15,
				},
				{
					Name:     "foo",
					Restarts: 5,
				},
			},
			fields: []pods.FieldOrder{
				{Field: "name", Order: "asc"},
				{Field: "restarts", Order: "asc"},
			},
			want: []*pods.Pod{
				{
					Name:     "foo",
					Restarts: 5,
				},
				{
					Name:     "foo",
					Restarts: 15,
				},
			},
		},
		"sort fields name, restarts, and age asc": {
			pods: []*pods.Pod{
				{
					Name:     "foo",
					Restarts: 15,
					Age:      4,
				},
				{
					Name:     "foo",
					Restarts: 5,
					Age:      5,
				},
				{
					Name:     "foo",
					Restarts: 5,
					Age:      2,
				},
				{
					Name:     "bar",
					Restarts: 10,
				},
				{
					Name:     "zaas",
					Restarts: 20,
				},
			},
			fields: []pods.FieldOrder{
				{Field: "name", Order: "asc"},
				{Field: "restarts", Order: "asc"},
				{Field: "age", Order: "asc"},
			},
			want: []*pods.Pod{
				{
					Name:     "bar",
					Restarts: 10,
				},
				{
					Name:     "foo",
					Restarts: 5,
					Age:      2,
				},
				{
					Name:     "foo",
					Restarts: 5,
					Age:      5,
				},
				{
					Name:     "foo",
					Restarts: 15,
					Age:      4,
				},
				{
					Name:     "zaas",
					Restarts: 20,
				},
			},
		},
		"sort fields age, restarts, name asc": {
			pods: []*pods.Pod{
				{
					Name:     "bb",
					Restarts: 15,
					Age:      4,
				},
				{
					Name:     "foo",
					Restarts: 5,
					Age:      5,
				},
				{
					Name:     "aa",
					Restarts: 5,
					Age:      4,
				},
				{
					Name:     "bar",
					Restarts: 10,
					Age:      8,
				},
				{
					Name:     "zaas",
					Restarts: 1,
					Age:      5,
				},
			},
			fields: []pods.FieldOrder{
				{Field: "age", Order: "asc"},
				{Field: "restarts", Order: "asc"},
				{Field: "name", Order: "asc"},
			},
			want: []*pods.Pod{
				{
					Name:     "aa",
					Restarts: 5,
					Age:      4,
				},
				{
					Name:     "bb",
					Restarts: 15,
					Age:      4,
				},
				{
					Name:     "zaas",
					Restarts: 1,
					Age:      5,
				},
				{
					Name:     "foo",
					Restarts: 5,
					Age:      5,
				},
				{
					Name:     "bar",
					Restarts: 10,
					Age:      8,
				},
			},
		},
		"sort fields name desc": {
			pods: []*pods.Pod{
				{
					Name:     "foo",
					Restarts: 15,
					Age:      4,
				},
				{
					Name:     "bar",
					Restarts: 5,
					Age:      5,
				},
				{
					Name:     "zas",
					Restarts: 5,
					Age:      4,
				},
			},
			fields: []pods.FieldOrder{
				{Field: "name", Order: "desc"},
			},
			want: []*pods.Pod{
				{
					Name:     "zas",
					Restarts: 5,
					Age:      4,
				},
				{
					Name:     "foo",
					Restarts: 15,
					Age:      4,
				},
				{
					Name:     "bar",
					Restarts: 5,
					Age:      5,
				},
			},
		},
		"sort fields name and restarts desc": {
			pods: []*pods.Pod{
				{
					Name:     "foo",
					Restarts: 15,
					Age:      4,
				},
				{
					Name:     "zas",
					Restarts: 15,
					Age:      5,
				},
				{
					Name:     "zas",
					Restarts: 5,
					Age:      4,
				},
			},
			fields: []pods.FieldOrder{
				{Field: "name", Order: "desc"},
				{Field: "restarts", Order: "desc"},
			},
			want: []*pods.Pod{
				{
					Name:     "zas",
					Restarts: 15,
					Age:      5,
				},
				{
					Name:     "zas",
					Restarts: 5,
					Age:      4,
				},
				{
					Name:     "foo",
					Restarts: 15,
					Age:      4,
				},
			},
		},
		"sort fields restarts, age and name desc": {
			pods: []*pods.Pod{
				{
					Name:     "foo",
					Restarts: 2,
					Age:      4,
				},
				{
					Name:     "bar",
					Restarts: 10,
					Age:      5,
				},
				{
					Name:     "foobar",
					Restarts: 10,
					Age:      100,
				},
				{
					Name:     "a",
					Restarts: 1,
					Age:      14,
				},
				{
					Name:     "b",
					Restarts: 1,
					Age:      14,
				},
				{
					Name:     "c",
					Restarts: 15,
					Age:      50,
				},
			},
			fields: []pods.FieldOrder{
				{Field: "restarts", Order: "desc"},
				{Field: "age", Order: "desc"},
				{Field: "name", Order: "desc"},
			},
			want: []*pods.Pod{
				{
					Name:     "c",
					Restarts: 15,
					Age:      50,
				},
				{
					Name:     "foobar",
					Restarts: 10,
					Age:      100,
				},
				{
					Name:     "bar",
					Restarts: 10,
					Age:      5,
				},
				{
					Name:     "foo",
					Restarts: 2,
					Age:      4,
				},
				{
					Name:     "b",
					Restarts: 1,
					Age:      14,
				},
				{
					Name:     "a",
					Restarts: 1,
					Age:      14,
				},
			},
		},
		"sort fields restarts desc and age asc": {
			pods: []*pods.Pod{
				{
					Name:     "foo",
					Restarts: 50,
					Age:      4,
				},
				{
					Name:     "bar",
					Restarts: 100,
					Age:      5,
				},
				{
					Name:     "foobar",
					Restarts: 50,
					Age:      1,
				},
			},
			fields: []pods.FieldOrder{
				{Field: "restarts", Order: "desc"},
				{Field: "age", Order: "asc"},
			},
			want: []*pods.Pod{
				{
					Name:     "bar",
					Restarts: 100,
					Age:      5,
				},
				{
					Name:     "foobar",
					Restarts: 50,
					Age:      1,
				},
				{
					Name:     "foo",
					Restarts: 50,
					Age:      4,
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			s := sort{}
			if got := s.Sort(tc.pods, tc.fields); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Sort(pods, fields) got %v; want %v", got, tc.want)
			}
		})
	}
}
