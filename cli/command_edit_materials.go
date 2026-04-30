package cli

import (
	"fmt"

	"github.com/Bralimus/save_inspector/app"
	"github.com/Bralimus/save_inspector/utils"
)

func EditMaterials(app *app.App, args []string) {
	if len(args) < 5 {
		fmt.Println("Usage: save-inspector edit-materials <slot> <materialID> <quantity>")
		return
	}

	slot := args[2]
	materialID := args[3]
	var quantity int
	_, err := fmt.Sscanf(args[4], "%d", &quantity)
	if err != nil {
		fmt.Println("Invalid quantity")
		return
	}

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
}
