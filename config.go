package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Path struct {
	Path        string
	EnvCfgPatch string
}

type cfg struct {
	EnvConfig      string `json:"env_config"`
	ConfigFiles    `json:"config_files"`
	VariablesGroup `json:"variables_group"`
}

type ConfigFiles []ConfigPath
type ConfigPath struct {
	TemplatePath       string `json:"template_path"`
	FormedTemplatePath string `json:"formed_template_path"`
}

type VariablesGroup []VariableGroup
type VariableGroup struct {
	GroupName string   `json:"group_name"`
	Variables []string `json:"variables"`
}

func (p *Path) SetCfgPath(args []string) error {
	lastIndex := len(args) - 1
	for i := range args {
		if args[i] == "-conf" {
			configPathIndex := i + 1
			if lastIndex < configPathIndex {
				return errors.New("config file patch not defined")
			}
			p.Path = args[configPathIndex]
			return nil
		}
	}
	return errors.New("config file required")
}

func (p *Path) SetEnvCfgPath(args []string) error {
	lastIndex := len(args) - 1
	for i := range args {
		if args[i] == "-env_cfg" {
			configPathIndex := i + 1
			if lastIndex < configPathIndex {
				return errors.New("config file patch not defined")
			}
			p.Path = args[configPathIndex]
			return nil
		}
	}
	return errors.New("config file required")
}

func (p *Path) SetCfg() (*cfg, error) {
	file, err := ioutil.ReadFile(p.Path)
	if err != nil {
		fmt.Println("can't read cfg file" + err.Error())
		return &cfg{}, err
	}
	cfg := cfg{}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		fmt.Println("cant unmarshall" + err.Error())
	}

	return &cfg, nil
}

func (c *cfg) GetReplacerVariable() (res []string, err error) {
	for i := range c.VariablesGroup {
		for k := range c.VariablesGroup[i].Variables {
			a := strings.SplitN(c.VariablesGroup[i].Variables[k], "=", 2)
			if len(a) != 2 {
				return nil, err
			}
			variableTempName := "{{" + c.VariablesGroup[i].GroupName + "." + a[0] + "}}"
			value := a[1]
			vType := Variable{}
			err := vType.SetVariable(value)
			if err != nil {
				return nil, err
			}
			vType.EnvCfgPath = c.EnvConfig
			vStr, err := vType.GetValue()
			if err != nil {
				return nil, err
			}

			res = append(res, variableTempName, vStr)
		}
	}
	return res, nil
}
