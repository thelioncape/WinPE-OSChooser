package getoslist

import (
	"encoding/json"
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

// GetOSList Returns the list of operating systems gathered from http /mount/Windows.json
func GetOSList() Osdata {
	gw, _ := gateway.DiscoverGateway()
	list := downloadOSList(getNextServer(gw.String()))
	data := Osdata{}
	err := json.Unmarshal(list, &data)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	return data
}

func getNextServer(gw string) string {
	gw = gw[:len(gw)-1]
	gw = gw + "2"
	return gw
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
