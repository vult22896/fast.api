package mongo_models

import "gopkg.in/mgo.v2/bson"

type Post struct {
	_ID                      bson.ObjectId `bson:"_idomitempty" json:"_id"`
	ID                       int           `json:"id"`
	Title                    string        `json:"title"`
	Content                  string        `json:"content"`
	ContentHeadings          []string      `json:"contentHeadings"`
	NumberLike               int           `json:"numberLike"`
	NumberComment            int           `json:"numberComment"`
	NumberSaved              int           `json:"numberSaved"`
	User                     User          `json:"user"`
	Topics                   []Topic       `json:"topics"`
	Products                 []Product     `json:"products"`
	Images                   []Image       `json:"images"`
	Sensitive                int           `json:"sensitive"`
	Type                     int           `json:"type"`
	ContentType              int           `json:"ContentType"`
	Url                      string        `json:"url"`
	Videos                   []string      `json:"videos"`
	CreatedAt                int           `json:"createdAt"`
	QuestionAnonymous        int           `json:"questionAnonymous"`
	Color                    []string      `json:"color"`
	BackgroundId             int           `json:"backgroundId"`
	AlgoVersion              int           `json:"algoVersion"`
	Group                    Group         `json:"group"`
	QuestionableType         string        `json:"questionableType"`
	QuestionableId           int           `json:"questionableId"`
	ExtraData                string        `json:"extraData"`
	CommunityPointByQuestion int           `json:"communityPointByQuestion"`
	Source                   int           `json:"source"`
	EditMode                 int           `json:"editMode"`
	SeoTitle                 string        `json:"seoTitle"`
	ShortDescription         string        `json:"shortDescription"`
	Sapo                     string        `json:"sapo"`
	MetaDescription          string        `json:"metaDescription"`
	MetaKeywords             string        `json:"metaKeywords"`
	Comments                 []Comment     `json:"comments"`
	StageFilterValueId       int           `json:"stageFilterValueId"`
	IsLike                   int           `json:"isLike"`
	IsSave                   int           `json:"isSave"`
}

type User struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	Avatar         string         `json:"avatar"`
	Type           int            `json:"type"`
	Url            string         `json:"url"`
	CommunityLevel CommunityLevel `json:"communityLevel" bson:"communityLevel"`
	IsFollow       int            `json:"isFollow"`
}

type CommunityLevel struct {
	Name             string `json:"name"`
	Image            string `json:"image"`
	Point            int    `json:"point"`
	PointMin         int    `json:"pointMin"`
	PointMax         int    `json:"pointMax"`
	PointToNextLevel int    `json:"pointToNextLevel"`
	UserRank         int    `json:"userRank"`
}

type Topic struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsFollow int    `json:"isFollow"`
}

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Promotion Promotion `json:"promotion"`
}

type Promotion struct {
	InfoBestPrice struct {
		BestPrice        int `json:"bestPrice"`
		PromotionElement struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Type        int    `json:"type"`
			PromotionId int    `json:"promotionId"`
			StartAt     int    `json:"startAt"`
			EndAt       int    `json:"endAt"`
			Banner      string `json:"banner"`
			WaterMark   string `json:"waterMark"`
		} `json:"promotionElement"`
	} `json:"infoBestPrice"`

	InfoBulk   []InfoBulk   `json:"infoBulk"`
	InfoBundle []InfoBundle `json:"infoBundle"`
	InfoGift   []InfoGift   `json:"infoGift"`
}

type InfoGift struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	GiftPackage []struct {
		ID       int `json:"id"`
		Products []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Image       string `json:"image"`
			Price       int    `json:"price"`
			ImportPrice int    `json:"importPrice"`
			OriginPrice int    `json:"originPrice"`
			NumberItem  int    `json:"numberItem"`
		} `json:"products"`
	} `json:"giftPackage"`
}

type InfoBundle struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Products []struct {
		ID             int      `json:"id"`
		Name           string   `json:"name"`
		Image          string   `json:"image"`
		Price          int      `json:"price"`
		NumberApplyMin int      `json:"numberApplyMin"`
		OriginPrice    int      `json:"originPrice"`
		Number         int      `json:"number"`
		Variants       []string `json:"variants"`
	} `json:"products"`
}

type InfoBulk struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Product struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Image          string `json:"image"`
		Price          int    `json:"price"`
		OriginPrice    int    `json:"originPrice"`
		ImportPrice    int    `json:"importPrice"`
		NumberApplyMin int    `json:"numberApplyMin"`
	} `json:"product"`
}

type Group struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsFollow int    `json:"isFollow"`
}

type Comment struct {
	ID                     int    `json:"id"`
	Content                string `json:"content"`
	Image                  string `json:"image"`
	NumberLike             int    `json:"numberLike"`
	NumberReply            int    `json:"numberReply"`
	CommunityPointByAnswer int    `json:"communityPointByAnswer"`
	User                   User   `json:"user"`
}
type Image struct {
	Image string  `json:"image"`
	Ratio float64 `json:"ratio"`
}
