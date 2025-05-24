package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
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

func main() {
	// fmt.Println(Addtask("test task"))
	// fmt.Println(Removetask(2))
	fmt.Println(Updatetask("updateInProgress", 3))
	// fmt.Println(data)
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

func Addtask(description string) bool {
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
		return true
	}
	return false
}

func Removetask(id int) bool {
	tasklist := Loadjson()

	for k, v := range tasklist.Todo {
		if id == v.Id {
			tasklist.Todo = append(tasklist.Todo[:k], tasklist.Todo[k+1:]...)
			return Savejson(tasklist)
		}
	}
	for k, v := range tasklist.Inprogress {
		if id == v.Id {
			tasklist.Inprogress = append(tasklist.Inprogress[:k], tasklist.Inprogress[k+1:]...)
			return Savejson(tasklist)
		}
	}
	for k, v := range tasklist.Done {
		if id == v.Id {
			tasklist.Done = append(tasklist.Done[:k], tasklist.Done[k+1:]...)
			return Savejson(tasklist)
		}
	}
	return false
}

func Updatetask(operation string, id int, description ...string) bool {
	tasklist := Loadjson()
	desc := ""
	if len(description) > 0 {
		desc = description[0]
	}
	if operation == "update" {
		for k, v := range tasklist.Todo {
			if id == v.Id {
				tasklist.Todo[k].Description = desc
				return Savejson(tasklist)
			}
		}
		for k, v := range tasklist.Inprogress {
			if id == v.Id {
				tasklist.Inprogress[k].Description = desc
				return Savejson(tasklist)
			}
		}
		for k, v := range tasklist.Done {
			if id == v.Id {
				tasklist.Done[k].Description = desc
				return Savejson(tasklist)
			}
		}
	} else if operation == "updateInProgress" {
		for k, v := range tasklist.Todo {
			if id == v.Id {
				temp := tasklist.Todo[k]
				Removetask(id)
				tasklist.Todo = append(tasklist.Todo[:k], tasklist.Todo[k+1:]...)
				tasklist.Inprogress = append(tasklist.Inprogress, temp)
				return Savejson(tasklist)
			}
		}
		fmt.Printf("Unable to find task with id: %d in in-progress list!!", id)
		return false
	} else if operation == "updateDone" {
		for k, v := range tasklist.Todo {
			if id == v.Id {
				temp := tasklist.Todo[k]
				Removetask(id)
				tasklist.Todo = append(tasklist.Todo[:k], tasklist.Todo[k+1:]...)
				tasklist.Done = append(tasklist.Done, temp)
				return Savejson(tasklist)
			}
		}
		for k, v := range tasklist.Inprogress {
			if id == v.Id {
				temp := tasklist.Inprogress[k]
				Removetask(id)
				tasklist.Inprogress = append(tasklist.Inprogress[:k], tasklist.Inprogress[k+1:]...)
				tasklist.Done = append(tasklist.Done, temp)
				return Savejson(tasklist)
			}
		}
		for _, v := range tasklist.Done {
			if id == v.Id {
				fmt.Println("Task already marked as Done!!")
				return false
			}
		}

	}
	return false
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
