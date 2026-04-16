package models

type SaveData struct {
	Gold           int        `json:"gold"`
	CurrentDungeon string     `json:"currentDungeon"`
	Party          []Champion `json:"party"`

	Raw map[string]interface{} `json:"-"`
}

type Champion struct {
	ID    string `json:"id"`
	Level int    `json:"level"`
	HP    int    `json:"hp"`
}
