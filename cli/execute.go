package cli

import (
	"github.com/Bralimus/save_inspector/app"
)

func Execute(app *app.App, rawArgs []string) {
	if len(rawArgs) < 2 {
		PrintUsage()
		return
	}

	args := ParseArgs(app, rawArgs)

	command := args[0]

	switch command {
	case "list":
		List(app, args)
	case "view":
		View(app, args)
	case "view-champion":
		ViewChampion(app, args)
	case "view-inventory":
		ViewInventory(app, args)
	case "edit":
		Edit(app, args)
	case "edit-champion":
		EditChampion(app, args)
	case "edit-items":
		EditItems(app, args)
	case "edit-materials":
		EditMaterials(app, args)
	default:
		PrintUsage()
	}
}

func PrintUsage() {
	println("Usage:")
	println("  save-inspector list [--path <save_directory>]")
	println("  save-inspector edit <slot_number> [--path <save_directory>]")
}

func ParseArgs(a *app.App, rawArgs []string) []string {
	var args []string

	for i := 1; i < len(rawArgs); i++ {
		if rawArgs[i] == "--path" && i+1 < len(rawArgs) {
			a.OverridePath = rawArgs[i+1]
			i++
		} else {
			args = append(args, rawArgs[i])
		}
	}

	return args
}
