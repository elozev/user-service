package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)


func defineUserRoutes(rg *gin.RouterGroup) {
	usersGroup := rg.Group("/users")

	usersGroup.GET("/", getUsers)
	usersGroup.GET("/:id", getUserById)
	usersGroup.PUT("/:id", updateUserById)
	usersGroup.POST("/", createUser)
	usersGroup.DELETE("/:id", deleteUser)
}

func getUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"users": "me",
	})
}

func getUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"user_with_id": id,
	})
}

func updateUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"updating_user_with_id": id,
	})
}

func createUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	fmt.Println("POST /users/")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to read body",
		})
		}

			ctx.JSON(http.StatusOK, gin.H{
				"message": "Body read successfully",
				"body": string(body),
			})
}

func deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprint("deleted user ", id),
	})
}