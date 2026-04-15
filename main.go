package main

import (
	"fmt"
	"os"

	"github.com/Bralimus/save_inspector/parser"
	"github.com/Bralimus/save_inspector/utils"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: save-inspector view <file>")
		return
	}

	command := os.Args[1]
	path := os.Args[2]

	switch command {
	case "view":
		data, _, err := parser.LoadSave(path)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		utils.PrintSummary(data)
	default:
		fmt.Println("Unknown command")
	}
}

// On step 5
