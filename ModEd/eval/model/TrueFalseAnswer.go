package model

type TrueFalseAnswer struct {
	BaseAnswer
	Boolean bool `json:"answer"`
}
