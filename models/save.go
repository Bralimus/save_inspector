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
	ID                 string `json:"championID"`
	Level              int    `json:"level"`
	Experience         int    `json:"currentEXP"`
	HP                 int    `json:"currentHealth"`
	Shield             int    `json:"currentShield"`
	BonusDefensePoints int    `json:"bonusDefensePoints"`
	BonusHealthPoints  int    `json:"bonusHealthPoints"`
	BonusPowerPoints   int    `json:"bonusPowerPoints"`
	BonusSpeedPoints   int    `json:"bonusSpeedPoints"`
	EquippedArmor      Item   `json:"equippedArmor"`
	EquippedWeapon     Item   `json:"equippedWeapon"`
	EquippedTrinket    Item   `json:"equippedTrinket"`
	StatPoints         int    `json:"statPoints"`
	TalentPoints       int    `json:"talentPoints"`
	Unlocked           bool   `json:"isUnlocked"`
	IsInParty          bool   `json:"isInParty"`
}

type Item struct {
	ID           string `json:"itemID"`
	UpgradeLevel int    `json:"upgradeLevel"`
}

func (s *SaveData) Validate() error {
	if s.Gold < 0 {
		return fmt.Errorf("gold must be a positive integer")
	}

	for _, champ := range s.Party {
		if champ.Level < 0 {
			return fmt.Errorf("champion %s has invalid level", champ.ID)
		}
		if champ.HP < 0 {
			return fmt.Errorf("champion %s has negative HP", champ.ID)
		}
	}

	return nil
}
