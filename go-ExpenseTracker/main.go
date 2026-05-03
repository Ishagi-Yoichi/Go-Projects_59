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

	case "update":
		if len(args) < 2 {
			fmt.Println("Usage: add 'Exp description and Amount'")
			return
		}
		id, _ := strconv.Atoi(args[1])
		amount, _ := strconv.Atoi(args[2])
		description := args[3]
		exp := controller.UpdateExpense(id, amount, description)
		fmt.Printf("Expense Update Successfully with Amount %v and Desc %v", exp.Amount, exp.Description)

	case "delete":
		if len(args) < 1 {
			fmt.Println("Usage: add 'Exp description and Amount'")
			return
		}
		id, _ := strconv.Atoi(args[1])
		_, msg := controller.DeleteExpense(id)
		if msg != "success" {
			panic("Expense doesn't exists! ")

		}
		fmt.Println("Deleted Successfully")
		return

	default:
		fmt.Print("Unkown Command:")
	}
}
