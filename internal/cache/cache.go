package cache

import (
	"cache/internal/dao"
	"cache/internal/model"
	"cache/pkg/utils"
)

type Cache struct {
	UserCache   *utils.Cache[int, model.User]
	PersonCache *utils.Cache[int, model.Organization]

	dao *dao.Dao
}

func NewCache(dao *dao.Dao) *Cache {
	return &Cache{
		UserCache:   nil,
		PersonCache: nil,
		dao:         dao,
	}
}
