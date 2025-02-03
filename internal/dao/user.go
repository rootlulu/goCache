package dao

import (
	"cache/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserFilter struct {
	ID   *int32
	Name *string
}

func (d *Dao) GetUser(filter *UserFilter) ([]model.User, error) {
	var rows []model.User
	var res *gorm.DB
	if filter.ID != nil && filter.Name != nil {
		res = d.Mysql.Debug().Where("id = ? AND name = ?", *filter.ID, *filter.Name).Find(&rows)
	} else if filter.ID != nil {
		res = d.Mysql.Debug().Where("id = ?", *filter.ID).Find(&rows)
	} else if filter.Name != nil {
		res = d.Mysql.Debug().Where("name = ?", *filter.Name).Find(&rows)
	} else {
		res = d.Mysql.Debug().Find(&rows)
	}

	if res.Error != nil {
		return nil, res.Error
	}

	logrus.Infoln(rows)
	return rows, nil
}
