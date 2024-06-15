package types

type Email struct {
	Record    string `json:"record"`
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}
