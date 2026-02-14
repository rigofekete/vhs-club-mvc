package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/handler"
	"github.com/rigofekete/vhs-club-mvc/internal/middleware"
	"github.com/rigofekete/vhs-club-mvc/repository"
	"github.com/rigofekete/vhs-club-mvc/service"
)

func main() {
	config.Load()
	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Dependency Injections

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	userHandler.RegisterRoutes(router)

	tapeRepository := repository.NewTapeRepository()
	tapeService := service.NewTapeService(tapeRepository)
	tapeHandler := handler.NewTapeHandler(tapeService)
	tapeHandler.RegisterRoutes(router)

	rentalRepository := repository.NewRentalRepository()
	rentalService := service.NewRentalService(rentalRepository)
	rentalHandler := handler.NewRentalHandler(rentalService)
	rentalHandler.RegisterRoutes(router)

	_ = router.Run(":8080")
}
