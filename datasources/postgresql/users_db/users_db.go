package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	Client *sql.DB
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	username := os.Getenv("mysql_users_username")
	password := os.Getenv("mysql_users_password")
	host := os.Getenv("mysql_users_host")
	schema := os.Getenv("mysql_users_schema")
	dataSourceName := fmt.Sprintf("%s:%s@%s/%s", username, password, host, schema)
	Client, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully connected")
}
