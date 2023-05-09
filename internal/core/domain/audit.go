package domain

import "time"

type Audit struct {
	Action    string //TODO: convert to enum
	UserID    ExternalID
	EntityID  ExternalID
	Change    string
	Entity    string //TODO: convert to enum
	Timestamp time.Time
	Message   string
}
