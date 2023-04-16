package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type JsonContent struct {
	Data interface{}
}

func loadFileDemo(filename string) {
	jc, err := parsePoetryFile(filename)
	if err != nil {
		log.Fatal(err)
		return
	}

	iterContent(jc.Data)
}

func parsePoetryFile(jsonFile string) (*JsonContent, error) {
	raw, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}
	var data = new(interface{})
	err = json.Unmarshal(raw, data)
	if err != nil {
		return nil, err
	}

	return &JsonContent{*data}, nil
}

// https://github.com/ChimeraCoder/gojson
func iterContent(content interface{}) {
	switch content.(type) {
	case []interface{}:
		val := content.([]interface{})
		fmt.Printf("DBG: [%d] item\n", len(val))
		for _, child := range val {
			iterContent(child)
		}
	case map[interface{}]interface{}:
		val := content.(map[interface{}]interface{})
		for k, v := range val {
			fmt.Printf("Key: [%v]\n", k)
			iterContent(v)
		}
	case map[string]interface{}:
		val := content.(map[string]interface{})
		for k, v := range val {
			fmt.Printf("Key: [%v]\n", k)
			iterContent(v)
		}
	case string:
		fmt.Printf("sTr:[%s]\n", content.(string))
	case int:
		fmt.Printf("iNt:[%s]\n", content.(string))
	case float64:
		fmt.Printf("fLt:[%v]\n", content.(float64))
	default:
		// https://stackoverflow.com/questions/6372474/how-to-determine-an-interface-values-real-type
		otherType := reflect.TypeOf(content)
		fmt.Printf("Unknown: %+v\n", otherType)
	}
}

/*
{
    "title": "登驪山高頂寓目",
    "author": "李顯",
    "biography": "",
    "paragraphs": [
        "四郊秦漢國，八水帝王都。閶闔雄裏閈，城闕壯規模。",
        "貫渭稱天邑，含岐實奧區。金門披玉館，因此識皇圖。"
    ],
    "notes": [
        ""
    ],
    "volume": "卷二",
    "no#": 10
}
*/
type Poetry struct {
	Title      string   `json:"title"`
	Author     string   `json:"author"`
	Biography  string   `json:"biography"`
	Paragraphs []string `json:"paragraphs"`
	Notes      []string `json:"notes"`
	Volume     string   `json:"volume"`
	No         float64  `json:"no#"`
}

func (p *Poetry) getDesc() string {
	return fmt.Sprintf("[%s]%s: Vol [%s]:%v", p.Title, p.Author, p.Volume, p.No)
}

func (p *Poetry) convId() string {
	return fmt.Sprintf("%d_%d", ChineseToNumber(p.Volume), int(p.No))
}

func (p *Poetry) ComparePoem(poem *ChinesePoem) bool {
	if poem == nil {
		return false
	}

	alltext := ""
	for _, text := range p.Paragraphs {
		alltext += text
	}

	alltext = removePunctuation(alltext)
	poemtext := removePunctuation(poem.AllText)

	if alltext != poemtext {
		/*
			fmt.Println("[" + alltext)
			fmt.Println("--------------------")
			fmt.Println(poemtext + "]")
		*/
		fmt.Println(poem.ID + "|" + p.Title + "|" + compString(alltext, poemtext))
		fmt.Println("------------------------")
		return false
	}

	return true
}

func removePunctuation(s string) string {
	rs := []rune(s)
	rsnew := []rune{}

	for _, r := range rs {
		if !IsPunctuation(r) {
			rsnew = append(rsnew, r)
		}
	}

	return string(rsnew)
}

type Poetries struct {
	AllPoets []Poetry
}

func (p *Poetries) loadPoetryFile(filename string) {
	jsonFile, err := os.Open(filename)
	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var poetries []Poetry
	err = json.Unmarshal(byteValue, &poetries)
	if err != nil {
		log.Printf("[WARN]File: %s error: %v!\n", filename, err)
		return
	}

	if len(poetries) == 0 {
		return
	}

	//fmt.Printf("[INFO] %d poetries.\n", len(poetries))
	p.AllPoets = append(p.AllPoets, poetries...)
}

func (p *Poetries) loadPoetryPath(fullpath string) {
	err := filepath.Walk(fullpath, func(path1 string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path1) != ".json" {
			return nil
		}

		p.loadPoetryFile(path1)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[INFO]Total %d poetries.\n", len(p.AllPoets))
}

func (p *Poetries) comparePoemsOld() {
	id2poet := make(map[string]*Poetry)

	for _, pt := range p.AllPoets {
		id := pt.convId()

		id2poet[id] = &pt
		fmt.Printf("[DBG]ID: %s: Poetry: %s\n", id, pt.Title)
	}

	fmt.Printf("[INFO]Parsed %d poetries.\n", len(id2poet))

	var qtsInst Qts
	qtsInst.Init()
	qtsInst.ReadFile("qts_zht.txt")

	for id, poetry := range id2poet {
		fmt.Printf("[DBG]ID: %s: Poetry: %s\n", id, poetry.Title)
		chpoem := qtsInst.GetPoem(id)
		poetry.ComparePoem(chpoem)
	}
}

func (p *Poetries) comparePoems() {
	var qtsInst Qts
	qtsInst.Init()
	qtsInst.ReadFile("qts_zht.txt")

	samecount := 0
	for _, pt := range p.AllPoets {
		id := pt.convId()
		chpoem := qtsInst.GetPoem(id)

		if pt.ComparePoem(chpoem) {
			samecount++
		}
	}

	fmt.Printf("[INFO]Parsed %d poetries. Same: %d\n", len(p.AllPoets), samecount)
}

func LoadPoetries(fullpath string) {
	var pt Poetries
	pt.loadPoetryPath(fullpath)
	pt.comparePoems()
}

func ChineseToNumber(chnStr string) int {
	rs := []rune(chnStr)
	totalVal := 0
	curVal := 0
	for i := 0; i < len(rs); i++ {
		switch rs[i] {
		case '一':
			curVal = 1
		case '二':
			curVal = 2
		case '三':
			curVal = 3
		case '四':
			curVal = 4
		case '五':
			curVal = 5
		case '六':
			curVal = 6
		case '七':
			curVal = 7
		case '八':
			curVal = 8
		case '九':
			curVal = 9
		case '零', '卷':
			curVal = 0
		case '百':
			{
				if curVal > 0 {
					totalVal += curVal * 100
					curVal = 0
				} else {
					log.Printf("[INFO]Format error 1: %s!\n", chnStr)
				}
			}
		case '十':
			{
				if curVal > 0 {
					totalVal += curVal * 10
					curVal = 0
				} else {
					log.Printf("[INFO]Format error 2: %s!\n", chnStr)
				}
			}
		default:
			log.Printf("[INFO]Format error 3: %s!\n", chnStr)
			return -1
		}
	}

	if curVal != 0 {
		totalVal += curVal
	}

	return totalVal
}

const (
	ChnNumbers = `一二三四五六七八九`
)

func isChnNumber(r rune) bool {
	return strings.ContainsRune(ChnNumbers, r)
}

func Chn2NumSimple(r rune) int {
	switch r {
	case '一':
		return 1
	case '二':
		return 2
	case '三':
		return 3
	case '四':
		return 4
	case '五':
		return 5
	case '六':
		return 6
	case '七':
		return 7
	case '八':
		return 8
	case '九':
		return 9
	case '零':
		return 0
	default:
		return -1
	}
}
