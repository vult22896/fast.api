package models

type TopicLike struct {
	Id         int `json:"id" gorm:"column:qa_topic_like_id"`
	UserId     int `json:"user_id"`
	TopicId    int `json:"topic_id" gorm:"column:qa_topic_id"`
	IsLike     int `json:"is_like" gorm:"column:qa_topic_like_is_like"`
	TimeAction int `json:"time_action" gorm:"column:qa_topic_like_created_at"`
}

func (TopicLike) TableName() string {
	return "qa_topic_like"
}
