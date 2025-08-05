package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/tinnoha/pet-progect/app/models"
)

func handlGetList(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		fmt.Fprintf(writer, "Get from localhost:8080%s \n", request.URL)

		DB, err := sql.Open("postgres", models.ConnStr)
		if err != nil {
			log.Fatal(err)
		}

		defer DB.Close()

		err = DB.Ping()
		if err != nil {
			log.Fatal(err)
		}

		rezult, err := DB.Query("Select * from users")
		if err != nil {
			log.Fatal(err)
		}
		for rezult.Next() {
			var user_id int
			var full_name string
			var username string
			var email string
			var password_hash string
			var role string
			var created_at time.Time
			err := rezult.Scan(&user_id, &full_name, &username, &email, &password_hash, &role, &created_at)
			if err != nil {
				log.Fatal(err)
			}
			strochka := strconv.Itoa(user_id) + " " + full_name + " " + username + " " + email + " " + password_hash + " " + role
			fmt.Fprint(writer, strochka, "\n")

		}

	} else if request.Method == http.MethodPost {
		fmt.Fprint(writer, "Give your data")
	}
}

func Run() {
	http.HandleFunc("/list", handlGetList)

	http.ListenAndServe(":8080", nil)
}
