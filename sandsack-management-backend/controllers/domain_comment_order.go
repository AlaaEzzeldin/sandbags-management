package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

func (a *App) CommentOrder(c *gin.Context) {
	var input models.CommentInput

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil{
		log.Println("CommentOrder error: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "incorrect request",
		})
		return
	}

	claims, err := GetClaims(c)
	if err != nil {
		log.Println("GetClaims error", err.Error())
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			ErrCode: http.StatusUnauthorized,
			ErrMessage: "no access",
		})
		return
	}

	permissions, err := service.GetUserOrderPermissions(a.DB, claims.Id, input.OrderId)

	flag := 0
	for _, i := range permissions {
		if i == models.DictPermissionName["CAN COMMENT"]{
			flag = 1
			break
		}
	}

	if flag == 0 {
		log.Println("No access for this action error")
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			ErrCode: http.StatusForbidden,
			ErrMessage: "no access",
		})
		return
	}


	var comments []models.Comment
	for _, row := range input.Comments {
		comment := models.Comment{
			CommentText: row,
		}
		comments = append(comments, comment)
	}


	if len(comments) == 0 {
		log.Println("Length of comments is 0", len(comments))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode: http.StatusBadRequest,
			ErrMessage: "there should be comments",
		})
		return
	}


	if err := service.InsertComments(a.DB, claims.Id, input.OrderId, comments); err != nil {
		return
	}

	c.JSON(http.StatusOK, input)
	return
}
