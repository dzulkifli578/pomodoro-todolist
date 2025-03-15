# Pomodoro CLI

Pomodoro CLI is a command-line interface (CLI) application for managing tasks and running a Pomodoro timer. This application is built with the Go programming language and uses the [Cobra](https://github.com/spf13/cobra) library for command management.

## Features

- **Task Management**  
  - `create`: Create a new task  
  - `read`: Display all tasks  
  - `delete`: Delete a task  

- **Pomodoro Timer**  
  - Runs a Pomodoro session for 25 minutes  
  - Provides options for a short break (5 minutes) and a long break (15 minutes)  
  - Allows canceling the Pomodoro session by typing `cancel` during the countdown

## Project Structure

```md
pomodoro-cli/
├── cmd/
│ ├── root.go # Root command and Cobra initialization
│ ├── task.go # Subcommands for task management (create, read, delete)
│ └── pomodoro.go # Subcommand for running the Pomodoro timer
├── internal/
│ └── task/ # Internal logic for task management
├── pkg/
│ └── timer/ # Timer package for Pomodoro logic
├── go.mod # Go module file
└── main.go # Application entry point
```

## Installation

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/username/pomodoro-cli.git
    cd pomodoro-cli
    ```
   
2. **Install Dependencies:**

    Make sure you have Go installed (minimum version 1.20). Run the following command:
    
    ```bash
    go mod tidy
    ```

3. **Build the Application:**

    ```bash
    go build -o pomodoro-cli
    ```

## Usage

Run the application with the command:

```bash
./pomodoro-cli
```

When no arguments are provided, the application displays a welcome message:

```bash
Welcome To Pomodoro CLI!
```

Example Commands

- Create a New Task:

    ```bash
    ./pomodoro-cli task create
    ```
  
- Read All Tasks:

    ```bash
    ./pomodoro-cli task read
    ```
  
- Delete a Task:

    ```bash
    ./pomodoro-cli task delete
    ```

- Start the Pomodoro Timer:

    ```bash
    ./pomodoro-cli pomodoro
    ```

> **Note:**  
> During a Pomodoro session, you can cancel the timer by typing `cancel` during the countdown.

## Contributing

Contributions are welcome! Please fork this repository, create a feature or bugfix branch, and submit a pull request. Be sure to update the documentation if significant changes are made.

## License

This project is distributed under the [MIT License](LICENSE).
