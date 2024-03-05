package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	name     string
	dirPath  string
	expected Environment
}

func TestReadDir(t *testing.T) {
	expected := make(Environment)
	expected["BAR"] = EnvValue{"bar", false}
	expected["EMPTY"] = EnvValue{"", false}
	expected["FOO"] = EnvValue{}
	expected["HELLO"] = EnvValue{"hello", false}
	expected["UNSET"] = EnvValue{"", true}

	testCaseSuccsess := testCase{
		name:     "success",
		dirPath:  "testdata/env/",
		expected: expected,
	}

	t.Run(testCaseSuccsess.name, func(t *testing.T) {
		environments, err := ReadDir(testCaseSuccsess.dirPath)
		fmt.Println(err)
		for k, v := range environments {
			fmt.Println(k, v)
		}
	})
}
