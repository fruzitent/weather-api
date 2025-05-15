package entity

type Subscription struct {
	Email     string `json:"email"`     // Email address
	City      string `json:"city"`      // City for weather updates
	Frequency string `json:"frequency"` // Frequency of updates
	Confirmed bool   `json:"confirmed"` // Whether the subscription is confirmed
}
