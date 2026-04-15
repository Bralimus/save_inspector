package utils

import (
	"fmt"

	"github.com/Bralimus/save_inspector/models"
)

func PrintSummary(data *models.SaveData) {
	fmt.Println("=== SAVE SUMMARY ===")
	fmt.Println("Gold:", data.Gold)
	fmt.Println("Dungeon:", data.CurrentDungeon)

	fmt.Println("\nParty:")
	for _, c := range data.Party {
		fmt.Printf("- %s (Lvl %d, HP %d)\n", c.ID, c.Level, c.HP)
	}
}
