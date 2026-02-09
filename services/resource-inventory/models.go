package main

import "time"

type Resource struct {
	ID                     string           `json:"id"`
	Href                   string           `json:"href"`
	Name                   string           `json:"name"`
	Category               string           `json:"category"` // MSISDN, SIM
	Description            string           `json:"description"`
	StartOperatingDate     time.Time        `json:"startOperatingDate"`
	EndOperatingDate       time.Time        `json:"endOperatingDate"`
	ResourceStatus         string           `json:"resourceStatus"` // Available, InUse, Reserved, Retired
	ResourceCharacteristic []Characteristic `json:"resourceCharacteristic"`
}

type Characteristic struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
