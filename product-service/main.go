package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"product-service/graph"
	"product-service/graph/generated"
	"product-service/internal/product/database"
	"product-service/internal/product/repository"
	"product-service/internal/product/service"
)

const (
	defaultPort = "8080"
	defaultDB = "postgres://admin:password@postgres:5432/productdb"
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

	// ---------------- DB ----------------
	ctx := context.Background()
	dbPool, err := database.NewPostgresPool(ctx, dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbPool.Close()

	// ------------- Layers ---------------
	productRepo := repository.NewProductRepository(dbPool)
	productService := service.NewProductService(productRepo)

	// --------- GraphQL Resolver ----------
	resolver := &graph.Resolver{
		ProductService: productService,
	}

	// -------- GraphQL Server -------------
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: resolver},
		),
	)

	// ------------ Routes ----------------
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("🚀 GraphQL running at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
