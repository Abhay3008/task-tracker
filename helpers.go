package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func Add(args []string) {
	if len(args) > 1 {
		Error("Too many arguements for add")
		Help()
		os.Exit(1)
	} else if len(args) < 1 {
		Error("add requires an arguement")
		Help()
		os.Exit(1)
	}
	id, err := Addtask(args[0])
	if err != nil {
		log.Println(err)
		Error("Unable to add task to the to-do list, Please retry!!")
		os.Exit(1)
	}
	fmt.Printf("Task added successfully (ID: %d)\n", id)

}

func Update(args []string) {
	if len(args) > 2 {
		Error("Too many arguements for update")
		Help()
		os.Exit(1)
	} else if len(args) < 2 {
		Error("update requires 2 arguement")
		Help()
		os.Exit(1)
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println(err)
		Error("Got error while converting id, Please retry with appropriate arguements.")
		Help()
		os.Exit(1)
	}
	err = Updatetask("update", id, args[1])
	if err != nil {
		log.Println(err)
		Error("Unable to update Task, Please retry!")
	}
	fmt.Printf("Task updated successfully (ID: %d)\n", id)
}

func Delete(args []string) {
	if len(args) > 1 {
		Error("Too many arguements for update")
		Help()
		os.Exit(1)
	} else if len(args) < 1 {
		Error("update requires an arguement")
		Help()
		os.Exit(1)
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println(err)
		Error("Got error while converting id, Please retry with appropriate arguements.")
		Help()
		os.Exit(1)
	}
	err = Removetask(id)
	if err != nil {
		log.Println(err)
		Error("Unable to remove Task, Please retry!")
	}
	fmt.Printf("Task removed successfully (ID: %d)\n", id)

}

func MarkInProgress(args []string) {
	if len(args) > 1 {
		Error("Too many arguements for mark-in-progress")
		Help()
		os.Exit(1)
	} else if len(args) < 1 {
		Error("mark-in-progress requires an arguement")
		Help()
		os.Exit(1)
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println(err)
		Error("Got error while converting id, Please retry with appropriate arguements.")
		Help()
		os.Exit(1)
	}
	err = Updatetask("updateInProgress", id)
	if err != nil {
		log.Println(err)
		Error("Unable to mark-in-progress Task, Please retry!")
	}
	fmt.Printf("Task moved to in-progress (ID: %d)\n", id)
}

func MarkDone(args []string) {
	if len(args) > 1 {
		Error("Too many arguements for mark-done")
		Help()
		os.Exit(1)
	} else if len(args) < 1 {
		Error("mark-done requires an arguement")
		Help()
		os.Exit(1)
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println(err)
		Error("Got error while converting id, Please retry with appropriate arguements.")
		Help()
		os.Exit(1)
	}
	err = Updatetask("updateDone", id)
	if err != nil {
		log.Println(err)
		Error("Unable to mark-done Task, Please retry!")
	}
	fmt.Printf("Task moved to Done (ID: %d)\n", id)
}

func List(args []string) {
	if len(args) > 1 {
		Error("Too many arguements for command list")
		Help()
		os.Exit(1)
	}
	if len(args) == 0 {
		ShowAllTasks()
		os.Exit(0)
	}

	switch args[0] {
	case "todo":
		ShowToDoTask()
	case "in-progress":
		ShowInProgressTask()
	case "done":
		ShowDoneTasks()
	default:
		fmt.Printf("Invalid option: %v for command list\n", args[0])
		Help()
	}
}
