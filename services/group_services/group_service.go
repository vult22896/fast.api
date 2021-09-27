package services

import (
	"fast.bibabo.vn/models"
	"fast.bibabo.vn/repositories"
)

type GroupService interface {
	repositories.GroupRepository
}

type groupService struct {
	groupRepository repositories.GroupRepository
}

func InstanceGroupService(groupRepository repositories.GroupRepository) GroupService {
	return &groupService{groupRepository: groupRepository}
}

func (s *groupService) Create(group models.Group) models.Group {
	return s.groupRepository.Create(group)
}

func (s *groupService) Find(id int) models.Group {

	return s.groupRepository.Find(id)
}

func (s *groupService) FindAll(offset int, limit int) []models.Group {

	return s.groupRepository.FindAll(offset, limit)
}

func (s *groupService) Update(id int, group models.Group) models.Group {
	return group
}

func (s *groupService) Delete(int) bool {
	return true
}
