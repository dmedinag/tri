/*
Copyright © 2025 Dani Medina <dani@medinag.me>
*/
package cmd

import (
	"errors"
	"log"
	"strconv"

	"tri/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	targetId int
	newName  string
	newPri   int
	items    []todo.Item
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a TODO",
	Run:   runEdit,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		items, err = todo.ReadItems(viper.GetString("dataFile"))
		if err != nil {
			return err
		}
		if targetId < 0 || targetId > len(items) {
			return errors.New("invalid id: " + strconv.Itoa(targetId) + " greater than max index " + strconv.Itoa(len(items)))
		}
		if cmd.Flags().Changed("name") && newName == "" {
			return errors.New("the new content for a TODO cannot be empty")
		}
		if cmd.Flags().Changed("priority") && (newPri < 1 || newPri > 3) {
			return errors.New("the new priority for a TODO needs to be within {1, 2, 3}")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().IntVar(&targetId, "id", 0, "id of the TODO to be edited")
	if err := editCmd.MarkFlagRequired("id"); err != nil {
		panic(err)
	}

	editCmd.Flags().StringVar(&newName, "name", "", "updated text for the TODO")
	editCmd.Flags().IntVarP(&newPri, "priority", "p", -1, "new priority for the T panic(err)")

	editCmd.MarkFlagsOneRequired("name", "priority")
}

func runEdit(cmd *cobra.Command, args []string) {
	targetItem := &items[targetId-1]
	if cmd.Flags().Changed("name") {
		prev := targetItem.Text
		targetItem.Text = newName
		log.Printf("Content of TODO#%v changed from %q to %q\n", targetId, prev, newName)
	}
	if cmd.Flags().Changed("priority") {
		prev := targetItem.Priority
		targetItem.Priority = newPri
		log.Printf("Priority of TODO#%v changed from '%v' to '%v'\n", targetId, prev, newPri)
	}

	err := todo.SaveItems(viper.GetString("dataFile"), items)
	if err != nil {
		log.Fatalln(err)
	}
}
