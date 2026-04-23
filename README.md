# Evershard Save Inspector

A CLI tool built in Go for inspecting and safely editing Unity JSON save files for Evershard: Heroes of Gallan's Landing.

## Features
 - View save file data in a readable format
 - Edit values of save file (ex. Gold)
 - Automatic backup before modification
 - Validation to prevent invalid game states
 - Preserves untouched JSON fields (no data loss)

## Commands
 - View list of save files <br>
    `go run main.go list`
 - View editable fields within save file <br>
    `go run main.go view <slot #>`
 - View full item and material inventory <br>
    `go run main.go view-inventory`
 - View champions within specified save file <br>
    `go run main.go view-champion <slot #> <championID>`
 - Edit field within save file to desired value <br>
    `go run main.go edit <slot #> <field> <value>`
 - Edit champion specific fields to desired value <br>
    `go run main.go edit-champion <slot #> <championID> <field> <value>`
 - Add, remove, or upgrade items <br>
    `go run main.go edit-items <slot #> <itemID> <action>`
 - Set material quantity
    `go run main.go edit-materials <slot #> <materialID> <quantity>`

## Notes
 - Database of all available items and materials are under data/

## How It Works
 - Loads the full save file into a raw JSON map and extracts editable fields into structured models.
 - Changes are applied to structured models and synced back into the raw JSON map, then written back to disk

## Future Improvements
 - Add UI frontend
