package entities

type Payload struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	ChannelID string `json:"channelID"`
}
