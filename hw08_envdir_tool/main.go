package main

import (
	"os"
	"strings"
)

func main() {
	environmentsDir, err := ReadDir(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	environments := make(Environment)

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		environments[pair[0]] = EnvValue{Value: pair[1]}
	}

	for k, v := range environmentsDir {
		if v.NeedRemove {
			delete(environments, k)
		}

		environments[k] = v
	}

	os.Exit(RunCmd(os.Args[2:], environments))
}
