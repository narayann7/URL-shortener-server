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
	RateRemaining int           `json:"rate_remaining"`
}

type AppError struct {
	Message       string `json:"message"`
	ErrorCode     int    `json:"error_code"`
	ErrorLocation string `json:"error_location"`
}