package models

type Post struct {
	ID               int    `json:"id" gorm:"column:question_id"`
	Title            string `json:"title" gorm:"column:question_question"`
	Description      string `json:"description" gorm:"column:question_desc"`
	PostType         int    `json:"postType" gorm:"column:post_type"`
	UserId           int    `json:"userId" gorm:"column:user_id"`
	User             User   `json:"user"`
	QuestionCreated  int    `json:"createdAt" gorm:"column:question_created_at"`
	Status           int    `json:"status" gorm:"column:question_status"`
	Sensitive        int    `json:"sensitive" gorm:"column:question_sensitive"`
	Anonymous        int    `json:"anonymous" gorm:"column:question_anonymous"`
	BibaboId         int    `json:"bibaboId" gorm:"column:bibabo_id"`
	CourseVideoId    int    `json:"courseVideoId" gorm:"column:course_video_id"`
	BibaboCourseId   int    `json:"bibaboCourseId" gorm:"column:bibabo_course_id"`
	BibaboCourseName string `json:"bibaboCourseName" gorm:"column:bibabo_course_name"`
	QaClubId         int    `json:"qaClubId" gorm:"column:qa_club_id"`
	TypeContent      int    `json:"typeContent" gorm:"column:type_content"`
}

func (Post) TableName() string {
	return "question"
}
