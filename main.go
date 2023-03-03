package main

import (
	"fmt"
	"postcases/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectToPostgres()
	fmt.Println(("app is listening"))

	r.Run(":8080")
}
