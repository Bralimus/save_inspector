package utils

import (
	"encoding/json"
	"os"
)

func Save(path string, original []byte, raw map[string]interface{}) error {
	// Backup original file
	err := os.WriteFile(path+".bak", original, 0644)
	if err != nil {
		return err
	}

	// Save updated file
	updated, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, updated, 0644)
	if err != nil {
		return err
	}

	return nil
}
