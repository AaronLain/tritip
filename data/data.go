package data

type OrderRecordInput struct {
	OrderNum     string  `csv:"Order - Number"`
	CustomField3 string  `csv:"CustomField3"`
	AvgTemp      float64 `csv:"AvgTemp"`
	City         string  `csv:"Ship To - City"`
	State        string  `csv:"Ship To - State"`
	PostalCode   string  `csv:"Ship To - Postal Code"`
	CountryCode  string  `csv:"Shipping Country"`
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
	AddressVerified string
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
	AddressVerified string
}

type Weight struct {
	Value       float64
	Units       string
	WeightUnits float64
}

type Options struct {
	Name  string
	Value string
}

type Dimensions struct {
	Units  string
	Length float64
	Width  float64
	Height float64
}

type InsuranceOptions struct {
	Provider       string
	InsureShipment bool
	InsuredValue   float64
}

type InternationalOptions struct {
	Contents     string
	CustomsItems string
	NonDelivery  string
}

type AdvancedOptions struct {
	WarehouseId       float64
	NonMachinable     bool
	SaturdayDelivery  bool
	ContainsAlcohol   bool
	MergedOrSplit     bool
	MergedIds         []string
	ParentId          string
	StoreId           float64
	CustomField1      string
	CustomField2      string
	CustomField3      string
	Source            string
	BillToParty       string
	BillToAccount     string
	BilltoPostalCode  string
	BillToCountryCode string
}

type LineItem struct {
	OrderItemId       float64
	LineItemKey       string
	Sku               string
	Name              string
	ImageUrl          string
	Weight            Weight
	Quantity          float64
	UnitPrice         float64
	TaxAmount         float64
	ShippingAmount    float64
	WarehouseLocation string
	Options           []Options
	ProductId         float64
	FulfillmentSku    string
	Adjustment        bool
	Upc               string
	CreateDate        string
	ModifyDate        string
}

type LineItemOptions struct {
	Name  string
	Value string
}

type OrderRecordOutput struct {
	OrderId                  float64
	OrderNumber              string
	OrderKey                 string
	OrderDate                string
	CreateDate               string
	ModifyDate               string
	PaymentDate              string
	ShipByDate               string
	OrderStatus              string
	CustomerId               float64
	CustomerUsername         string
	CustomerEmail            string
	BillTo                   BillTo
	ShipTo                   ShipTo
	Items                    []LineItem
	OrderTotal               float64
	AmountPaid               float64
	TaxAmount                float64
	ShippingAmount           float64
	CustomerNotes            string
	InternalNotes            string
	Gift                     bool
	GiftMessage              string
	PaymentMethod            string
	RequestedShippingService string
	CarrierCode              string
	ServiceCode              string
	Confirmation             string
	ShipDate                 string
	HoldUntilDate            string
	Weight                   Weight
	Dimensions               Dimensions
	InsuranceOptions         InsuranceOptions
	AdvancedOptions          AdvancedOptions
	TagIds                   []int
	UserId                   string
	ExternallyFulfilled      bool
	ExternallFulfilledBy     string
}

type OrderRecordOutputResp struct {
	Orders []OrderRecordOutput
}

type Tag struct {
	TagId float64
	Name  string
	Color string
}
