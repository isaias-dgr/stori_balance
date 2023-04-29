package dto

import "github.com/google/uuid"

type EventNotify struct {
	Records []Record `json:"Records"`
}

type Record struct {
	MessagesID uuid.UUID `json:"messageID"`
	Body       string    `json:"body"`
}
