package main

import (
	"net/http"
)

func SingleHost(handler http.Handler, allowHost string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		println(r.Host)
		if r.Host == allowHost {
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
		}
	}
	return http.HandlerFunc(fn)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func main() {
	single := SingleHost(http.HandlerFunc(myHandler), "example.com")
	http.ListenAndServe(":8080", single)
}
