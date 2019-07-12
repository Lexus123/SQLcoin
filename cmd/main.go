package main

import (
	"sqlcoin/services/blockreader"
	"sqlcoin/services/fileopener"
	"sqlcoin/services/flagger"
	// "sqlcoin/server"
)

func main() {
	// server.Host()
	_, _, blkPath := flagger.ReturnFlags()
	file := fileopener.OpenFile(blkPath)
	blockreader.StartReading(file)
}
