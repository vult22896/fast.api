package models

type Ufriend struct {
	ID              int `json:"id"`
	UserId          int `json:"userId"`
	UfriendFriendId int `json:"ufriendFriendId"`
	UfriendStatus   int `json:"ufriendStatus"`
	IsFake          int `json:"isFake"`
}

func (Ufriend) TableName() string {
	return "ufriend"
}
