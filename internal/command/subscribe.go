package command

type Subscribe struct {
	Email     string // Email address
	City      string // City for weather updates
	Frequency string // Frequency of updates
	Confirmed bool   // Whether the subscription is confirmed
}

type SubscribeRes struct{}
