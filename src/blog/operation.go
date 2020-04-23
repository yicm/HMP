package blog

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

var timeMap = make(map[string]int)

func CreateMetaFile(meta *TMeta, path string, filename string) error {
	CreateDirs(path)
	fullname := path + "./" + filename
	f, err := os.Create(fullname)
	if err != nil {
		log.Printf("create file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	// convert data
	mp, err := TMeta2Map(meta)
	if err != nil {
		log.Printf("meta to map error: %v\n", err)
		return err
	}

	// sort data by key
	var keys []string
	for k := range mp {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// write data
	fmt.Fprintln(w, "---")
	for _, k := range keys {
		lineStr := ""
		if k == T_Tags || k == T_Categories || k == T_SyncLinks {
			lineKStr := fmt.Sprintf("%s:", k)
			fmt.Fprintln(w, lineKStr)
			values := strings.Split(mp[k], ",")

			for _, v := range values {
				if len(v) > 0 {
					lineVStr := fmt.Sprintf("- %s", v)
					fmt.Fprintln(w, lineVStr)
				}
			}
		} else {
			lineStr = fmt.Sprintf("%s: %s", k, mp[k])
			fmt.Fprintln(w, lineStr)
		}
	}
	fmt.Fprintln(w, "---")

	return w.Flush()
}

func GenKey() TKey {
	t := time.Now()
	return t.Format("20060102150405")
}

func GetAllFiles(path string) []string {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic("failed to read directory: " + path)
	}

	var names []string
	for _, f := range files {
		if f.IsDir() {
			//fmt.Println("dir: ", f.Name())
			// nothing to do
		} else {
			//fmt.Println("file: ", f.Name())
			names = append(names, f.Name())
		}
	}

	return names
}

func GetMeta(filename string) (key string, meta TMeta) {
	suffix := strings.HasSuffix(filename, ".md")
	if !suffix {
		return
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Cannot open text file: %s, err: [%v]", filename, err)
		panic("failed to open file: " + filename)
	}
	defer file.Close()

	metaMap := make(map[string]string)
	scanner := bufio.NewScanner(file)
	endLineCount := 0
	isMultiValue := false
	multiValueKey := ""
	for scanner.Scan() {
		line := scanner.Text()
		if T_SplitLine == line {
			endLineCount += 1
			//fmt.Println("split line", endLineCount)
		}
		if 2 == endLineCount {
			break
		}

		if T_SplitLine != line {
			//fmt.Printf("%s\n", line)
			splitRes := strings.Split(strings.Trim(line, " "), ":")
			//fmt.Printf("%v, len=%v\n", splitRes, len(splitRes))
			if len(splitRes) == 1 && isMultiValue {
				if splitRes[0][0:1] == "-" {
					//fmt.Println("heiheihei")
					if len(metaMap[multiValueKey]) > 0 {
						metaMap[multiValueKey] = metaMap[multiValueKey] + "," + strings.Trim(splitRes[0][1:], " ")
						//fmt.Println("mutil v:", metaMap[multiValueKey])
					} else {
						metaMap[multiValueKey] = strings.Trim(splitRes[0][1:], " ")
						//fmt.Println("mutil v1:", metaMap[multiValueKey])
					}
				} else {
					isMultiValue = false
				}
			}
			if len(splitRes) == 2 && len(strings.Trim(splitRes[1], " ")) == 0 {
				if strings.Trim(splitRes[0], " ") == T_Categories ||
					strings.Trim(splitRes[0], " ") == T_Tags ||
					strings.Trim(splitRes[0], " ") == T_SyncLinks {
					//fmt.Println("is multi value")
					isMultiValue = true
					multiValueKey = strings.Trim(splitRes[0], " ")
					metaMap[multiValueKey] = ""
				} else {
					//fmt.Println("key  = ", splitRes[0])
					metaMap[splitRes[0]] = strings.Trim(splitRes[1], " ")
				}
			} else if T_Cover == splitRes[0] || T_Permalink == splitRes[0] {
				metaMap[splitRes[0]] = strings.Trim(splitRes[1]+":"+splitRes[2], " ")
				isMultiValue = false
			} else if T_Date == splitRes[0] || T_Updated == splitRes[0] {
				metaMap[splitRes[0]] = strings.Trim(splitRes[1]+":"+splitRes[2]+":"+splitRes[3], " ")
				isMultiValue = false
			} else if len(splitRes) != 1 {
				metaMap[splitRes[0]] = strings.Trim(splitRes[1], " ")
				isMultiValue = false
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Cannot scanner text file: %s, err: [%v]", filename, err)
		panic("cannot scanner blog file :" + filename)
	}

	key = metaMap[T_Key]
	meta = Map2TMeta(metaMap)

	return
}

var BuildMetaMapBuf sync.Map

func HandleBlogMeta(filename string) func() {
	return func() {
		//fmt.Println("==handle ", filename)
		key, meta := GetMeta(filename)
		//fmt.Println(meta.WxEnable)
		if meta.WxEnable {
			BuildMetaMapBuf.Store(key, meta)
		} else {
			fmt.Println(filename, " disable in mini-program or not markdown file.")
		}
	}
}
