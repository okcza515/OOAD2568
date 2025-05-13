// MEP-1006

package model

type QuizStatus string

const (
	QuizDraft     QuizStatus = "draft"
	QuizPublished QuizStatus = "published"
	QuizHidden    QuizStatus = "hidden"
)
