package main

import (
	"fmt"
	"testing"
)

func TestCmd_RunCmd(t *testing.T) {
	v := Variable{}
	err := v.SetVariable("$$CMD(echo $(hostname -i):9000)")

	if err != nil {
		fmt.Println(err)
	}
	val, _ := v.GetValue()
	fmt.Println(val)
}
