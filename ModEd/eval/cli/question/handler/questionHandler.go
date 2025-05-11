// MEP-1007
package handler
import (
	"fmt"
	"ModEd/core"
	"ModEd/eval/model"
)

type QuestionHandler struct {
	controller *core.BaseController[*model.Question]
}

func NewQuestionHandler(controller *core.BaseController[*model.Question]) *QuestionHandler {
	return &QuestionHandler{controller}
}

func (h *QuestionHandler) List() {
	questions, err := h.controller.List(nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, q := range questions {
		fmt.Printf("ID: %d | SectionID: %d | Score: %.2f | Type: %s | Question: %s\n",
			q.ID, q.SectionID, q.Score, q.QuestionType, q.ActualQuestion)
	}
}

func (h *QuestionHandler) Retrieve() {
	var id uint = 1 // mock ID
	q, err := h.controller.RetrieveByID(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("ID: %d\nSectionID: %d\nScore: %.2f\nType: %s\nQuestion: %s\n",
		q.ID, q.SectionID, q.Score, q.QuestionType, q.ActualQuestion)
}

func (h *QuestionHandler) Create() {
	q := &model.Question{
		SectionID:      1,
		Score:          5.0,
		ActualQuestion: "What is the capital of Thailand?",
		QuestionType:   model.MultipleChoiceQuestion,
	}

	if err := h.controller.Insert(q); err != nil {
		fmt.Println("Insert failed:", err)
	} else {
		fmt.Println("✅ Mock question created successfully.")
	}
}

func (h *QuestionHandler) CreateMockData() {
	mockQuestions := []*model.Question{
		{
			SectionID:      1,
			Score:          2.0,
			ActualQuestion: "What is 2 + 2?",
			QuestionType:   model.ShortAnswerQuestion,
		},
		{
			SectionID:      1,
			Score:          3.5,
			ActualQuestion: "Is the Earth flat?",
			QuestionType:   model.TrueFalseQuestion,
		},
		{
			SectionID:      2,
			Score:          5.0,
			ActualQuestion: "Who wrote the Thai national anthem?",
			QuestionType:   model.MultipleChoiceQuestion,
		},
		{
			SectionID:      2,
			Score:          4.0,
			ActualQuestion: "Define Newton's second law.",
			QuestionType:   model.ShortAnswerQuestion,
		},
		{
			SectionID:      3,
			Score:          1.0,
			ActualQuestion: "Is water wet?",
			QuestionType:   model.TrueFalseQuestion,
		},
	}

	for _, q := range mockQuestions {
		err := h.controller.Insert(q)
		if err != nil {
			fmt.Printf("❌ Failed to insert mock question: %s\n", q.ActualQuestion)
		} else {
			fmt.Printf("✅ Inserted: %s\n", q.ActualQuestion)
		}
	}
}

func (h *QuestionHandler) Update() {
	var id uint = 1 // mock ID
	q, err := h.controller.RetrieveByID(id)
	if err != nil {
		fmt.Println("Not found:", err)
		return
	}

	q.SectionID = 2
	q.Score = 4.5
	q.ActualQuestion = "UPDATED: What is the capital city of Thailand?"
	q.QuestionType = model.MultipleChoiceQuestion

	if err := h.controller.UpdateByID(q); err != nil {
		fmt.Println("Update failed:", err)
	} else {
		fmt.Println("✅ Updated successfully.")
	}
}

func (h *QuestionHandler) Delete() {
	var id uint = 1 // mock ID
	if err := h.controller.DeleteByID(id); err != nil {
		fmt.Println("Delete failed:", err)
	} else {
		fmt.Println("✅ Deleted successfully.")
	}
}
