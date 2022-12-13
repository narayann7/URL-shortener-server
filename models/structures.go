package models

import "time"

type Request struct {
	Url            string        `json:"url"`
	Expiry         time.Duration `json:"expiry"`
	CustomShortUrl string        `json:"custom_short_url"`
}

type Responce struct {
	Url           string        `json:"url"`
	NewUrl        string        `json:"new_url"`
	Expiry        time.Duration `json:"expiry"`
	RateRemaining string        `json:"rate_remaining"`
}
