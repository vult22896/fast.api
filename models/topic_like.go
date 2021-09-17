package models

type TopicLike struct {
	Id         int `json:"id" gorm:"column:qa_topic_like_id"`
	UserId     int `json:"userId"`
	TopicId    int `json:"topicId" gorm:"column:qa_topic_id"`
	IsLike     int `json:"isLike" gorm:"column:qa_topic_like_is_like"`
	TimeAction int `json:"timeAction" gorm:"column:qa_topic_like_created_at"`
}

func (TopicLike) TableName() string {
	return "qa_topic_like"
}
