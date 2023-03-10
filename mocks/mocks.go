package mocks

import pods "github.com/mig-elgt/okteto-pods"

type PodListerMock struct {
	TotalFn func(namespace string) (int, error)
	ListFn  func(namespace string) ([]*pods.Pod, error)
}

func (p PodListerMock) Total(namespace string) (int, error) {
	return p.TotalFn(namespace)
}

func (p PodListerMock) List(namespace string) ([]*pods.Pod, error) {
	return p.ListFn(namespace)
}

type SorterMock struct {
	SortFn func(pods []*pods.Pod, fields []pods.FieldOrder) []*pods.Pod
}

func (s *SorterMock) Sort(pods []*pods.Pod, fields []pods.FieldOrder) []*pods.Pod {
	return s.SortFn(pods, fields)
}
