package cli

import (
	"fmt"

	"github.com/Bralimus/save_inspector/app"
	"github.com/Bralimus/save_inspector/utils"
)

func View(app *app.App, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: save-inspector view <slot>")
		return
	}

	slot := args[1]

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

	utils.PrintSummary(data)
}
