package main

import (
	"encoding/json"
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
	case "edit":
		if len(os.Args) < 5 {
			fmt.Println("Usage: save-inspector edit <file> <field> <value>")
			return
		}

		field := os.Args[3]
		value := os.Args[4]

		data, original, err := parser.LoadSave(path)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		switch field {
		case "gold":
			var newGold int
			_, err := fmt.Sscanf(value, "%d", &newGold)
			if err != nil {
				fmt.Println("Invalid number for gold")
				return
			}
			data.Raw["gold"] = newGold

		default:
			fmt.Println("Unknown field:", field)
			return
		}

		err = data.Validate()
		if err != nil {
			fmt.Println("Validation error:", err)
			return
		}

		// Backup original file
		err = os.WriteFile(path+".bak", original, 0644)
		if err != nil {
			fmt.Println("Backup failed:", err)
			return
		}

		// Save updated file
		updated, err := json.MarshalIndent(data.Raw, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling:", err)
			return
		}

		err = os.WriteFile(path, updated, 0644)
		if err != nil {
			fmt.Println("Write failed:", err)
			return
		}

		fmt.Println("Save updated successfully")
	default:
		fmt.Println("Unknown command")
	}
}

// On step 9
