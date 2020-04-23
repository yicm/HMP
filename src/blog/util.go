package blog

import (
	"errors"
	_ "fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CreateDirs(path string) error {
	if !pathExists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		return err
	}
	return nil
}

func slice2Str(slice []string) (ret string) {
	ret = ""
	for i := 0; i < len(slice); i++ {
		if i == (len(slice) - 1) {
			ret += slice[i]
		} else {
			ret += slice[i] + ","
		}
	}
	return
}

func Contain(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			//fmt.Println("比对", targetValue.Index(i).Interface(), obj)
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

func InitLogger() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func TMeta2Map(meta *TMeta) (map[string]string, error) {
	tt := reflect.TypeOf(*meta)
	e := reflect.ValueOf(meta).Elem()

	mp := make(map[string]string)
	for i := 0; i < tt.NumField(); i++ {
		fieldType := tt.Field(i)
		if "uint32" == fieldType.Type.Name() {
			// uint32
			mp[fieldType.Tag.Get("json")] = strconv.FormatUint(e.Field(i).Uint(), 10)
		} else if "string" == fieldType.Type.Name() {
			// string
			mp[fieldType.Tag.Get("json")] = e.Field(i).String()
		} else if "bool" == fieldType.Type.Name() {
			// bool
			value := "false"
			if e.Field(i).Bool() {
				value = "true"
			}
			mp[fieldType.Tag.Get("json")] = value
		} else {
			// slice
			slice, ok := e.Field(i).Interface().([]string)
			if !ok {
				return mp, errors.New("value not a []string")
			}
			mp[fieldType.Tag.Get("json")] = slice2Str(slice)
		}
	}

	return mp, nil
}

func Map2TMeta(mp map[string]string) TMeta {
	meta := NewDefaultTMeta()
	for k, v := range mp {
		switch k {
		case T_Layout:
			meta.Layout = v
		case T_Key:
			meta.Key = v
		case T_Title:
			meta.Title = v
		case T_Tags:
			tags := strings.Split(v, ",")
			meta.Tags = tags
		case T_Categories:
			categories := strings.Split(v, ",")
			meta.Categories = categories
		case T_Top:
			top, _ := strconv.ParseUint(v, 10, 32)
			meta.Top = uint32(top)
		case T_Description:
			meta.Description = v
		case T_Permalink:
			meta.Permalink = v
		case T_SyncLinks:
			synclinks := strings.Split(v, ",")
			meta.SyncLinks = synclinks
		case T_Cover:
			meta.Cover = v
		case T_ContentFirstImg:
			meta.ContentFirstImg = v
		case T_PostBgImg:
			meta.PostBgImg = v
		case T_Author:
			meta.Author = v
		case T_Date:
			meta.Date = v
		case T_Updated:
			meta.Updated = v
		case T_Keywords:
			meta.Keywords = v
		case T_Password:
			meta.Password = v
		case T_PasswordDesc:
			meta.PasswordDesc = v
		case T_Label:
			meta.Label = v
		case T_WxEnable:
			if v == "false" {
				meta.WxEnable = false
			} else {
				meta.WxEnable = true
			}
		case T_Gank:
			if v == "false" {
				meta.Gank = false
			} else {
				meta.Gank = true
			}
		case T_Limit:
			if v == "false" {
				meta.Limit = false
			} else {
				meta.Limit = true
			}
		case T_Donate:
			if v == "false" {
				meta.Donate = false
			} else {
				meta.Donate = true
			}
		}

	}
	return meta
}
