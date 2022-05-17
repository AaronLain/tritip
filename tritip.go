package main

import (
	data "ajl/tritip/data"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	fmt.Printf("tags: %v\n", tags)

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

		fmt.Printf("request error: %v\n", err)
	}

	respJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read i/o error ::\n", err)
	}

	err = json.Unmarshal([]byte(respJSON), &orders)
	if err != nil {
		fmt.Printf("json unmarshalling error :: %v\n", err)
	}

	defer resp.Body.Close()

	return orders, nil
}

func firstFiveZip(zip string) string {
	counter := 0
	for i := range zip {
		if counter == 5 {
			zip = zip[:i]
		}
		counter++
	}

	return zip
}

func iceProfileAssignment(zips []*data.OrderRecordInput) ([]data.OrderRecordOutput, error) {
	// Adds Ice Profile to Orders
	ssOrders, err := getOrders()
	if err != nil {

		fmt.Printf("couldn't get orders:: %v\n", err)
	}
	updatedOrders := []data.OrderRecordOutput{}
	for _, thisOrder := range ssOrders.Orders {
		// add the special API tag so we know the order was touched
		thisOrder.TagIds = append(thisOrder.TagIds, 122060)
		// make sure the order is in the queue when it updates
		thisOrder.OrderStatus = "awaiting_shipment"
		for _, zip := range zips {
			firstFive := firstFiveZip(thisOrder.ShipTo.PostalCode)
			if firstFive == zip.PostalCode {
				// assign ice profile
				thisOrder.AdvancedOptions.CustomField3 = zip.CustomField3
				updatedOrders = append(updatedOrders, thisOrder)
			}
		}
	}

	// fmt.Printf("updated orders: %v\n length: %v\n", updatedOrders, len(updatedOrders))

	return updatedOrders, nil
}

func postOrders(ordersQueue []data.OrderRecordOutput) (int, error) {
	client := &http.Client{}
	orders, err := json.Marshal(ordersQueue)
	if err != nil {
		fmt.Printf("json marshal error: %v\n", err)
	}
	key, secret := data.GetApiSecret()

	req, err := http.NewRequest(http.MethodPost, "https://ssapi.shipstation.com/orders/createorders", bytes.NewBuffer(orders))

	req.SetBasicAuth(key, secret)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {

		fmt.Printf("request error: %v\n", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {

			fmt.Printf("failed to read response: %v\n", err)
		}

		jsonStr := string(body)
		fmt.Printf("Response: %v\n", jsonStr)

		return 201, nil

	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {

			fmt.Printf("failed to read response: %v\n", err)
		}

		jsonStr := string(body)
		fmt.Printf("Response: %v\n", jsonStr)

		return resp.StatusCode, err
	}

}

func updateOrders(orders []data.OrderRecordOutput) error {
	ordersQueue := []data.OrderRecordOutput{}
	for i, order := range orders {
		orderCount := 0
		ordersQueue = append(ordersQueue, order)
		fmt.Printf("count: %v\n", orderCount)
		// API limit is 100 orders
		if orderCount == 99 || i == len(ordersQueue)-1 {
			// update orders en masse
			status, err := postOrders(ordersQueue)
			if err != nil {

				fmt.Printf("failed to post Orders: %v\n", err)
				fmt.Printf("recieved status code: %v\n", status)
			}
			orderCount = 0
		}
	}
	return nil
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

		fmt.Printf("Can't assign ice profiles:: %v\n", err)
	}

	if err := updateOrders(orders); err != nil {

		fmt.Printf("Update Failed:: %v\n", err)
	}

}

func main() {

	initialize()

}
