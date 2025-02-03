package cache

import (
	"cache/internal/model"
	"cache/pkg/utils"

	"github.com/sirupsen/logrus"
)

// TODO: judge the cache is nil, if nil, create the cache.
// use the model.Filter but not id only.
func (c *Cache) loadUserCache() error {
	c.UserCache = utils.NewCache[int, model.User]()
	var rows []model.User
	db := c.dao.Mysql.Find(&rows)
	logrus.Infoln(rows)
	if db.Error != nil {
		logrus.Errorf("failed to get user cache: %s", db.Error)
		return db.Error
	}
	for _, row := range rows {
		logrus.Infoln(row)
		c.UserCache.Set(row.ID, row)
	}
	return nil
}

func (c *Cache) GetUserCacheBool(id int) (*model.User, bool) {
	if c.UserCache == nil {
		err := c.loadUserCache()
		if err != nil {
			return nil, false
		}
	}
	row, ok := c.UserCache.GetBool(id)
	return &row, ok
}

func (c *Cache) GetUserCache(id int) *model.User {
	if c.UserCache == nil {
		err := c.loadUserCache()
		if err != nil {
			return nil
		}
	}
	row := c.UserCache.Get(id)
	return &row
}

func (c *Cache) SetUserCache(id int, updater model.User) error {
	if c.UserCache == nil {
		err := c.loadUserCache()
		if err != nil {
			return nil
		}
	}
	var row model.User
	if db := c.dao.Mysql.Where("id = ?", id).First(&row); db.Error != nil {
		return db.Error
	}
	if db := c.dao.Mysql.Model(&row).Updates(&updater); db.Error != nil {
		return db.Error
	}
	c.UserCache.Set(row.ID, updater)
	return nil
}

func (c *Cache) DelUserCache(id int) error {
	if c.UserCache == nil {
		err := c.loadUserCache()
		if err != nil {
			return nil
		}
	}
	if db := c.dao.Mysql.Where("id = ?", id).Delete(&model.User{}); db.Error != nil {
		return db.Error
	}
	c.UserCache.Del(id)
	return nil
}
