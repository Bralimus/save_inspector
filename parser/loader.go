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

	if val, ok := raw["sceneName"].(string); ok {
		data.SceneName = val
	}

	selectedChampions := map[string]bool{}

	if selRaw, ok := raw["selectedChampionIDs"].([]interface{}); ok {
		for _, s := range selRaw {
			if id, ok := s.(string); ok {
				selectedChampions[id] = true
			}
		}
	}

	if champsRaw, ok := raw["ownedChampions"].([]interface{}); ok {
		for _, c := range champsRaw {
			if cMap, ok := c.(map[string]interface{}); ok {
				champ := models.Champion{}

				if id, ok := cMap["championID"].(string); ok {
					champ.ID = id
				}
				if lvl, ok := cMap["level"].(float64); ok {
					champ.Level = int(lvl)
				}
				if hp, ok := cMap["currentHealth"].(float64); ok {
					champ.HP = int(hp)
				}
				if unlocked, ok := cMap["isUnlocked"].(bool); ok {
					champ.Unlocked = unlocked
				}

				champ.IsInParty = selectedChampions[champ.ID]

				data.All = append(data.All, champ)

				if champ.IsInParty && champ.Unlocked {
					data.Party = append(data.Party, champ)
				}
			}
		}
	}

	return data, file, nil
}
