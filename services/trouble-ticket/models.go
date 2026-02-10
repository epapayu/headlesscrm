package main

import "time"

type TroubleTicket struct {
	ID                 string          `json:"id"`
	Href               string          `json:"href"`
	Description        string          `json:"description"`
	Severity           string          `json:"severity"`
	Type               string          `json:"type"`
	CreationDate       time.Time       `json:"creationDate"`
	Status             string          `json:"status"`
	StatusChangeReason string          `json:"statusChangeReason"`
	RelatedParty       []RelatedParty  `json:"relatedParty"`
	Note               []Note          `json:"note"`
}

type RelatedParty struct {
	ID   string `json:"id"`
	Href string `json:"href"`
	Role string `json:"role"`
	Name string `json:"name"`
}

type Note struct {
	Date   time.Time `json:"date"`
	Author string    `json:"author"`
	Text   string    `json:"text"`
}
