package main

import (
	"go-playground/app_config"
	"go-playground/firebase_client"
	"go-playground/handler"
	"go-playground/middlewares"
	"go-playground/redis_client"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	e := echo.New()
	redis_client.InitRedisClient()
	firebase_client.InitFirebaseClient()

	e.Use(middleware.Secure())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     app_config.CorsValidDomains,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
	// 	Skipper: func(echo.Context) bool {
	// 		return true
	// 	},
	// 	Timeout:      10 * time.Second,
	// 	ErrorMessage: "Timeout",
	// }))

	g := e.Group("/chat", middlewares.CheckAccessToken)
	// endpoint for the chat stream
	g.GET("/with", handler.ChatWithEndpoint)
	g.GET("/user", handler.ChatEndpoint)
	e.GET("/setup/key/:apiKey", handler.SetupKeyEndpoint)
	e.GET("/setup/model/:model", handler.SetupModelEndpoint)

	h2s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}

	s := http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(e, h2s),
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
