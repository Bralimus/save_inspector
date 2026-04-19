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
	}
}
