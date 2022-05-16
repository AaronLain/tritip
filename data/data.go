package data

type OrderRecordInput struct {
	OrderNum      string  `csv:"Order - Number"`
	CustomField3  string  `csv:"CustomField3"`
	AvgTemp       float64 `csv:"AvgTemp"`
	City          string  `csv:"Ship To - City"`
	State         string  `csv:"Ship To - State"`
	PostalCode    string  `csv:"Ship To - Postal Code"`
	CountryCode   string  `csv:"Shipping Country"`
	ItemSKU       string  `csv:"Lineitem sku"`
	ItemName      string  `csv:"Lineitem name"`
	ItemUnitPrice string  `csv:"Lineitem price"`
}

type BillTo struct {
	Name            string
	Company         string
	Street1         string
	Street2         string
	Street3         string
	City            string
	State           string
	PostalCode      string
	Country         string
	Phone           string
	Residential     bool
	AddressVerified bool
}

type ShipTo struct {
	Name            string
	Company         string
	Street1         string
	Street2         string
	Street3         string
	City            string
	State           string
	PostalCode      string
	Country         string
	Phone           string
	Residential     bool
	AddressVerified bool
}

type Weight struct {
}

type Dimensions struct {
}

type InternationalOptions struct {
}

type AdvancedOptions struct {
}

type LineItem struct {
}

type LineItemOptions struct {
	Name  string
	Value string
}

type OrderRecordOutput struct {
	OrderId          string
	OrderNumber      string
	OrderKey         string
	OrderDate        string
	CreateDate       string
	ModifyDate       string
	PaymentDate      string
	ShipByDate       string
	OrderStatus      string
	CustomerId       string
	CustomerUsername string
	CustomerEmail    string
	BillTo           BillTo
	ShipTo           ShipTo
}
