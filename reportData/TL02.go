package reportData

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"mujiCheck/utils"
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
			if i == 0 {
				CheckTxnStartCorrect(lastTxn[6:16], each_ln[6:16])
				f.WriteString(each_ln + "\n")
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
							// Add checking - if Txn exists in TxnRecord_Saved.txt return
							if !utils.CheckIfItIsRepost(miss) {
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
							} else {
								continue
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
	// CheckTL02ALL()
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
	// fileSaved := "TL02/8042/TxnRecord_Saved.txt"
	fileSaved := "TL02/8042/TxnRecord_Sort.txt"
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
			// log.Println(txnAfter, txnFirst, i, cal)
			// miss := txnAfter - txnFirst
			// fmt.Println(miss)
			if cal != 1 {
				for a := 1; a < cal; a++ {
					// log.Println("Miss record ARRRRRRRRRRRRRRRRRRRRRRRRRR")
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
					f, err := os.OpenFile("SaveMissingTxn.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						log.Println(err)
					}
					fileMissSaveFile := "SaveMissingTxn.txt" // JUST for Record/References, not used
					fileMissSave, err := ioutil.ReadFile(fileMissSaveFile)
					if err != nil {
						return
					}
					s := string(fileMissSave)
					if strings.Contains(s, missRecord) {
						continue
					} else {
						f.WriteString(missRecord + "\n")
					}
				}
			}
			txnFirst = txnAfter
		}
	}
	file.Close()
	log.Println("----------End Checking All Txn Record----------")
}

// func CheckIfItIsRepost(txnNo int) bool {
// 	var text []string
// 	var repostData string
// 	txnNoLen := len(strconv.Itoa(txnNo))
// 	if txnNoLen == 3 {
// 		repostData = "30101-0000000" + strconv.Itoa(txnNo)
// 	}
// 	if txnNoLen == 4 {
// 		repostData = "30101-000000" + strconv.Itoa(txnNo)
// 	}
// 	if txnNoLen == 5 {
// 		repostData = "30101-00000" + strconv.Itoa(txnNo)
// 	}
// 	if txnNoLen == 6 {
// 		repostData = "30101-0000" + strconv.Itoa(txnNo)
// 	}
// 	if txnNoLen == 7 {
// 		repostData = "30101-000" + strconv.Itoa(txnNo)
// 	}
// 	if txnNoLen == 8 {
// 		repostData = "30101-00" + strconv.Itoa(txnNo)
// 	}
// 	// fileName := "TL02/8042/TxnRecord_Saved.txt"
// 	fileName := "TL02/8042/TxnRecord_Sort.txt"
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanLines)
// 	for scanner.Scan() {
// 		message := scanner.Text()
// 		text = append(text, message)
// 	}
// 	for _, each_ln := range text {
// 		if strings.EqualFold(each_ln, repostData) {
// 			return true
// 		}
// 	}
// 	file.Close()
// 	return false
// }

