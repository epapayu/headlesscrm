package main

import "time"

// Money represents a monetary value
type Money struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// TimePeriod represents a validity period
type TimePeriod struct {
	StartDateTime time.Time `json:"startDateTime"`
	EndDateTime   time.Time `json:"endDateTime"`
}

// AccountBalance represents a user's balance
type AccountBalance struct {
	ID           string       `json:"id"`
	Href         string       `json:"href"`
	Amount       Money        `json:"amount"`
	ValidFor     TimePeriod   `json:"validFor"`
	Status       string       `json:"status"` // Active, Suspended
	RelatedParty []RelatedParty `json:"relatedParty"`
}

// RelatedParty represents a linked entity (Customer)
type RelatedParty struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

// BalanceAdjustment represents a credit/debit action
type BalanceAdjustment struct {
	ID            string      `json:"id"`
	Type          string      `json:"type"` // adjustment, topup, deduction
	Amount        Money       `json:"amount"`
	Description   string      `json:"description"`
	ReasonCode    string      `json:"reasonCode"`
	RelatedWallet RelatedEntity `json:"relatedWallet"`
}

type RelatedEntity struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
