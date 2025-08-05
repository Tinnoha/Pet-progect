package handlers

import (
	"database/sql"
	"encoding/json"
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
			var full_name, username, email, password_hash, role string
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

		petya := models.User{}

		err := json.NewDecoder(request.Body).Decode(&petya)

		if err != nil {
			log.Fatal(err)
		}

		DB, err := sql.Open("postgres", models.ConnStr)

		if err != nil {
			log.Fatal(err)
		}

		err = DB.Ping()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(petya)

		_, err = DB.Exec(
			"INSERT INTO polzovately (id,first_name,middle_name,last_name,PasswordHash, email, balance) values ($1,$2, $3, $4, $5, $6, $7)",
			petya.Id,
			petya.First_name,
			petya.Middle_name,
			petya.Last_name,
			petya.PasswordHash,
			petya.Email,
			petya.Balance,
		)
		if err != nil {
			log.Fatal(err)
		}
		DB.Close()

		/*
			create table polzovately ( Таблица юзеров
				id serial,
				first_name varchar[40],
				middle_name varchar[40],
				last_name varchar[40],
				PasswordHash varchar,
				email varchar [100],
				balance bigint default 0
			);
		*/

	}
}

func Run() {
	http.HandleFunc("/list", handlGetList)

	http.ListenAndServe(":8080", nil)
}
