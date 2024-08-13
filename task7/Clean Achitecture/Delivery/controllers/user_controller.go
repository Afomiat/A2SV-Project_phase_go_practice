package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"task1.go/task7/task_manager/Domain"
	"task1.go/task7/task_manager/Infrastructure"
	"task1.go/task7/task_manager/Usecase"
)

type UserController struct {
	UserUsecase *Usecase.UserUsecase
}

func NewUserController(userUsecase *Usecase.UserUsecase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}

func (uc *UserController) RegisterUser(c *gin.Context) {

	var user Domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.UserUsecase.RegisterUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var user Domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserUsecase.LoginUser(user.Username, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := Infrastructure.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user logged in successfully", "Token": token})
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	fmt.Println("inhere*****************************")

	users, err := uc.UserUsecase.GetAllUsers()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) DeleteUserID(c *gin.Context) {
	var user Domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	Role_user := c.GetString("role")

	fmt.Println(" admin   ****************************************************", Role_user)

	if Role_user != "admin" {

		new_user, err := uc.UserUsecase.GetUserByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		user_id := c.GetString("user_id")
		newID, err := primitive.ObjectIDFromHex(string(user_id))

		fmt.Println(newID, user.ID, "DERTYUIO:****************************************************", newID)
		if new_user.ID != newID {
			fmt.Println("DERTYUIO:****************************************************")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
	}

	user, err := uc.UserUsecase.DeleteUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully", "user": user})
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := uc.UserUsecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
