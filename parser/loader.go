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

	// Load full JSON into raw map
	var raw map[string]interface{}
	err = json.Unmarshal(file, &raw)
	if err != nil {
		return nil, nil, err
	}

	// Extract known fields safely
	data := &models.SaveData{
		Raw: raw,
	}

	if val, ok := raw["gold"].(float64); ok {
		data.Gold = int(val)
	}

	if val, ok := raw["currentDungeon"].(string); ok {
		data.CurrentDungeon = val
	}

	if partyRaw, ok := raw["party"].([]interface{}); ok {
		for _, p := range partyRaw {
			if pMap, ok := p.(map[string]interface{}); ok {
				champ := models.Champion{}

				if id, ok := pMap["id"].(string); ok {
					champ.ID = id
				}
				if lvl, ok := pMap["level"].(float64); ok {
					champ.Level = int(lvl)
				}
				if hp, ok := pMap["hp"].(float64); ok {
					champ.HP = int(hp)
				}

				data.Party = append(data.Party, champ)
			}
		}
	}

	return data, file, nil
}
