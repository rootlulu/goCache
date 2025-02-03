package model

import (
	"fmt"
	"strconv"
)

type User struct {
	ID    int    `gorm:"column:id;PRIMARY_KEY" mapstructure:"id"`        // primary key
	Name  string `gorm:"column:name;size:255" mapstructure:"name"`       // set field size to 255
	Sex   bool   `gorm:"column:sex;default:true" mapstructure:"sex"`     // set field default value to true: male, false: female
	OrgID int    `gorm:"index;column:orgId" mapstructure:"orgid"`        // create index for this field
	Age   int    `gorm:"index;column:age;default:18" mapstructure:"age"` // create index and set field default value to 18
}

func (User) TableName() string {
	return "user"
}

func (u *User) String() string {
	return fmt.Sprintf(`User: 
	ID: " + %s + ", Name: " + %s + ",	OrgID: " + %s + ", Age: " + %s + "`,
		strconv.Itoa(u.ID),
		u.Name,
		strconv.Itoa(u.OrgID),
		strconv.Itoa(u.Age))
}
