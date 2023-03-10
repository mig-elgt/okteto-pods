package kubernetes

import (
	"context"
	"time"

	pods "github.com/mig-elgt/okteto-pods"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// kubernetes implements a PodLister interface.
type kubernetes struct {
	client *k8s.Clientset
}

// New creates new kubernetes instance and stores a clientset to Kubernetes API
func New() (*kubernetes, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, errors.Wrap(err, "could not create cluster config")
	}
	clientset, err := k8s.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "could not create client set")
	}
	return &kubernetes{client: clientset}, nil
}

// Total returns the number of Pods running in a namespace.
func (k *kubernetes) Total(namespace string) (int, error) {
	pods, err := k.client.CoreV1().Pods("mig-elgt").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return 0, errors.Wrap(err, "could not list pods")
	}
	return len(pods.Items), nil
}

// List gets a set of Pods objects given a namespace.
func (k *kubernetes) List(namespace string) ([]*pods.Pod, error) {
	podsList, err := k.client.CoreV1().Pods("mig-elgt").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "could not list pods")
	}
	var result []*pods.Pod
	for _, pod := range podsList.Items {
		var containerRestars int32
		for _, cs := range pod.Status.ContainerStatuses {
			containerRestars += cs.RestartCount
		}
		podStatus := "Running"
		for _, condition := range pod.Status.Conditions {
			if condition.Status != "True" {
				podStatus = "Pending"
				break
			}
		}
		age := time.Since(pod.CreationTimestamp.Time).Round(time.Second)
		result = append(result, &pods.Pod{
			Name:     pod.Name,
			Status:   podStatus,
			Age:      age,
			Restarts: containerRestars,
		})
	}
	return result, nil
}
