package cli

import (
	"fmt"

	"github.com/Bralimus/save_inspector/app"
	"github.com/Bralimus/save_inspector/utils"
)

func ViewInventory(app *app.App, args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: save-inspector view-inventory <slot>")
		return
	}

	slot := args[2]

	path, err := utils.GetSavePathFromSlot(slot, app.OverridePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data, _, err := utils.LoadSave(path)
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
}
