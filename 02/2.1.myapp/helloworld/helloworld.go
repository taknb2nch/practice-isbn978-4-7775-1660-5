package helloworld

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(rw, "Hello world")
	fmt.Fprintf(rw, "Hello Google App Engine")
}
