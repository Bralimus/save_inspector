package utils

import (
	"fmt"

	"github.com/Bralimus/save_inspector/models"
)

func PrintSummary(data *models.SaveData) {
	fmt.Println("=== SAVE SUMMARY ===")
	fmt.Println("Gold:", data.Gold)
	fmt.Println("Scene:", data.SceneName)

	PrintChampions(data.Party, data.All)
}

func PrintChampions(party []models.Champion, all []models.Champion) {
	fmt.Println("\n=== PARTY ===")
	for i, c := range party {
		fmt.Printf("[%d] %-8s - Lvl %d - HP %d\n", i, c.ID, c.Level, c.HP)
	}

	fmt.Println("\n=== ALL CHAMPIONS ===")
	for _, c := range all {
		lock := ""
		if !c.Unlocked {
			lock = " - LOCKED"
		}

		check := "[ ]"
		if c.IsInParty {
			check = "[✔]"
		}

		fmt.Printf("%s %-8s - Lvl %d - HP %d%s\n",
			check, c.ID, c.Level, c.HP, lock)

		fmt.Println()
		fmt.Printf("Experience: %d\n", c.Experience)
		fmt.Printf("Shield: %d\n", c.Shield)
		fmt.Printf("Bonus Defense: %d\n", c.BonusDefensePoints)
		fmt.Printf("Bonus Health: %d\n", c.BonusHealthPoints)
		fmt.Printf("Bonus Power: %d\n", c.BonusPowerPoints)
		fmt.Printf("Bonus Speed: %d\n", c.BonusSpeedPoints)

		fmt.Printf("Equipped Armor: %s (Upgrade Level %d)\n", c.EquippedArmor.ID, c.EquippedArmor.UpgradeLevel)
		fmt.Printf("Equipped Weapon: %s (Upgrade Level %d)\n", c.EquippedWeapon.ID, c.EquippedWeapon.UpgradeLevel)
		fmt.Printf("Equipped Trinket: %s (Upgrade Level %d)\n", c.EquippedTrinket.ID, c.EquippedTrinket.UpgradeLevel)

		fmt.Printf("Stat Points: %d\n", c.StatPoints)
		fmt.Printf("Talent Points: %d\n", c.TalentPoints)
		fmt.Println()
	}
}
