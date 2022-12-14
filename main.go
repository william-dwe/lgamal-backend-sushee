package main

import (
	"fmt"

	"final-project-backend/db"
	"final-project-backend/server"
)

func main() {
	dbErr := db.Connect()
	if dbErr != nil {
		fmt.Println("error connecting to DB")
	}
	server.Init()
}
