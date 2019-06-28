package flagger

import (
	"flag"
)

/*
ReturnFlags ...
*/
func ReturnFlags() (string, string, string) {
	var dbCreds, dbName, blkPath string

	flag.StringVar(&dbCreds, "dbcreds", "", "(REQUIRED) For example: user:password")
	flag.StringVar(&dbName, "dbname", "", "(REQUIRED) For example: databasename")
	flag.StringVar(&blkPath, "blockfile", "", "(REQUIRED) For example: ~/bitcoin/blocks/blk00001.dat")

	flag.Parse()
	return dbCreds, dbName, blkPath
}
