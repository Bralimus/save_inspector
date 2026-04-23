package main

import (
	"fmt"
	"os"

	"github.com/Bralimus/save_inspector/models"
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

	case "view-champion":
		if len(os.Args) < 4 {
			fmt.Println("Usage: save-inspector view-champion <slot> <championID>")
			return
		}

		slot := os.Args[2]
		champID := os.Args[3]

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

		var champ *models.Champion
		for _, c := range data.All {
			if c.ID == champID {
				champ = &c
				break
			}
		}

		if champ == nil {
			fmt.Printf("Champion '%s' not found in save", champID)
			return
		}

		utils.PrintChampion(*champ)

	case "view-inventory":
		if len(os.Args) < 3 {
			fmt.Println("Usage: save-inspector view-inventory <slot>")
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

		fmt.Println("=== ITEM INVENTORY ===")
		for _, item := range data.Items {
			fmt.Printf("- %s (Upgrade Level: %d)\n", item.ID, item.UpgradeLevel)
		}

		fmt.Println("\n=== MATERIAL INVENTORY ===")
		for _, mat := range data.Materials {
			fmt.Printf("- %s (Qty: %d)\n", mat.ID, mat.Quantity)
		}

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

	case "edit-champion":
		if len(os.Args) < 6 {
			fmt.Println("Usage: save-inspector edit-champion <slot> <championID> <field> <value>")
			return
		}

		slot := os.Args[2]
		champID := os.Args[3]
		field := os.Args[4]
		value := os.Args[5]

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

	case "edit-items":
		if len(os.Args) < 5 {
			fmt.Println("Usage: save-inspector edit-items <slot> <itemID> <action>")
			return
		}

		slot := os.Args[2]
		itemID := os.Args[3]
		action := os.Args[4]

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

	case "edit-materials":
		if len(os.Args) < 5 {
			fmt.Println("Usage: save-inspector edit-materials <slot> <materialID> <quantity>")
			return
		}

		slot := os.Args[2]
		materialID := os.Args[3]
		var quantity int
		_, err := fmt.Sscanf(os.Args[4], "%d", &quantity)
		if err != nil {
			fmt.Println("Invalid quantity")
			return
		}

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

		matsRaw, ok := data.Raw["materialInventory"].([]interface{})
		if !ok {
			fmt.Println("Invalid materials format")
			return
		}

		found := false
		for _, mat := range matsRaw {
			if mMap, ok := mat.(map[string]interface{}); ok {
				if id, ok := mMap["materialID"].(string); ok && id == materialID {
					mMap["quantity"] = quantity
					found = true
					break
				}
			}
		}

		if !found {
			if !data.ValidMaterial(materialID) {
				fmt.Printf("Invalid material ID: '%s'\n", materialID)
				return
			}

			newMat := map[string]interface{}{
				"materialID": materialID,
				"quantity":   quantity,
			}
			matsRaw = append(matsRaw, newMat)
			data.Raw["materialInventory"] = matsRaw

			fmt.Printf("Material '%s' added with quantity %d\n", materialID, quantity)
		} else {
			fmt.Printf("Material '%s' updated with quantity %d\n", materialID, quantity)
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

		fmt.Println("Materials updated successfully")

	default:
		fmt.Println("Unknown command")
	}
}
