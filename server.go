package main

import (
	"fmt"
	"log"
	"ming_backend/graph"
	"ming_backend/graph/generated"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jinzhu/gorm"
	"ming_backend/graph/model"
)

const defaultPort = "8080"

var db *gorm.DB

func initDB() {
	host := os.Getenv("DB_HOST")
	ussername := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	var err error
	dataSourceName := `root:@tcp(66.42.111.28:3306)/?parseTime=True`
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.LogMode(true)

	//// Create the database. This is a one-time step.
	//// Comment out if running multiple times - You may see an error otherwise
	//db.Exec("CREATE DATABASE test_db")
	//db.Exec("USE test_db")
	//
	//// Migration to create tables for Order and Item schema
	//db.AutoMigrate(&model.Invoice{}, &model.Item{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	initDB()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
