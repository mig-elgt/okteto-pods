package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.Int("p", 8080, "service port")
	flag.Parse()

	http.HandleFunc("/healthy", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	fmt.Println("Server running al localhost: ", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		panic(err)
	}
}
