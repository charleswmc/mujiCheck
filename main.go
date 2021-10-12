package main

import "mujiCheck/reportData"

func main() {
	reportData.CheckTL02()
	reportData.SaveToSortTxnRecordFile()
	reportData.SortTxnRecordFile()
	reportData.CheckTL02ALL()
	// reportData.CheckTL05()
}
