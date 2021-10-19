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
	}
}