func SaveToSortTxnRecordFile() { // Save Txn Record to TxnRecord_Sort.txt
	var text []string
	fileName := "TL02/8042/TxnRecord.txt"
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
	length := len(text)
	fileSort := "TL02/8042/TxnRecord_Sort.txt"
	f, err := os.OpenFile(fileSort, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	for i := 0; i < length; i++ {
		// fmt.Println(utils.SortASC(text)[i])
		if !utils.FindExistString(fileSort, utils.SortASC(text)[i]) {
			f.WriteString(utils.SortASC(text)[i] + "\n")
		}
	}
}

func SortTxnRecordFile() {
	var text []string
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
	length := len(text)
	// fileSort := "TL02/8042/TxnRecord_Sort.txt"
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	for i := 0; i < length; i++ {
		// fmt.Println(utils.SortASC(text)[i])
		f.WriteString(utils.SortASC(text)[i] + "\n")
		// if !utils.FindExistString(fileSort, utils.SortASC(text)[i]) {
		// f.WriteString(utils.SortASC(text)[i] + "\n")
		// }
	}
}

// func SaveMissingTxn() {
// 	file, err := os.Open("SaveMissingTxn.txt")
// 	if err != nil {
// 		log.Println(err)
// 	}

// }

func CheckTL02FileSize() {
	path := "TL02/SFTP_file"
	// os.ReadDir(path)
	dir, err := ioutil.ReadDir(path)
	// dir, err := os.ReadDir(path)
	var counter int
	if err != nil {
		log.Println("Cannot read directory when update octopus upload")
		return
	}
	for _, f := range dir {
		name := f.Name()
		// fmt.Println(name)
		info, err := os.Stat(path + "/" + name)
		if err != nil {
			log.Println(err)
		}
		fileSize := info.Size()
		// fmt.Println(fileSize)
		if fileSize < 693 {
			log.Println("File Size is abnormal, Please Check the file: ", name)
		}
		counter++ // 1 day should have 96 files, if do checking at 2300, should have 100 files. files from 2205 to another day 2250
	}
	if counter != 100 {
		log.Println("No. of Files is not matched. There is missing files")
	}

}

func CheckTL02MissingFile() {
	path := "TL02/SFTP_file"
	dir, err := ioutil.ReadDir(path)
	var counter, counter1, counter2, counter3, counter4 int
	var namePrevious05, namePrevious20, namePrevious35, namePrevious50 string
	var nameNow05, nameNow20, nameNow35, nameNow50 string
	var fn string
	var cal int
	first := "22"
	// last := "22"
	if err != nil {
		log.Println("Cannot read directory when update octopus upload")
		return
	}
	os.Remove("TL02/MissingFiles.txt") //Remove before checking another day
	missingFiles := "TL02/MissingFiles.txt"
	fileName, err := os.OpenFile(missingFiles, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Open missingFiles.txt failed. ", err)
	}
	// if the first/last file is missed, can't println the file name
	for _, f := range dir {
		name := f.Name()
		if strings.Contains(name, "05.csv") {
			nameNow05 = name
			counter1++
			if counter1 > 1 {
				now, err := strconv.Atoi(nameNow05[23:25])
				if err != nil {
					log.Println("convert now to int failed")
				}
				previous, err := strconv.Atoi(namePrevious05[23:25])
				if err != nil {
					log.Println("convert previous to int failed")
				}
				if now != 0 {
					if now-previous != 1 {
						for a := 1; a < now-previous; a++ {
							if previous < 10 {
								if (previous + a) < 10 {
									fn = name[0:23] + "0" + strconv.Itoa(previous+a) + "05.csv"
									// log.Println("Missing files: ", fn)
									fileName.WriteString(fn + "\n")
								} else {
									fn = name[0:23] + strconv.Itoa(previous+a) + "05.csv"
									// log.Println("Missing files: ", fn)
									fileName.WriteString(fn + "\n")
								}
							} else {
								fn = name[0:23] + strconv.Itoa(previous+a) + "05.csv"
								// log.Println("Missing files: ", fn)
								fileName.WriteString(fn + "\n")
							}
						}
					}
				} else { // case: now == 0
					if previous-now == 22 {
						today, err := strconv.Atoi(name[14:22])
						if err != nil {
							log.Println(err)
						}
						yesterday := strconv.Itoa(today - 1)
						fn = name[0:14] + yesterday + "_2305.csv"
						// log.Println("Missing files: ", fn)
						fileName.WriteString(fn + "\n")
					}
				}
			}
			if counter1 == 1 && !strings.Contains(name, first+"05.csv") {
				hourName, err := strconv.Atoi(nameNow05[23:25])
				if err != nil {
					log.Println("convert hourName from string to int failed", err)
				}
				if hourName == 23 {
					fn = nameNow05[0:23] + first + "05.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				} else {
					// case if 0
					// fmt.Println(nameNow05[14:22])
					lastDay, err := strconv.Atoi(nameNow05[14:22])
					if err != nil {
						log.Println(err)
					}
					lastDayString := strconv.Itoa(lastDay - 1)
					fn = nameNow05[0:14] + lastDayString + "_" + first + "05.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
					fn = nameNow05[0:14] + lastDayString + "_" + "23" + "05.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
					if hourName > 0 {
						for i := 0; i < hourName; i++ {
							fnHour := strconv.Itoa(hourName - (hourName - i))
							if (hourName - (hourName - i)) < 10 {
								fn = nameNow05[0:23] + "0" + fnHour + "05.csv"
								// log.Println("Missing files: ", nameNow05[0:23]+"0"+fnHour+"05.csv")
								fileName.WriteString(fn + "\n")
							} else {
								fn = nameNow05[0:23] + fnHour + "05.csv"
								// log.Println("Missing files: ", nameNow05[0:23]+fnHour+"05.csv")
								fileName.WriteString(fn + "\n")
							}
						}
					}
				}
			}
		} else if strings.Contains(name, "20.csv") {
			nameNow20 = name
			counter2++
			if counter2 > 1 {
				now, err := strconv.Atoi(nameNow20[23:25])
				if err != nil {
					log.Println("convert now to int failed")
				}
				previous, err := strconv.Atoi(namePrevious20[23:25])
				if err != nil {
					log.Println("convert previous to int failed")
				}
				if now != 0 {
					if now-previous != 1 {
						for a := 1; a < now-previous; a++ {
							if previous < 10 {
								if (previous + a) < 10 {
									fn = name[0:23] + "0" + strconv.Itoa(previous+a) + "20.csv"
									// log.Println("Missing files: ", fn)
									fileName.WriteString(fn + "\n")
								} else {
									fn = name[0:23] + strconv.Itoa(previous+a) + "20.csv"
									// log.Println("Missing files: ", fn)
									fileName.WriteString(fn + "\n")
								}
							} else {
								fn = name[0:23] + strconv.Itoa(previous+a) + "20.csv"
								// log.Println("Missing files: ", fn)
								fileName.WriteString(fn + "\n")
							}
						}
					}
				} else { // case: now == 0
					if previous-now == 22 {
						today, err := strconv.Atoi(name[14:22])
						if err != nil {
							log.Println(err)
						}
						yesterday := strconv.Itoa(today - 1)
						fn = name[0:14] + yesterday + "_2320.csv"
						// log.Println("Missing files: ", fn)
						fileName.WriteString(fn + "\n")
					}
				}
			}
			if counter2 == 1 && !strings.Contains(name, first+"20.csv") {
				hourName, err := strconv.Atoi(nameNow20[23:25])
				if err != nil {
					log.Println("convert hourName from string to int failed", err)
				}
				if hourName == 23 {
					fn = nameNow20[0:23] + first + "20.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				} else {
					// case if 0
					// fmt.Println(nameNow05[14:22])
					lastDay, err := strconv.Atoi(nameNow20[14:22])
					if err != nil {
						log.Println(err)
					}
					lastDayString := strconv.Itoa(lastDay - 1)
					fn = nameNow20[0:14] + lastDayString + "_" + first + "20.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
					fn = nameNow20[0:14] + lastDayString + "_" + "23" + "20.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
					if hourName > 0 {
						for i := 0; i < hourName; i++ {
							fnHour := strconv.Itoa(hourName - (hourName - i))
							if (hourName - (hourName - i)) < 10 {
								fn = nameNow20[0:23] + "0" + fnHour + "20.csv"
								// log.Println("Missing files: ", nameNow20[0:23]+"0"+fnHour+"20.csv")
								fileName.WriteString(fn + "\n")
							} else {
								fn = nameNow20[0:23] + fnHour + "20.csv"
								// log.Println("Missing files: ", nameNow20[0:23]+fnHour+"20.csv")
								fileName.WriteString(fn + "\n")
							}
						}
					}
				}
			}
		} else if strings.Contains(name, "35.csv") {
			nameNow35 = name
			counter3++
			if counter3 > 1 {
				now, err := strconv.Atoi(nameNow35[23:25])
				if err != nil {
					log.Println("convert now to int failed")
				}
				previous, err := strconv.Atoi(namePrevious35[23:25])
				if err != nil {
					log.Println("convert previous to int failed")
				}
				if now != 0 {
					if now-previous != 1 {
						for a := 1; a < now-previous; a++ {
							if previous < 10 {
								if (previous + a) < 10 {
									fn = name[0:23] + "0" + strconv.Itoa(previous+a) + "35.csv"
									// log.Println("Missing files: ", fn)
									fileName.WriteString(fn + "\n")
								} else {
									fn = name[0:23] + strconv.Itoa(previous+a) + "35.csv"
									// log.Println("Missing files: ", fn)
									fileName.WriteString(fn + "\n")
								}
							} else {
								fn = name[0:23] + strconv.Itoa(previous+a) + "35.csv"
								// log.Println("Missing files: ", fn)
								fileName.WriteString(fn + "\n")
							}
						}
					}
				} else { // case: now == 0
					if previous-now == 22 {
						today, err := strconv.Atoi(name[14:22])
						if err != nil {
							log.Println(err)
						}
						yesterday := strconv.Itoa(today - 1)
						fn = name[0:14] + yesterday + "_2335.csv"
						// log.Println("Missing files: ", fn)
						fileName.WriteString(fn + "\n")
					}
				}
			}
			if counter3 == 1 && !strings.Contains(name, first+"35.csv") {
				hourName, err := strconv.Atoi(nameNow35[23:25])
				if err != nil {
					log.Println("convert hourName from string to int failed", err)
				}
				if hourName == 23 {
					fn = nameNow35[0:23] + first + "35.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				} else {
					// case if 0
					// fmt.Println(nameNow05[14:22])
					lastDay, err := strconv.Atoi(nameNow35[14:22])
					if err != nil {
						log.Println(err)
					}
					lastDayString := strconv.Itoa(lastDay - 1)
					fn = nameNow35[0:14] + lastDayString + "_" + first + "35.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
					fn = nameNow35[0:14] + lastDayString + "_" + "23" + "35.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
					if hourName > 0 {
						for i := 0; i < hourName; i++ {
							fnHour := strconv.Itoa(hourName - (hourName - i))
							if (hourName - (hourName - i)) < 10 {
								fn = nameNow35[0:23] + "0" + fnHour + "35.csv"
								// log.Println("Missing files: ", nameNow35[0:23]+"0"+fnHour+"35.csv")
								fileName.WriteString(fn + "\n")
							} else {
								fn = nameNow35[0:23] + fnHour + "35.csv"
								// log.Println("Missing files: ", nameNow35[0:23]+fnHour+"35.csv")
								fileName.WriteString(fn + "\n")
							}
						}
					}
				}
			}
		} else if strings.Contains(name, "50.csv") {
			nameNow50 = name
			counter4++
			if counter4 > 1 {
				now, err := strconv.Atoi(nameNow50[23:25])
				if err != nil {
					log.Println("convert now to int failed")
				}
				previous, err := strconv.Atoi(namePrevious50[23:25])
				if err != nil {
					log.Println("convert previous to int failed")
				}
				if now != 0 {
					if now-previous != 1 {
						for a := 1; a < now-previous; a++ {
							if previous < 10 {
								if (previous + a) < 10 {
									fn = name[0:23] + "0" + strconv.Itoa(previous+a) + "50.csv"
									// log.Println("Missing files: ", fn)
									fileName.WriteString(fn + "\n")
								} else {
									fn = name[0:23] + strconv.Itoa(previous+a) + "50.csv"
									// log.Println("Missing files: ", fn)
									fileName.WriteString(fn + "\n")
								}
							} else {
								fn = name[0:23] + strconv.Itoa(previous+a) + "50.csv"
								// log.Println("Missing files: ", fn)
								fileName.WriteString(fn + "\n")
							}
						}
					}

				} else { // case: now == 0
					if previous-now == 22 {
						today, err := strconv.Atoi(name[14:22])
						if err != nil {
							log.Println(err)
						}
						yesterday := strconv.Itoa(today - 1)
						// fn = name[0:14] + yesterday + "_2250.csv"
						// log.Println("Missing files: ", fn)
						// fileName.WriteString(fn + "\n")
						fn = name[0:14] + yesterday + "_2350.csv"
						// log.Println("Missing files: ", fn)
						fileName.WriteString(fn + "\n")
					}
				}
			}
			if counter4 == 1 && !strings.Contains(name, first+"50.csv") {
				hourName, err := strconv.Atoi(nameNow50[23:25])
				if err != nil {
					log.Println("convert hourName from string to int failed", err)
				}
				if hourName == 23 {
					fn = nameNow50[0:23] + first + "50.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				} else {
					// case if 0
					// fmt.Println(nameNow05[14:22])
					lastDay, err := strconv.Atoi(nameNow50[14:22])
					if err != nil {
						log.Println(err)
					}
					lastDayString := strconv.Itoa(lastDay - 1)
					fn = nameNow50[0:14] + lastDayString + "_" + first + "50.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
					fn = nameNow50[0:14] + lastDayString + "_" + "23" + "50.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
					if hourName > 0 {
						for i := 0; i < hourName; i++ {
							fnHour := strconv.Itoa(hourName - (hourName - i))
							if (hourName - (hourName - i)) < 10 {
								fn = nameNow50[0:23] + "0" + fnHour + "50.csv"
								// log.Println("Missing files: ", nameNow50[0:23]+"0"+fnHour+"50.csv")
								fileName.WriteString(fn + "\n")
							} else {
								fn = nameNow50[0:23] + fnHour + "50.csv"
								// log.Println("Missing files: ", nameNow50[0:23]+fnHour+"50.csv")
								fileName.WriteString(fn + "\n")
							}
						}
					}
				}
			}
		}
		counter++
		namePrevious05 = nameNow05
		namePrevious20 = nameNow20
		namePrevious35 = nameNow35
		namePrevious50 = nameNow50
	}
	if counter1 < 25 {
		if counter1 == 0 {
			log.Println("All 05.csv is missing")
		}
		hour, err := strconv.Atoi(namePrevious05[23:25])
		if err != nil {
			log.Println("Convert hour from string to int failed. ", err)
		}
		if hour != 22 {
			for i := 1; i <= 22-hour; i++ {
				hourName := strconv.Itoa(hour + i)
				if hour+i < 10 {
					fn = namePrevious05[0:23] + "0" + hourName + "05.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				} else {
					fn = namePrevious05[0:23] + hourName + "05.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				}
			}
		}
		cal = 25 - counter1
		fmt.Println("There are ", cal, " 05.csv is missing")
	}
	if counter2 < 25 {
		if counter2 == 0 {
			log.Println("All 20.csv is missing")
		}
		hour, err := strconv.Atoi(namePrevious20[23:25])
		if err != nil {
			log.Println("Convert hour from string to int failed. ", err)
		}
		if hour != 22 {
			for i := 1; i <= 22-hour; i++ {
				hourName := strconv.Itoa(hour + i)
				if hour+i < 10 {
					fn = namePrevious20[0:23] + "0" + hourName + "20.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				} else {
					fn = namePrevious20[0:23] + hourName + "20.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				}
			}
		}
		cal = 25 - counter2
		fmt.Println("There are ", cal, " 20.csv is missing")
	}
	if counter3 < 25 {
		if counter3 == 0 {
			log.Println("All 35.csv is missing")
		}
		hour, err := strconv.Atoi(namePrevious35[23:25])
		if err != nil {
			log.Println("Convert hour from string to int failed. ", err)
		}
		if hour != 22 {
			for i := 1; i <= 22-hour; i++ {
				hourName := strconv.Itoa(hour + i)
				if hour+i < 10 {
					fn = namePrevious35[0:23] + "0" + hourName + "35.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				} else {
					fn = namePrevious35[0:23] + hourName + "35.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				}
			}
		}
		cal = 25 - counter3
		fmt.Println("There are ", cal, " 35.csv is missing")
	}
	if counter4 < 25 {
		if counter4 == 0 {
			log.Println("All 50.csv is missing")
		}
		hour, err := strconv.Atoi(namePrevious50[23:25])
		if err != nil {
			log.Println("Convert hour from string to int failed. ", err)
		}
		if hour != 22 {
			for i := 1; i <= 22-hour; i++ {
				hourName := strconv.Itoa(hour + i)
				if hour+i < 10 {
					fn = namePrevious50[0:23] + "0" + hourName + "50.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				} else {
					fn = namePrevious50[0:23] + hourName + "50.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				}
			}
		}
		cal = 25 - counter4
		fmt.Println("There are ", cal, " 50.csv is missing")
	}
}

func SortTL02MissingFile() {
	var text []string
	fileName := "TL02/MissingFiles.txt"
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
	length := len(text)
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	for i := 0; i < length; i++ {
		f.WriteString(utils.SortASC(text)[i] + "\n")
	}
}

func GetTL02MissingFile(fn string) (fileName string) {
	path := "TL02/SFTP_file"
	dir, err := ioutil.ReadDir(path)
	var counter, hourCounter int
	exist := false
	if err != nil {
		log.Println("Cannot read directory when update octopus upload")
		return
	}
	hour := 22
	if fn == "05.csv" {
		counter = 0
	} else if fn == "20.csv" {
		counter = 1
	} else if fn == "35.csv" {
		counter = 2
	} else if fn == "50.csv" {
		counter = 3
	}
	for _, f := range dir {
		name := f.Name()
		if hourCounter%5 == 4 {
			counter = 0
			hourCounter = 0
			hour++
		}
		if hour == 24 {
			hour = 0
		}
		str := strconv.Itoa(hour)
		if strings.Contains(name, str+fn) {
			// fmt.Println("file exist: ", name)
			exist = true
		} else if !strings.Contains(name, str+fn) {
			exist = false
		}
		if exist {
			// fmt.Println("file not exist: ", name)
			fmt.Println("file exist: ", name, counter, hourCounter)
		} else if !exist && counter == hourCounter {
			fmt.Println("file not exist: ", name, counter, hourCounter)
		}
		hourCounter++
		counter++
		exist = false
	}
	return fileName
}
