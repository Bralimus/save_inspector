package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Bralimus/save_inspector/parser"
	"github.com/Bralimus/save_inspector/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: save-inspector list")
		return
	}

	var overridePath string
	var args []string

	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--path" && i+1 < len(os.Args) {
			overridePath = os.Args[i+1]
			i++
		} else {
			args = append(args, os.Args[i])
		}
	}

	if len(args) == 0 {
		fmt.Println("No command provided")
		return
	}

	command := args[0]

	switch command {
	case "list":
		dir, err := utils.GetSaveDirectory(overridePath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for i := 0; i < 3; i++ {
			path := fmt.Sprintf("%s/save_slot_%d.json", dir, i)
			if _, err := os.Stat(path); err == nil {
				fmt.Printf("[%d] Slot %d - EXISTS\n", i, i)
			} else {
				fmt.Printf("[%d] Slot %d - EMPTY\n", i, i)
			}
		}
	case "view":
		if len(os.Args) < 3 {
			fmt.Println("Usage: save-inspector view <slot>")
			return
		}

		slot := os.Args[2]

		path, err := utils.GetSavePathFromSlot(slot, overridePath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		data, _, err := parser.LoadSave(path)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		utils.PrintSummary(data)
	case "edit":
		if len(os.Args) < 5 {
			fmt.Println("Usage: save-inspector edit <slot> <field> <value>")
			return
		}

		slot := os.Args[2]
		field := os.Args[3]
		value := os.Args[4]

		path, err := utils.GetSavePathFromSlot(slot, overridePath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

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
			fmt.Println("Unknown field: %s", field)
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
