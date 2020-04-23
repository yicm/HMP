package blog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "reflect"
	"sort"
	"strconv"
	"strings"
)

/*
time/top/pages
*/
func sortByDateAndTop(metas []TMeta) {
	sort.SliceStable(metas, func(i, j int) bool {
		mi, mj := metas[i], metas[j]
		switch {
		case mi.Date != mj.Date:
			if mi.Top != mj.Top {
				return mi.Top > mj.Top
			} else {
				return mi.Date > mj.Date
			}
		default:
			return mi.Top > mj.Top
		}
	})
}

func sortByTCategoryConfig(items []TCategoryConfig) {
	// var items interface{}
	// switch input.(type) {
	// case []TKeyCountItem:
	// 	fmt.Printf("type []blog.TKeyCountItem")
	// 	items := input.([]TKeyCountItem)
	// case []TCategoryConfig:
	// 	fmt.Printf("type []blog.TCategoryConfig")
	// 	items := input.([]TCategoryConfig)
	// default:
	// 	return errors.New("do not support this type for sorting")
	// }

	sort.SliceStable(items, func(i, j int) bool {
		mi, mj := items[i], items[j]
		switch {
		case mi.Count != mj.Count:
			return mi.Count > mj.Count
		default:
			return mi.Count > mj.Count
		}
	})
}

func sortByTKeyCountItem(items []TKeyCountItem) {
	// var items interface{}
	// switch input.(type) {
	// case []TKeyCountItem:
	// 	fmt.Printf("type []blog.TKeyCountItem")
	// 	items := input.([]TKeyCountItem)
	// case []TCategoryConfig:
	// 	fmt.Printf("type []blog.TCategoryConfig")
	// 	items := input.([]TCategoryConfig)
	// default:
	// 	return errors.New("do not support this type for sorting")
	// }

	sort.SliceStable(items, func(i, j int) bool {
		mi, mj := items[i], items[j]
		switch {
		case mi.Count != mj.Count:
			return mi.Count > mj.Count
		default:
			return mi.Count > mj.Count
		}
	})
}

func outputTimeTopPages(path string, pageNum int, metas []TMeta) {
	homeRootDir := path + "./" + Target_Blog + "./" + Target_Blog_Time_Top_Pages + "./"
	doPages(homeRootDir, "", pageNum, metas)
	/*
		sortByDateAndTop(metas)
		var totalPages int
		if len(metas)%pageNum == 0 {
			totalPages = len(metas) / pageNum
		} else {
			totalPages = len(metas)/pageNum + 1
		}
		var pageMetas []TMeta
		homeRootDir := path + "./" + Target_Blog + "./" + Target_Blog_Time_Top_Pages + "./"
		for i, j := 0, 0; i < len(metas); i, j = i+pageNum, j+1 {
			if j == totalPages-1 {
				// last page
				pageMetas = metas[i:]
				pageInfo := TPageInfo{PageNum: totalPages}
				j, _ := json.MarshalIndent(&pageInfo, "", "    ")
				_ = ioutil.WriteFile(homeRootDir+"page.json", j, 0644)
			} else {
				pageMetas = metas[i : i+pageNum]
			}
			if len(pageMetas) > 0 {
				data, err := json.MarshalIndent(pageMetas, "", "    ")
				if err != nil {
					fmt.Printf("json.marshal failed, err:", err)
					return
				}

				CreateDirs(homeRootDir)
				_ = ioutil.WriteFile(homeRootDir+strconv.FormatInt(int64(j), 10)+".json", data, 0644)
			}
		}
	*/
}

