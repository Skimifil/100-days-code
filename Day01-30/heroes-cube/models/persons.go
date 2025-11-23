package models

type Person struct {
	Id              string `json:"id" gorm:"primary_key;not null;type:uuid"`
	RaceName        string `json:"race_name" gorm:"not null"`
	Race            Race   `json:"race" gorm:"foreignkey:Racename;references:Name"`
	PtsStrength     int    `json:"pts_strength" gorm:"not null;type:int"`
	PtsDexterity    int    `json:"pts_dexterity" gorm:"not null:type:int"`
	PtsIntelligence int    `json:"pts_intelligence" gorm:"not null;type:int"`
	PtsWisdom       int    `json:"pts_wisdom" gorm:"not null;type:int"`
}

func (p *Person) GetStrength() int {
	return p.PtsStrength + p.Race.ModStrength
}

func (p *Person) GetDexterity() int {
	return p.PtsDexterity + p.Race.ModDexterity
}

func (p *Person) GetIntelligence() int {
	return p.PtsIntelligence + p.Race.ModIntelligence
}

func (p *Person) GetWisdom() int {
	return p.PtsWisdom + p.Race.ModWisdom
}
