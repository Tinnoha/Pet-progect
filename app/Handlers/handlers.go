package handlers

import (
	"fmt"
	"net/http"
)

func handlGetList(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		fmt.Fprintf(writer, "Hello world from localhost:8080%s", request.URL)
	} else if request.Method == http.MethodPost {
		fmt.Fprint(writer, "Give your data")
	}
}

func Run() {
	http.HandleFunc("/list", handlGetList)

	http.ListenAndServe(":8080", nil)
}
