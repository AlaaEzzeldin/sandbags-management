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

	r.GET("/api-doc/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/users/login", a.Login)

	auth := r.Group("/users")
	auth.Use(AuthorizeJWT())
	auth.POST("/refresh", a.RefreshAccessToken)
	auth.POST("/", a.GetUserList)
	auth.POST("/create_user", a.CreateUser)
	auth.POST("/me", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "in development",
		})
	})
	auth.POST("/logout", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "in development",
		})
	})

	r.POST("users/forgot_password", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "in development",
		})
	})
	auth.POST("/change_password", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "in development",
		})
	})
	r.POST("/users/verify", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "in development",
		})
	})
	r.POST("/users/activation", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "in development",
		})
	})
	r.POST("/users/email_verification", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "in development",
		})
	})

	r.POST("/order", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "in development",
		})
	})

	r.POST("/order/upgrade", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "in development",
		})
	})
	r.POST("/order/cancel", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "in development",
		})
	})

	//authorized := r.Use(AuthorizeJWT())
	//authorized.POST("/registration", a.Registration)

	_ = r.Run(port)
}

