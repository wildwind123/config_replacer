package main

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"regexp"
)

const (
	VTypeCMD    = "CMD"
	VTypeEnvCfg = "ENV_CFG"
	VTypeString = "STRING"
)

type Variable struct {
	EnvCfgPath string
	Type       string
	Value      string
}

func (v *Variable) SetVariable(variable string) (err error) {
	reg := regexp.MustCompile("^\\$\\$([A-Za-z]*)\\((.+)\\)$")
	ress := reg.FindAllStringSubmatch(variable, 2)

	if ress == nil || len(ress[0]) < 3 {
		v.Type = VTypeString
		v.Value = variable
		return nil
	}
	v.Type = ress[0][1]
	v.Value = ress[0][2]

	return nil
}

func (v *Variable) GetValue() (string, error) {
	if v.Type == VTypeCMD {
		c := cmd{}
		cRes, err := c.RunCmd(v.Value)
		if err != nil {
			return "", err
		}
		return cRes, nil
	} else if v.Type == VTypeString {
		return v.Value, nil
	} else if v.Type == VTypeEnvCfg {
		err := godotenv.Load(v.EnvCfgPath)
		if err != nil {
			return "", errors.New("error: " + fmt.Sprintf("%v", v) + " : " + err.Error())
		}
		res := os.Getenv(v.Value)
		if res == "" {
			fmt.Println("this variable does not set " + fmt.Sprintf("%v", v))
		}
		return res, nil
	}

	return "", errors.New("unknown type " + fmt.Sprintf("%v", v))
}
