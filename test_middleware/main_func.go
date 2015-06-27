package main

import (
	"net/http"
)

type SingleHost struct {
	handler   http.Handler
	allowHost string
}

func (this *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Host == this.allowHost {
		this.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func main() {
	single := &SingleHost{
		handler:   http.HandlerFunc(myHandler), //标准库内部将func 转成handler
		allowHost: "example.com",
	}
	http.ListenAndServe(":8080", single)
}
