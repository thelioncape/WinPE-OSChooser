package getoslist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jackpal/gateway"
)

// Osdata types the JSON input from mount/Windows.json
type Osdata struct {
	OperatingSystems []struct {
		Name     string `json:"Name"`
		Location string `json:"Location"`
	} `json:"Operating Systems"`
}

// PrintOSList prints the list of Operating Systems on to the console
func PrintOSList() Osdata {
	data := getOSList()
	var index uint8
	for _, element := range data.OperatingSystems {
		index++
		fmt.Printf("%d. %s\n", index, element.Name)
	}
	return data
}

// GetOSList Returns the list of operating systems gathered from http /mount/Windows.json
func getOSList() Osdata {
	list := downloadOSList(GetNextServer())
	data := Osdata{}
	err := json.Unmarshal(list, &data)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	return data
}

// GetNextServer displays the .2 of the current subnet assuming the gateway is .1
func GetNextServer() string {
	gw, _ := gateway.DiscoverGateway()
	strgw := gw[:len(gw)-1].String()
	strgw = strgw + "2"
	return strgw
}

func downloadOSList(ns string) []byte {
	connstring := "http://" + ns + "/mount/Windows.json"

	res, err := http.Get(connstring)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	list, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("An error occurred reading the response from the server:")
		log.Fatal(err)
	}

	return list
}
