package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func FindExistString(fileName string, str string) bool {
	var text []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		message := scanner.Text()
		text = append(text, message)
	}
	for _, each_ln := range text {
		if strings.EqualFold(each_ln, str) {
			return true
		}

	}
	return false
}
