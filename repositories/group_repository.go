package repositories

import (
	"strconv"
	"time"

	"fast.bibabo.vn/lib"
	"fast.bibabo.vn/models"
	"github.com/go-redis/cache/v8"
	"gorm.io/gorm"
)

type GroupRepository interface {
	Create(group models.Group) models.Group
	Find(id int) models.Group
	FindAll(offset int, limit int) []models.Group
	Update(id int, group models.Group) models.Group
	Delete(id int) bool
}

type groupRrpository struct {
	db    *gorm.DB
	cache *cache.Cache
}

func InstanceGroupRepository(db *gorm.DB, cache *cache.Cache) GroupRepository {
	return &groupRrpository{
		db:    db,
		cache: cache,
	}
}

func (r *groupRrpository) Create(group models.Group) models.Group {
	r.db.Create(&group)
	return group
}

func (r *groupRrpository) Find(id int) models.Group {
	var group models.Group
	r.db.Where("id", id).First(&group)
	return group
}

func (r *groupRrpository) FindAll(offset int, limit int) []models.Group {
	var groups []models.Group
	key := "find_all_groups:o" + strconv.Itoa(offset) + ":l" + strconv.Itoa(limit)
	error := r.cache.Once(&cache.Item{
		Key:   key,
		Value: &groups,
		TTL:   time.Minute * 5,
		Do: func(i *cache.Item) (interface{}, error) {
			r.db.Offset(offset).Limit(limit).Find(&groups)
			for key, group := range groups {
				groups[key].ImageAvatar = lib.GetInstanceCdn().GetImage(group.ImageAvatar, lib.FOLDER_QUESTION, lib.SIZE_LANGE, lib.SIZE_LANGE)
				groups[key].ImageCover = lib.GetInstanceCdn().GetImage(group.ImageCover, lib.FOLDER_QUESTION, lib.SIZE_LANGE, lib.SIZE_LANGE)
			}

			return &groups, nil
		},
	})

	if error != nil {
		lib.SlackLog("find all group error")
	}
	go func() {
		time.Sleep(time.Minute)
		lib.SlackLog("Send background message")
	}()
	return groups
}

func (r *groupRrpository) Update(id int, group models.Group) models.Group {
	return group
}

func (r *groupRrpository) Delete(id int) bool {
	r.db.Delete(&models.Group{}, id)
	return true
}
