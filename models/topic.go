package models

type Topic struct {
	ID             int    `json:"id" gorm:"column:qa_topic_id"`
	IsParent       int    `json:"isParent" gorm:"column:qa_topic_is_parent"`
	ParentId       int    `json:"ParentId" gorm:"column:qa_topic_parent_id"`
	DeepLevel      int    `json:"deepLevel" gorm:"column:qa_topic_deep_level"`
	Name           string `json:"name" gorm:"column:qa_topic_name"`
	Desc           string `json:"desc" gorm:"column:qa_topic_desc"`
	NumberPost     int    `json:"numberPost" gorm:"column:qa_topic_number_question"`
	NumberAnswer   int    `json:"numberAnswer" gorm:"column:qa_topic_number_answer"`
	NumberLike     int    `json:"numberLike" gorm:"column:qa_topic_number_like"`
	Position       int    `json:"position" gorm:"column:position"`
	IsPromoting    int    `json:"isPromoting" gorm:"column:is_promoting"`
	Status         int    `json:"status" gorm:"column:qa_topic_status"`
	Image          string `json:"image" gorm:"column:qa_topic_image"`
	SubImage       string `json:"subImage" gorm:"column:qa_topic_sub_image"`
	BrandId        int    `json:"brandId" gorm:"column:brand_id"`
	BibaboId       int    `json:"bibaboId" gorm:"column:bibabo_id"`
	BibaboCourseId int    `json:"bibaboCourseId" gorm:"column:bibabo_course_id"`
	Type           int    `json:"type" gorm:"column:qa_topic_type"`
	IsHighlight    int    `json:"isHighlight" gorm:"column:qa_topic_is_highlight"`
}

func (Topic) TableName() string {
	return "qa_topic"
}
