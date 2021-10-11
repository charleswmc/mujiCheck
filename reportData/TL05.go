package reportData

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

func CheckTL05() {
	// lastTxn := getLastTxnForLastCheck()
	var (
		text []string
	)
	file, err := os.Open("sortedAll2One_TL05.csv")
	if err != nil {
		log.Println("Failed to open file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		message := scanner.Text()
		text = append(text, message)
	}
	timestamp := time.Now().Unix()
	datetime := time.Unix(timestamp, 0)
	date := datetime.Format("20060102")
	// fileName := "TL02\\backup\\" + date + "_sortedAll2One_TL02.csv"	//Windows
	fileName := "TL05/backup/" + date + "_sortedAll2One_TL05.csv" //Linux
	if err = os.Remove(fileName); err != nil {
		// log.Println(err)
	} else {
		log.Println("Duplicate file deleted")
	}
	log.Println("sortedAll2One.csv is Created!")
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	for _, each_ln := range text {

		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		each_ln = each_ln[1:17]
		if strings.Contains(each_ln, `SAL`) {

		} else {
			continue
		}
	}

	file.Close()
}
