package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
			log.Printf("%s: %s", name, h)
		}
	}

	fmt.Fprintf(w, "hello\n")

}

func hc(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "200\n")

}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/hc", hc)

	http.ListenAndServe(":8090", nil)
}
