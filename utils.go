package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const file = "./tasks.json"

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type Tasklist struct {
	Todo       []Task `json:"todo"`
	Inprogress []Task `json:"inprogress"`
	Done       []Task `json:"done"`
	TotalId    int    `json:"totalid"`
}

func Loadjson() Tasklist {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return Tasklist{
			Todo:       []Task{},
			Inprogress: []Task{},
			Done:       []Task{},
			TotalId:    0,
		}
	} else {
		readfile, _ := os.ReadFile(file)
		var list Tasklist
		json.Unmarshal(readfile, &list)
		return list
	}
}

func Addtask(description string) (int, error) {
	tasklist := Loadjson()
	currtime := time.Now().Format("2006-01-02 15:04:05")
	newtask := Task{
		Id:          tasklist.TotalId + 1,
		Description: description,
		CreatedAt:   currtime,
		UpdatedAt:   currtime,
	}
	tasklist.Todo = append(tasklist.Todo, newtask)
	tasklist.TotalId++
	if Savejson(tasklist) {
		return newtask.Id, nil
	}
	return 0, errors.New("got some internal error while adding task")
}

func Removetask(id int) error {
	tasklist := Loadjson()

	for k, v := range tasklist.Todo {
		if id == v.Id {
			tasklist.Todo = append(tasklist.Todo[:k], tasklist.Todo[k+1:]...)
			if Savejson(tasklist) {
				return nil
			} else {
				return errors.New("internal error occured while updating json file")
			}
		}
	}
	for k, v := range tasklist.Inprogress {
		if id == v.Id {
			tasklist.Inprogress = append(tasklist.Inprogress[:k], tasklist.Inprogress[k+1:]...)
			if Savejson(tasklist) {
				return nil
			} else {
				return errors.New("internal error occured while updating json file")
			}
		}
	}
	for k, v := range tasklist.Done {
		if id == v.Id {
			tasklist.Done = append(tasklist.Done[:k], tasklist.Done[k+1:]...)
			if Savejson(tasklist) {
				return nil
			} else {
				return errors.New("internal error occured while updating json file")
			}
		}
	}
	return errors.New("Unable to find task with id:" + strconv.Itoa(id) + "in task list!!")
}

func Updatetask(operation string, id int, description ...string) error {
	tasklist := Loadjson()
	desc := ""
	if len(description) > 0 {
		desc = description[0]
	}
	currtime := time.Now().Format("2006-01-02 15:04:05")
	if operation == "update" {
		for k, v := range tasklist.Todo {
			if id == v.Id {
				tasklist.Todo[k].Description = desc
				tasklist.Todo[k].UpdatedAt = currtime
				if Savejson(tasklist) {
					return nil
				} else {
					return errors.New("internal error occured while updating json file")
				}
			}
		}
		for k, v := range tasklist.Inprogress {
			if id == v.Id {
				tasklist.Inprogress[k].Description = desc
				tasklist.Inprogress[k].UpdatedAt = currtime
				if Savejson(tasklist) {
					return nil
				} else {
					return errors.New("internal error occured while updating json file")
				}
			}
		}
		for k, v := range tasklist.Done {
			if id == v.Id {
				tasklist.Done[k].Description = desc
				tasklist.Done[k].UpdatedAt = currtime
				if Savejson(tasklist) {
					return nil
				} else {
					return errors.New("internal error occured while updating json file")
				}
			}
		}
	} else if operation == "updateInProgress" {
		for k, v := range tasklist.Todo {
			if id == v.Id {
				temp := tasklist.Todo[k]
				temp.UpdatedAt = currtime
				Removetask(id)
				tasklist.Todo = append(tasklist.Todo[:k], tasklist.Todo[k+1:]...)
				tasklist.Inprogress = append(tasklist.Inprogress, temp)
				if Savejson(tasklist) {
					return nil
				} else {
					return errors.New("internal error occured while updating json file")
				}
			}
		}
		fmt.Printf("Unable to find task with id: %d in in-progress list!!\n", id)
		return errors.New("Unable to find task with id:" + strconv.Itoa(id) + "in in-progress list!!")
	} else if operation == "updateDone" {
		for k, v := range tasklist.Todo {
			if id == v.Id {
				temp := tasklist.Todo[k]
				temp.UpdatedAt = currtime
				Removetask(id)
				tasklist.Todo = append(tasklist.Todo[:k], tasklist.Todo[k+1:]...)
				tasklist.Done = append(tasklist.Done, temp)
				if Savejson(tasklist) {
					return nil
				} else {
					return errors.New("internal error occured while updating task")
				}
			}
		}
		for k, v := range tasklist.Inprogress {
			if id == v.Id {
				temp := tasklist.Inprogress[k]
				temp.UpdatedAt = currtime
				Removetask(id)
				tasklist.Inprogress = append(tasklist.Inprogress[:k], tasklist.Inprogress[k+1:]...)
				tasklist.Done = append(tasklist.Done, temp)
				if Savejson(tasklist) {
					return nil
				} else {
					return errors.New("internal error occured while updating task")
				}
			}
		}
		for _, v := range tasklist.Done {
			if id == v.Id {
				fmt.Println("Task already marked as Done!!")
				if Savejson(tasklist) {
					return nil
				} else {
					return errors.New("internal error occured while updating task")
				}
			}
		}

	}
	return errors.New("Invalid operation:" + operation)
}

