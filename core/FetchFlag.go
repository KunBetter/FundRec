package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	FetchFlagDir = "FetchFlag"
)

type FetchFlag struct {
	MD5       uint64
	LatestDay string
}

func GetFetchFlag(tableName string) *FetchFlag {
	buf, err := ioutil.ReadFile("FetchFlag/" + tableName)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	lines := strings.Split(string(buf), "\n")
	if 2 == len(lines) {
		md5, _ := strconv.ParseUint(lines[0], 10, 64)
		fetchFlag := &FetchFlag{
			MD5:       md5,
			LatestDay: lines[1],
		}
		return fetchFlag
	}
	return nil
}

func WriteFlag(tableName string, fetchFlag *FetchFlag) {
	_, err := os.Stat(FetchFlagDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(FetchFlagDir, 0777)
			if err != nil {
				fmt.Print(err)
			}
		}
	}

	cont := fmt.Sprintf("%d\n%s", fetchFlag.MD5, fetchFlag.LatestDay)
	err = ioutil.WriteFile(FetchFlagDir+"/"+tableName, []byte(cont), 0666)
	if err != nil {
		fmt.Print(err)
	}
}
