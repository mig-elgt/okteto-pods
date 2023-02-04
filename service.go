package pods

// Pod describes a Kubernetes Pod resource.
type Pod struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	Restarts int    `json:"restarts"`
	Age      string `json:"age"`
}

// PodLister describes the behavior to perform PODs operations
type PodLister interface {
	// Total returns the number of Pods running in a namespace.
	Total(namespace string) (int, error)
	// List gets a set of Pods objects given a namespace.
	List(namespace string) ([]*Pod, error)
}
