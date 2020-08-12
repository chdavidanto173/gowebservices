package main

import (
	"log"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	var err error
	readData("books.csv")
	switch request.Method {
	case "GET":
		err = handleGet(writer, request)
	case "POST":
		err = handlePost(writer, request)
	case "PUT":
		err = handlePut(writer, request)
	case "DELETE":
		err = handleDelete(writer, request)
	}
	writeData("books.csv")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/book/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
