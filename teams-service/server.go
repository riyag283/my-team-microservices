package main

import (
	"context"
	"log"
	"teams/db"
	"teams/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GinContextKey string

const ginContextKey GinContextKey = "ginContext"

// Defining the Graphql handler
func graphqlHandler(logger *zap.Logger) gin.HandlerFunc {
	h := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: &graph.Resolver{Logger: logger},
		}),
	)

	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ginContextKey, c)
		h.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	// Deferring the logger sync
	defer logger.Sync()

	// Logging a message
	logger.Info("Logger initialized successfully")

	db.InitialiseDBConnection()

	// Setting up Gin
	r := gin.Default()

	// Log a message to indicate that the server has started
	logger.Info("Server started")

	// Set up the GraphQL handler
	r.POST("/my-team", graphqlHandler(logger))

	// Set up the Playground handler
	r.GET("/", playgroundHandler())

	// Start the server
	if err := r.Run(":8080"); err != nil {
		logger.Error("Server failed to start", zap.Error(err))
	}
	r.Run()
}
