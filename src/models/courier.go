package models

type Courier struct {
	ID      int    `gorm:"index;type:int" json:"id"`
	Name    string `gorm:"type:text" json:"name"`
	Created int    `gorm:"index;type:bigint" json:"created"`
	Updated int    `gorm:"index;type:bigint" json:"updated"`
	Busy    int8   `gorm:"type:tinyint" json:"busy"`
}

func (s Courier) TableName() string { return "couriers" }
