package reportData

import (
	"bufio"
	"io/ioutil"
	"log"
	"mujiCheck/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

func CheckTL05() {
	log.Println("----------Start Checking TL05 Daily----------")
	CheckTL05ileSize()
	// lastTxn := getLastTxnForLastCheck()
	var (
		text                                                   []string
		salCounter, refCounter, refItems, adjCounter, adjItems int
		ref, refNow, adj, adjNow                               string
		// sal, salNow                                            string
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
		if strings.Contains(each_ln, `SAL`) { // need to check sal txnNo is continous
			// if salCounter == 0 {
			// 	sal = each_ln[6:17]
			// 	// log.Println("sal: "+sal[6:16], salNow)
			// 	salCounter++
			// } else {
			// 	salNow = each_ln[6:17]
			// 	salInt, err := strconv.Atoi(sal)
			// 	if err != nil {
			// 		log.Println(err)
			// 	}
			// 	salNowInt, err := strconv.Atoi(salNow)
			// 	if err != nil {
			// 		log.Println(err)
			// 	}
			// 	if salNowInt-salInt != 1 {
			// 		for i := 0; i <= 1; i++ {
			// 		}
			// 	}
			// }
			f.WriteString(each_ln + "\n")
			salCounter++
		} else if strings.Contains(each_ln, `REF`) {
			if refCounter == 0 {
				ref = each_ln[1:17]
				refCounter++
			} else {
				refNow = each_ln[1:17]
				if refNow != ref {
					refCounter++
					ref = refNow
				}
			}
			refItems++
			f.WriteString(each_ln + "\n")
		} else if strings.Contains(each_ln, `ADJ`) {
			if refCounter == 0 {
				adj = each_ln[1:17]
				adjCounter++
			} else {
				adjNow = each_ln[1:17]
				if adjNow != adj {
					adjCounter++
					adj = adjNow
				}
			}
			adjItems++
			f.WriteString(each_ln + "\n")
		}
	}
	log.Println("salCounter: " + strconv.Itoa(salCounter))
	log.Println("refCounter: " + strconv.Itoa(refCounter) + ", refItems: " + strconv.Itoa(refItems))
	log.Println("adjCounter: " + strconv.Itoa(adjCounter) + ", adjItems: " + strconv.Itoa(adjItems))
	file.Close()
	CheckTL05SAL()
	log.Println("----------End Checking TL05 Daily----------")
}

//Match TL02 len,
func CheckTL05SAL() {
	var TL02text, TL05text []string
	// var match bool
	timestamp := time.Now().Unix()
	datetime := time.Unix(timestamp, 0)
	date := datetime.Format("20060102")
	TL02fileName := "TL02/backup/" + date + "_sortedAll2One_TL02.csv"
	TL02file, err := os.Open(TL02fileName)
	if err != nil {
		log.Println(err)
	}
	TL02scanner := bufio.NewScanner(TL02file)
	TL02scanner.Split(bufio.ScanLines)
	for TL02scanner.Scan() {
		message := TL02scanner.Text()
		// log.Println(message)
		TL02text = append(TL02text, message)
	}

	TL05fileName := "TL05/backup/" + date + "_sortedAll2One_TL05.csv"
	TL05file, err := os.Open(TL05fileName)
	if err != nil {
		log.Println(err)
	}
	TL05scanner := bufio.NewScanner(TL05file)
	TL05scanner.Split(bufio.ScanLines)
	for TL05scanner.Scan() {
		message := TL05scanner.Text()
		if strings.Contains(message, "SAL") {
			TL05text = append(TL05text, message[1:17])
		}
	}
	log.Println("TL05text: ", TL05text)
	log.Println("TL02text: ", TL02text)
	log.Println("TL02 sal count: ", len(TL02text))
	log.Println("TL05 sal count: ", len(TL05text))
	match := false
	if len(TL05text) == len(TL02text) {
		log.Println("TL05 SAL records len match with TL02")

		for i := 0; i < len(TL02text); i++ {
			//
			if TL02text[i] == TL05text[i] {
				match = true
				break
			} else {
				log.Println("Not Match [", i, "] ", "TL02text[i]:", TL02text[i], "TL05text[i]:", TL05text[i])
			}
		}
	} else {
		log.Println("TL05 SAL records len not match with TL02")
	}
	if match {
		log.Println("TL05 SAL records match with TL02")
	}
	if !match {
		log.Println("TL05 SAL records not match with TL02")
	}
	// check / list not match record
	if len(TL02text) > len(TL05text) {
		for i := 0; i < len(TL02text); i++ {
			if utils.SliceContains(TL05text, TL02text[i]) {
			} else {
				log.Println("SAL records: " + TL02text[i] + " in TL02 but not in TL05")
			}
		}
	} else if len(TL05text) > len(TL02text) {
		for i := 0; i < len(TL05text); i++ {
			if utils.SliceContains(TL02text, TL05text[i]) {
			} else {
				log.Println("SAL records: " + TL05text[i] + " in TL05 but not in TL02")
			}
		}
	}
}

func CheckTL05ileSize() {
	path := "TL05/SFTP_download"
	dir, err := ioutil.ReadDir(path)
	var counter int
	if err != nil {
		log.Println("Cannot read directory when update octopus upload")
		return
	}
	for _, f := range dir {
		name := f.Name()
		info, err := os.Stat(path + "/" + name)
		if err != nil {
			log.Println(err)
		}
		fileSize := info.Size()
		if fileSize < 163 {
			log.Println("File Size is abnormal, Please Check the file: ", name)
		}
		counter++
	}
	if counter != 50 {
		log.Println("No. of Files is not matched. There is missing files")
		// CheckTL05MissingFile() //check/list out missing files
	}
}

func CheckTL05MissingFile() {
	path := "TL05/SFTP_download"
	dir, err := ioutil.ReadDir(path)
	var counter, counter1, counter2 int
	var namePrevious05, namePrevious35 string
	var nameNow05, nameNow35 string
	var fn string
	var cal int
	first := "22"
	if err != nil {
		log.Println("Cannot read directory when update octopus upload")
		return
	}
	os.Remove("TL05/MissingFiles.txt") //Remove before checking another day
	missingFiles := "TL05/MissingFiles.txt"
	fileName, err := os.OpenFile(missingFiles, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Open missingFiles.txt failed. ", err)
	}
	for _, f := range dir {
		name := f.Name()
		if strings.Contains(name, "05.csv") {
			nameNow05 = name
			counter1++
			if counter1 > 1 {
				now, err := strconv.Atoi(nameNow05[38:40])
				if err != nil {
					log.Println("convert now to int failed")
				}
				previous, err := strconv.Atoi(namePrevious05[38:40])
				if err != nil {
					log.Println("convert previous to int failed")
				}
				if now != 0 {
					if now-previous != 1 {
						for a := 1; a < now-previous; a++ {
							if previous < 10 {
								if (previous + a) < 10 {
									fn = name[0:38] + "0" + strconv.Itoa(previous+a) + "05.csv"
									fileName.WriteString(fn + "\n")
								} else {
									fn = name[0:38] + strconv.Itoa(previous+a) + "05.csv"
									fileName.WriteString(fn + "\n")
								}
							} else {
								fn = name[0:38] + strconv.Itoa(previous+a) + "05.csv"
								fileName.WriteString(fn + "\n")
							}
						}
					}
				} else { // case: now == 0
					if previous-now == 22 {
						today, err := strconv.Atoi(name[29:37])
						if err != nil {
							log.Println(err)
						}
						yesterday := strconv.Itoa(today - 1)
						fn = name[0:29] + yesterday + "_2305.csv"
						fileName.WriteString(fn + "\n")
					}
				}
			}
			if counter1 == 1 && !strings.Contains(name, first+"05.csv") {
				hourName, err := strconv.Atoi(nameNow05[38:40])
				if err != nil {
					log.Println("convert hourName from string to int failed", err)
				}
				if hourName == 23 {
					fn = nameNow05[0:38] + first + "05.csv"
					fileName.WriteString(fn + "\n")
				} else {
					// case if 0
					lastDay, err := strconv.Atoi(nameNow05[29:37])
					if err != nil {
						log.Println(err)
					}
					lastDayString := strconv.Itoa(lastDay - 1)
					fn = nameNow05[0:29] + lastDayString + "_" + first + "05.csv"
					fileName.WriteString(fn + "\n")
					fn = nameNow05[0:29] + lastDayString + "_" + "23" + "05.csv"
					fileName.WriteString(fn + "\n")
					if hourName > 0 {
						for i := 0; i < hourName; i++ {
							fnHour := strconv.Itoa(hourName - (hourName - i))
							if (hourName - (hourName - i)) < 10 {
								fn = nameNow05[0:38] + "0" + fnHour + "05.csv"
								fileName.WriteString(fn + "\n")
							} else {
								fn = nameNow05[0:38] + fnHour + "05.csv"
								fileName.WriteString(fn + "\n")
							}
						}
					}
				}
			}
		} else if strings.Contains(name, "35.csv") {
			nameNow35 = name
			counter2++
			if counter2 > 1 {
				now, err := strconv.Atoi(nameNow35[38:40])
				if err != nil {
					log.Println("convert now to int failed")
				}
				previous, err := strconv.Atoi(namePrevious35[38:40])
				if err != nil {
					log.Println("convert previous to int failed")
				}
				if now != 0 {
					if now-previous != 1 {
						for a := 1; a < now-previous; a++ {
							if previous < 10 {
								if (previous + a) < 10 {
									fn = name[0:38] + "0" + strconv.Itoa(previous+a) + "35.csv"
									// log.Println("Missing files: ", fn)
									fileName.WriteString(fn + "\n")
								} else {
									fn = name[0:38] + strconv.Itoa(previous+a) + "35.csv"
									// log.Println("Missing files: ", fn)
									fileName.WriteString(fn + "\n")
								}
							} else {
								fn = name[0:38] + strconv.Itoa(previous+a) + "35.csv"
								// log.Println("Missing files: ", fn)
								fileName.WriteString(fn + "\n")
							}
						}
					}
				} else { // case: now == 0
					if previous-now == 22 {
						today, err := strconv.Atoi(name[29:37])
						if err != nil {
							log.Println(err)
						}
						yesterday := strconv.Itoa(today - 1)
						fn = name[0:29] + yesterday + "_2335.csv"
						// log.Println("Missing files: ", fn)
						fileName.WriteString(fn + "\n")
					}
				}
			}
			if counter2 == 1 && !strings.Contains(name, first+"35.csv") {
				hourName, err := strconv.Atoi(nameNow35[38:40])
				if err != nil {
					log.Println("convert hourName from string to int failed", err)
				}
				if hourName == 23 {
					fn = nameNow35[0:38] + first + "35.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				} else {
					// case if 0
					// fmt.Println(nameNow05[14:22])
					lastDay, err := strconv.Atoi(nameNow35[29:37])
					if err != nil {
						log.Println(err)
					}
					lastDayString := strconv.Itoa(lastDay - 1)
					fn = nameNow35[0:29] + lastDayString + "_" + first + "35.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
					fn = nameNow35[0:29] + lastDayString + "_" + "23" + "35.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
					if hourName > 0 {
						for i := 0; i < hourName; i++ {
							fnHour := strconv.Itoa(hourName - (hourName - i))
							if (hourName - (hourName - i)) < 10 {
								fn = nameNow35[0:38] + "0" + fnHour + "35.csv"
								// log.Println("Missing files: ", nameNow35[0:23]+"0"+fnHour+"35.csv")
								fileName.WriteString(fn + "\n")
							} else {
								fn = nameNow35[0:38] + fnHour + "35.csv"
								// log.Println("Missing files: ", nameNow35[0:23]+fnHour+"35.csv")
								fileName.WriteString(fn + "\n")
							}
						}
					}
				}
			}
		}
		counter++
		namePrevious05 = nameNow05
		namePrevious35 = nameNow35
	}
	if counter1 < 25 {
		if counter1 == 0 {
			log.Println("All 05.csv is missing")
		}
		hour, err := strconv.Atoi(namePrevious05[38:40])
		if err != nil {
			log.Println("Convert hour from string to int failed. ", err)
		}
		if hour != 22 {
			for i := 1; i <= 22-hour; i++ {
				hourName := strconv.Itoa(hour + i)
				if hour+i < 10 {
					fn = namePrevious05[0:38] + "0" + hourName + "05.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				} else {
					fn = namePrevious05[0:38] + hourName + "05.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				}
			}
		}
		cal = 25 - counter1
		log.Println("There are ", cal, " 05.csv is missing")
	}
	if counter2 < 25 {
		if counter2 == 0 {
			log.Println("All 35.csv is missing")
		}
		hour, err := strconv.Atoi(namePrevious35[38:40])
		if err != nil {
			log.Println("Convert hour from string to int failed. ", err)
		}
		if hour != 22 {
			for i := 1; i <= 22-hour; i++ {
				hourName := strconv.Itoa(hour + i)
				if hour+i < 10 {
					fn = namePrevious35[0:38] + "0" + hourName + "35.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				} else {
					fn = namePrevious35[0:38] + hourName + "35.csv"
					// log.Println("Missing files: ", fn)
					fileName.WriteString(fn + "\n")
				}
			}
		}
		cal = 25 - counter2
		log.Println("There are ", cal, " 35.csv is missing")
	}
}

func SortTL05MissingFile() {
	var text []string
	fileName := "TL05/MissingFiles.txt"
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

func PrintTL05MissingFile() {
	var text []string
	fileName := "TL05/MissingFiles.txt"
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
		log.Println("Missing files: " + each_ln)
	}
}

// func CheckTL05MissingFile() {
// 	// path := "TL02/SFTP_file"
// 	path := "TL05/SFTP_download"
// 	dir, err := ioutil.ReadDir(path)
// 	var counter, counter1, counter3 int
// 	var namePrevious05, namePrevious35 string
// 	var nameNow05, nameNow35 string
// 	var fn string
// 	var cal int
// 	first := "22"
// 	// last := "22"
// 	if err != nil {
// 		log.Println("Cannot read directory when update octopus upload")
// 		return
// 	}
// 	os.Remove("TL05/MissingFiles.txt") //Remove before checking another day
// 	missingFiles := "TL05/MissingFiles.txt"
// 	fileName, err := os.OpenFile(missingFiles, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Println("Open missingFiles.txt failed. ", err)
// 	}
// 	// if the first/last file is missed, can't println the file name
// 	for _, f := range dir {
// 		name := f.Name()
// 		if strings.Contains(name, "05.csv") {
// 			nameNow05 = name
// 			counter1++
// 			if counter1 > 1 {
// 				now, err := strconv.Atoi(nameNow05[23:25])
// 				if err != nil {
// 					log.Println("convert now to int failed")
// 				}
// 				previous, err := strconv.Atoi(namePrevious05[23:25])
// 				if err != nil {
// 					log.Println("convert previous to int failed")
// 				}
// 				if now != 0 {
// 					if now-previous != 1 {
// 						for a := 1; a < now-previous; a++ {
// 							if previous < 10 {
// 								if (previous + a) < 10 {
// 									fn = name[0:23] + "0" + strconv.Itoa(previous+a) + "05.csv"
// 									// log.Println("Missing files: ", fn)
// 									fileName.WriteString(fn + "\n")
// 								} else {
// 									fn = name[0:23] + strconv.Itoa(previous+a) + "05.csv"
// 									// log.Println("Missing files: ", fn)
// 									fileName.WriteString(fn + "\n")
// 								}
// 							} else {
// 								fn = name[0:23] + strconv.Itoa(previous+a) + "05.csv"
// 								// log.Println("Missing files: ", fn)
// 								fileName.WriteString(fn + "\n")
// 							}
// 						}
// 					}
// 				} else { // case: now == 0
// 					if previous-now == 22 {
// 						today, err := strconv.Atoi(name[14:22])
// 						if err != nil {
// 							log.Println(err)
// 						}
// 						yesterday := strconv.Itoa(today - 1)
// 						fn = name[0:14] + yesterday + "_2305.csv"
// 						// log.Println("Missing files: ", fn)
// 						fileName.WriteString(fn + "\n")
// 					}
// 				}
// 			}
// 			if counter1 == 1 && !strings.Contains(name, first+"05.csv") {
// 				hourName, err := strconv.Atoi(nameNow05[23:25])
// 				if err != nil {
// 					log.Println("convert hourName from string to int failed", err)
// 				}
// 				if hourName == 23 {
// 					fn = nameNow05[0:23] + first + "05.csv"
// 					// log.Println("Missing files: ", fn)
// 					fileName.WriteString(fn + "\n")
// 				} else {
// 					// case if 0
// 					// fmt.Println(nameNow05[14:22])
// 					lastDay, err := strconv.Atoi(nameNow05[14:22])
// 					if err != nil {
// 						log.Println(err)
// 					}
// 					lastDayString := strconv.Itoa(lastDay - 1)
// 					fn = nameNow05[0:14] + lastDayString + "_" + first + "05.csv"
// 					// log.Println("Missing files: ", fn)
// 					fileName.WriteString(fn + "\n")
// 					fn = nameNow05[0:14] + lastDayString + "_" + "23" + "05.csv"
// 					// log.Println("Missing files: ", fn)
// 					fileName.WriteString(fn + "\n")
// 					if hourName > 0 {
// 						for i := 0; i < hourName; i++ {
// 							fnHour := strconv.Itoa(hourName - (hourName - i))
// 							if (hourName - (hourName - i)) < 10 {
// 								fn = nameNow05[0:23] + "0" + fnHour + "05.csv"
// 								// log.Println("Missing files: ", nameNow05[0:23]+"0"+fnHour+"05.csv")
// 								fileName.WriteString(fn + "\n")
// 							} else {
// 								fn = nameNow05[0:23] + fnHour + "05.csv"
// 								// log.Println("Missing files: ", nameNow05[0:23]+fnHour+"05.csv")
// 								fileName.WriteString(fn + "\n")
// 							}
// 						}
// 					}
// 				}
// 			}
// 		} else if strings.Contains(name, "35.csv") {
// 			nameNow35 = name
// 			counter3++
// 			if counter3 > 1 {
// 				now, err := strconv.Atoi(nameNow35[23:25])
// 				if err != nil {
// 					log.Println("convert now to int failed")
// 				}
// 				previous, err := strconv.Atoi(namePrevious35[23:25])
// 				if err != nil {
// 					log.Println("convert previous to int failed")
// 				}
// 				if now != 0 {
// 					if now-previous != 1 {
// 						for a := 1; a < now-previous; a++ {
// 							if previous < 10 {
// 								if (previous + a) < 10 {
// 									fn = name[0:23] + "0" + strconv.Itoa(previous+a) + "35.csv"
// 									// log.Println("Missing files: ", fn)
// 									fileName.WriteString(fn + "\n")
// 								} else {
// 									fn = name[0:23] + strconv.Itoa(previous+a) + "35.csv"
// 									// log.Println("Missing files: ", fn)
// 									fileName.WriteString(fn + "\n")
// 								}
// 							} else {
// 								fn = name[0:23] + strconv.Itoa(previous+a) + "35.csv"
// 								// log.Println("Missing files: ", fn)
// 								fileName.WriteString(fn + "\n")
// 							}
// 						}
// 					}
// 				} else { // case: now == 0
// 					if previous-now == 22 {
// 						today, err := strconv.Atoi(name[14:22])
// 						if err != nil {
// 							log.Println(err)
// 						}
// 						yesterday := strconv.Itoa(today - 1)
// 						fn = name[0:14] + yesterday + "_2335.csv"
// 						// log.Println("Missing files: ", fn)
// 						fileName.WriteString(fn + "\n")
// 					}
// 				}
// 			}
// 			if counter3 == 1 && !strings.Contains(name, first+"35.csv") {
// 				hourName, err := strconv.Atoi(nameNow35[23:25])
// 				if err != nil {
// 					log.Println("convert hourName from string to int failed", err)
// 				}
// 				if hourName == 23 {
// 					fn = nameNow35[0:23] + first + "35.csv"
// 					// log.Println("Missing files: ", fn)
// 					fileName.WriteString(fn + "\n")
// 				} else {
// 					// case if 0
// 					// fmt.Println(nameNow05[14:22])
// 					lastDay, err := strconv.Atoi(nameNow35[14:22])
// 					if err != nil {
// 						log.Println(err)
// 					}
// 					lastDayString := strconv.Itoa(lastDay - 1)
// 					fn = nameNow35[0:14] + lastDayString + "_" + first + "35.csv"
// 					// log.Println("Missing files: ", fn)
// 					fileName.WriteString(fn + "\n")
// 					fn = nameNow35[0:14] + lastDayString + "_" + "23" + "35.csv"
// 					// log.Println("Missing files: ", fn)
// 					fileName.WriteString(fn + "\n")
// 					if hourName > 0 {
// 						for i := 0; i < hourName; i++ {
// 							fnHour := strconv.Itoa(hourName - (hourName - i))
// 							if (hourName - (hourName - i)) < 10 {
// 								fn = nameNow35[0:23] + "0" + fnHour + "35.csv"
// 								// log.Println("Missing files: ", nameNow35[0:23]+"0"+fnHour+"35.csv")
// 								fileName.WriteString(fn + "\n")
// 							} else {
// 								fn = nameNow35[0:23] + fnHour + "35.csv"
// 								// log.Println("Missing files: ", nameNow35[0:23]+fnHour+"35.csv")
// 								fileName.WriteString(fn + "\n")
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 		counter++
// 		namePrevious05 = nameNow05
// 		namePrevious35 = nameNow35
// 	}
// 	if counter1 < 25 {
// 		if counter1 == 0 {
// 			log.Println("All 05.csv is missing")
// 		}
// 		hour, err := strconv.Atoi(namePrevious05[23:25])
// 		if err != nil {
// 			log.Println("Convert hour from string to int failed. ", err)
// 		}
// 		if hour != 22 {
// 			for i := 1; i <= 22-hour; i++ {
// 				hourName := strconv.Itoa(hour + i)
// 				if hour+i < 10 {
// 					fn = namePrevious05[0:23] + "0" + hourName + "05.csv"
// 					// log.Println("Missing files: ", fn)
// 					fileName.WriteString(fn + "\n")
// 				} else {
// 					fn = namePrevious05[0:23] + hourName + "05.csv"
// 					// log.Println("Missing files: ", fn)
// 					fileName.WriteString(fn + "\n")
// 				}
// 			}
// 		}
// 		cal = 25 - counter1
// 		log.Println("There are ", cal, " 05.csv is missing")
// 	}
// 	if counter3 < 25 {
// 		if counter3 == 0 {
// 			log.Println("All 35.csv is missing")
// 		}
// 		hour, err := strconv.Atoi(namePrevious35[23:25])
// 		if err != nil {
// 			log.Println("Convert hour from string to int failed. ", err)
// 		}
// 		if hour != 22 {
// 			for i := 1; i <= 22-hour; i++ {
// 				hourName := strconv.Itoa(hour + i)
// 				if hour+i < 10 {
// 					fn = namePrevious35[0:23] + "0" + hourName + "35.csv"
// 					// log.Println("Missing files: ", fn)
// 					fileName.WriteString(fn + "\n")
// 				} else {
// 					fn = namePrevious35[0:23] + hourName + "35.csv"
// 					// log.Println("Missing files: ", fn)
// 					fileName.WriteString(fn + "\n")
// 				}
// 			}
// 		}
// 		cal = 25 - counter3
// 		log.Println("There are ", cal, " 35.csv is missing")
// 	}
// }
