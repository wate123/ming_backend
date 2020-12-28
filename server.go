package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"ming_backend/graph"
	"ming_backend/graph/generated"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const defaultPort = "8080"

var db *gorm.DB

func initDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", username, pass, host, port, dbName)
	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

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
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
