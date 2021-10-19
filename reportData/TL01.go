package reportData

import (
	"io/ioutil"
	"log"
	"os"
)

func CheckTL01() {
	CheckTL01FileSize()
}

func CheckTL01FileSize() {
	log.Println()
	log.Println("----------Start Checking TL01----------")
	defer log.Println("----------End Checking TL01----------")
	path := "TL01/SFTP_download"
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
		log.Println("No TL01 Master Report")
	}
}
