package main

import (
	"fmt"

	"github.com/thelioncape/WinPE-OSChooser/getoslist"
)

func main() {
	fmt.Println("Choose an OS")
	getoslist.PrintOSList()
	fmt.Print("Please enter your choice and then press enter: ")
	var choice string
	fmt.Scanf("%s", &choice)
	fmt.Println("Booting option", choice)
}
