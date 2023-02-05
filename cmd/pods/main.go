package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/mig-elgt/okteto-pods/handler"
	"github.com/mig-elgt/okteto-pods/kubernetes"
	"github.com/mig-elgt/okteto-pods/sort"
	log "github.com/sirupsen/logrus"
)

func main() {
	port := flag.Int("p", 8080, "service port")
	flag.Parse()

	k8s, err := kubernetes.New()
	if err != nil {
		log.Fatalf("could not create kubernetes client: %v", err)
	}
	h := handler.New(k8s, sort.New())

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: h,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Infof("Server running al localhost: %v", *port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
