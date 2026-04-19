package models

import "fmt"

type SaveData struct {
	Gold      int        `json:"gold"`
	SceneName string     `json:"sceneName"`
	Party     []Champion `json:"party"`
	All       []Champion `json:"all"`

	Raw map[string]interface{} `json:"-"`
}

type Champion struct {
	ID        string `json:"championID"`
	Level     int    `json:"level"`
	HP        int    `json:"currentHealth"`
	Unlocked  bool   `json:"isUnlocked"`
	IsInParty bool   `json:"isInParty"`
}

func (s *SaveData) Validate() error {
	if s.Gold < 0 {
		return fmt.Errorf("gold must be a positive integer")
	}

	for _, champ := range s.Party {
		if champ.Level < 1 {
			return fmt.Errorf("champion %s has invalid level", champ.ID)
		}
		if champ.HP < 0 {
			return fmt.Errorf("champion %s has negaitve HP", champ.ID)
		}
	}

	return nil
}
