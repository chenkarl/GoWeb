package main

import (
	"encoding/json"
	"net/http"
)

//Post 格式
type Post struct {
	User string
	Sex  string
}

func reloactionExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://127.0.0.1:8080/json")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "Chen YuZhao",
		Sex:  "Male",
	}
	j, _ := json.Marshal(post)
	w.Write(j)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/relocation", reloactionExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
