/*
Copyright © 2025 Dani Medina <dani@medinag.me>
*/
package cmd

import (
	"log"
	"tri/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new TODO",
	Long:  `Add will create a new TODO item to the list`,
	Run:   addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: int in {1, 2, 3}")
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("dataFile"))
	if err != nil {
		log.Fatalln(err)
	}

	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}

	err = todo.SaveItems(viper.GetString("dataFile"), items)
	if err != nil {
		log.Fatalln(err)
	}
}
