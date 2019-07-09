package fileopener

import (
	"os"
	"sqlcoin/services/custommodels"
	"sqlcoin/services/errorchecker"
)

/*
OpenFile ...
*/
func OpenFile(path string) custommodels.File {
	fullpath := "./testfiles/" + path
	blk, err := os.Open(fullpath)
	customBlk := custommodels.File{
		blk,
	}
	errorchecker.CheckFileError(err)
	return customBlk
}
