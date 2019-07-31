package main

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/thelioncape/WinPE-OSChooser/getoslist"
)

func main() {
	var location string
	if location == "" {
		fmt.Println("Choose an OS")
		data := getoslist.PrintOSList()
		choice, data := getChoice(data)
		intchoice := int(choice)
		for index, element := range data.OperatingSystems {
			if intchoice == (index + 1) {
				location = element.Location
			}
		}
		fmt.Println("Location:", location)
	}

	server := getoslist.GetNextServer()

	cmd := exec.Command("net use W: \\\\" + server + "\\mount")
	cmd.Run()

	cmd = exec.Command(location + "\\setup.exe")
	cmd.Run()
}

func getChoice(data getoslist.Osdata) (uint8, getoslist.Osdata) {
	fmt.Print("Please enter your choice and then press enter: ")
	var choice string
	fmt.Scanf("%s", &choice)
	intchoice, _ := strconv.ParseInt(choice, 10, 8)

	uintchoice := uint8(intchoice)

	fmt.Println("Attempting to boot option", uintchoice)

	return uintchoice, data
}
