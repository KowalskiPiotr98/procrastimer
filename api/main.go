package main

import (
	"github.com/KowalskiPiotr98/gotabase"
	_ "github.com/lib/pq"
	"procrastimer/database"
)

func main() {
	initDatabase()
}

func initDatabase() {
	//todo: read this from ENV
	err := gotabase.InitialiseConnection("user=postgres dbname=procrastimer password=postgres sslmode=disable", "postgres")
	if err != nil {
		panic(err)
	}
	defer gotabase.CloseConnection()
	if err = database.RunMigrations(gotabase.GetConnection()); err != nil {
		panic(err)
	}
}
