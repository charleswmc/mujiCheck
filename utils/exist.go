package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func CheckIfItIsRepost(txnNo int) bool {
	var text []string
	var repostData string
	txnNoLen := len(strconv.Itoa(txnNo))
	if txnNoLen == 3 {
		repostData = "30101-0000000" + strconv.Itoa(txnNo)
	}
	if txnNoLen == 4 {
		repostData = "30101-000000" + strconv.Itoa(txnNo)
	}
	if txnNoLen == 5 {
		repostData = "30101-00000" + strconv.Itoa(txnNo)
	}
	if txnNoLen == 6 {
		repostData = "30101-0000" + strconv.Itoa(txnNo)
	}
	if txnNoLen == 7 {
		repostData = "30101-000" + strconv.Itoa(txnNo)
	}
	if txnNoLen == 8 {
		repostData = "30101-00" + strconv.Itoa(txnNo)
	}
	// fileName := "TL02/8042/TxnRecord_Saved.txt"
	fileName := "TL02/8042/TxnRecord_Sort.txt"
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
		if strings.EqualFold(each_ln, repostData) {
			return true
		}
	}
	file.Close()
	return false
}