func doPages(rootDir string, subDir string, pageNum int, metas []TMeta) {
	sortByDateAndTop(metas)
	var totalPages int
	if len(metas)%pageNum == 0 {
		totalPages = len(metas) / pageNum
	} else {
		totalPages = len(metas)/pageNum + 1
	}
	var pageMetas []TMeta

	//fmt.Println(subDir, len(metas))
	for i, j := 0, 0; i < len(metas); i, j = i+pageNum, j+1 {
		if j == totalPages-1 {
			CreateDirs(rootDir + "./" + subDir)
			// last page
			pageMetas = metas[i:]
			pageInfo := TPageInfo{PageNum: totalPages}
			j, _ := json.MarshalIndent(&pageInfo, "", "    ")
			_ = ioutil.WriteFile(rootDir+"./"+subDir+"./"+"page.json", j, 0644)
		} else {
			pageMetas = metas[i : i+pageNum]
		}
		if len(pageMetas) > 0 {
			data, err := json.MarshalIndent(pageMetas, "", "    ")
			if err != nil {
				fmt.Printf("json.marshal failed, err: %s", err)
				return
			}

			CreateDirs(rootDir + "./" + subDir)
			_ = ioutil.WriteFile(rootDir+"./"+subDir+"./"+strconv.FormatInt(int64(j), 10)+".json", data, 0644)
		}
	}
}

/*
time/top&category/pages
*/
var buildTagsMapBuf = make(map[string]int)
var buildCategoriesMapBuf = make(map[string]int)
var buildTagsCategoriesMapBuf = make(map[string]int)
var buildYearMapBuf = make(map[string]int)
var buildOpenSourceMapBuf = make(map[string]int)

func updateCategoryesJsonContent() {
	for category, count := range buildCategoriesMapBuf {
		if _, ok := CategoriesConfig[category]; ok {
			// 种类已存在，则更新该种类下文章数目
			CategoriesConfig[category].Count = count
		} else {
			item := NewDefaultTCategoryConfig()
			item.Count = count
			item.Category = category
			item.Alias = category
			CategoriesConfig[category] = &item
		}
	}
}

