package models

type PostLike struct {
	ID        int `json:"id" gorm:"column:question_like_id"`
	PostId    int `json:"postId"`
	UserId    int `json:"UserId"`
	CreatedAt int `json:"CreatedAt" gorm:"column:question_like_created_at"`
	IsLike    int `json:"isLike" gorm:"column:question_like_is_like"`
}

func (PostLike) TableName() string {
	return "question_like"
}
