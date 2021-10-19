package reportData

import (
	"io/ioutil"
	"log"
	"os"
)

func CheckTL04() {
	CheckTL04FileSize()
}

func CheckTL04FileSize() {
	log.Println()
	log.Println("----------Start Checking TL04----------")
	defer log.Println("----------End Checking TL04----------")
	path := "TL04/SFTP_download"
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
		if fileSize == 0 {
			log.Println("File Size is empty, Please Check the file: ", name)
		}
		counter++
	}
	if counter < 1 {
		log.Println("No TL04 Daily Inventory Report")
	}
}
