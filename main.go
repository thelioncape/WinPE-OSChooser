package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jackpal/gateway"
)

func main() {
	fmt.Println("Please choose an OS")
	gw, _ := gateway.DiscoverGateway()
	ns := getNextServer(gw.String())
	getOSList(ns)

}

func getNextServer(gw string) string {
	gw = gw[:len(gw)-1]
	gw = gw + "2"
	return gw
}

func getOSList(ns string) map[string]string {
	gw, _ := gateway.DiscoverGateway()
	list := downloadOSList(getNextServer(gw.String()))
	data := make(map[string]string)
	err := json.Unmarshal(list, &data)

	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	return data
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
