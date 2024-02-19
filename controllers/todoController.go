package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nikkbh/todo-cli/initializers"
	"github.com/nikkbh/todo-cli/models"
)

func New(item string) {
	task := models.Todo{Task: item, Status: models.Pending}
	result := initializers.DB.Create(&task)

	if result.Error != nil {
		log.Fatal("Error adding a new task.")
		return
	}
	fmt.Println("TODO task created: ", task.ID)
}

func List(status string) {
	var todos []models.Todo
	switch status {
	case "all":
		result := initializers.DB.Find(&todos)
		if len(todos) == 0 {
			log.Println("No records found.")
		} else if result.Error != nil {
			log.Println("Error executing query: ", result.Error)
		} else {
			// Process the todos
			prettyJSON, err := json.MarshalIndent(todos, "", "    ")
			if err != nil {
				log.Fatal(err)
			}
			// Print the pretty JSON
			fmt.Println(string(prettyJSON))
		}
	case string(models.Pending):
		result := initializers.DB.Where("status = ?", "pending").Find(&todos)
		if len(todos) == 0 {
			log.Println("No records found.")
		} else if result.Error != nil {
			log.Println("Error executing query: ", result.Error)
		} else {
			// Process the todos
			prettyJSON, err := json.MarshalIndent(todos, "", "    ")
			if err != nil {
				log.Fatal(err)
			}
			// Print the pretty JSON
			fmt.Println(string(prettyJSON))
		}
	case string(models.Done):
		result := initializers.DB.Where("status = ?", "done").Find(&todos)
		if len(todos) == 0 {
			log.Println("No records found.")
		} else if result.Error != nil {
			log.Println("Error executing query: ", result.Error)
		} else {
			// Process the todos
			prettyJSON, err := json.MarshalIndent(todos, "", "    ")
			if err != nil {
				log.Fatal(err)
			}
			// Print the pretty JSON
			fmt.Println(string(prettyJSON))
		}
	}
}

func Done(id int) {
	var todo models.Todo
	initializers.DB.First(&todo, id)
	result := initializers.DB.Model(&todo).Updates(models.Todo{Status: models.Done})

	if result.Error != nil {
		log.Fatal("Error updating the status of the task.")
		return
	}
	prettyJSON, err := json.MarshalIndent(todo, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	// Print the pretty JSON
	fmt.Println(string(prettyJSON))
}

func Delete(id int) {
	result := initializers.DB.Delete(&models.Todo{}, id)

	if result.Error != nil {
		log.Fatal("Error deleteing the TODO item.")
		return
	}
	fmt.Println("Deleted TODO item with ID:", id)
}
