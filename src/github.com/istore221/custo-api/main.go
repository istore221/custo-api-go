package main

import (
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}


