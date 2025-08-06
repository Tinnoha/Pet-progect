package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"log"
	"net/http"
	"strconv"

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

		rezult, err := DB.Query("Select * from polzovately")
		if err != nil {
			log.Fatal(err)
		}

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

		for rezult.Next() {
			var id int
			var first_name, middle_name, last_name, PasswordHash, email string
			var balance int64
			err := rezult.Scan(&id, &first_name, &middle_name, &last_name, &PasswordHash, &email, &balance)
			if err != nil {
				log.Fatal(err)
			}
			strochka := strconv.Itoa(id) + " " + first_name + " " + middle_name + " " + last_name + " " + PasswordHash + " " + email + " " + strconv.FormatInt(balance, 10)
			fmt.Fprint(writer, strochka, "\n")

		}
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
			"INSERT INTO polzovately (first_name,middle_name,last_name,PasswordHash, email, balance) values ($1,$2, $3, $4, $5, $6)",
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

		var count = 0

		sqlRequest := "UPDATE polzovately"

		if vanya.First_name != "" {
			sqlRequest = sqlRequest + fmt.Sprintf("\nSET first_name = '%s'", vanya.First_name)
			count++
		}
		if vanya.Middle_name != "" {
			if count != 0 {
				sqlRequest += fmt.Sprintf("\n, middle_name = '%s'", vanya.Middle_name)
			} else {
				sqlRequest = sqlRequest + fmt.Sprintf("\nSET middle_name = '%s'", vanya.Middle_name)
				count++
			}
		}
		if vanya.Last_name != "" {
			if count != 0 {
				sqlRequest += fmt.Sprintf("\n, last_name = '%s'", vanya.Last_name)
			} else {
				sqlRequest = sqlRequest + fmt.Sprintf("\nSET last_name = '%s'", vanya.Last_name)
				count++
			}
		}
		if vanya.Email != "" {
			if count != 0 {
				sqlRequest += fmt.Sprintf("\n, email = '%s'", vanya.Email)
			} else {
				sqlRequest = sqlRequest + fmt.Sprintf("\nSET email = '%s'", vanya.Email)
				count++
			}
		}
		if vanya.PasswordHash != "" {
			if count != 0 {
				sqlRequest += fmt.Sprintf("\n, passwordHash = '%s'", vanya.PasswordHash)
			} else {
				sqlRequest = sqlRequest + fmt.Sprintf("\nSET passwordHash = '%s'", vanya.PasswordHash)
				count++
			}
		}
		if vanya.Balance != 0 {
			balik := strconv.Itoa(int(vanya.Balance))
			if count != 0 {
				sqlRequest += fmt.Sprintf("\n, balance = '%s'", balik)
			} else {
				sqlRequest = sqlRequest + fmt.Sprintf("\nSET balance = %s", balik)
				count++
			}
		}
		// id_user, err := strconv.Atoi(id)

		// if err != nil {
		// 	log.Fatal(err)
		// }

		sqlRequest = sqlRequest + fmt.Sprintf(" WHERE id =  %s;", id)

		fmt.Println(sqlRequest)

		DB, err := sql.Open("postgres", models.ConnStr)

		if err != nil {
			log.Fatal(err)
		}
		defer DB.Close()

		err = DB.Ping()

		if err != nil {
			log.Fatal(err)
		}

		_, err = DB.Exec(sqlRequest)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Fprintf(writer, "Данный запрос не поддерживается")
	}
}

func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /list", handlGetList)
	mux.HandleFunc("POST /list", handlePostList)

	mux.HandleFunc("PUT /list/{id}", handlPutList)

	http.ListenAndServe(":8080", mux)
}
