# Evershard Save Inspector

A CLI tool built in Go for inspecting and safely editing Unity JSON save files for Evershard: Heroes of Gallan's Landing.

## Features
 - View save file data in a readable format
 - Edit values of save file (ex. Gold)
 - Automatic backup before modification
 - Validation to prevent invalid game states
 - Preserves untouched JSON fields (no data loss)

## Usage
1. View editable fields of a save file
    go run main.go view testdata/save.json
2. Edit field to desired value
    go run main.go edit testdata/save.json gold 999

## How It Works
 - Loads the full save file into a raw JSON map and extracts editable fields into structured models.
 - Changes are applied to structured models and synced back into the raw JSON map, then written back to disk

## Future Improvements
 - Support editing more fields
 - Support editing nested fields (ex. champion stats)
 - Add UI frontend