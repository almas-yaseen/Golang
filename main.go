package main

import (
	"fmt"
	"ginapp/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	route := gin.Default()
	database.ConnectDatabase()
	route.POST("/add", addUser)
	route.GET("/users", getAllUsers)
	route.DELETE("/users/:id", userDelete)

	route.Run(":8080")
}

func addUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad input"})
		return
	}

	response := gin.H{
		"message": "User successfully created",
	}
	ctx.JSON(http.StatusOK, response)

}

func getAllUsers(ctx *gin.Context) {
	// Perform a database query to fetch all users
	rows, err := database.Db.Query("SELECT username, password FROM users") // fetching all user details
	if err != nil {
		fmt.Println("Error querying database:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Username, &user.Password)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Return the users as a JSON response
	ctx.JSON(http.StatusOK, users)
}

func userDelete(ctx *gin.Context) {
	userIDParam := ctx.Param("id")

	// Convert the "id" parameter to an integer
	userID, err := strconv.Atoi(userIDParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	_, err = database.Db.Exec("DELETE FROM users where id=$1", userID)
	if err != nil {
		fmt.Println("Error delete user", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete users"})
		return

	}

	response := gin.H{
		"message": "User successfully deleteted",
	}
	ctx.JSON(http.StatusOK, response)
}