func getTagsCategoriesNumStatistics(mode string, path string, metas []TMeta) {
	for _, meta := range metas {
		for _, tag := range meta.Tags {
			if Build_Mode_Tag == mode {
				buildTagsMapBuf[tag] += 1
			} else if Build_Mode_All == mode {
				buildTagsCategoriesMapBuf[tag] += 1
			} else if Build_Mode_Open_Source == mode {
				if strings.Index(tag, "开源") >= 0 {
					buildOpenSourceMapBuf[tag] += 1
				} else {
					// nothing to do
				}

			}
		}
		for _, category := range meta.Categories {
			if Build_Mode_Category == mode {
				buildCategoriesMapBuf[category] += 1
			} else if Build_Mode_All == mode {
				buildTagsCategoriesMapBuf[category] += 1
			} else if Build_Mode_Open_Source == mode {
				// 仅以tag作为“开源”项目标志
				// nothing to do
			}
		}
		if Build_Mode_Year == mode {
			year := strings.Split(strings.Trim(meta.Date, " "), "-")[0]
			buildYearMapBuf[year] += 1
		}
	}

	/*
		for k, v := range BuildTagsMapBuf {
			fmt.Println(k, v)
		}
		for k, v := range BuildCategoriesMapBuf {
			fmt.Println(k, v)
		}
	*/
	// output json file: tags.json/categories.json
	if Build_Mode_Tag == mode {
		tagsRootDir := path + "./" + Target_Blog + "./" + Target_Blog_Time_Tags_Pages + "./"
		CreateDirs(tagsRootDir)
		// change map[string]int to []{}
		var tagsOutputInfo []TKeyCountItem
		index := 0
		for k, count := range buildTagsMapBuf {
			tempItem := TKeyCountItem{
				ID:    index,
				Key:   k,
				Count: count,
			}
			index++
			tagsOutputInfo = append(tagsOutputInfo, tempItem)
		}
		// sort by count
		sortByTKeyCountItem(tagsOutputInfo)

		dataTags, err := json.MarshalIndent(tagsOutputInfo, "", "    ")
		if err != nil {
			fmt.Printf("json.marshal failed, err: %s", err)
			return
		}
		_ = ioutil.WriteFile(tagsRootDir+"tags.json", dataTags, 0644)
	}
	if Build_Mode_Category == mode {
		categoriesRootDir := path + "./" + Target_Blog + "./" + Target_Blog_Time_Categories_Pages + "./"
		CreateDirs(categoriesRootDir)
		dataCategories, err := json.MarshalIndent(buildCategoriesMapBuf, "", "    ")
		if err != nil {
			fmt.Printf("json.marshal failed, err: %s", err)
			return
		}

		_ = ioutil.WriteFile(categoriesRootDir+"categories.json", dataCategories, 0644)

		// 更新category计数
		updateCategoryesJsonContent()
		// 导出原始数据
		categoriesInfo, err := json.MarshalIndent(CategoriesConfig, "", "    ")
		if err != nil {
			fmt.Printf("json.marshal failed, err: %s", err)
			return
		}
		_ = ioutil.WriteFile(categoriesRootDir+"categories_info.json", categoriesInfo, 0644)
		// 将CategoriesConfig中数据转化为[]TCategoryConfig,作为最后使用数据
		var categoriesConfigOutput []TCategoryConfig
		for _, data := range CategoriesConfig {
			categoriesConfigOutput = append(categoriesConfigOutput, *data)
		}
		if nil != categoriesConfigOutput {
			// sort by count
			sortByTCategoryConfig(categoriesConfigOutput)
			categoriesInfoOutput, err := json.MarshalIndent(categoriesConfigOutput, "", "    ")
			if err != nil {
				fmt.Printf("json.marshal failed, err: %s", err)
				return
			}
			_ = ioutil.WriteFile(categoriesRootDir+"categories_info_output.json", categoriesInfoOutput, 0644)
		}
	}
	if Build_Mode_All == mode {
		tagsCategoriesRootDir := path + "./" + Target_Blog + "./" + Target_Blog_Time_Tags_Categories_Pages + "./"
		CreateDirs(tagsCategoriesRootDir)

		// change map[string]int to []{}
		var tagsCategoriesOutputInfo []TKeyCountItem
		index := 0
		for k, count := range buildTagsCategoriesMapBuf {
			tempItem := TKeyCountItem{
				ID:    index,
				Key:   k,
				Count: count,
			}
			index++
			tagsCategoriesOutputInfo = append(tagsCategoriesOutputInfo, tempItem)
		}
		// sort by count
		sortByTKeyCountItem(tagsCategoriesOutputInfo)

		dataAll, err := json.MarshalIndent(tagsCategoriesOutputInfo, "", "    ")
		if err != nil {
			fmt.Printf("json.marshal failed, err: %s", err)
			return
		}
		_ = ioutil.WriteFile(tagsCategoriesRootDir+"labels.json", dataAll, 0644)
	}

	if Build_Mode_Year == mode {
		yearRootDir := path + "./" + Target_Blog + "./" + Target_Blog_Year_Time_Pages + "./"
		CreateDirs(yearRootDir)
		dataAll, err := json.MarshalIndent(buildYearMapBuf, "", "    ")
		if err != nil {
			fmt.Printf("json.marshal failed, err: %s", err)
			return
		}
		_ = ioutil.WriteFile(yearRootDir+"years.json", dataAll, 0644)
	}

	if Build_Mode_Open_Source == mode {
		opensourceRootDir := path + "./" + Target_Blog + "./" + Target_Blog_Time_Top_OpenSource + "./"
		CreateDirs(opensourceRootDir)
		dataOpensource, err := json.MarshalIndent(buildOpenSourceMapBuf, "", "    ")
		if err != nil {
			fmt.Printf("json.marshal failed, err: %s", err)
			return
		}
		_ = ioutil.WriteFile(opensourceRootDir+"opensource.json", dataOpensource, 0644)
	}
}

func getTimeLabelMetas(mode string, label string, metas []TMeta) []TMeta {
	var output []TMeta
	for _, v := range metas {
		if Build_Mode_Tag == mode {
			if Contain(label, v.Tags) {
				output = append(output, v)
			}
		} else if Build_Mode_Category == mode {
			if Contain(label, v.Categories) {
				output = append(output, v)
			}
		} else if Build_Mode_All == mode {
			// TODO: tag & category重叠，需根据同id去重
			if Contain(label, v.Tags) {
				output = append(output, v)
			}
			if Contain(label, v.Categories) {
				output = append(output, v)
			}
		} else if Build_Mode_Year == mode {
			// BUG
			year := strings.Split(strings.Trim(v.Date, " "), "-")[0]
			if label == year {
				output = append(output, v)
			}
		} else if Build_Mode_Open_Source == mode {
			if Contain(label, v.Tags) {
				output = append(output, v)
			}
		}
	}
	return output
}

