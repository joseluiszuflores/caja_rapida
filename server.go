package main

import (
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/configs"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/infrastructure/db/postgres"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/infrastructure/db/postgres/migration"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/graph"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/graph/generated"
)

const defaultPort = "80"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db := postgres.GetDB()
	migration.InitTables(db)
	configs.LoadConfig()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
