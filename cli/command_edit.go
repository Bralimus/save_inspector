package cli

import (
	"fmt"

	"github.com/Bralimus/save_inspector/app"
	"github.com/Bralimus/save_inspector/utils"
)

func Edit(app *app.App, args []string) {
	if len(args) < 5 {
		fmt.Println("Usage: save-inspector edit <slot> <field> <value>")
		return
	}

	slot := args[2]
	field := args[3]
	value := args[4]

	path, err := utils.GetSavePathFromSlot(slot, app.OverridePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data, original, err := utils.LoadSave(path)
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
		fmt.Println("Unknown field: ", field)
		return
	}

	err = data.Validate()
	if err != nil {
		fmt.Println("Validation error:", err)
		return
	}

	err = utils.Save(path, original, data.Raw)
	if err != nil {
		fmt.Println("Error saving:", err)
		return
	}

	fmt.Println("Save updated successfully")
}
