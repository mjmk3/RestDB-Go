package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

type User struct {
	ID   string `json: "id"`
	Name string `json: "name"`
	Age  int    `json: "age"`
}

var Users []User // if we use it without {} will return null but this will return []

func main() {
	r := gin.Default()

	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/", getUsers)
		userRoutes.POST("/add", CreateUser)
		userRoutes.PUT("/:id", UpdateUser)
		userRoutes.DELETE("/:id", RemoveUser)
	}

	if err := r.Run(":5500"); err != nil {
		log.Fatal(err.Error())
	}
}

func getUsers(c *gin.Context) {
	c.JSON(200, Users)
}

func CreateUser(c *gin.Context) {
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	reqBody.ID = uuid.New().String()
	Users = append(Users, reqBody)

	c.JSON(200, gin.H{
		"error": false,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	for i, u := range Users {
		if u.ID == id {
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}

func RemoveUser(c *gin.Context) {
	id := c.Param("id")

	for i, u := range Users {
		if u.ID == id {
			Users = append(Users[:i], Users[i+1:]...)

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}
