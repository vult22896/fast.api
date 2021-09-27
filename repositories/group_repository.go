package repositories

import (
	"fmt"

	"fast.bibabo.vn/models"
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
	db *gorm.DB
}

func InstanceGroupRepository(db *gorm.DB) GroupRepository {
	return &groupRrpository{
		db: db,
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
	fmt.Println(offset)
	fmt.Println(limit)
	var groups []models.Group
	r.db.Offset(offset).Limit(limit).Find(&groups)

	return groups
}

func (r *groupRrpository) Update(id int, group models.Group) models.Group {
	return group
}

func (r *groupRrpository) Delete(int) bool {
	return true
}
