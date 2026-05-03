package controller

import (
	"encoding/json"
	"fmt"
	"go-ExpenseTracker/model"
	"os"
	"time"
)

const filepath = "data/expenses.json"

func LoadExp() ([]model.Expense, error) {

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return []model.Expense{}, nil
	}

	data, error := os.ReadFile(filepath)
	if error != nil {
		fmt.Printf("Failed to read file: %v", error)
		return nil, error
	}

	if len(data) == 0 {
		return []model.Expense{}, nil
	}

	var ExpFromFile []model.Expense
	error = json.Unmarshal(data, &ExpFromFile)
	return ExpFromFile, error
}

func SaveExp(exp []model.Expense) error {
	data, err := json.MarshalIndent(exp, "", " ")
	if err != nil {
		fmt.Printf("Error Saving expense: %v", err)
		return err
	}

	return os.WriteFile(filepath, data, 0644)
}

func AddExpense(description string, amount int) model.Expense {
	fmt.Printf("Adding expense with description: '%s' and amount: %d\n", description, amount)
	exps, err := LoadExp()
	if err != nil {
		fmt.Printf("Error loading expenses: %s", err)
		return model.Expense{}
	}
	nextId := 1
	if len(exps) > 0 {
		nextId = exps[len(exps)-1].ID + 1
	}
	newExp := model.Expense{
		ID:          nextId,
		Description: description,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	exps = append(exps, newExp)
	err = SaveExp(exps)
	if err != nil {
		fmt.Printf("Error creating new task: %s", err)
	}
	return newExp
}
