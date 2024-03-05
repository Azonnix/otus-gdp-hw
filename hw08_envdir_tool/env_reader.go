package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	entrysDir, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	environments := make(Environment)

	for _, entryDir := range entrysDir {
		if strings.Contains(entryDir.Name(), "=") {
			continue
		}

		content, err := os.Open(dir + "/" + entryDir.Name())
		if err != nil {
			return nil, err
		}

		stat, err := content.Stat()
		if err != nil {
			return environments, err
		}

		if stat.Size() == 0 {
			environments[entryDir.Name()] = EnvValue{NeedRemove: true}
			continue
		}

		contentReader := bufio.NewReader(content)

		line, _, err := contentReader.ReadLine()
		if err != nil {
			return nil, err
		}

		if len(line) == 0 {
			environments[entryDir.Name()] = EnvValue{}
			continue
		}

		environments[entryDir.Name()] = EnvValue{
			Value: strings.TrimRight(string(bytes.ReplaceAll(line, []byte("\x00"), []byte("\n"))), " "),
		}
	}

	return environments, nil
}
