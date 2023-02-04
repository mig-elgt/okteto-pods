package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/mig-elgt/okteto-pods/handler"
)

func main() {
	port := flag.Int("p", 8080, "service port")
	flag.Parse()

	h := handler.New()

	fmt.Println("Server running al localhost: ", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), h); err != nil {
		panic(err)
	}
}
