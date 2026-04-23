package models

import (
	"fmt"

	"github.com/Bralimus/save_inspector/data"
)

type SaveData struct {
	Gold      int        `json:"gold"`
	SceneName string     `json:"sceneName"`
	Items     []Item     `json:"itemInventory"`
	Materials []Material `json:"materialInventory"`
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

type Material struct {
	ID       string `json:"materialID"`
	Quantity int    `json:"quantity"`
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

func (s *SaveData) ValidItem(itemID string) bool {
	_, exists := data.ValidItems[itemID]
	return exists
}

func (s *SaveData) ValidMaterial(materialID string) bool {
	_, exists := data.ValidMaterials[materialID]
	return exists
}
