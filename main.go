package main

import (
	"fmt"

	"github.com/thelioncape/WinPE-OSChooser/getoslist"
)

func main() {
	fmt.Println("Please choose an OS")
	getoslist.GetOSList()

}
