package env

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
 * Env File: /data/apps/appenv
 * Content: env=test or env=prod
 */

type Env int32

const (
	Test Env = 0
	Prod Env = 1

	AppEnvFilePath = "/data/apps/appenv"
)

func GetCurEnv() Env {
	buf, err := ioutil.ReadFile(AppEnvFilePath)
	if err != nil {
		fmt.Print(err)
	}
	lines := strings.Split(string(buf), "\n")
	for i := 0; i < len(lines) && lines[i] != ""; i++ {
		line := lines[i]
		data := strings.Split(line, "=")
		if 2 == len(data) && data[0] == "env" && data[1] == "prod" {
			return Prod
		}
	}

	return Test
}
