package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/tinnoha/pet-progect/app/models"
)

func handlGetList(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		fmt.Fprintf(writer, "Get from localhost:8080%s", request.URL)

		DB, err := sql.Open("postgres", models.ConnStr)
		if err != nil {
			log.Fatal(err)
		}

		defer DB.Close()

		err = DB.Ping()
		if err != nil {
			log.Fatal(err)
		}

	} else if request.Method == http.MethodPost {
		fmt.Fprint(writer, "Give your data")
	}
}

func Run() {
	http.HandleFunc("/list", handlGetList)

	http.ListenAndServe(":8080", nil)
}
