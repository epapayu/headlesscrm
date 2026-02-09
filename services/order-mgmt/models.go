package main

import "time"

type ProductOrder struct {
	ID               string             `json:"id"`
	Href             string             `json:"href"`
	ExternalID       string             `json:"externalId"`
	State            string             `json:"state"` // Acknowledged, InProgress, Completed
	OrderDate        time.Time          `json:"orderDate"`
	Description      string             `json:"description"`
	RelatedParty     []RelatedParty     `json:"relatedParty"`
	ProductOrderItem []ProductOrderItem `json:"productOrderItem"`
}

type ProductOrderItem struct {
	ID              string             `json:"id"`
	Action          string             `json:"action"` // add, modify, delete
	Quantity        int                `json:"quantity"`
	ProductOffering ProductOfferingRef `json:"productOffering"`
	Product         Product            `json:"product"`
}

type ProductOfferingRef struct {
	ID   string `json:"id"`
	Href string `json:"href"`
	Name string `json:"name"`
}

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RelatedParty struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	Name string `json:"name"`
}
