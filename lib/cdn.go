package lib

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	stringHelper "fast.bibabo.vn/helpers"
	"github.com/joho/godotenv"
)

const FOLDER_DEFAULT_DEV = "dev/d/"
const FOLDER_DEFAULT = "def/"
const FOLDER_QUESTION = "qs/"
const FOLDER_TICKET = "tk/"
const FOLDER_CHAT = "ch/"
const FOLDER_USER = "us/"
const FOLDER_PRODUCT = "pr/"
const FOLDER_OTHER = "ot/"

const SIZE_TINY = 20
const SIZE_SMALLER = 50
const SIZE_SMALL = 70
const SIZE_MEDIUM = 100
const SIZE_LANGE = 200
const SIZE_LARGER = 300
const SIZE_BIG = 500
const SIZE_BIGER = 700
const SIZE_BIGGEST = 1500

const IMAGE_LAYOUT_PORTRAIT = "portrait"
const IMAGE_LAYOUT_LANDSCAPE = "landscape"

type Cdn interface {
	GetImage(image string, folder string, width int, height int) string
	GetHost() string
}
type cdn struct{}

func GetInstanceCdn() Cdn {
	return &cdn{}
}

func (c cdn) GetImage(image string, folder string, width int, height int) string {
	if len(image) == 0 {
		return ""
	}
	if image[:7] == "http://" && image[:8] == "https://" {
		return image
	}
	extention := stringHelper.After(image, ".")
	fileName := stringHelper.Before(image, ".")
	firstChar := image[:1]
	first2Char := image[:2]

	path := c.GetHost() + "/uploads/bo/vi/" + folder + firstChar + "/" + first2Char + "/" + fileName + "-" + strconv.Itoa(width) + "x" + strconv.Itoa(height) + "-resize." + extention
	return path
}

func (c cdn) GetHost() string {
	error := godotenv.Load()
	if error != nil {
		panic("Failed load env file")
	}
	videoEdgeServerUrl := os.Getenv("VIDEO_EDGE_SERVER_URL")
	parseSlice := strings.Split(videoEdgeServerUrl, ",")
	rand.Seed(time.Now().UnixNano())
	return parseSlice[rand.Intn(len(parseSlice))]
}
