package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/mig-elgt/okteto-pods/handler"
	"github.com/mig-elgt/okteto-pods/kubernetes"
	"github.com/sirupsen/logrus"
)

func main() {
	port := flag.Int("p", 8080, "service port")
	flag.Parse()

	k8s, err := kubernetes.New()
	if err != nil {
		logrus.Fatalf("could not create kubernetes client: %v", err)
	}
	h := handler.New(k8s)

	fmt.Println("Server running al localhost: ", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), h); err != nil {
		panic(err)
	}
}
