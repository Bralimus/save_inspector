package cli

import (
	"fmt"

	"github.com/Bralimus/save_inspector/app"
	"github.com/Bralimus/save_inspector/models"
	"github.com/Bralimus/save_inspector/utils"
)

func EditItems(app *app.App, args []string) {
	if len(args) < 5 {
		fmt.Println("Usage: save-inspector edit-items <slot> <itemID> <action>")
		return
	}

	slot := args[2]
	itemID := args[3]
	action := args[4]

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

	itemsRaw, ok := data.Raw["itemInventory"].([]interface{})
	if !ok {
		fmt.Println("Invalid inventory format")
		return
	}

	switch action {
	case "add":
		if !data.ValidItem(itemID) {
			fmt.Printf("Invalid item ID: '%s'\n", itemID)
			return
		}

		newItem := models.Item{
			ID:           itemID,
			UpgradeLevel: 0,
		}
		itemsRaw = append(itemsRaw, newItem)
		data.Raw["itemInventory"] = itemsRaw

	case "remove":
		var updated []interface{}

		for _, i := range itemsRaw {
			iMap := i.(map[string]interface{})
			if iMap["itemID"] != itemID {
				updated = append(updated, i)
			}
		}

		data.Raw["itemInventory"] = updated

	case "upgrade":
		found := false

		for _, i := range itemsRaw {
			iMap := i.(map[string]interface{})
			if iMap["itemID"] == itemID {
				if level, ok := iMap["upgradeLevel"].(float64); ok {
					iMap["upgradeLevel"] = int(level) + 1
					found = true
					break
				}
			}
		}

		if !found {
			fmt.Printf("Item '%s' not found\n", itemID)
			return
		}

	default:
		fmt.Println("Unknown inventory action: ", action)
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

	fmt.Println("Inventory updated successfully")
}
