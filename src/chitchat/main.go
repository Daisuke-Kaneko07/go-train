package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	// "time"
)

type Post struct {
	User string
	Threads []string
}

func process(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
	fmt.Fprintln(w, r.MultipartForm)
}

func writeExample (w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go Web Programming</title></head>
	<body><h1>Hello World</h1></body>
	</html>`
	w.Write([]byte(str));
}

func writeHeaderExample (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501);
	fmt.Fprintln(w, "そのようなサービスはありません。ほかを当たってください。")
}

func headerExample (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func jsonExample (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "Daisuke Kaneko",
		Threads: []string{"1番目", "2番目", "3番目"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	// starting up the server
	server := http.Server{
	// server := &http.Server{
		Addr:           "127.0.0.1:8080",
		// Addr:           config.Address,
		// Handler:        mux,
		// ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		// WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		// MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
