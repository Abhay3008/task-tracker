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

// func main() {
// 	// fmt.Println(Addtask("test task"))
// 	// fmt.Println(Removetask(2))
// 	fmt.Println(Updatetask("updateInProgress", 3))
// 	// fmt.Println(data)
// }

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
	return 0, errors.New("Got some Internal Error while adding Task!!")
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
	if operation == "update" {
		for k, v := range tasklist.Todo {
			if id == v.Id {
				tasklist.Todo[k].Description = desc
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
		fmt.Printf("Unable to find task with id: %d in in-progress list!!", id)
		return errors.New("Unable to find task with id:" + strconv.Itoa(id) + "in in-progress list!!")
	} else if operation == "updateDone" {
		for k, v := range tasklist.Todo {
			if id == v.Id {
				temp := tasklist.Todo[k]
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

func Help() {
	fmt.Println("this is help!!")
}
