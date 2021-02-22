package main

import (
	"testing"
)

func TestFiles_ListFiles(t *testing.T) {
	cp := Path{
		Path: "./config_test.json",
	}
	cfg, err := cp.SetCfg()
	if err != nil {
		t.Error(err)
	}
	f := Files{}
	files, err := f.pathTmplFiles(cfg.ConfigFiles[0])
	if len(files) != 1 {
		t.Error("file count wrong")
	}
}

func TestFiles_GetFormedTemplate(t *testing.T) {
	cp := Path{
		Path: "./config_test.json",
	}
	cfg, err := cp.SetCfg()
	if err != nil {
		t.Error(err)
	}
	f := Files{}
	values, err := cfg.GetReplacerVariable()
	if err != nil {
		t.Error(err)
	}
	_, err = f.GetFormedTemplate("./templates_test/test.conf.tmpl", values)
	if err != nil {
		t.Error(err)
	}
}

func TestCfg_GetVariablesByString(t *testing.T) {
	cp := Path{
		Path: "./config_test.json",
	}
	cfg, err := cp.SetCfg()
	if err != nil {
		t.Error(err)
	}
	r, _ := cfg.GetReplacerVariable()
	if r[0] != "{{mysql.port}}" {
		t.Error("Wrong result")
	}
}

func TestFiles_WriteFormedTemplate(t *testing.T) {
	cp := Path{
		Path: "./config_test.json",
	}
	cfg, err := cp.SetCfg()
	if err != nil {
		t.Error(err)
	}
	f := Files{}
	values, err := cfg.GetReplacerVariable()
	if err != nil {
		t.Error(err)
	}
	r, err := f.GetFormedTemplate("./templates_test/test.conf.tmpl", values)
	if err != nil {
		t.Error(err)
	}

	err = f.WriteFormedTemplate([]byte(r), "./formed_test/formed.conf")
	if err != nil {
		t.Error(err)
	}
}

func TestFiles_GetFormedFileName(t *testing.T) {
	f := Files{}
	r, err := f.GetFormedFileName("/home/rick/test.conf.tmpl")
	if err != nil {
		t.Error(err)
	}
	if r != "test.conf" {
		t.Error("wrong result")
	}
}
