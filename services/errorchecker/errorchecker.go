package errorchecker

import (
	"fmt"
	"io"
)

/*
CheckFileError ...
*/
func CheckFileError(err error) {
	if err != nil {
		if err == io.EOF {
			fmt.Println("End of file")
		}
		panic(err)
	}
}
