package main

import (
	data "ajl/tritip/data"
	"fmt"

	"os"

	"strings"

	"github.com/gocarina/gocsv"
)

func getOrders() {
	// get orders for update
}

func orderUpdate(o []*data.OrderRecordInput) {
	// Shipstation API Call
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
}
