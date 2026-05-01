package controller

import (
	"encoding/json"
	"go-TaskTracker/internal/model"
	"go-TaskTracker/pkg/utils"
	"os"
	"time"
)

const filepath = "data/task.json"

func LoadTask() ([]model.Task, error) {

	//check if file doesn't exists in filepath, then return empty task
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return []model.Task{}, nil
	}

	//reading from file
	data, err := os.ReadFile(filepath)
	utils.CheckNilError(err)
	if len(data) == 0 {
		return []model.Task{}, nil
	}

	//unmarshal file and return if file is not empty
	var taskFromFile []model.Task
	utils.CheckNilError(err)
	err = json.Unmarshal(data, &taskFromFile)
	return taskFromFile, nil

}

// take task object and write on existing file
func SaveTask(task []model.Task) error {
	data, err := json.MarshalIndent(task, "", " ")
	utils.CheckNilError(err)

	return os.WriteFile(filepath, data, 0644)
}

// to add tasks into task array inside task file.
func AddTask(description string) model.Task {
	//get all exisiting tasks
	tasks, err := LoadTask()
	utils.CheckNilError(err)

	nextId := 1
	if len(tasks) > 0 {
		//get new task id
		nextId = tasks[len(tasks)-1].ID + 1
	}

	//create an object for new task.
	newTask := model.Task{
		ID:          nextId,
		Description: description,
		Status:      model.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	err = SaveTask(tasks)

	utils.CheckNilError(err)

	//save the tasks array
	return newTask

}
