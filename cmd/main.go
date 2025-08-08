package main

import (
	database "github.com/tinnoha/pet-progect/app/Database"
	handlers "github.com/tinnoha/pet-progect/app/Handlers"
)

func main() {
	database.ChangeBalace(10, "3")
	handlers.Run()

}
