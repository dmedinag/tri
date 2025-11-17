/*
Copyright © 2025 Dani Medina <dani@medinag.me>
*/
package cmd

import (
	"log"
	"sort"
	"strconv"
	"tri/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark item as done",
	Run:     doneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("dataFile"))
	if err != nil {
		log.Fatalln(err)
	}
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}
	if i > 0 && i <= len(items) {
		items[i-1].Done = true
		log.Printf("%q %v\n", items[i-1].Text, "marked done")
		sort.Sort(todo.ByPri(items))
		err := todo.SaveItems(viper.GetString("dataFile"), items)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Println(i, "doesn't match any items")
	}
}
