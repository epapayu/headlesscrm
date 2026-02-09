package main

// Customer represents a TMF Customer
type Customer struct {
	ID            string          `json:"id"`
	Href          string          `json:"href"`
	Name          string          `json:"name"`
	Status        string          `json:"status"`       // Active, Suspended, Initial
	StatusReason  string          `json:"statusReason"` // e.g., "KYC Pending"
	Account       []AccountRef    `json:"account"`
	EngagedParty  PartyRef        `json:"engagedParty"`
	ContactMedium []ContactMedium `json:"contactMedium"`
	Characteristic []Characteristic `json:"characteristic"`
}

type AccountRef struct {
	ID   string `json:"id"`
	Href string `json:"href"`
	Name string `json:"name"`
}

type PartyRef struct {
	ID   string `json:"id"`
	Href string `json:"href"`
	Name string `json:"name"`
}

type ContactMedium struct {
	MediumType     string               `json:"mediumType"` // Email, Phone
	Preferred      bool                 `json:"preferred"`
	Characteristic MediumCharacteristic `json:"characteristic"`
}

type MediumCharacteristic struct {
	EmailAddress string `json:"emailAddress,omitempty"`
	PhoneNumber  string `json:"phoneNumber,omitempty"`
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
}

type Characteristic struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
