package main

import (
	"github.com/nikkbh/todo-cli/initializers"
	"github.com/nikkbh/todo-cli/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Todo{})
}
