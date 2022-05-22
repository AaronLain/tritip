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
	"time"

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
		// make sure the order is in the queue when it updates
		thisOrder.OrderStatus = "awaiting_shipment"

		for _, zip := range zips {
			firstFive := firstFiveZip(thisOrder.ShipTo.PostalCode)
			if firstFive == zip.PostalCode {
				// assign ice profile
				thisOrder.AdvancedOptions.CustomField3 = zip.CustomField3
				if zip.CustomField3 == "Profile 5" {
					thisOrder.ServiceCode = "ups_next_day_air_saver"
					thisOrder.AdvancedOptions.CustomField3 = "Profile 4"
				}
				// add the special API tag so we know the order was touched
				thisOrder.TagIds = append(thisOrder.TagIds, 122060)
				updatedOrders = append(updatedOrders, thisOrder)
			}
		}

	}

	// fmt.Printf("updated orders: %v\n length: %v\n", updatedOrders, len(updatedOrders))

	return updatedOrders, nil
}

func postOrders(ordersQueue data.OrderRecordOutput) (int, error) {
	client := &http.Client{}
	orders, err := json.Marshal(ordersQueue)
	if err != nil {
		fmt.Printf("json marshal error: %v\n", err)
	}
	key, secret := data.GetApiSecret()

	req, err := http.NewRequest(http.MethodPost, "https://ssapi.shipstation.com/orders/createorder", bytes.NewBuffer(orders))

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

	} else if resp.StatusCode == 429 {
		time.Sleep(50 * time.Second)
		fmt.Println("API Rate limit exceeded, sleeping...")
		return resp.StatusCode, err
	} else {
		fmt.Printf("Response Code: %v\n", resp.StatusCode)

		return resp.StatusCode, err
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func sleepAlert() {
	time.Sleep(1510 * time.Millisecond)
	fmt.Println("Sleeping...")
}

func updateOrders(orders []data.OrderRecordOutput) error {
	for _, order := range orders {
		sleepAlert()
		status, err := postOrders(order)
		if err != nil {
			fmt.Printf("failed to post Orders: %v\n", err)
			fmt.Printf("recieved status code: %v\n", status)
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
