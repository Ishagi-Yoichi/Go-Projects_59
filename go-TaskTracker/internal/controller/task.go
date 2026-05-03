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

func MarkTask(id int, status string) []model.Task {
	tasks, err := LoadTask()
	utils.CheckNilError(err)

	for i, task := range tasks {
		if task.ID == id {
			switch status {
			case "inprogress":
				tasks[i].Status = "in-progress"
				tasks[i].UpdatedAt = time.Now()

			case "done":
				tasks[i].Status = "done"
				tasks[i].UpdatedAt = time.Now()

			default:
				tasks[i].Status = task.Status
			}

			err := SaveTask(tasks)
			utils.CheckNilError(err)

			return tasks
		}
	}
	return []model.Task{}
}

// This function will take a status string and convert it into status model.
func ParseTaskStatus(s string) (model.TaskStatus, bool) {
	switch s {
	case string(model.StatusTodo):
		return model.StatusTodo, true
	case string(model.StatusInProgress):
		return model.StatusInProgress, true
	case string(model.StatusDone):
		return model.StatusDone, true
	default:
		return "", false
	}
}

func UpdateTask(id int, description string) model.Task {
	tasks, err := LoadTask()
	utils.CheckNilError(err)

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()

			err := SaveTask(tasks)
			utils.CheckNilError(err)

			return tasks[i]
		}
	}

	return model.Task{}
}

func DeleteTask(id int) ([]model.Task, string) {
	tasks, err := LoadTask()
	utils.CheckNilError(err)

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			err := SaveTask(tasks)
			utils.CheckNilError(err)

			return tasks, "success"
		}
	}

	return tasks, "Task not found"
}
