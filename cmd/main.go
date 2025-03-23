package main

import (
	"fmt"
	"os"

	"github.com/UseVader/vader/internal/commands"
)

func main() {
	args := os.Args[1:]

	fmt.Println("Welcome to Vader")

	commands.RunCmd(args)
}
