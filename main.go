package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// TodoItem represents a single TODO item
type TodoItem struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// TodoList represents a collection of TODO items
type TodoList struct {
	Items []TodoItem `json:"items"`
	NextID int       `json:"next_id"`
}

const dataFile = "todos.json"

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if addCmd.NArg() == 0 {
			fmt.Println("Error: Please provide a task description")
			os.Exit(1)
		}
		task := addCmd.Args()[0]
		addTodo(task)

	case "list":
		listCmd.Parse(os.Args[2:])
		listTodos()

	case "complete":
		completeCmd.Parse(os.Args[2:])
		if completeCmd.NArg() == 0 {
			fmt.Println("Error: Please provide a task ID")
			os.Exit(1)
		}
		id, err := strconv.Atoi(completeCmd.Args()[0])
		if err != nil {
			fmt.Printf("Error: Invalid task ID: %s\n", completeCmd.Args()[0])
			os.Exit(1)
		}
		completeTodo(id)

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if deleteCmd.NArg() == 0 {
			fmt.Println("Error: Please provide a task ID")
			os.Exit(1)
		}
		id, err := strconv.Atoi(deleteCmd.Args()[0])
		if err != nil {
			fmt.Printf("Error: Invalid task ID: %s\n", deleteCmd.Args()[0])
			os.Exit(1)
		}
		deleteTodo(id)

	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("TODO List CLI Tool")
	fmt.Println("\nUsage:")
	fmt.Println("  todo add <task>        Add a new TODO item")
	fmt.Println("  todo list              List all TODO items")
	fmt.Println("  todo complete <id>     Mark a TODO item as completed")
	fmt.Println("  todo delete <id>      Delete a TODO item")
	fmt.Println("\nExamples:")
	fmt.Println("  todo add \"Buy groceries\"")
	fmt.Println("  todo list")
	fmt.Println("  todo complete 1")
	fmt.Println("  todo delete 2")
}

func loadTodos() *TodoList {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &TodoList{Items: []TodoItem{}, NextID: 1}
		}
		fmt.Printf("Error reading todos: %v\n", err)
		os.Exit(1)
	}

	var todoList TodoList
	if err := json.Unmarshal(data, &todoList); err != nil {
		fmt.Printf("Error parsing todos: %v\n", err)
		os.Exit(1)
	}

	return &todoList
}

func saveTodos(todoList *TodoList) {
	data, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		fmt.Printf("Error encoding todos: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(dataFile, data, 0644); err != nil {
		fmt.Printf("Error saving todos: %v\n", err)
		os.Exit(1)
	}
}

func addTodo(task string) {
	todoList := loadTodos()
	
	item := TodoItem{
		ID:        todoList.NextID,
		Task:      task,
		Completed: false,
		CreatedAt: time.Now(),
	}
	
	todoList.Items = append(todoList.Items, item)
	todoList.NextID++
	
	saveTodos(todoList)
	fmt.Printf("Added TODO #%d: %s\n", item.ID, task)
}

func listTodos() {
	todoList := loadTodos()
	
	if len(todoList.Items) == 0 {
		fmt.Println("No TODO items found.")
		return
	}

	fmt.Println("\nTODO List:")
	fmt.Println(strings.Repeat("-", 60))
	for _, item := range todoList.Items {
		status := "[ ]"
		if item.Completed {
			status = "[✓]"
		}
		fmt.Printf("%s #%d: %s\n", status, item.ID, item.Task)
		fmt.Printf("    Created: %s\n", item.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	fmt.Println(strings.Repeat("-", 60))
}

func completeTodo(id int) {
	todoList := loadTodos()
	
	found := false
	for i := range todoList.Items {
		if todoList.Items[i].ID == id {
			if todoList.Items[i].Completed {
				fmt.Printf("TODO #%d is already completed.\n", id)
				return
			}
			todoList.Items[i].Completed = true
			found = true
			break
		}
	}
	
	if !found {
		fmt.Printf("Error: TODO #%d not found.\n", id)
		os.Exit(1)
	}
	
	saveTodos(todoList)
	fmt.Printf("Marked TODO #%d as completed.\n", id)
}

func deleteTodo(id int) {
	todoList := loadTodos()
	
	found := false
	for i, item := range todoList.Items {
		if item.ID == id {
			todoList.Items = append(todoList.Items[:i], todoList.Items[i+1:]...)
			found = true
			break
		}
	}
	
	if !found {
		fmt.Printf("Error: TODO #%d not found.\n", id)
		os.Exit(1)
	}
	
	saveTodos(todoList)
	fmt.Printf("Deleted TODO #%d.\n", id)
}
