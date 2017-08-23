package model

import "time"

//APIAIRequest ...API.AIからのリクエストをパースする
type APIAIRequest struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Lang      string    `json:"lang"`
	Result    struct {
		Source           string `json:"source"`
		ResolvedQuery    string `json:"resolvedQuery"`
		Speech           string `json:"speech"`
		Action           string `json:"action"`
		ActionIncomplete bool   `json:"actionIncomplete"`
		Parameters       struct {
		} `json:"parameters"`
		Contexts []struct {
			Name       string `json:"name"`
			Parameters struct {
			} `json:"parameters"`
			Lifespan int `json:"lifespan"`
		} `json:"contexts"`
		Metadata struct {
			IntentID                  string `json:"intentId"`
			WebhookUsed               string `json:"webhookUsed"`
			WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
			IntentName                string `json:"intentName"`
		} `json:"metadata"`
		Fulfillment struct {
			Speech   string `json:"speech"`
			Messages []struct {
				Type   int    `json:"type"`
				Speech string `json:"speech"`
			} `json:"messages"`
		} `json:"fulfillment"`
		Score float64 `json:"score"`
	} `json:"result"`
	Status struct {
		Code      int    `json:"code"`
		ErrorType string `json:"errorType"`
	} `json:"status"`
	SessionID string `json:"sessionId"`
}
