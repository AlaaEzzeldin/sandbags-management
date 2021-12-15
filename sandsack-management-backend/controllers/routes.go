package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"io"
	"log"
	"net/http"
	"os"
	"team2/sandsack-management-backend/docs"
	_ "team2/sandsack-management-backend/docs"
)

const defaultPort = ":8000"

func (a *App) RunAllRoutes(){
	var port = defaultPort

	r := gin.Default()
	f, err := os.Create("gin.log")


	if err != nil {
		fmt.Println("file create error", err.Error())
	}

	log.SetOutput(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Swagger
	docs.SwaggerInfo.Title = "ASPD API Documentation"
	docs.SwaggerInfo.Description = "This page provides overview of all API endpoints and necessary details"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// unauthorized endpoints
	r.GET("/api-doc/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Admin endpoints
	admin := r.Group("/admin")
	admin.Use(a.AuthorizeAdmin())
	admin.POST("/user", a.CreateUser)
	admin.POST("/user/email_verification", a.SendVerifyEmail)

	// Authentication and user profile endpoints
	auth := r.Group("/users")
	auth.Use(AuthorizeJWT())

	auth.GET("/", a.GetUserList)
	auth.POST("/login", a.Login)
	auth.POST("/activation", a.VerifyEmail)
	auth.POST("/forgot_password", a.SendRecoveryPassword)
	auth.POST("/recovery_password", a.RecoveryPassword)
	auth.POST("/refresh", a.RefreshAccessToken)
	auth.POST("/logout", a.Logout)
	auth.POST("/change_password", a.ChangePassword)
	auth.PATCH("/me", a.PatchProfile)

	//order
	auth.POST("/order", a.CreateOrder)
	auth.GET("/order", a.ListOrder)
	auth.POST("/order/cancel", a.DeclineOrder)
	auth.POST("/order/accept", a.AcceptOrder)

	auth.PATCH("/order/upgrade", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "in development",
		})
	})




	_ = r.Run(port)
}

