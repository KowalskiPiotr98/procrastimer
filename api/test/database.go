package test

import (
	"fmt"
	"github.com/KowalskiPiotr98/gotabase"
	"github.com/gofrs/uuid"
	"os"
	"procrastimer/database"
	"testing"

	_ "github.com/lib/pq"
)

func GetDatabase() (gotabase.Connector, string) {
	dbName, _ := uuid.NewV4()
	baseConnectionString := getBaseConnectionString()
	err := gotabase.InitialiseConnection(baseConnectionString+dbName.String(), "postgres")
	if err != nil {
		PanicOnErr(gotabase.InitialiseConnection(baseConnectionString+"postgres", "postgres"))
		_, err = gotabase.GetConnection().Exec(fmt.Sprintf("create database \"%s\"", dbName.String()))
		PanicOnErr(err)
		PanicOnErr(gotabase.CloseConnection())
		PanicOnErr(gotabase.InitialiseConnection(baseConnectionString+dbName.String(), "postgres"))
	}
	PanicOnErr(database.RunMigrations(gotabase.GetConnection()))
	return gotabase.GetConnection(), dbName.String()
}

func DropDatabase(dbName string) {
	PanicOnErr(gotabase.CloseConnection())
	PanicOnErr(gotabase.InitialiseConnection(getBaseConnectionString()+"postgres", "postgres"))
	_, err := gotabase.GetConnection().Exec(fmt.Sprintf("drop database \"%s\"", dbName))
	PanicOnErr(err)
	PanicOnErr(gotabase.CloseConnection())
}

func GetDatabaseWithCleanup(t *testing.T) gotabase.Connector {
	db, name := GetDatabase()
	t.Cleanup(func() { DropDatabase(name) })
	return db
}

func getBaseConnectionString() string {
	baseConnectionString := os.Getenv("TEST_POSTGRES")
	if baseConnectionString == "" {
		baseConnectionString = "user=postgres password=postgres sslmode=disable dbname="
	}
	return baseConnectionString
}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
