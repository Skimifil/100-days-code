package models

type Race struct {
	Name            string `json:"name" gorm:"primary_key"`
	Description     string `json:"description" gorm:"not null;type:string"`
	ModStrength     int    `json:"mod_strength" gorm:"not null;type:int"`
	ModDexterity    int    `json:"mod_dexterity" gorm:"not null:type:int"`
	ModIntelligence int    `json:"mod_intelligence" gorm:"not null;type:int"`
	ModWisdom       int    `json:"mod_wisdom" gorm:"not null;type:int"`
}
