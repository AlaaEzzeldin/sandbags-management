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
	"team2/sandsack-management-backend/middleware"
)

const defaultPort = ":8000"

func (a *App) RunAllRoutes() {
	var port = defaultPort

	r := gin.New()
	r.Use(middleware.CORSMiddleware())
	f, err := os.Create("gin.log")

	if err != nil {
		fmt.Println("file create error", err.Error())
	}

	log.SetOutput(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Swagger
	docs.SwaggerInfo.Title = "ASPD API Documentation"
	docs.SwaggerInfo.Description = "This page provides overview of all API endpoints and necessary details"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// unauthorized endpoints
	r.GET("/api-doc/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	notAuthUsers := r.Group("/users")
	notAuthUsers.POST("/login", a.Login)
	notAuthUsers.POST("/activation", a.VerifyEmail)
	notAuthUsers.POST("/forgot_password", a.SendRecoveryPassword)
	notAuthUsers.POST("/recovery_password", a.RecoveryPassword)
	notAuthUsers.POST("/refresh", a.RefreshAccessToken)

	// Admin endpoints
	admin := r.Group("/admin")
	admin.Use(a.AuthorizeAdmin())
	admin.POST("/create_user", a.CreateUser)
	admin.POST("/email_verification", a.SendVerifyEmail)
	admin.GET("/orders", a.AdminAllOrders)

	// Authentication and user profile endpoints
	auth := r.Group("/")
	auth.Use(AuthorizeJWT())

	users := auth.Group("users")
	users.GET("/", a.GetUserList)
	users.POST("/logout", a.Logout)
	users.PATCH("/change_password", a.ChangePassword)
	users.GET("/me", a.GetProfile)
	users.PATCH("/change_password", a.ChangePassword)
	users.PATCH("/me", a.PatchProfile)

	// order
	order := auth.Group("order")
	order.POST("/", a.CreateOrder)
	order.GET("/", a.ListOrder)
	order.POST("/cancel", a.DeclineOrder)
	order.POST("/accept", a.AcceptOrder)
	order.POST("/dispatch", a.DispatchOrder)
	order.POST("/comment", a.CommentOrder)
	order.PATCH("/edit", a.EditOrder)
	order.GET("/stats", a.GetStatistics)
	order.PATCH("/upgrade", func(context *gin.Context) {
		context.JSON(http.StatusNoContent, gin.H{
			"message": "in development",
		})
	})
	order.POST("/delivery/confirm", a.ConfirmDelivery)

	// core
	core := auth.Group("core")
	core.GET("/equipment", a.GetEquipment)
	core.GET("/priority", a.GetPriority)
	core.PATCH("/equipment/return", a.AddEquipmentQuantity)
	core.PATCH("/equipment/update", a.UpdateEquipment)
	core.POST("/priority/add", a.AddPriority)
	core.POST("/equipment/add", a.AddEquipment)
	core.POST("/driver/add", a.AddDriver)


	_ = r.Run(port)
}
