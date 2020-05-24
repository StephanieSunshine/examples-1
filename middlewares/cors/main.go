package main

import (
	"github.com/atreugo/cors"
	"github.com/savsgio/atreugo/v11"
)

func main() {
	config := atreugo.Config{
		Addr: "0.0.0.0:8001",
	}
	server := atreugo.New(config)

	cors := cors.New(cors.Config{
		AllowedOrigins:   []string{"http://localhost:8001", "null"},
		AllowedHeaders:   []string{"Content-Type", "X-Custom"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		ExposedHeaders:   []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		AllowMaxAge:      5600,
	})

	server.UseAfter(cors)

	server.GET("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.JSONResponse(atreugo.JSON{"Method": "GET"})
	})

	server.POST("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.JSONResponse(atreugo.JSON{"Method": "POST"})
	})

	server.PUT("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.JSONResponse(atreugo.JSON{"Method": "PUT"})
	})

	server.DELETE("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.JSONResponse(atreugo.JSON{"Method": "DELETE"})
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
