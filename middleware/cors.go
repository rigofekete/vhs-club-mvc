package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Needed to bypass CORS (Cross-Origin Resource Sharing) security from the browser side, when sending requests through the React app
func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:5174",
		},
		AllowMethods: []string{"GET", "POST", "UPDATE", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	})
}
