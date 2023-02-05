package sort

import (
	ss "sort"

	pods "github.com/mig-elgt/okteto-pods"
)

// sort implements Sorter interface to compute a Pod sorted list.
type sort struct{}

// New creates a new sort instance.
func New() *sort {
	return &sort{}
}

// Sort performs the Sort algorithm for a list of Pods and a set
// of fields values to get the order.
func (s *sort) Sort(pods []*pods.Pod, fields []pods.FieldOrder) []*pods.Pod {
	if len(fields) == 0 {
		return pods
	}
	ss.Slice(pods, func(i, j int) bool {
		for _, f := range fields {
			if f.Field == "name" {
				if pods[i].Name != pods[j].Name {
					if f.Order == "asc" {
						return pods[i].Name < pods[j].Name
					}
					return pods[i].Name > pods[j].Name
				}
				continue
			}
			if f.Field == "restarts" {
				if pods[i].Restarts != pods[j].Restarts {
					if f.Order == "asc" {
						return pods[i].Restarts < pods[j].Restarts
					}
					return pods[i].Restarts > pods[j].Restarts
				}
				continue
			}
			if f.Field == "age" {
				if pods[i].Age != pods[j].Age {
					if f.Order == "asc" {
						return pods[i].Age < pods[j].Age
					}
					return pods[i].Age > pods[j].Age
				}
			}
		}
		return false
	})
	return pods
}
