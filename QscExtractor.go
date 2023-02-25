package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func ExportQsc(qscfile, exportDir string) {
	var qsce QscExtractor
	qsce.Export(qscfile, exportDir)
}

type QscExtractor struct {
	toDir  string
	cipais Cipais
}

func (p *QscExtractor) Export(qscfile, exportDir string) {
	content, err := ReadTextFile(qscfile)
	if err != nil {
		fmt.Printf("Cannot read file: %s!\n", qscfile)
		return
	}

	p.toDir = exportDir
	p.cipais.Init("CiPai.txt")

	arr := strings.Split(content, "\n\n")
	for _, s := range arr {
		p.parsePoems(s)
	}
	//fmt.Printf("Poets: %d\n", len(arr))
}

const BlankChars = "　 \t\r\n" // Contains a Zh Blank

func (p *QscExtractor) parsePoems(s string) {
	sn := strings.TrimLeft(s, BlankChars)
	lines := strings.Split(sn, "\n")

	if len(lines) == 0 {
		return
	}

	author := lines[0]
	authorclean := strings.Trim(author, BlankChars)
	childdir := p.toDir
	if authorclean != "" {
		childdir = path.Join(p.toDir, authorclean)

		// https://zetcode.com/golang/directory/
		if _, err := os.Stat(childdir); os.IsNotExist(err) {
			os.Mkdir(childdir, os.ModePerm)
		} else {
			fmt.Printf("Directory %s already exists!\n", childdir)
		}
	} else {
		//fmt.Printf("[WARN]Cannot parse parser for [%s]!\n", s)
		return
	}

	lines = lines[1:]
	if len(lines) == 0 {
		fmt.Printf("[WARN]Empty content for %s!\n", s)
		return
	}

	content := []string{}
	curTitle := ""

	for _, line := range lines {
		isTitle, _ := p.cipais.HasCipai(line)
		if isTitle {
			if len(content) > 0 {
				fullpath := combineFullPath(childdir, curTitle)
				writeOnePoem(content, fullpath)
				content = []string{}
			} else {

			}

			curTitle = line
		} else {
			content = append(content, line)
		}
	}

	if len(content) > 0 {
		fullpath := combineFullPath(childdir, curTitle)
		writeOnePoem(content, fullpath)
	}
}

func combineFullPath(childdir, title string) string {
	if title == "" {
		return ""
	}

	return path.Join(childdir, title+".txt")
}

func writeOnePoem(lines []string, fullpath string) {
	if fullpath == "" {
		fmt.Printf("No title for %+v!\n", lines)
		return
	}

	if !FileExists(fullpath) {
		// file does not exist
		WriteLines(lines, fullpath)
		return
	}

	//fmt.Printf("File exists: %s!\n", fullpath)
	oldlines, err := ReadLines(fullpath)
	if err != nil {
		fmt.Printf("Cannot read file: %s!\n", fullpath)
		return
	}

	if len(oldlines) <= 1 {
		return
	}

	firstline := lines[0]
	firstold := oldlines[0]
	if prefixCharSame(firstline, firstold, 3) {
		fmt.Printf("Repeat poem in %s!\n", fullpath)
		return
	}

	curId := findPoemId(fullpath)
	if curId < 0 {
		return
	}

	maxfilename := ""
	for {
		maxfilename = combineNewFile(fullpath, curId+1)
		if FileExists(maxfilename) {
			curId++
		} else {
			break
		}
	}

	WriteLines(lines, maxfilename)
}

func prefixCharSame(str1, str2 string, nchar int) bool {
	actualLen := nchar
	if actualLen > len(str1) {
		actualLen = len(str1)
	}
	if actualLen > len(str2) {
		actualLen = len(str2)
	}
	if actualLen == 0 {
		// ?
		return false
	}

	for i := 0; i < actualLen; i++ {
		if str1[i] != str2[i] {
			return false
		}
	}

	return true
}

const NumberString = "0123456789"

// `d:\aaa\西江月3.txt` ==> 3
func findPoemId(filename string) int {
	basefile := filepath.Base(filename)
	purefile := fileNameWithoutExt(basefile)

	// https://networkbit.ch/golang-regular-expression/
	pattern := regexp.MustCompile("\\d+")
	arr := pattern.FindAllString(purefile, -1)
	if len(arr) == 0 {
		return 0
	}

	lastitem := arr[len(arr)-1]
	v, err := strconv.Atoi(lastitem)
	if err != nil {
		fmt.Printf("WARN: Cannot get id for %s!\n", filename)
		return -1
	}
	return v
}

// https://freshman.tech/snippets/go/filename-no-extension/
func fileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func combineNewFile(filename string, id int) string {
	fpath := filepath.Dir(filename)
	basefile := filepath.Base(filename)
	purefile := fileNameWithoutExt(basefile)
	fileext := filepath.Ext(filename)

	filenew := fmt.Sprintf("%s%d", purefile, id)
	return filepath.Join(fpath, filenew+fileext)
}
