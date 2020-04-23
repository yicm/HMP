package blog

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestStringsContain(t *testing.T) {
	testStr := "开源项目"
	if strings.Index(testStr, "开源") >= 0 {
		fmt.Println("包含")
	} else {
		fmt.Println("未包含")
	}
}

func TestSortByCount(t *testing.T) {
	meta := TCategoryConfig{
		ID:    0,
		Count: 2,
	}
	var input []TCategoryConfig
	input = append(input, meta)
	sortByTCategoryConfig(input)
}

func TestNewTMeta(t *testing.T) {
	meta := NewDefaultTMeta()

	// j, _ := json.MarshalIndent(&meta, "", "    ")
	// fmt.Println(string(j))
	tt := reflect.TypeOf(meta)
	fmt.Println(tt)
	e := reflect.ValueOf(&meta).Elem()

	for i := 0; i < tt.NumField(); i++ {
		fieldType := tt.Field(i)
		fmt.Printf("Type: %v Name: %v  Tag_json: %v Tag_bson: %v  Value: %v\n",
			fieldType.Type.Name(),
			fieldType.Name,
			fieldType.Tag.Get("json"),
			fieldType.Tag.Get("bson"),
			e.Field(i).Interface())
	}
}
