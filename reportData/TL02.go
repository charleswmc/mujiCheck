package reportData

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func CheckTL02() {
	log.Println("----------Start Checking TL02 Daily----------")
	lastTxn := GetLastTxnForLastCheck()
	var (
		text                []string
		missRecord          string
		counter, txn1, txn2 int
	)
	file, err := os.Open("sortedAll2One_TL02.csv")
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
	fileName := "TL02/backup/" + date + "_sortedAll2One_TL02.csv" //Linux
	if err = os.Remove(fileName); err != nil {
		// log.Println(err)
	} else {
		log.Println("Duplicate file deleted")
	}
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.Println(date + "_sortedAll2One_TL02.csv is Created!")
	for i, each_ln := range text {

		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		each_ln = each_ln[1:17]
		if strings.Contains(each_ln, `-`) {
			SaveTxnNo(each_ln[6:16])
			if i == 1 {
				CheckTxnStartCorrect(lastTxn[6:16], each_ln[6:16])

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
						}
					}
					txn1 = txnNoInInt
				}
			}
		} else {
			continue
		}
	}
	LastTxnForLastCheck(strconv.Itoa(txn1))
	file.Close()
	log.Println("----------End Checking TL02 Daily----------")
}

func LastTxnForLastCheck(txn string) {
	var txnNo string
	fileName := "LastTxnForLastCheck.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	if len(txn) == 3 {
		txnNo = "30101-0000000" + txn
	}
	if len(txn) == 4 {
		txnNo = "30101-000000" + txn
	}
	if len(txn) == 5 {
		txnNo = "30101-00000" + txn
	}
	if len(txn) == 6 {
		txnNo = "30101-0000" + txn
	}
	if len(txn) == 7 {
		txnNo = "30101-000" + txn
	}
	if len(txn) == 8 {
		txnNo = "30101-00" + txn
	}
	if txnNo != "" { // no TxnNo
		if _, err = file.WriteString(txnNo + "\n"); err != nil {
			log.Println(err)
		}
	}
	file.Close()
}

func GetLastTxnForLastCheck() (lastTxn string) {
	file, err := os.Open("LastTxnForLastCheck.txt")
	if err != nil {
		log.Println("Failed to open file")
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		message := scanner.Text()
		lastTxn = message
	}
	fmt.Println("lastTxn: ", lastTxn)
	return lastTxn
}

func CheckTxnStartCorrect(lastTxn string, txnStart string) {
	log.Println("----------Start CheckTxnStartCorrect----------")
	var txnNo string
	lastTxnInt, err := strconv.Atoi(lastTxn)
	if err != nil {
		log.Println(err)
	}
	txnStartInt, err := strconv.Atoi(txnStart)
	if err != nil {
		log.Println(err)
	}
	result := txnStartInt - lastTxnInt
	if result > 1 {
		log.Println("TxnStart is not correct, missing txn Record")
		for i := 1; i < result; i++ {
			txn := strconv.Itoa(lastTxnInt + i)

			if len(txn) == 3 {
				txnNo = "30101-0000000" + txn
			}
			if len(txn) == 4 {
				txnNo = "30101-000000" + txn
			}
			if len(txn) == 5 {
				txnNo = "30101-00000" + txn
			}
			if len(txn) == 6 {
				txnNo = "30101-0000" + txn
			}
			if len(txn) == 7 {
				txnNo = "30101-000" + txn
			}
			if len(txn) == 8 {
				txnNo = "30101-00" + txn
			}

			log.Println("Missing Txn Record: = ", txnNo)
		}
	} else if result < 1 {
		log.Println("Please Confirm the Input Date, you are inputing Previous Date")
	}
	log.Println("----------End CheckTxnStartCorrect----------")
}

func SaveTxnNo(txn string) {
	fileName := "TL02/8042/TxnRecord.txt"
	fileSaved := "TL02/8042/TxnRecord_Saved.txt"
	file, err := ioutil.ReadFile(fileSaved)
	if err != nil {
		return
	}
	s := string(file)
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("SaveTxnNo Failed: ", err)
	}
	defer f.Close()
	if strings.Contains(s, txn) {
		return
	} else {
		vmId := "30101"
		f.WriteString(vmId + "-" + txn + "\n")
	}
}

func CheckTL02ALL() {
	log.Println("----------Start Checking All Txn Record----------")
	var text []string
	var txnFirst, txnAfter int
	var missRecord string
	fileName := "TL02/8042/TxnRecord_Saved.txt"
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
	for i, each_ln := range text {
		if i == 0 {
			txnFirst, err = strconv.Atoi(each_ln[6:16])
			if err != nil {
				log.Println(err)
				return
			}
		} else {
			txnAfter, err = strconv.Atoi(each_ln[6:16])
			if err != nil {
				log.Println(err)
				return
			}
			cal := txnAfter - txnFirst
			// log.Println(i, cal)
			// miss := txnAfter - txnFirst
			if cal != 1 {
				for a := 1; a < cal; a++ {
					miss := txnFirst + a
					if len(strconv.Itoa(txnAfter)) == 3 {
						missRecord = "30101-0000000" + strconv.Itoa(miss)
					}
					if len(strconv.Itoa(txnAfter)) == 4 {
						missRecord = "30101-000000" + strconv.Itoa(miss)
					}
					if len(strconv.Itoa(txnAfter)) == 5 {
						missRecord = "30101-00000" + strconv.Itoa(miss)
					}
					if len(strconv.Itoa(txnAfter)) == 6 {
						missRecord = "30101-0000" + strconv.Itoa(miss)
					}
					if len(strconv.Itoa(txnAfter)) == 7 {
						missRecord = "30101-000" + strconv.Itoa(miss)
					}
					if len(strconv.Itoa(txnAfter)) == 8 {
						missRecord = "30101-00" + strconv.Itoa(miss)
					}
					log.Println("Missing Txn Record: ", missRecord)
				}
			}
			txnFirst = txnAfter
		}
	}
	file.Close()
	log.Println("----------End Checking All Txn Record----------")
}