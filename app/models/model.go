package models

const ConnStr string = "host=localhost port=5432 user=postgres password=1234 dbname=postgres sslmode=disable"

type User struct {
	Id           int    `json:"id"`
	First_name   string `json:"first_name"`
	Middle_name  string `json:"middle_name"`
	Last_name    string `json:"last_name"`
	PasswordHash string `json:"passwordHash"`
	Email        string `json:"email"`
	Balance      int64  `json:"balance"`
}

// {
//     "id": 1,
//     "first_name": "Иван",
//     "middle_name": "Иванович",
//     "last_name": "Иванов",
//     "PasswordHash": "abc123",
//     "email": "ivan@example.com",
//     "balance": 1000
// }
