
# Task Tracker CLI

A simple command-line to-do/task manager written in Go. Tasks are stored in a local JSON file and organized by status: **To Do**, **In Progress**, and **Done**.

## Features

- Add new tasks
- Update task descriptions
- Delete tasks
- Mark tasks as in-progress or done
- List all tasks or filter by status

## Usage

```bash
task-tracker <command> [arguments]
```

## Commands

### Add a task

```bash
task-tracker add "Task description"
```

### Update a task

```bash
task-tracker update <id> "New description"
```

### Delete a task

```bash
task-tracker delete <id>
```

### Mark a task as In Progress

```bash
task-tracker mark-in-progress <id>
```



### List tasks

- List all tasks:
```bash
task-tracker list
```
- List all tasks by status:
```bash
task-tracker list todo
task-tracker list in-progress
task-tracker list done
```

### Help

```bash
task-tracker help
```

## Task Storage

All tasks are saved in a local tasks.json file in the same directory.

## Build
To build the binary:
```bash
go build -o task-tracker
```

To build and run with Docker:
```bash
docker build -t task-tracker .
docker run --rm -v $(pwd)/tasks.json:/app/tasks.json task-tracker list
```

### Example
```bash
$ task-tracker add "Finish writing README"
Task added successfully (ID: 1)

$ task-tracker list
```

