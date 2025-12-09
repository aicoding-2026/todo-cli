# TODO CLI

A simple and efficient TODO List terminal tool written in Go.

## Features

- ✅ Add new TODO items
- 📋 List all TODO items with status
- ✓ Mark TODO items as completed
- 🗑️ Delete TODO items
- 💾 Persistent storage using JSON file

## Installation

### Prerequisites

- Go 1.21 or higher

### Build from Source

```bash
git clone https://github.com/aicoding-2026/todo-cli.git
cd todo-cli
go build -o todo main.go
```

For Windows:
```bash
go build -o todo.exe main.go
```

## Usage

### Add a TODO item

```bash
todo add "Your task description"
```

### List all TODO items

```bash
todo list
```

### Mark a TODO item as completed

```bash
todo complete <id>
```

### Delete a TODO item

```bash
todo delete <id>
```

## Examples

```bash
# Add some tasks
todo add "Buy groceries"
todo add "Write documentation"
todo add "Fix bugs"

# List all tasks
todo list

# Mark task #1 as completed
todo complete 1

# Delete task #2
todo delete 2

# List again to see the changes
todo list
```

## Data Storage

All TODO items are stored in a `todos.json` file in the current directory. This file is automatically created when you add your first TODO item.

## License

MIT License
