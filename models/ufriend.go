package models

type Ufriend struct {
	ID              int `json:"id"`
	UserId          int `json:"userId"`
	UfriendFriendId int `json:"ufriend_friend_id"`
	UfriendStatus   int `json:"ufriend_status"`
	IsFake          int `json:"is_fake"`
}

func (Ufriend) TableName() string {
	return "u_friend"
}
