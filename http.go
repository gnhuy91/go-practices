package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string `json:"msg"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("Welcome, %s!", r.URL.Path[1:])
	// if composed the string in a struct, json.Marshal can return json (key&value)
	// based on struct field name (if no `json:field` specified) and string content.
	j, _ := json.Marshal(&Message{s})
	w.Write(j)
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	m := Message{"go API, build v0.0.001.992."}
	j, _ := json.Marshal(&m)
	w.Write(j)
}

func hello(w http.ResponseWriter, r *http.Request) {
	// If a raw value, Marshal doesn't seem to do anything.
	j, _ := json.Marshal("Hello from API")
	w.Write(j)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", about)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
