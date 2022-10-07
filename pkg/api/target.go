package api

// Target represents a GitHub actions workflow target to send webhooks to
type Target struct {
	URL       string
	Token     string
	EventType string
}
