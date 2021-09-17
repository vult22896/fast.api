package models

type UserGroupFollow struct {
	Id       int `json:"id"`
	UserId   int `json:"userId"`
	GroupId  int `json:"groupId"`
	Type     int `json:"type"`
	SendChat int `json:"sendChat"`
}

func (UserGroupFollow) TableName() string {
	return "user_group_follow"
}
