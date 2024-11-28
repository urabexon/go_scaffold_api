package main

import {
	"io"
	"log"
	"net/http"
}

func main(){
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World!")
	}

	http.HandleFunc("/hello", helloHander)

	log.Println("server start at port: 8080")
	log.Fatal(http.ListenAndServer(":8080", nil))
}
