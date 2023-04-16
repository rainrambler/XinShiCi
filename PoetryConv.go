package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

/*
[
  {
    "author": "石孝友",
    "paragraphs": [
      "扁舟破浪鸣双橹。",
      "岁晚客心分万绪。",
      "香红漠漠落梅村，愁碧萋萋芳草渡。",
      "汉皋佩失诚相误。",
      "楚峡云归无觅处。",
      "一天明月缺还圆，千里伴人来又去。"
    ],
    "rhythmic": "玉楼春"
  }
]
*/
type CiPoetry struct {
	Author     string   `json:"author"`
	Paragraphs []string `json:"paragraphs"`
	Rhythmic   string   `json:"rhythmic"`
}

/*
{
    "description": "--(1014－1079)字子政，一作子正。宋城(今河南商丘)人。",
    "name": "蔡挺",
    "short_description": ""
 }
*/
type CiPoet struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ShortDesc   string `json:"short_description"`
}

func ExportPoetry(qscdir, exportDir string) {
	var pt CiPoetries
	pt.LoadPoetryPath(qscdir)

	fmt.Printf("Total %d poets.\n", len(pt.AllPoets))

	pt.exportToDir(exportDir)
}

type CiPoetries struct {
	AllPoets        []CiPoetry
	AllCiPoets      []CiPoet
	authorRhy2Count map[string]int
}

func (p *CiPoetries) LoadPoets(filename string) {
	jsonFile, err := os.Open(filename)
	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var items []CiPoet
	err = json.Unmarshal(byteValue, &items)
	if err != nil {
		log.Printf("[WARN]Poet File: %s error: %v!\n", filename, err)
		return
	}

	if len(items) == 0 {
		return
	}

	fmt.Printf("[INFO]Total %d poets.\n", len(items))
	p.AllCiPoets = append(p.AllCiPoets, items...)
}

func (p *CiPoetries) loadPoetryFile(filename string) {
	jsonFile, err := os.Open(filename)
	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var poetries []CiPoetry
	err = json.Unmarshal(byteValue, &poetries)
	if err != nil {
		log.Printf("[WARN]File: %s error: %v!\n", filename, err)
		return
	}

	if len(poetries) == 0 {
		return
	}

	//fmt.Printf("[INFO] %d poetries.\n", len(poetries))
	for _, item := range poetries {
		key := item.Author + "|" + item.Rhythmic
		v, exists := p.authorRhy2Count[key]

		if exists {
			item.Rhythmic = fmt.Sprintf("%s%d", item.Rhythmic, v+1)
			p.authorRhy2Count[key] = v + 1
		} else {
			p.authorRhy2Count[key] = 0
		}

		p.AllPoets = append(p.AllPoets, item)
	}
	//p.AllPoets = append(p.AllPoets, poetries...)
}

func (p *CiPoetries) LoadPoetryPath(fullpath string) {
	p.authorRhy2Count = make(map[string]int)
	p.LoadPoets(path.Join(fullpath, "author.song.json"))

	err := filepath.Walk(fullpath, func(path1 string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if isPoetryFile(path1) {
			p.loadPoetryFile(path1)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[INFO]Total %d poetries.\n", len(p.AllPoets))
}

// filename should be "ci.song.21000.json"
func isPoetryFile(filename string) bool {
	basefile := filepath.Base(filename)
	purefile := fileNameWithoutExt(basefile)
	fileext := filepath.Ext(filename)

	if fileext != ".json" {
		return false
	}

	fmt.Printf("filename: %s\n", purefile)
	return strings.HasPrefix(purefile, "ci")
}

func (p *CiPoetries) CreatePathes(dirname string) {
	for _, author := range p.AllCiPoets {
		authorname := author.Name

		if authorname != "" {
			childdir := path.Join(dirname, authorname)
			// https://zetcode.com/golang/directory/
			if _, err := os.Stat(childdir); os.IsNotExist(err) {
				os.Mkdir(childdir, os.ModePerm)
			} else {
				fmt.Printf("Directory %s already exists!\n", childdir)
			}
		}
	}
}

func (p *CiPoetries) exportToDir(dirname string) {
	p.CreatePathes(dirname)

	for _, item := range p.AllPoets {
		childdir := path.Join(dirname, item.Author, item.Rhythmic+".txt")
		p.exportOnePoetry(&item, childdir)
	}
}

func (p *CiPoetries) exportOnePoetry(item *CiPoetry, dirname string) {
	WriteLines(item.Paragraphs, dirname)
}
