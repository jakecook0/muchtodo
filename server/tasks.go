package main

import (
	"fmt"
)

var lists = make([]List, 0)

var testlists = []List{
	{
		Name: "Test Tasks",
		TaskItems: []Task{
			{Value: "test one", ID: 0, Done: false},
			{Value: "test two", ID: 1, Done: false},
			// Task{Value: "test three", ID: 2, Done: true},
		},
		NumTasks: 2,
	},
	{
		Name: "Demo Tasks",
		TaskItems: []Task{
			{Value: "demo 1", ID: 0, Done: false},
			{Value: "demo 2", ID: 1, Done: false},
			{Value: "demo 3", ID: 2, Done: true},
		},
		NumTasks: 3,
	},
}

type Task struct {
	Value string
	ID    int
	Done  bool
}

type List struct {
	Name      string
	TaskItems []Task
	NumTasks  int
}

func getLists() []List {
	// return just the names of each task as json string

	return testlists
}

func addTaskToList(listname string, text string) error {
	listIndex, err := getListIndex(listname)
	if err != nil {
		println("Failed to update list", err)
		return err
	}

	currNumTasks := testlists[listIndex].NumTasks

	newTask := Task{
		Value: text,
		ID:    currNumTasks + 1,
		Done:  false,
	}

	listTasks := testlists[listIndex].TaskItems
	// update task list
	listTasks = append(listTasks, newTask)
	// Insert updated tasks and task numbers back into parent list
	testlists[listIndex].TaskItems = listTasks
	testlists[listIndex].NumTasks = currNumTasks + 1

	return nil
}

func createList(listname string) error {
	lists = append(lists, List{Name: listname, TaskItems: make([]Task, 0), NumTasks: 0})
	return nil
}

// Debug func only
func getListIndex(listname string) (int, error) {
	for index, value := range testlists {
		if value.Name == listname {
			return index, nil
		}
	}
	// list not found
	return -1, fmt.Errorf("List not found with name: %s", listname)
}
