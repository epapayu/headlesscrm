package main

import "time"

type ServiceOrder struct {
	ID               string             `json:"id"`
	Href             string             `json:"href"`
	ExternalID       string             `json:"externalId"`
	State            string             `json:"state"` // Acknowledged, InProgress, Completed
	OrderDate        time.Time          `json:"orderDate"`
	Description      string             `json:"description"`
	ServiceOrderItem []ServiceOrderItem `json:"serviceOrderItem"`
}

type ServiceOrderItem struct {
	ID      string  `json:"id"`
	Action  string  `json:"action"` // add, modify, delete
	Service Service `json:"service"`
}

type Service struct {
	Name                  string           `json:"name"`
	ServiceCharacteristic []Characteristic `json:"serviceCharacteristic"`
}

type Characteristic struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
