package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const templateExt = ".tmpl"

type Files struct {
}

func (f *Files) ExportTemplates(c cfg) error {
	return nil
}

func (f *Files) pathTmplFiles(cf ConfigPath) (filesPath []string, err error) {
	files, err := ioutil.ReadDir(cf.TemplatePath)
	if err != nil {
		fmt.Println("cant get files" + err.Error())
		return nil, err
	}
	for i := range files {
		if files[i].IsDir() {
			continue
		}
		fileName := files[i].Name()
		extension := filepath.Ext(fileName)
		if extension != templateExt {
			continue
		}
		filesPath = append(filesPath, cf.TemplatePath+"/"+fileName)

	}
	return filesPath, nil
}

func (f *Files) GetFormedTemplate(filePath string, variables []string) (str string, err error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	r := strings.NewReplacer(variables...)

	result := r.Replace(string(file))

	return result, err
}

func (f *Files) WriteFormedTemplate(b []byte, path string) (err error) {
	return ioutil.WriteFile(path, b, 0644)
}

func (f *Files) GetFormedFileName(s string) (str string, err error) {
	file := filepath.Base(s)
	out := strings.Replace(file, templateExt, "", -1)
	return out, err
}

func (f *Files) FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
