package todo_test

import (
	"dmedinag/tri/todo"
	"fmt"
	"os"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabel(t *testing.T) {
	item := &todo.Item{}
	expected := "0. "

	assert.Equal(t, item.Label(), expected)
}

func TestPrettyDone(t *testing.T) {
	item := &todo.Item{Done: false}
	expected := ""
	assert.Equal(t, item.PrettyDone(), expected)
}

func TestPrettyUnDone(t *testing.T) {
	item := &todo.Item{Done: true}
	expected := "X"
	assert.Equal(t, item.PrettyDone(), expected)
}

func TestPrettyP(t *testing.T) {
	printedPriorities := []int{1, 3}
	item := &todo.Item{}
	for priority := range 3 {
		item.SetPriority(priority)
		var expected string
		if slices.Contains(printedPriorities, priority) {
			expected = fmt.Sprintf("(%d)", priority)
		} else {
			expected = " "
		}

		assert.Equal(t, item.PrettyP(), expected, "Pretty P didn't print what it should")
	}
}

func TestIO(t *testing.T) {
	filename := "test_db.json"

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	assert.Nil(t, err, "Error opening file for testing")
	file.Close()

	items := []todo.Item{
		{
			Text:     "least important",
			Priority: 3,
			Done:     false,
		},
		{
			Text:     "why did we bother adding this",
			Priority: 3,
			Done:     true,
		},
		{
			Text:     "truly critical",
			Priority: 1,
			Done:     false,
		},
		{
			Text:     "good to have",
			Priority: 2,
			Done:     false,
		},
		{
			Text:     "good that we have it",
			Priority: 2,
			Done:     true,
		},
	}
	err = todo.SaveItems(filename, items)
	if err != nil {
		t.Error("Error saving items:", err)
	}
	retrievedItems, err := todo.ReadItems(filename)
	if err != nil {
		t.Error("Error retrieving items:", err)
	}
	assert.EqualExportedValues(t, items, retrievedItems, "Saved and retrieved items differ")
	err = os.Remove(filename)
	assert.Nilf(t, err, "Couldn't clean up test by deleting %q", filename)
}
