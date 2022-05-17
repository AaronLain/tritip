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
	Name            string `json:"name"`
	Company         string `json:"company"`
	Street1         string `json:"street1"`
	Street2         string `json:"street2"`
	Street3         string `json:"street3"`
	City            string `json:"city"`
	State           string `json:"state"`
	PostalCode      string `json:"postalCode"`
	Country         string `json:"country"`
	Phone           string `json:"phone"`
	Residential     bool   `json:"residential"`
	AddressVerified string `json:"addressVerified"`
}

type ShipTo struct {
	Name            string `json:"name"`
	Company         string `json:"company"`
	Street1         string `json:"street1"`
	Street2         string `json:"street2"`
	Street3         string `json:"street3"`
	City            string `json:"city"`
	State           string `json:"state"`
	PostalCode      string `json:"postalCode"`
	Country         string `json:"country"`
	Phone           string `json:"phone"`
	Residential     bool   `json:"residential"`
	AddressVerified string `json:"addressVerified"`
}

type Weight struct {
	Value       float64 `json:"value"`
	Units       string  `json:"units"`
	WeightUnits float64 `json:"WeightUnits"`
}

type Options struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Dimensions struct {
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Units  string  `json:"units"`
}

type InsuranceOptions struct {
	Provider       string  `json:"provider"`
	InsureShipment bool    `json:"insureShipment"`
	InsuredValue   float64 `json:"insuredValue"`
}

type InternationalOptions struct {
	Contents     string `json:"contents"`
	CustomsItems string `json:"customsItems"`
	NonDelivery  string `json:"nonDelivery"`
}

type AdvancedOptions struct {
	WarehouseId          float64  `json:"warehouseId"`
	NonMachinable        bool     `json:"nonMachinable"`
	SaturdayDelivery     bool     `json:"saturdayDelivery"`
	ContainsAlcohol      bool     `json:"containsAlcohol"`
	StoreId              int      `json:"storeId"`
	CustomField1         string   `json:"customField1"`
	CustomField2         string   `json:"customField2"`
	CustomField3         string   `json:"customField3"`
	Source               string   `json:"source"`
	MergedOrSplit        bool     `json:"mergedOrSplit"`
	MergedIds            []string `json:"mergedIds"`
	ParentId             string   `json:"parentId"`
	BillToParty          string   `json:"billToParty"`
	BillToAccount        string   `json:"billToAccount"`
	BilltoPostalCode     string   `json:"billToPostalCode"`
	BillToCountryCode    string   `json:"billToCountryCode"`
	BillToMyOtherAccount string   `json:"billToMyOtherAccount"`
}

type LineItem struct {
	OrderItemId       float64   `json:"orderItemId"`
	LineItemKey       string    `json:"lineItemKey"`
	Sku               string    `json:"sku"`
	Name              string    `json:"name"`
	ImageUrl          string    `json:"imageUrl"`
	Weight            Weight    `json:"weight"`
	Quantity          float64   `json:"quantity"`
	UnitPrice         float64   `json:"unitPrice"`
	TaxAmount         float64   `json:"taxAmount"`
	ShippingAmount    float64   `json:"shippingAmount"`
	WarehouseLocation string    `json:"warehouseLocation"`
	Options           []Options `json:"options"`
	ProductId         float64   `json:"productId"`
	FulfillmentSku    string    `json:"fulfillmentSku"`
	Adjustment        bool      `json:"adjustment"`
	Upc               string    `json:"upc"`
	CreateDate        string    `json:"createDate"`
	ModifyDate        string    `json:"modifyDate"`
}

type LineItemOptions struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type OrderRecordOutput struct {
	OrderId                  float64              `json:"orderId"`
	OrderNumber              string               `json:"orderNumber"`
	OrderKey                 string               `json:"orderKey"`
	OrderDate                string               `json:"orderDate"`
	CreateDate               string               `json:"createDate"`
	ModifyDate               string               `json:"modifyDate"`
	ShipByDate               string               `json:"paymentDate"`
	PaymentDate              string               `json:"shipByDate"`
	OrderStatus              string               `json:"orderStatus"`
	CustomerId               float64              `json:"customerId"`
	CustomerUsername         string               `json:"customerUsername"`
	CustomerEmail            string               `json:"customerEmail"`
	BillTo                   BillTo               `json:"billTo"`
	ShipTo                   ShipTo               `json:"shipTo"`
	Items                    []LineItem           `json:"items"`
	OrderTotal               float64              `json:"orderTotal"`
	AmountPaid               float64              `json:"amountPaid"`
	TaxAmount                float64              `json:"taxAmount"`
	ShippingAmount           float64              `json:"shippingAmount"`
	CustomerNotes            string               `json:"customerNotes"`
	InternalNotes            string               `json:"internalNotes"`
	Gift                     bool                 `json:"gift"`
	GiftMessage              string               `json:"giftMessage"`
	PaymentMethod            string               `json:"paymentMethod"`
	RequestedShippingService string               `json:"requestedShippingService"`
	CarrierCode              string               `json:"carrierCode"`
	ServiceCode              string               `json:"serviceCode"`
	Confirmation             string               `json:"confirmation"`
	PackageCode              string               `json:"packageCode"`
	ShipDate                 string               `json:"shipDate"`
	HoldUntilDate            string               `json:"holdUntilDate"`
	Weight                   Weight               `json:"weight"`
	Dimensions               Dimensions           `json:"dimensions"`
	InsuranceOptions         InsuranceOptions     `json:"insuranceOptions"`
	InternationalOptions     InternationalOptions `json:"internationalOptions"`
	AdvancedOptions          AdvancedOptions      `json:"advancedOptions"`
	TagIds                   []int                `json:"tagIds"`
	UserId                   string               `json:"userId"`
	ExternallyFulfilled      bool                 `json:"externallyFulfilled"`
	ExternallFulfilledBy     string               `json:"externallyFulfilledBy"`
}

type OrderRecordOutputResp struct {
	Orders []OrderRecordOutput
}

type Tag struct {
	TagId float64 `json:"tagId"`
	Name  string  `json:"name"`
}
