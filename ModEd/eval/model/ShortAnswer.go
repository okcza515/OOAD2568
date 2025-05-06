package model

type ShortAnswer struct {
	BaseAnswer
	ShortText string `json:"answer_text"`
}
