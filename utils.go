package main


import (
	"bytes"
	"io/ioutil"
	"regexp"
	"strings"
	"fmt"
	"log"
	"github.com/fatih/color"
)


const (
    Newline = "\n"
)

func ReadFile(path string) []byte {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return []byte{}
	}
	return data
}


func RunGrep(pattern, dataFromFile []byte, e bool) {
	if e {
		regGrep(pattern, dataFromFile)
	} else {
		grep(pattern, dataFromFile)
	}
}

func grep(pattern, dataFromFile []byte) {
	for _, line := range bytes.Split(dataFromFile, []byte(Newline)) {
		if bytes.Contains(line, pattern) {
			fmt.Println(strings.ReplaceAll(string(line), string(pattern), color.RedString("%s", pattern)))
		}
	}
}

func regGrep(pattern, dataFromFile []byte) {
	reg := regexp.MustCompile(string(pattern))
	
	for _, line := range bytes.Split(dataFromFile, []byte(Newline)) {	
		coloredLine := string(line)
		if reg.Match(line) {
			var allIndex = reg.FindAllIndex(line, -1)
			for _, index := range allIndex {
				begin := index[0]
				end := index[1]

				machedString := line[begin:end]
				coloredLine = strings.ReplaceAll(coloredLine, string(machedString), color.RedString("%s", machedString))
			}
			fmt.Println(string(coloredLine))
		}
	}
}

