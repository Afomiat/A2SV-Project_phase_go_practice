package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/middleware"
	"task_manager/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var jwtSecret = []byte("your_secret_key")

func Register(c *gin.Context) {
	var creds models.Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user := models.User{
		Username: creds.Username,
		Password: string(hashedPassword),
		Role:     "user", // Default role
	}

	if err := data.AddUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}

func Login(c *gin.Context) {
	var creds models.Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := data.GetUserByUsername(creds.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.Claims{
		Username: user.Username,
		Role:     user.Role,
		// StandardClaims: jwt.StandardClaims{
		// 	ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		// },
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// GetTasks retrieves tasks for the logged-in user or admin
func GetTasks(c *gin.Context) {
	role, _ := c.Get("role")
	username, _ := c.Get("username")

	var filter bson.M
	if role == "user" {
		filter = bson.M{"username": username}
	} else {
		filter = bson.M{}
	}

	tasks, err := data.GetTasks(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTaskById retrieves a task by its ID
func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	role, _ := c.Get("role")
	username, _ := c.Get("username")

	task, err := data.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	if role == "user" && task.Username != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient permissions"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func PostTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	username, _ := c.Get("username")
	task.Username = username.(string)
	task.ID = primitive.NewObjectID()
	// task.Status = "pending"

	if err := data.AddTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add task"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "task added successfully"})
}

// UpdateTask updates a task by its ID
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	existingTask, err := data.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	role, _ := c.Get("role")

	username, _ := c.Get("username")

	if role == "user" && existingTask.Username != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient permissions"})
		return
	}
	
	task.Username = username.(string)

	if err := data.UpdateTask(id, task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task updated successfully"})
}

// DeleteTask removes a task by its ID
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	role, _ := c.Get("role")
	username, _ := c.Get("username")

	task, err := data.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	if role == "user" && task.Username != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient permissions"})
		return
	}

	if err := data.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}
