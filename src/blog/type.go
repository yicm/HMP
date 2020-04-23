package blog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type TLayout = string
type TCommentCtrl = string
type TDonateCtrl = bool
type TKey = string
type TTitle = string
type TTags = []string
type TCategories = []string
type TTop = uint32
type TDesc = string
type TPermalink = string
type TSyncLinks = []string
type TAuthor = string
type TImg = string
type TDate = string
type TLabel = string
type TKeywords = string
type TPassword = string
type TPasswordDesc = string
type TLimit = bool
type TWxEnable = bool
type TGank = bool

const (
	Build_Mode_Tag         string = "tag"
	Build_Mode_Category    string = "category"
	Build_Mode_All         string = "all"
	Build_Mode_Year        string = "year"
	Build_Mode_Open_Source string = "open_source"
)

const (
	Label_Original    string = "original"
	Label_Forward     string = "forward"
	Label_Translation string = "translation"
)

const (
	LAYOUT_POST  TLayout = "post"
	LAYOUT_PAGE  TLayout = "page"
	LAYOUT_DRAFT TLayout = "draft"
)

const (
	COMMENT_CTRL_DISABLE TCommentCtrl = "disable"
	COMMENT_CTRL_ENABLE  TCommentCtrl = "enable"
	COMMENT_CTRL_HIDE    TCommentCtrl = "hide"
	COMMENT_CTRL_NORMAL  TCommentCtrl = "normal"
	COMMENT_CTRL_TRUE    TCommentCtrl = "true"
	COMMENT_CTRL_FALSE   TCommentCtrl = "false"
)

const (
	Target_Blog                            string = ""
	Target_Blog_Time_Top_Pages             string = "blog_time_top_pages"
	Target_Blog_Time_Tags_Pages            string = "blog_time_tags_pages"
	Target_Blog_Time_Categories_Pages      string = "blog_time_categories_pages"
	Target_Blog_Time_Tags_Categories_Pages string = "blog_time_tags_categories_pages"
	Target_Blog_Year_Time_Pages            string = "blog_year_time_pages"
	Target_Blog_Time_Top_OpenSource        string = "blog_time_top_opensource_pages"
)

const (
	T_SplitLine       string = "---"
	T_Layout          string = "layout"
	T_Comments        string = "comments"
	T_Donate          string = "donate"
	T_Key             string = "key"
	T_Title           string = "title"
	T_Tags            string = "tags"
	T_Categories      string = "categories"
	T_Top             string = "top"
	T_Description     string = "description"
	T_Permalink       string = "permalink"
	T_SyncLinks       string = "sync_links"
	T_Cover           string = "cover"
	T_ContentFirstImg string = "content_first_img"
	T_PostBgImg       string = "post_bg_img"
	T_Author          string = "author"
	T_Date            string = "date"
	T_Updated         string = "updated"
	T_Label           string = "label"
	T_Keywords        string = "keywords"
	T_Password        string = "password"
	T_PasswordDesc    string = "password_desc"
	T_Limit           string = "limit"
	T_WxEnable        string = "wx_enable"
	T_Gank            string = "gank"
)

type TMeta struct {
	Layout          TLayout       `json:"layout" bson:"layout"`
	Comments        TCommentCtrl  `json:"comments" bson:"comments"`
	Donate          TDonateCtrl   `json:"donate" bson:"donate"`
	Key             TKey          `json:"key" bson:"key"`
	Title           TTitle        `json:"title" bson:"title"`
	Label           TLabel        `json:"label" bson:"label"`
	Tags            TTags         `json:"tags" bson:"tags"`
	Categories      TCategories   `json:"categories" bson:"categories"`
	Top             TTop          `json:"top" bson:"top"`
	Description     TDesc         `json:"description" bson:"description"`
	Permalink       TPermalink    `json:"permalink" bson:"permalink"`
	SyncLinks       TSyncLinks    `json:"sync_links" bson:"sync_links"`
	Cover           TImg          `json:"cover" bson:"cover"`
	ContentFirstImg TImg          `json:"content_first_img" bson:"content_first_img"`
	PostBgImg       TImg          `json:"post_bg_img" bson:"post_bg_img"`
	Author          TAuthor       `json:"author" bson:"author"`
	Date            TDate         `json:"date" bson:"date"`
	Updated         TDate         `json:"updated" bson:"updated"`
	Keywords        TKeywords     `json:"keywords" bson:"keywords"`
	Password        TPassword     `json:"password" bson:"password"`
	PasswordDesc    TPasswordDesc `json:"password_desc" bson:"password_desc"`
	Limit           TLimit        `json:"limit" bson:"limit"`
	WxEnable        TWxEnable     `json:"wx_enable" bson:"wx_enable"`
	Gank            TGank         `json:"gank" bson:"gank"`
}

