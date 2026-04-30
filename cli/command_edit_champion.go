package cli

import (
	"fmt"

	"github.com/Bralimus/save_inspector/app"
	"github.com/Bralimus/save_inspector/utils"
)

func EditChampion(app *app.App, args []string) {
	if len(args) < 6 {
		fmt.Println("Usage: save-inspector edit-champion <slot> <championID> <field> <value>")
		return
	}

	slot := args[2]
	champID := args[3]
	field := args[4]
	value := args[5]

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

	champsRaw, ok := data.Raw["ownedChampions"].([]interface{})
	if !ok {
		fmt.Println("Invalid save format")
		return
	}

	var target map[string]interface{}
	for _, c := range champsRaw {
		if cMap, ok := c.(map[string]interface{}); ok {
			if id, ok := cMap["championID"].(string); ok && id == champID {
				target = cMap
				break
			}
		}
	}

	if target == nil {
		fmt.Printf("Champion '%s' not found\n", champID)
		return
	}

	switch field {
	case "level":
		var newLevel int
		_, err := fmt.Sscanf(value, "%d", &newLevel)
		if err != nil {
			fmt.Println("Invalid level")
			return
		}
		target["level"] = newLevel

	case "hp":
		var newHP int
		_, err := fmt.Sscanf(value, "%d", &newHP)
		if err != nil {
			fmt.Println("Invalid HP")
			return
		}
		target["currentHealth"] = newHP

	default:
		fmt.Println("Unknown champion field: ", field)
		return
	}

	utils.Save(path, original, data.Raw)

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

	fmt.Println("Champion updated successfully")
}
