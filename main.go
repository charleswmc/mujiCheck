package main

import (
	"log"
	"mujiCheck/reportData"
)

func main() {
	reportData.CheckTL01()

	log.Println()
	log.Println("----------Start Checking TL02 files----------")
	reportData.CheckTL02FileSize()
	reportData.CheckTL02MissingFile()
	reportData.SortTL02MissingFile()
	reportData.PrintTL02MissingFile()
	log.Println("----------End Checking TL02 files----------")
	log.Println()
	// reportData.GetTL02MissingFile("05.csv")
	reportData.CheckTL02()
	reportData.SaveToSortTxnRecordFile()
	reportData.SortTxnRecordFile()
	reportData.CheckTL02ALL()

	reportData.CheckTL04()

	log.Println()
	log.Println("----------Start Checking TL05 files----------")
	reportData.CheckTL05ileSize()
	reportData.CheckTL05MissingFile()
	reportData.SortTL05MissingFile()
	reportData.PrintTL05MissingFile()
	log.Println("----------End Checking TL05 files----------")
	log.Println()
	reportData.CheckTL05()
	// reportData.CheckTL05SAL()
}
