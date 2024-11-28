package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/GigaDesk/eardrum-server/auth"
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-server/database/postgreutils"
	"github.com/GigaDesk/eardrum-server/graph"
	"github.com/GigaDesk/eardrum-server/phoneutils"
	"github.com/GigaDesk/eardrum-server/pkg/jwt"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

var (
	postgresInstance postgreutils.PostgresInstance
	neo4jInstance  neo4jutils.Neo4jInstance
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

	postgresInstance.Init(os.Getenv("POSTGRES_DBURL"))
	neo4jerr:=neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))

	if neo4jerr !=nil {
		log.Println("problem initializing neo4j database: ", neo4jerr)
	}
	
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)

	port := defaultPort
	router := chi.NewRouter()
	router.Use(auth.Middleware())

	//srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Sql: &dborm}})) //.... <- here set dborm to resolver
	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Sql: &postgresInstance.Dborm, Neo4j: &neo4jInstance }}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
