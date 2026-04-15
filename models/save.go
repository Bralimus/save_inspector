package models

type SaveData struct {
	Gold           int        `json:"gold"`
	CurrentDungeon string     `json:"currentDungeon"`
	Party          []Champion `json:"party"`
}

type Champion struct {
	ID    string `json:"id"`
	Level int    `json:"level"`
	HP    int    `json:"hp"`
}
