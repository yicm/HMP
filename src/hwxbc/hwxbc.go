// main project main.go
package main

import (
	"blog"
	"fmt"
)

func callBuildAll(input string, output string, prefix string, suffix string) {
	//fmt.Printf("%s, %s, %s\n", input, prefix, suffix)
	// load config
	blog.LoadConfig("./config", "page_config.json", &blog.Config)
	blog.LoadConfig("./config", "categories_config.json", &blog.CategoriesConfig)
	//fmt.Println(blog.Config.CategoryPageInterval)
	//fmt.Println(blog.Config.TagCategoryPageInterval)
	//fmt.Println(blog.Config.YearPageInterval)
	//fmt.Println(blog.CategoriesConfig["职场"].Title)
	// data preprocessing
	files := blog.GetAllFiles(input)
	dataProcessPool := blog.NewSimplePoll(20)
	for _, file := range files {
		dataProcessPool.Add(blog.HandleBlogMeta(input + "./" + file))
	}
	dataProcessPool.Run()
	// sort and output
	length := 0
	var metas []blog.TMeta
	blog.BuildMetaMapBuf.Range(func(_, v interface{}) bool {
		metas = append(metas, v.(blog.TMeta))
		length++
		return true
	})
	//fmt.Println("meta buffer size: ", length)
	// for _, member := range metas {
	// 	fmt.Println(member)
	// }

	//fmt.Println("-----------------------")
	output_path := "../source/_data/target/"
	if len(output) > 0 {
		output_path = output
	}
	blog.OutputPages(blog.Build_Mode_All, output_path, metas)
	fmt.Println("Build success!")
}

func callBuildType(build_type string, input string, output string, prefix string, suffix string) {
	fmt.Printf("%s, %s, %s, %s\n", build_type, input, prefix, suffix)
}

func callCreateArticle(categories string, output string, tags string, author string, title string) {
	key := blog.GenKey()
	meta := blog.NewDefaultTMeta()
	meta.SetKey(&meta, key)
	meta.SetTags(&meta, tags)
	meta.SetCategories(&meta, categories)
	meta.SetAuthor(&meta, author)
	meta.SetTitle(&meta, title)
	output_path := "../source/_posts/"
	if len(output) > 0 {
		output_path = output
	}
	err := blog.CreateMetaFile(&meta, output_path, key+".md")
	if err == nil {
		colorTip("Create Success!")
	}
}

func callCleanBuild() {
	fmt.Println("call clean build")
}

func main() {
	cliInit()
}