func outputTimeTagPages(path string, pageNum int, metas []TMeta) {
	getTagsCategoriesNumStatistics(Build_Mode_Tag, path, metas)
	for k, _ := range buildTagsMapBuf {
		tagMetas := getTimeLabelMetas(Build_Mode_Tag, k, metas)
		// create tag directory
		tagDir := path + "./" + Target_Blog + "./" + Target_Blog_Time_Tags_Pages
		doPages(tagDir, k, pageNum, tagMetas)
	}
}

func outputTimeCategoryPages(path string, pageNum int, metas []TMeta) {
	getTagsCategoriesNumStatistics(Build_Mode_Category, path, metas)
	for k, _ := range buildCategoriesMapBuf {
		categoryMetas := getTimeLabelMetas(Build_Mode_Category, k, metas)
		fmt.Println("==========")
		// create tag directory
		categoryDir := path + "./" + Target_Blog + "./" + Target_Blog_Time_Categories_Pages
		doPages(categoryDir, k, pageNum, categoryMetas)
		fmt.Println("----------")
	}
}

func outputTimeAllLabelPages(path string, pageNum int, metas []TMeta) {
	getTagsCategoriesNumStatistics(Build_Mode_All, path, metas)
	for k, _ := range buildTagsCategoriesMapBuf {
		labelMetas := getTimeLabelMetas(Build_Mode_All, k, metas)
		newLabelMetas := removeRepeatMeta(labelMetas)
		// create label directory
		labelDir := path + "./" + Target_Blog + "./" + Target_Blog_Time_Tags_Categories_Pages
		doPages(labelDir, k, pageNum, newLabelMetas)
	}
}

func outputYearTimePages(path string, pageNum int, metas []TMeta) {
	getTagsCategoriesNumStatistics(Build_Mode_Year, path, metas)
	for k, _ := range buildYearMapBuf {
		yearMetas := getTimeLabelMetas(Build_Mode_Year, k, metas)
		labelDir := path + "./" + Target_Blog + "./" + Target_Blog_Year_Time_Pages
		doPages(labelDir, k, pageNum, yearMetas)
	}
}

func outputOpenSourcePages(path string, pageNum int, metas []TMeta) {
	getTagsCategoriesNumStatistics(Build_Mode_Open_Source, path, metas)
	for k, _ := range buildOpenSourceMapBuf {
		opensourceMetas := getTimeLabelMetas(Build_Mode_Open_Source, k, metas)
		labelDir := path + "./" + Target_Blog + "./" + Target_Blog_Time_Top_OpenSource
		doPages(labelDir, k, pageNum, opensourceMetas)
	}
}

func OutputPages(mode string, path string, metas []TMeta) {
	switch mode {
	case Build_Mode_All:
		//fmt.Println(Config.HomePageInterval)
		outputTimeTopPages(path, Config.HomePageInterval, metas)
		// tag & category
		outputTimeAllLabelPages(path, Config.TagCategoryPageInterval, metas)
		outputTimeTagPages(path, Config.TagPageInterval, metas)
		outputTimeCategoryPages(path, Config.CategoryPageInterval, metas)
		// split blog with year directory
		outputYearTimePages(path, Config.YearPageInterval, metas)
		outputOpenSourcePages(path, Config.OpenSourceProjectInterval, metas)
	case Build_Mode_Tag:
		outputTimeTagPages(path, Config.TagPageInterval, metas)
	case Build_Mode_Category:
		outputTimeCategoryPages(path, Config.CategoryPageInterval, metas)
	default:
		outputTimeAllLabelPages(path, Config.TagCategoryPageInterval, metas)
	}
}
