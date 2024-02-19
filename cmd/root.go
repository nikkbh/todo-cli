/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/nikkbh/todo-cli/controllers"
	"github.com/nikkbh/todo-cli/initializers"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A TODO CLI application that stores data to PostgreSQL DB.",
	Long: `TODO CLI is an easy to use TODO application for managing your day-to-day tasks.
			For example:
			todo --new "Pick Laundry"`,
	Run: todo,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	rootCmd.Flags().StringP("new", "n", "", "Add a New TODO item.")
	rootCmd.Flags().StringP("list", "l", "", "List the status of TODOS. 'all', 'pending', 'done'")
	rootCmd.Flags().IntP("done", "d", 0, "Mark the TODO item to Done")
	rootCmd.Flags().IntP("delete", "x", 0, "Delete the TODO item")
}

func todo(cmd *cobra.Command, args []string) {
	new, _ := cmd.Flags().GetString("new")
	list, _ := cmd.Flags().GetString("list")
	updateID, _ := cmd.Flags().GetInt("done")
	deleteID, _ := cmd.Flags().GetInt("delete")

	switch {
	case new != "":
		controllers.New(new)
	case list != "":
		controllers.List(list)
	case updateID != 0:
		controllers.Done(updateID)
	case deleteID != 0:
		controllers.Delete(deleteID)
	default:
		fmt.Println("Invalid flag argument")
	}
}