type TConfig struct {
	HomePageInterval          int `json:"home_page_interval"`
	TagPageInterval           int `json:"tag_page_interval"`
	CategoryPageInterval      int `json:"category_page_interval"`
	TagCategoryPageInterval   int `json:"tag_category_page_interval"`
	YearPageInterval          int `json:"year_page_interval"`
	OpenSourceProjectInterval int `json:"open_source_project_interval"`
}

type TCategoryConfig struct {
	ID          string `json:"id"`
	Category    string `json:"category"`
	Alias       string `json:"alias"`
	Title       string `json:"title"`
	Count       int    `json:"count"`
	Show        bool   `json:"show"`
	Logo        string `json:"logo"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
}

type TPageInfo struct {
	PageNum int `json:"total_pages"`
}

type TKeyCountItem struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Count int    `json:"count"`
}

func removeRepeatMeta(input []TMeta) []TMeta {
	result := []TMeta{}
	tempMap := map[string]byte{}
	for _, e := range input {
		ll := len(tempMap)
		tempMap[e.Key] = 0
		if len(tempMap) != ll {
			result = append(result, e)
		}
	}
	return result
}

func NewDefaultTMeta() TMeta {
	return TMeta{
		Layout:          LAYOUT_POST,
		Comments:        COMMENT_CTRL_TRUE,
		Donate:          true,
		Key:             "",
		Title:           "",
		Label:           Label_Original,
		Tags:            []string{"Git", "C/C++"},
		Categories:      nil,
		Top:             0,
		Description:     "",
		Permalink:       "",
		SyncLinks:       nil,
		Cover:           "",
		ContentFirstImg: "",
		PostBgImg:       "",
		Author:          "Ethan",
		Date:            time.Now().Format("2006-01-02 15:04:05"),
		Updated:         time.Now().Format("2006-01-02 15:04:05"),
		Keywords:        "",
		Password:        "",
		PasswordDesc:    "Please input the password!",
		Limit:           true,
		WxEnable:        true,
		Gank:            false,
	}
}

func NewDefaultTCategoryConfig() TCategoryConfig {
	return TCategoryConfig{
		ID:          time.Now().Format("20060102150405"),
		Category:    "",
		Alias:       "",
		Title:       "",
		Count:       0,
		Show:        true,
		Logo:        "",
		Cover:       "https://ae01.alicdn.com/kf/H159be3b332024195917804fee84df459I.png",
		Description: "",
	}
}

func (*TMeta) SetKey(meta *TMeta, key TKey) {
	meta.Key = key
}

func (*TMeta) SetTags(meta *TMeta, tags string) {
	if len(tags) > 0 {
		tmp := strings.Split(tags, "，")
		if len(tmp) > 1 {
			panic("You can't use commas as splitters")
		}

		stags := strings.Split(tags, ",")
		for i, _ := range stags {
			stags[i] = strings.Trim(stags[i], " ")
		}

		meta.Tags = stags
	}
}

func (*TMeta) SetCategories(meta *TMeta, categories string) {
	if len(categories) > 0 {
		tmp := strings.Split(categories, "，")
		if len(tmp) > 1 {
			panic("You can't use commas as splitters")
		}
		scategories := strings.Split(categories, ",")
		for i, _ := range scategories {
			scategories[i] = strings.Trim(scategories[i], " ")
		}

		meta.Categories = scategories
	}
}

func (*TMeta) SetAuthor(meta *TMeta, author string) {
	meta.Author = strings.Trim(author, " ")
}

func (*TMeta) SetKeywords(meta *TMeta, keywords string) {
	meta.Keywords = strings.Trim(keywords, " ")
}

func (*TMeta) SetTitle(meta *TMeta, title string) {
	meta.Title = strings.Trim(title, " ")
}

func (*TMeta) SetPassword(meta *TMeta, password string) {
	meta.Password = strings.Trim(password, " ")
}

func (*TMeta) SetPasswordDesc(meta *TMeta, password_desc string) {
	meta.PasswordDesc = strings.Trim(password_desc, " ")
}

const (
	ERR_Params           string = "Parameters cannot be empty! Please check!"
	ERR_Params_Output    string = "Ouput parameters cannot be empty! Please check!"
	ERR_Params_Input     string = "Input directory cannot be empty! Please check!"
	ERR_Type_Not_Support string = "Do not support the build type: "
)

const (
	BUILD_Type_Time_Top_Page      string = "time_top_page"
	BUILD_Type_Time_Tag_Page      string = "time_tag_page"
	BUILD_Type_Time_Category_Page string = "time_category_page"
)

var BUILD_TYPES = []string{BUILD_Type_Time_Top_Page,
	BUILD_Type_Time_Tag_Page,
	BUILD_Type_Time_Category_Page}

var Config TConfig = TConfig{}
var CategoriesConfig map[string]*TCategoryConfig

func LoadConfig(path string, filename string, config interface{}) {
	file, err := ioutil.ReadFile(path + "./" + filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	_ = json.Unmarshal([]byte(file), config)
}
