package main

// ProductOffering represents a sellable item
type ProductOffering struct {
	ID                   string                  `json:"id"`
	Href                 string                  `json:"href"`
	Name                 string                  `json:"name"`
	Description          string                  `json:"description"`
	Status               string                  `json:"status"` // Active, Retired
	ProductSpecification ProductSpecificationRef `json:"productSpecification"`
	Category             []CategoryRef           `json:"category"`
	ProductOfferingPrice []ProductOfferingPrice  `json:"productOfferingPrice"`
}

type ProductSpecificationRef struct {
	ID   string `json:"id"`
	Href string `json:"href"`
	Name string `json:"name"`
}

type CategoryRef struct {
	ID   string `json:"id"`
	Href string `json:"href"`
	Name string `json:"name"`
}

type ProductOfferingPrice struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	PriceType string `json:"priceType"` // oneTime, recurring
	Price     Money  `json:"price"`
}

type Money struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}
