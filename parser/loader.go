package parser

import (
	"encoding/json"
	"os"

	"github.com/Bralimus/save_inspector/models"
)

func LoadSave(path string) (*models.SaveData, []byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	var data models.SaveData
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, nil, err
	}

	return &data, file, nil
}
