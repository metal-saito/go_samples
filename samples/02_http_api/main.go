package main

import (
	"log"

	"reservation_api/internal/handler"
	"reservation_api/internal/repository"
	"reservation_api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewMemoryRepository()
	svc := service.NewReservationService(repo)
	h := handler.NewReservationHandler(svc)

	r := gin.Default()

	// ミドルウェア
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// エラーハンドリングミドルウェア
	r.Use(errorHandler())

	// ルーティング
	api := r.Group("/api/v1")
	{
		api.POST("/reservations", h.CreateReservation)
		api.GET("/reservations", h.ListReservations)
		api.DELETE("/reservations/:id", h.CancelReservation)
	}

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
	}
}