func ShowAllTasks() {
	tasklist := Loadjson()
	PrintMajorseparator()
	fmt.Println("To Do Tasks")
	PrintMajorseparator()
	if len(tasklist.Todo) < 1 {
		fmt.Println("No Task for now, Take some Rest!!")
	}
	for k, v := range tasklist.Todo {
		fmt.Printf(" ID: %v\n Description: %v\n", v.Id, v.Description)
		if k != len(tasklist.Todo)-1 {
			PrintMinorseparator()
		}
	}
	fmt.Println()

	PrintMajorseparator()
	fmt.Println("In Progress Tasks")
	PrintMajorseparator()
	if len(tasklist.Todo) < 1 {
		fmt.Println("No Task for now, Take some Rest!!")
	}
	for k, v := range tasklist.Inprogress {
		fmt.Printf(" ID: %v\n Description: %v\n", v.Id, v.Description)
		if k != len(tasklist.Inprogress)-1 {
			PrintMinorseparator()
		}
	}
	fmt.Println()

	PrintMajorseparator()
	fmt.Println("Completed Tasks")
	PrintMajorseparator()
	if len(tasklist.Done) < 1 {
		fmt.Println("No Task found!!")
	}
	for k, v := range tasklist.Done {
		fmt.Printf(" ID: %v\n Description: %v\n", v.Id, v.Description)
		if k != len(tasklist.Done)-1 {
			PrintMinorseparator()
		}
	}

}

func ShowToDoTask() {
	tasklist := Loadjson()
	PrintMajorseparator()
	fmt.Println("To Do Tasks")
	PrintMajorseparator()
	if len(tasklist.Todo) < 1 {
		fmt.Println("No Task for now, Take some Rest!!")
	}
	for k, v := range tasklist.Todo {
		fmt.Printf(" ID: %v\n Description: %v\n", v.Id, v.Description)
		if k != len(tasklist.Todo)-1 {
			PrintMinorseparator()
		}
	}
}

func ShowInProgressTask() {
	tasklist := Loadjson()
	PrintMajorseparator()
	fmt.Println("In Progress Tasks")
	PrintMajorseparator()
	if len(tasklist.Todo) < 1 {
		fmt.Println("No Task for now, Take some Rest!!")
	}
	for k, v := range tasklist.Inprogress {
		fmt.Printf(" ID: %v\n Description: %v\n", v.Id, v.Description)
		if k != len(tasklist.Inprogress)-1 {
			PrintMinorseparator()
		}
	}
	fmt.Println()
}

func ShowDoneTasks() {
	tasklist := Loadjson()
	PrintMajorseparator()
	fmt.Println("Completed Tasks")
	PrintMajorseparator()
	if len(tasklist.Done) < 1 {
		fmt.Println("No Task found!!")
	}
	for k, v := range tasklist.Done {
		fmt.Printf(" ID: %v\n Description: %v\n", v.Id, v.Description)
		if k != len(tasklist.Done)-1 {
			PrintMinorseparator()
		}
	}
}

func Savejson(tasklist Tasklist) bool {

	newfile, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
		return false
	}
	res, err := json.MarshalIndent(tasklist, "", "\t")
	if err != nil {
		log.Fatal(err)
		return false
	}
	io.Writer.Write(newfile, res)
	newfile.Close()
	return true
}

func Error(msg string) {
	fmt.Printf("**Error: %v\n", msg)
}

func PrintMajorseparator() {
	fmt.Println("---------------")
}

func PrintMinorseparator() {
	fmt.Println("-------")
}

func Help() {
	var helpText = `
Usage:
  task-tracker add "Task description"
    Adds a new task. Example:
    $ task-tracker add "Buy groceries"

  task-tracker update <id> "New description"
  task-tracker delete <id>
  task-tracker mark-in-progress <id>
  task-tracker mark-done <id>
  task-tracker list
  task-tracker list <status>
    (todo, in-progress, done)

  task-tracker help
    Displays this help message.
`
	fmt.Println(helpText)
}
