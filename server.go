package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/GigaDesk/eardrum-server/auth"
	"github.com/GigaDesk/eardrum-server/graph"
	"github.com/GigaDesk/eardrum-server/graph/db"
	"github.com/GigaDesk/eardrum-server/phoneutils"
	"github.com/GigaDesk/eardrum-server/pkg/jwt"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
    
	go phoneutils.InitializeTwilio()
	go jwt.InitializeJwtSecretKey()

	defaultPort := os.Getenv("DEFAULT_PORT")
	host := os.Getenv("POSTGRES_CLOUD_SQL_ACCOUNT_HOST")
	user := os.Getenv("POSTGRES_CLOUD_SQL_ACCOUNT_USER")
	password := os.Getenv("POSTGRES_CLOUD_SQL_ACCOUNT_PASSWORD")
	dbname := os.Getenv("POSTGRES_CLOUD_SQL_ACCOUNT_DBNAME")
	dbport := os.Getenv("POSTGRES_CLOUD_SQL_ACCOUNT_PORT")
	sslmode := os.Getenv("POSTGRES_CLOUD_SQL_ACCOUNT_SSLMODE")
	TimeZone := os.Getenv("POSTGRES_CLOUD_SQL_ACCOUNT_TIMEZONE")

	dbUrl := "host=" + host + " " + "user=" + user + " " + "password=" + password + " " + "dbname=" + dbname + " " + "port=" + dbport + " " + "sslmode=" + sslmode + " " + "TimeZone=" + TimeZone
	dbInstance, dbError := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if dbError != nil {
		panic(dbError)
	}
	dborm := db.NewAutoGqlDB(dbInstance)
	dborm.Init()

	port := defaultPort
	router := chi.NewRouter()
	router.Use(auth.Middleware())

	//srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Sql: &dborm}})) //.... <- here set dborm to resolver
	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Sql: &dborm}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
