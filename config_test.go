package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCfgPath_SetCfgPath(t *testing.T) {
	cp := Path{}
	agrPath := "/home/test.cfg"

	err := cp.SetCfgPath([]string{"-conf", agrPath})
	if err != nil {
		t.Error(err)
	}
	if cp.Path != agrPath {
		t.Error("path wrong")
	}

	err = cp.SetCfgPath([]string{"-conf"})
	if err == nil {
		t.Error("error should by empty")
	}
}

func TestCfgPath_SetCfg(t *testing.T) {
	cp := Path{
		Path: "./config_test.json",
	}
	cfg, err := cp.SetCfg()
	if err != nil {
		t.Error(err)
	}
	if len(cfg.ConfigFiles) != 2 {
		t.Error("wrong cfg")
	}
}

func Test_Full(t *testing.T) {
	cp := Path{
		Path: "./config_test.json",
	}
	cfg, err := cp.SetCfg()
	if err != nil {
		t.Error(err)
	}
	variables, err := cfg.GetReplacerVariable()
	if err != nil {
		t.Error(err)
	}
	f := Files{}
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
			t.Error(err)
		}
		for x := range fs {
			if !f.FileExists(fs[x]) {
				fmt.Println("file does not exist = " + fs[x])
				continue
			}

			str, err := f.GetFormedTemplate(fs[x], variables)
			if err != nil {
				t.Error(err)
			}
			formedFileName, err := f.GetFormedFileName(fs[x])
			if err != nil {
				t.Error(err)
			}

			err = f.WriteFormedTemplate([]byte(str), cfg.ConfigFiles[i].FormedTemplatePath+"/"+formedFileName)
			if err != nil {
				t.Error(err)
			}
		}
	}
}
