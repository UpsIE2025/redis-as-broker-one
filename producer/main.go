package main

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const channel = "message_queue"

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	defer rdb.Close()

	r := gin.Default()

	r.POST("/publish", func(ctx *gin.Context) {
		var req struct {
			Message string `json:"message"`
		}
		if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		err := rdb.Publish(ctx, channel, req.Message).Err()
		if err != nil {
			slog.Error(err.Error())
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error de conexión"})
			return
		}
		ctx.Status(http.StatusOK)
	})

	r.Run(":80")
}
