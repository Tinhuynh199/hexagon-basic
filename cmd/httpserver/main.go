package main

import (
	"github.com/gin-gonic/gin"
	"hexagonal/internal/core/service/gamesrv"
	"hexagonal/internal/handlers/gamehdl"
	"hexagonal/internal/repositories/gamesrepo"
	"hexagonal/pkg/uidgen"
)

func main() {
	gamesRepository := gamesrepo.NewMemKVS()
	gamesService := gamesrv.New(gamesRepository, uidgen.New())
	gamesHandler := gamehdl.NewHTTPHandler(gamesService)

	router := gin.New()
	router.GET("/games/:id", gamesHandler.Get)
	router.POST("/games", gamesHandler.Create)
	router.PUT("/games/:id", gamesHandler.RevealCell)

	router.Run(":8080")
}