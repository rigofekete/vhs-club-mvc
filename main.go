package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rigofekete/vhs-club-mvc/handler"
	"github.com/rigofekete/vhs-club-mvc/repository"
	"github.com/rigofekete/vhs-club-mvc/service"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Dependency Injections
	tapeRepository := repository.NewTapeRepository()
	tapeService := service.NewTapeService(tapeRepository)
	tapeHandler := handler.NewTapeHandler(tapeService)
	tapeHandler.RegisterRoutes(router)

	_ = router.Run(":8080")
}
