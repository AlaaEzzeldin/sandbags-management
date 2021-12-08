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
	r.POST("/users/activation", a.VerifyEmail)
	r.POST("/users/forgot_password", a.SendRecoveryPassword)
	r.POST("/users/recovery_password", a.RecoveryPassword)
	r.POST("users/refresh", a.RefreshAccessToken)

	admin := r.Group("/admin")
	admin.Use(a.AuthorizeAdmin())
	admin.POST("/user", a.CreateUser)
	admin.POST("/user/email_verification", a.SendVerifyEmail)

	auth := r.Group("/users")
	auth.Use(AuthorizeJWT())


	auth.GET("/", a.GetUserList)
	auth.PUT("/password", a.ChangePassword)
	auth.POST("/me", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "in development",
		})
	})
	auth.POST("/logout", a.Logout)

	auth.POST("/change_password", func(context *gin.Context) {
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

