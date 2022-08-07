package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/KazukiHayase/datastore-todo-app/config"
	"github.com/KazukiHayase/datastore-todo-app/graph/generated"
	"github.com/KazukiHayase/datastore-todo-app/graph/resolver"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.Environ()
	if err != nil {
		log.Fatalf("環境変数の設定に失敗しました: %v\n", err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &resolver.Resolver{
				Config: config,
			},
		},
	))

	r := gin.Default()
	r.POST("/graphql", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})
	r.GET("/", playgroundHandler())

	r.Run(":8888")
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
