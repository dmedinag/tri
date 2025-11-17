/*
Copyright © 2025 Dani Medina <dani@medinag.me>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"

	"tri/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	doneOpt    bool
	showAllOpt bool
	targetPrio []string
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List TODOs",
	Long:  `Listing all TODOs, formatted and following the given filters`,
	Run:   listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&doneOpt, "done", "d", false, "Show done TODOs")
	listCmd.Flags().BoolVarP(&showAllOpt, "all", "a", false, "Show all TODOs")
	listCmd.Flags().StringArrayVarP(&targetPrio, "priority", "p", []string{}, "Show TODOs with matching priority")
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("dataFile"))
	if err != nil {
		log.Fatalf("Could not read items: %v\n", err)
	}

	if len(items) == 0 {
		log.Println("No TODOs in database")
	}

	sort.Sort(todo.ByPri(items))

	hiddenItems := 0

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		if (showAllOpt || i.Done == doneOpt) && (len(targetPrio) == 0 || Contains(targetPrio, strconv.Itoa(i.Priority))) {
			if _, err := fmt.Fprintln(w, i.Label()+"\t"+i.PrettyDone()+"\t"+i.PrettyP()+"\t"+i.Text+"\t"); err != nil {
				log.Fatalln("Error printing items\n", err)
			}
		} else {
			hiddenItems += 1
		}
	}

	if len(items) == hiddenItems {
		log.Printf("No TODOs to show (%v hidden due to filters)\n", hiddenItems)
	}

	if err := w.Flush(); err != nil {
		log.Fatalf("Could not read items: %v\n", err)
	}
}

func Contains[T comparable](s []T, item T) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}
