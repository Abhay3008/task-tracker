package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

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

	// var command string

	// fmt.Scan(&command)

	var list Tasklist
	task := Task{
		Id:          1,
		Description: "temp task",
		CreatedAt:   "20250506223809",
		UpdatedAt:   "20250506223809",
	}
	fmt.Println(task)
	list.Todo = append(list.Todo, task)
	fmt.Println(list.Todo)
	res, _ := json.MarshalIndent(list, "", "\t")
	fmt.Printf("%s\n", res)

	file, _ := os.Create("./tasks.json")
	io.Writer.Write(file, res)
	fmt.Printf("file has been created: %T", res)
	file.Close()

}
