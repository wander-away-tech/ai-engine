package structures

type ItineraryRequestBody struct {
	Destination string `json:"destination"`
	Duration    string `json:"duration"`
	Preferences string `json:"preferences"`
}
