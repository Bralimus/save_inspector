package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetDefaultSavePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(home, "AppData", "LocalLow", "Herculean Studios", "Evershard_ Heroes of Gallan's Landing", "save_slot_0.json")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("save file not found at %s", path)
	}

	return path, nil
}

func GetSaveDirectory(override string) (string, error) {
	if override != "" {
		if _, err := os.Stat(override); os.IsNotExist(err) {
			return "", fmt.Errorf("provided path not found: %s", override)
		}
		return override, nil
	}

	if env := os.Getenv("SAVE-INSPECTOR_DIR"); env != "" {
		if _, err := os.Stat(env); err == nil {
			return env, nil
		}
	}

	switch runtime.GOOS {
	case "windows":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		dir := filepath.Join(home, "AppData", "LocalLow", "Herculean Studios", "Evershard_ Heroes of Gallan's Landing")
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			return "", fmt.Errorf("save directory not found at %s", dir)
		}
		return dir, nil
	case "linux":
		base := "/mnt/c/Users"

		entries, err := os.ReadDir(base)
		if err != nil {
			return "", fmt.Errorf("failed to read directory: %s", base)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				dir := filepath.Join(base, entry.Name(), "AppData", "LocalLow", "Herculean Studios", "Evershard_ Heroes of Gallan's Landing")
				if _, err := os.Stat(dir); err == nil {
					return dir, nil
				}
			}
		}
		return "", fmt.Errorf("save directory not found")
	default:
		return "", fmt.Errorf("unsupported platform")
	}
}

func GetSavePathFromSlot(slot string, override string) (string, error) {
	baseDir, err := GetSaveDirectory(override)
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("save_slot_%s.json", slot)
	fullPath := filepath.Join(baseDir, filename)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return "", fmt.Errorf("save file not found at %s", fullPath)
	}

	return fullPath, nil
}
