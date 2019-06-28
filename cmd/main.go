package main

import (
	"sqlcoin/services/blockreader"
	"sqlcoin/services/fileopener"
	"sqlcoin/services/flagger"
)

func main() {
	dbCreds, dbName, blkPath := flagger.ReturnFlags()
	file := fileopener.OpenFile(blkPath)
	blockreader.StartReading(file, dbCreds, dbName)
}
