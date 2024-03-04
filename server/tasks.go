package main

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

func addTaskToList(text string, listname string) string {
	listIndex := getListIndex(listname)
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

	return "Task Added"
}

// Debug func only
func getListIndex(listname string) int {
	for index, value := range testlists {
		if value.Name == listname {
			return index
		}
	}
	// list not found
	return -1
}
