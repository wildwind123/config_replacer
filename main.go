package main

import (
	"fmt"
	"os"
)

func main() {
	p := Path{}
	err := p.SetCfgPath(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
	process(p)
}

func process(cp Path) {
	cfg, err := cp.SetCfg()
	if err != nil {
		fmt.Println(err)
		return
	}
	variables, err := cfg.GetReplacerVariable()
	if err != nil {
		fmt.Println(err)
		return
	}
	f := Files{}
	if len(cfg.ConfigFiles) == 0 {
		fmt.Println("ConfigFiles empty")
		return
	}
	for i := range cfg.ConfigFiles {
		if _, err := os.Stat(cfg.ConfigFiles[i].TemplatePath); os.IsNotExist(err) {
			fmt.Println("path: " + cfg.ConfigFiles[i].TemplatePath + ". Does not exist")
			continue
		}
		if _, err := os.Stat(cfg.ConfigFiles[i].FormedTemplatePath); os.IsNotExist(err) {
			fmt.Println("path: " + cfg.ConfigFiles[i].FormedTemplatePath + ". Does not exist")
			continue
		}

		fs, err := f.pathTmplFiles(cfg.ConfigFiles[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		for x := range fs {
			if !f.FileExists(fs[x]) {
				fmt.Println("file does not exist = " + fs[x])
				continue
			}

			str, err := f.GetFormedTemplate(fs[x], variables)
			if err != nil {
				fmt.Println(err)
				return
			}
			formedFileName, err := f.GetFormedFileName(fs[x])
			if err != nil {
				fmt.Println(err)
				return
			}

			err = f.WriteFormedTemplate([]byte(str), cfg.ConfigFiles[i].FormedTemplatePath+"/"+formedFileName)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
