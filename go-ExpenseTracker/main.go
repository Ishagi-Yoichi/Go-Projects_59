package main

import (
	"fmt"
	"go-ExpenseTracker/controller"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	fmt.Printf("Received arguments: %+v\n", args)
	if len(args) == 0 {
		fmt.Println("No command passed !")
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			fmt.Println("Usage: add 'Exp description and Amount'")
			return
		}

		amount, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Invalid amount: %v", err)
			return
		}
		description := args[2]
		exp := controller.AddExpense(description, amount)
		fmt.Printf("Expense %v Added Succesfully of Amount %v", exp.ID, exp.Amount)

	}
}
