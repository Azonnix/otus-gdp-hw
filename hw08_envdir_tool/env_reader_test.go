package main

import (
	"testing"

	"github.com/stretchr/testify/require"
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
	expected["FOO"] = EnvValue{"   foo\nwith new line", false}
	expected["HELLO"] = EnvValue{"\"hello\"", false}
	expected["UNSET"] = EnvValue{"", true}

	testCaseSuccsess := testCase{
		name:     "success",
		dirPath:  "testdata/env/",
		expected: expected,
	}

	t.Run(testCaseSuccsess.name, func(t *testing.T) {
		environments, err := ReadDir(testCaseSuccsess.dirPath)
		require.NoError(t, err)
		for k, v := range environments {
			require.Equal(t, expected[k].NeedRemove, v.NeedRemove)
			require.Equal(t, expected[k].Value, v.Value)
		}
	})
}
