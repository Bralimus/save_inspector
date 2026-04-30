package cli

import (
	"fmt"

	"github.com/Bralimus/save_inspector/app"
	"github.com/Bralimus/save_inspector/models"
	"github.com/Bralimus/save_inspector/utils"
)

func ViewChampion(app *app.App, args []string) {
	if len(args) < 4 {
		fmt.Println("Usage: save-inspector view-champion <slot> <championID>")
		return
	}

	slot := args[2]
	champID := args[3]

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
}
