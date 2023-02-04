package pods

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
