package model

//APIAIResponse ...API.AIからのレスポンスをパースする
type APIAIResponse struct {
	Speech      string `json:"speech"`
	DisplayText string `json:"displayText"`
	Data        struct {
	} `json:"data"`
	ContextOut []interface{} `json:"contextOut"`
	Source     string        `json:"source"`
}
