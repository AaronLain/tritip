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

func getTags() {
	client := &http.Client{}
	key, secret := data.GetApiSecret()
	req, err := http.NewRequest(http.MethodGet, "https://ssapi.shipstation.com/accounts/listtags", nil)
	req.SetBasicAuth(key, secret)
	tags := []data.Tag{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("request error: %v\n", err)
	}

	respJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read i/o error ::\n", err)
	}

	err = json.Unmarshal([]byte(respJSON), &tags)
	if err != nil {
		fmt.Printf("json unmarshalling error :: %v\n", err)
	}

	fmt.Printf("tags: %v", tags)

	defer resp.Body.Close()
}

func getOrders() (data.OrderRecordOutputResp, error) {
	// get orders for update
	client := &http.Client{}
	orders := data.OrderRecordOutputResp{}
	key, secret := data.GetApiSecret()

	req, err := http.NewRequest(http.MethodGet, "https://ssapi.shipstation.com/orders?orderStatus=awaiting_shipment", nil)

	req.SetBasicAuth(key, secret)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("request error: %v\n", err)
	}

	respJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read i/o error ::\n", err)
	}

	err = json.Unmarshal([]byte(respJSON), &orders)
	fmt.Printf("orders: %v\n", orders)
	if err != nil {
		fmt.Printf("json unmarshalling error :: %v\n", err)
	}

	defer resp.Body.Close()

	return orders, err
}

func firstFiveZip(zip string) string {
	counter := 0
	for i := range zip {
		if counter == 5 {
			zip = zip[:i]
		}
		counter++
	}
	fmt.Printf("zip: %v\n", zip)

	return zip
}

func iceProfileAssignment(zips []*data.OrderRecordInput) ([]data.OrderRecordOutput, error) {
	// Adds Ice Profile to Orders
	ssOrders, err := getOrders()
	if err != nil {
		log.Fatal(err)
		fmt.Printf("couldn't get orders:: %v", err)
	}
	updatedOrders := []data.OrderRecordOutput{}
	for _, order := range ssOrders.Orders {
		thisOrder := order
		// add the special API tag so we know the order was touched
		thisOrder.TagIds = append(thisOrder.TagIds, 122060)
		for _, zip := range zips {
			firstFive := firstFiveZip(thisOrder.ShipTo.PostalCode)
			if firstFive == zip.PostalCode {
				thisOrder.AdvancedOptions.CustomField3 = zip.CustomField3
				updatedOrders = append(updatedOrders, thisOrder)
			}
		}
	}

	fmt.Printf("updated orders: %v\n", updatedOrders)

	return updatedOrders, nil
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

	return records, nil
}

func initialize() {
	localString := "./"
	input := strings.Join(os.Args[1:], "")
	fileName := localString + input

	records, err := csvReader(fileName)
	if err != nil {
		fmt.Println("Can't initialize reader ::", err)
	}

	orders, err := iceProfileAssignment(records)
	if err != nil {
		log.Fatal(err)
	}

	getTags()
	fmt.Printf("Orders: %v\n", orders)

}

func main() {

	initialize()

}
