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
	blk, err := os.Open(path)
	customBlk := custommodels.File{
		blk,
	}
	errorchecker.CheckFileError(err)
	return customBlk
}
