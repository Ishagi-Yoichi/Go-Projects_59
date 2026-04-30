package main

import (
	"fmt"
	"os"
)

type Tasks struct{
	status string
	name string
}


func main(){
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No command passed !")
		return
	}

	switch args[0]{
		case "add":
			if len(args)>2{
				fmt.Println("Usage: task-cli add 'Task description'")
				return
			}
		    description := args[1]
			task :=

	}

}
