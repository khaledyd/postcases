package app

import (
	"github.com/gin-gonic/gin"
	"github.com/khaledyd/postcases/internal/app/handlers"
	"github.com/khaledyd/postcases/internal/middleware"
)

func SetupRoutes(r *gin.Engine) {
	// define routes
	r.POST("/signup", handlers.SignUp)
	r.POST("/signin", middleware.AuthMiddleware(), handlers.SignIn)
	// r.GET("/users/:id", GetUser)
	// r.PUT("/users/:id", UpdateUser)
	// r.DELETE("/users/:id", DeleteUser)
}
