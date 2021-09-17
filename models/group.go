package models

type Group struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	ImageCover   string `json:"imageCover"`
	ImageAvatar  string `json:"imageAvatar"`
	NumberMember int    `json:"numberMember"`
	NumberPost   int    `json:"numberPost"`
}
