// should read: https://rcrowley.org/talks/strange-loop-2013.html#1

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Message struct {
	Text string `json:"msg"`
}

// my version copied from tsenart's, looks like more of a mess but it works!
func welcomeHandler(config string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprintf(w, "Welcome!, config: %s, user: %s", config, name)
	}
}

func middleware(l *log.Logger, next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		began := time.Now()
		next(w, r)
		l.Printf("%s: %s %s took %s", time.Now(), r.Method, r.URL, time.Since(began))
	}
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

func serve(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, ">>> A long task finished.")
	time.Sleep(5 * time.Second)
	// test: for i in $(seq 1 3); do (time curl -sS localhost:8080/serve &); done
}

// took from the awesome: https://gist.github.com/tsenart/5fc18c659814c078378d
func userHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprintf(w, "user: %s", name)
	})
}

func withMetrics(l *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		began := time.Now()
		next.ServeHTTP(w, r)
		l.Printf("%s: %s %s took %s", time.Now(), r.Method, r.URL, time.Since(began))
	})
}

func main() {
	// Create our logger
	logger := log.New(os.Stdout, "", 0)

	r := mux.NewRouter()

	// my version of 'HTTP closure'
	config := "my config"
	r.HandleFunc("/welcome/{name}", welcomeHandler(config))
	r.HandleFunc("/_welcome/{name}", middleware(logger, welcomeHandler(config)))

	// hanlders without closure
	r.HandleFunc("/about", about)
	r.HandleFunc("/hello", hello)
	r.HandleFunc("/serve", serve)

	// tsenart's version (y)
	r.Handle("/user/{name}", userHandler())
	r.Handle("/_user/{name}", withMetrics(logger, userHandler()))

	http.ListenAndServe(":8080", r)
}
