package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jackc/pgx/v5/pgxpool"

	"product-service/graph"
	"product-service/internal/database"
	"product-service/internal/product/repository"
	"product-service/internal/product/service"
)

const (
	defaultPort = "8080"
	defaultDB   = "postgres://admin:password@postgres:5432/productdb"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = defaultDB
	}

	// Initialize database connection
	ctx := context.Background()
	dbPool, err := database.NewPostgresPool(ctx, dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbPool.Close()

	// Initialize repository and service
	productRepo := repository.NewProductRepository(dbPool)
	productService := service.NewProductService(productRepo)

	// Create GraphQL resolver
	resolver := graph.NewResolver(productService)

	// Configure GraphQL server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	// Set up routes
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}