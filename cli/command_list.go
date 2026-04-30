package cli

import (
	"fmt"
	"os"

	"github.com/Bralimus/save_inspector/app"
	"github.com/Bralimus/save_inspector/utils"
)

func List(app *app.App, args []string) {
	dir, err := utils.GetSaveDirectory(app.OverridePath)
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
}
