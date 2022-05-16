package main

import (
	data "ajl/tritip/data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

func getOrders() (data.ListResponse, error) {
	// get orders for update
	client := &http.Client{}
	orders := data.ListResponse{}
	key, secret := data.GetApiSecret()

	req, err := http.NewRequest(http.MethodGet, "https://ssapi.shipstation.com/orders?orderStatus=awaiting_shipment", nil)

	req.SetBasicAuth(key, secret)

	fmt.Printf("req: %v\n", req)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("request error: %v\n", err)
	}

	respJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read i/o error ::", err)
	}

	err = json.Unmarshal([]byte(respJSON), &orders)
	fmt.Printf("orders: %v\n", orders)
	if err != nil {
		fmt.Printf("json unmarshalling error :: %v\n", err)
	}

	fmt.Printf("resp: %v", resp)

	defer resp.Body.Close()

	return orders, err
}

func orderUpdate(o []*data.OrderRecordInput) {
	// Shipstation API Call
	fmt.Printf("order update %v\n", o)
}

func csvReader(s string) ([]*data.OrderRecordInput, error) {
	recordFile, err := os.Open(s)
	if err != nil {
		fmt.Println("Reader Error occured! ::", err)
	}

	records := []*data.OrderRecordInput{}

	if err := gocsv.UnmarshalFile(recordFile, &records); err != nil {
		fmt.Println("Unmarshalling Error occured! ::", err)
	}
	defer recordFile.Close()

	return records, err
}

func initializeCSV() {
	localString := "./"
	input := strings.Join(os.Args[1:], "")
	fileName := localString + input

	records, err := csvReader(fileName)
	if err != nil {
		fmt.Println("Can't initialize reader ::", err)
	}

	orderUpdate(records)

}

func main() {

	initializeCSV()

	getOrders()
}
