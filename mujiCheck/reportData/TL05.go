package reportData

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func CheckTL05() {
	// lastTxn := getLastTxnForLastCheck()
	var (
		text                []string
		missRecord          string
		counter, txn1, txn2 int
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
	for i, each_ln := range text {

		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		each_ln = each_ln[1:17]
		if strings.Contains(each_ln, `-`) {
			if i == 1 {
				// checkTxnStartCorrect(lastTxn[6:16], each_ln[6:16])

			} else if _, err := f.WriteString(each_ln + "\n"); err != nil {
				log.Println(err)
			} else if txnNoInInt, err := strconv.Atoi(each_ln[6:16]); err != nil {
				log.Println(err)
			} else {
				counter++
				if counter == 1 {
					txn1 = txnNoInInt
				} else {
					txn2 = txnNoInInt
					cal := txn2 - txn1
					if cal != 1 {
						// log.Println("There is missing SAL txnNo. ", txn2, " ", txn1)
						// noOfTxn := strconv.Itoa(cal)
						for i := 1; i < cal; i++ {
							miss := txn2 - (cal - i)
							if len(strconv.Itoa(txn2)) == 3 {
								missRecord = "30101-0000000" + strconv.Itoa(miss)
								fmt.Println("Missing Txn Record: " + missRecord)
							}
							if len(strconv.Itoa(txn2)) == 4 {
								missRecord = "30101-000000" + strconv.Itoa(miss)
								fmt.Println("Missing Txn Record: " + missRecord)
							}
							if len(strconv.Itoa(txn2)) == 5 {
								missRecord = "30101-00000" + strconv.Itoa(miss)
								fmt.Println("Missing Txn Record: " + missRecord)
							}
							if len(strconv.Itoa(txn2)) == 6 {
								missRecord = "30101-0000" + strconv.Itoa(miss)
								fmt.Println("Missing Txn Record: " + missRecord)
							}
							if len(strconv.Itoa(txn2)) == 7 {
								missRecord = "30101-000" + strconv.Itoa(miss)
								fmt.Println("Missing Txn Record: " + missRecord)
							}
							if len(strconv.Itoa(txn2)) == 8 {
								missRecord = "30101-00" + strconv.Itoa(miss)
								fmt.Println("Missing Txn Record: " + missRecord)
							}
							// log.Println("Missing Txn Record: ", txn2-(cal-i))
						}
					}
					txn1 = txnNoInInt
				}
			}
		} else {
			continue
		}
	}
	// fmt.Println(strconv.Itoa(txn1))
	LastTxnForLastCheck(strconv.Itoa(txn1))
	file.Close()
}
