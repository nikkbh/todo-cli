# TODO CLI built using Go, Cobra, GORM, PostgreSQL & Docker

Simple TODO CLI built in <a href="https://go.dev/doc/" target="_blank" rel="noreferrer"><img  src="https://raw.githubusercontent.com/danielcranney/readme-generator/main/public/icons/skills/go-colored.svg" width="56" height="56" alt="Go" /></a>

## Description
A simple TODO CLI app built in Go. It supports adding new tasks, listing out tasks based on their status, marking the task as done, and deleting an item.
Used **PostgreSQL Docker Container** to store the TODO items as it speeds up the development and reduces the overhead of setting up on your local machine.

**NOTE:** The repository also contains the `compose.yaml` file to spin up your own PostgreSQL Docker container with PgAdmin. To use just run `docker-compose up -d`
in the project root.

## Technologies uses

- Go
- Cobra
- GORM
- PostgreSQL
- Docker

## Setup

After cloning the repository run the following command to build the CLI.
```bash
go build -o bin/todo.exe
```
The ```.exe```extension is used only to run this CLI on Windows. To build the CLI on Linux use the following command:
```
go build -o bin/todo
```
A reference to build for other OS environments [here.](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures)

To create the database table from the struct types, run the `migrate.go` file using `go run migrate.go`. (Remember to `cd` into the migrate directory)

## Usage
```
TODO CLI is an easy-to-use TODO application for managing your day-to-day tasks.
                        For example:
                        todo --new "Pick Laundry"

Usage:
  todo [flags]

Flags:
  -x, --delete int    Delete the TODO item
  -d, --done int      Mark the TODO item to Done
  -h, --help          help for todo
  -l, --list string   List the status of TODOS. 'all', 'pending', 'done'
  -n, --new string    Add a New TODO item.
```

## Troubleshooting
- If you get database-related errors, I would suggest you look at whether the docker container is up and running using `docker ps`
- I have mapped port 5431 of the HOST to 5432 of the Container as I already have Postgres installed on my machine, thus to avoid conflicting ports, I have used port 5431.
