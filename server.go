package main

import "net/http"

func main() {

	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)

}

func Hello(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Hello FullCycle - Kubernetes - Versão 02</h1>"))

}