package main

import (
	"fmt"
	"log"
	"net/http"
)

type srv struct{}

func (s *srv) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("当前路径为：%s\n", r.URL.Path)
	switch r.URL.Path {
	case "/user/abc":
		w.Write([]byte("Hello User >> abc"))
	case "/user":
		w.Write([]byte("Hello User"))
	default:
		w.Write([]byte("Hello Index"))
	}
}

func main() {
	s := new(srv)

	err := http.ListenAndServe(":8081", s)
	if err != nil {
		log.Fatal(err.Error())
	}
}
