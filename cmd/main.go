package main

import (
	// "sqlcoin/services/blockreader"
	// "sqlcoin/services/fileopener"
	// "sqlcoin/services/flagger"
	"sqlcoin/server"
)

func main() {
	server.Host()
	// dbCreds, dbName, blkPath := flagger.ReturnFlags()
	// file := fileopener.OpenFile(blkPath)
	// blockreader.StartReading(file, dbCreds, dbName)
}
