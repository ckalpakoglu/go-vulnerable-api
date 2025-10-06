package models

import "time"

type OTP struct {
	Value     string    `json:"value"`
	Expired   bool      `json:"expired"`
	CreatedAt time.Time `json:"ttl"`
}
