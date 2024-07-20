package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

func main() {
	router := echo.New()

	router.GET("/health-check", healthCheck)

	port := getEnv("PORT", "8000")

	if err := router.Start(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("server is not running %s", err)
	}
}

func healthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"message":  "--- OK ---",
		"message2": "--- OK2 ___",
	})
}

func getEnv(envVar, fallback string) string {
	if v := os.Getenv(envVar); v != "" {
		return v
	}

	return fallback
}
