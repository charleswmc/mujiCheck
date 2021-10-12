package reportData

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

func CheckTL05() {
	log.Println("----------Start Checking TL05 Daily----------")
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
	}
	log.Println("sortedAll2One.csv is Created!")
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	for _, each_ln := range text {
		// fmt.Println(each_ln)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if strings.Contains(each_ln, `SAL`) {
			f.WriteString(each_ln + "\n")
		} else if strings.Contains(each_ln, `REF`) {
			f.WriteString(each_ln + "\n")
		} else if strings.Contains(each_ln, `ADJ`) {
			f.WriteString(each_ln + "\n")
		}
	}

	file.Close()
	log.Println("----------End Checking TL02 Daily----------")
}
