package db

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() error {
 var err error
 err = godotenv.Load()
 if err != nil {
  log.Println("error loading .env file: ", err)
  return err
 }
 dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
 DB, err = sql.Open("postgres", dsn)
 if err != nil {
  log.Println("error while starting the psql server: ", err)
  return err
 }
 if err = DB.Ping(); err != nil {
  log.Println("error making a test ping to the server: ", err)
  return err
 }
 log.Println("Database connected successfully")
 return nil
}
