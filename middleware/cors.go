package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// TODO: Remove file since we are not using this bypass through the backend any longer
// Needed to bypass CORS (Cross-Origin Resource Sharing) security from the browser side, when sending requests through the React app
func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:5174",
		},
		AllowMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	})
}
