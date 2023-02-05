package pods

import (
	"errors"
	"time"
)

// Pod describes a Kubernetes Pod resource.
type Pod struct {
	Name     string
	Status   string
	Restarts int32
	Age      time.Duration
}

// PodLister describes the behavior to perform PODs operations
type PodLister interface {
	// Total returns the number of Pods running in a namespace.
	Total(namespace string) (int, error)

	// List gets a set of Pods objects given a namespace.
	List(namespace string) ([]*Pod, error)
}

type FieldOrder struct {
	// Field represents a Pods' field to use in the Sort algorithm
	// It could have the next range values: name, restarts and age.
	Field string

	// Order describes the a field order: asc or desc
	Order string
}

// Sorter defines an interface to perform a Sort Algorithm given a set of Pods and
// fields values.
type Sorter interface {
	// Sort performs the Sort algorithm for a list of Pods and a set
	// of fields values to get the order.
	Sort(pods []*Pod, fields []FieldOrder) []*Pod
}

var (
	ErrIncorrectSortValue = errors.New("Incorrect field order value, should be asc or desc.")
)
