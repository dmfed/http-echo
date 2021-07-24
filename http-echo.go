package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func echoRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Request protocol: %v, method: %v\n", r.Proto, r.Method)
	fmt.Fprintf(w, "To: %v\n", r.Host)
	fmt.Fprintf(w, "Url requested: %s\n", r.URL)
	fmt.Fprintf(w, "From: %v\n\n", r.RemoteAddr)
	fmt.Fprintln(w, "Request header contents:")
	for k, v := range r.Header {
		fmt.Fprintf(w, "%v\t%v\n", k, v)
	}
	if body, err := io.ReadAll(r.Body); err == nil {
		fmt.Fprintf(w, "\nRequest body:\n%v\n", string(body))
		r.Body.Close()
	} else {
		fmt.Fprintln(w, "\nCould not read request body")
	}
}

func main() {
	var (
		flagListen = flag.String("addr", ":8080", "address to listen on")
	)
	flag.Parse()
	http.HandleFunc("/", echoRequest)
	http.ListenAndServe(*flagListen, nil)
}
