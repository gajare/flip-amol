package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"user-serice/internal/user/database"
)

const (
	defaultPort = "8081"
	defaultDB   = "postgres://admin:password@postgres:5432/userdb"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultDB
	}
	url_db := os.Getenv("DATABASE_URL")
	if url_db == "" {
		url_db = defaultDB
	}
	// ---------------------- Database Connection ----------------------
	ctx := context.Background()
	db_pool, err := database.NewDBPool(ctx, url_db)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db_pool.Close()

	// ---------------------- Layers ----------------------

	// userRepo := repository.NewUserRepository(db_pool)
	// userService := service.NewUserService(userRepo)

	//---------------------- GraphQL Resolver ----------------------

	// resolver := &graphql.Resolver{
	// 	UserService: userService,
	// }

	// // ---------------------- GraphQL Server ----------------------
	// serv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}))

	// // ---------------------- Routes ----------------------

	// http.Handle("/", playgroud.Handler("GraphQL Playgorund", "/query"))
	// http.Handle("/query", nil)

	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
