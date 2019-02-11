package main

import (
	"fmt"
	"os"

	cmd "github.com/mvp-eXpress/3g-todo-fullstack/be/repository/db/pkg/cmd/server"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
