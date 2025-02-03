package model

type Organization struct {
	ID   int    `gorm:"column:id;PRIMARY_KEY" mapstructure:"id"`
	Name string `gorm:"column:name;size:255" mapstructure:"name"`
}

func (Organization) TableName() string {
	return "user"
}
