package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	path := "C:/Users/hasee/Desktop/文档"
	allFile, _ := GetAllFile(path)
	dir := dirTras(allFile) 
	for _, p := range dir {
		fmt.Println(p)
	}
}

func GetAllFile(pathname string) ([]string, error) {
    var res []string
    // var allDir string
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read directory fail:", err)
		return res, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			// allDir := pathname + "/" + fi.Name()
			allDir := pathname + "/" + fi.Name()
			res, err = GetAllFile(allDir)
			if err != nil {
				return res, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			res = append(res, fullName)
		}
	}
	return res, nil
}

func dirTras(path []string) []string {
	var res []string
	var tmp string
	for _, fPath := range allFile {
		for _, f := strings.Split(fPath[:len(fPath)-1], "/") {
			tmp += "*****" 
		}
		tmp += fPath[len(fPath)-1]
		res = append(res, tmp)
	}
	return res
}