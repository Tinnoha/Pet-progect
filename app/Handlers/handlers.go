package handlers

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	_ "github.com/lib/pq"
	database "github.com/tinnoha/pet-progect/app/Database"
	"github.com/tinnoha/pet-progect/app/models"
)

func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /list", handlGetList)
	mux.HandleFunc("POST /list", handlePostList)

	mux.HandleFunc("PUT /list/{id}", handlPutList)
	mux.HandleFunc("DELETE /list/{id}", handlDeleteList)

	http.ListenAndServe(":8080", mux)
}

func handlGetList(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		fmt.Fprintf(writer, "Get from localhost:8080%s \n", request.URL)

		database.PrintGetRows("polzovately", writer)

	} else {
		fmt.Fprintf(writer, "Данные метод не поддерживается")
	}
}

func handlePostList(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		fmt.Fprint(writer, "Give your data")

		petya := models.User{}

		err := json.NewDecoder(request.Body).Decode(&petya)

		if err != nil {
			log.Fatal(err)
		}

		database.InsertDataBase("polzovately", petya)

	} else {
		fmt.Fprintf(writer, "Данный запрос не поддерживается")
	}
}

func handlPutList(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPut {
		id := request.PathValue("id")
		fmt.Println(id)
		var vanya models.User

		err := json.NewDecoder(request.Body).Decode(&vanya)

		if err != nil {
			log.Fatal(err)
		}

		database.UpdateDataBase("polzovately", vanya, id)

	} else {
		fmt.Fprintf(writer, "Данный запрос не поддерживается")
	}
}
func handlDeleteList(w http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodDelete {
		id := request.PathValue("id")

		database.DeleteDataBase("polzovately", id)

	} else {
		fmt.Fprintf(w, "Данный метод не поддерживается")
	}
}
