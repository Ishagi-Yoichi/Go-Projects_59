package controller

import (
	"encoding/json"
	"go-TaskTracker/internal/model"
	"go-TaskTracker/pkg/utils"
	"os"
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
