package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestPlaygroundHandler(t *testing.T) {
	// Create a test router and add the Playground handler to it
	router := gin.Default()
	router.GET("/playground", playgroundHandler())

	// Create a new HTTP request to the Playground endpoint
	req, err := http.NewRequest("GET", "/playground", nil)
	assert.NoError(t, err)

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Perform the request and record the response
	router.ServeHTTP(w, req)

	// Check that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Check that the response body is not empty
	assert.NotEmpty(t, w.Body.String())
}

func TestGraphqlHandler(t *testing.T) {
    logger, err := zap.NewProduction()
    if err != nil {
        log.Fatalf("can't initialize zap logger: %v", err)
    }

    // Deferring the logger sync
    defer logger.Sync()

    // Logging a message
    logger.Info("Logger initialized successfully")

    // Create a test router and add the Playground and GraphQL handlers to it
    router := gin.Default()
    router.GET("/", playgroundHandler())
    router.POST("/query", graphqlHandler(logger))

    // Create a new HTTP request to the GraphQL endpoint
    req, err := http.NewRequest("POST", "/query", nil)
    assert.NoError(t, err)

    // Create a new HTTP response recorder
    w := httptest.NewRecorder()

    // Perform the request and record the response
    router.ServeHTTP(w, req)

    // Check that the response status code is 200 OK
    assert.Equal(t, http.StatusBadRequest, w.Code)

    // Check that the response body is not empty
    assert.NotEmpty(t, w.Body.String())
}

func TestMainFunction(t *testing.T) {
    logger, _ := zap.NewProduction()
    defer logger.Sync()

    go main()

    time.Sleep(2 * time.Second)

    resp, err := http.Get("http://localhost:8080/")
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)

    resp, err = http.Post("http://localhost:8080/my-team", "application/json", nil)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
}
