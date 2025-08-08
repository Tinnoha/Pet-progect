package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/tinnoha/pet-progect/app/models"
)

func connection(dataSource string) *sql.DB {
	DB, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return DB
}

func PrintGetRows(tableName string, w http.ResponseWriter) {
	Database := connection(models.ConnStr)
	defer Database.Close()

	rows, err := Database.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int
		var first_name, middle_name, last_name, email, PasswordHash string
		var balance int64

		err := rows.Scan(&id, &first_name, &middle_name, &last_name, &PasswordHash, &email, &balance)

		if err != nil {
			log.Fatal(err)
		}

		str := strconv.Itoa(id) + " " + first_name + " " + middle_name + " " + last_name + " " + PasswordHash + " " + email + " " + strconv.FormatInt(balance, 10) + "\n"
		fmt.Fprintf(w, str)

	}

}

func InsertDataBase(tableName string, petya models.User) {
	Database := connection(models.ConnStr)
	defer Database.Close()

	_, err := Database.Exec(fmt.Sprintf("INSERT INTO %s (first_name,middle_name,last_name,passwordhash,email,balance) VALUES ($1,$2,$3,$4,$5,$6)", tableName),
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
}

func UpdateDataBase(tableName string, vanya models.User, id string) {
	var count = 0
	sqlRequest := fmt.Sprintf("UPDATE %s", tableName)

	if vanya.First_name != "" {
		sqlRequest += fmt.Sprintf("\nSET first_name = '%s'", vanya.First_name)
		count++
	}
	if vanya.Middle_name != "" {
		if count == 0 {
			sqlRequest += fmt.Sprintf("\nSET middle_name = '%s'", vanya.Middle_name)
			count++
		} else {
			sqlRequest += fmt.Sprintf(",\n\tmiddle_name = '%s'", vanya.Middle_name)
		}
	}
	if vanya.Last_name != "" {
		if count == 0 {
			sqlRequest += fmt.Sprintf("\nSET last_name = '%s'", vanya.Last_name)
			count++
		} else {
			sqlRequest += fmt.Sprintf(",\n\tlast_name = '%s'", vanya.Last_name)
		}
	}
	if vanya.PasswordHash != "" {
		if count == 0 {
			sqlRequest += fmt.Sprintf("\nSET PasswordHash = '%s'", vanya.PasswordHash)
			count++
		} else {
			sqlRequest += fmt.Sprintf(",\n\tPasswordHash = '%s'", vanya.PasswordHash)
		}
	}
	if vanya.Email != "" {
		if count == 0 {
			sqlRequest += fmt.Sprintf("\nSET Email = '%s'", vanya.Email)
			count++
		} else {
			sqlRequest += fmt.Sprintf(",\n Email = '%s'", vanya.Email)
		}
	}
	sqlRequest += fmt.Sprintf("\nWHERE id = %s", id)

	DB := connection(models.ConnStr)
	defer DB.Close()

	_, err := DB.Exec(sqlRequest)

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteDataBase(tableName string, id string) {
	DataBase := connection(models.ConnStr)
	defer DataBase.Close()

	DataBase.Exec(fmt.Sprintf("delete from %s Where id = %s", tableName, id))
}

func ChangeBalace(count int, id string) {
	database := connection(models.ConnStr)

	if count == 0 {
		log.Fatal("Нет изменений")
	}

	_, err := database.Exec(fmt.Sprintf("UPDATE polzovately SET balance = balance + %d WHERE id = %s", count, id))

	if err != nil {
		log.Fatal(err)
	}
}

/*
create table polzovately (
	id serial primary key,
	first_name varchar(40),
	middle_name varchar(40) default 'Нет отчества',
	last_name varchar(40),
	PasswordHash varchar,
	email varchar (100),
	balance bigint default 0
);
*/
