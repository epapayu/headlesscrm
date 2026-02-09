package main

import "time"

type Product struct {
	ID                    string           `json:"id"`
	Href                  string           `json:"href"`
	Name                  string           `json:"name"`
	Description           string           `json:"description"`
	Status                string           `json:"status"` // Active, Suspended, Terminated
	StartDate             time.Time        `json:"startDate"`
	ProductOffering       ProductOfferingRef `json:"productOffering"`
	RelatedParty          []RelatedParty   `json:"relatedParty"`
	ProductCharacteristic []Characteristic `json:"productCharacteristic"`
}

type ProductOfferingRef struct {
	ID   string `json:"id"`
	Href string `json:"href"`
	Name string `json:"name"`
}

type RelatedParty struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	Name string `json:"name"`
}

type Characteristic struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
