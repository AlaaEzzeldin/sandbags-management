package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// todo
func (a *App) Logout(c *gin.Context) {
	header := c.Request.Header["Authorization"]
	token := strings.Split(header[0], " ")[1]

	fmt.Println(token)


}
