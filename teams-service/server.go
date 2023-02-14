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

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: &graph.Resolver{},
		}),
	)

	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "ginContext", c)
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

	defer logger.Sync()

	db.InitialiseDBConnection()

	// Setting up Gin
	r := gin.Default()
	r.POST("/my-team", graphqlHandler())
	r.GET("/", playgroundHandler())

	// Log a message when the server starts running
	logger.Info("Server started running on port 8080")

	if err := r.Run(":8080"); err != nil {
		// Log an error message if the server fails to start
		logger.Error("Failed to start server", zap.Error(err))
	}
}
