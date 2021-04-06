package main

import (
	"fmt"
	"io"
	"net/http"
)

func echoRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Request protocol: %v, method: %v\n", r.Proto, r.Method)
	fmt.Fprintf(w, "To: %v\n", r.Host)
	fmt.Fprintf(w, "From: %v\n", r.RemoteAddr)
	fmt.Fprintf(w, "Url requested: %s\n\n", r.URL)
	fmt.Fprintln(w, "Request header contents:")
	for k, v := range r.Header {
		fmt.Fprintf(w, "%v\t%v\n", k, v)
	}
	if body, err := io.ReadAll(r.Body); err == nil {
		fmt.Fprintf(w, "Request body:\n%v\n", string(body))
		r.Body.Close()
	} else {
		fmt.Fprintln(w, "Could not read request body")
	}
}

func main() {
	http.HandleFunc("/", echoRequest)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
