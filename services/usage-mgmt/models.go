package main

import "time"

type Usage struct {
	ID                  string              `json:"id"`
	Href                string              `json:"href"`
	Date                time.Time           `json:"date"`
	Type                string              `json:"type"`   // Voice, Data, SMS
	Description         string              `json:"description"`
	Status              string              `json:"status"` // Received, Rated, Billed
	UsageCharacteristic []UsageCharacteristic `json:"usageCharacteristic"`
	RelatedParty        []RelatedParty      `json:"relatedParty"`
	RatedProductUsage   []RatedProductUsage `json:"ratedProductUsage"`
}

type UsageCharacteristic struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type RelatedParty struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	Name string `json:"name"`
}

type RatedProductUsage struct {
	RatingAmountType        string `json:"ratingAmountType"`
	TaxIncludedRatingAmount Money  `json:"taxIncludedRatingAmount"`
}

type Money struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}
