package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping() func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	}
}

func Health() func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	}
}

func Greet(msg string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, msg)
	}
}
