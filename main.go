package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/KazukiHayase/datastore-todo-app/graph/generated"
	"github.com/KazukiHayase/datastore-todo-app/graph/resolver"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/healthz"}}))
	r.Use(gin.Recovery())

	r.POST("/graphql", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.GET("/healthz", func(c *gin.Context) {
		c.String(200, "I'm healthy")
	})

	r.Run(":8888")
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
