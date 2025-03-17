package model

type PresentationType string

const (
	PresentationTypeProposal PresentationType = "proposal"
	PresentationTypeMidterm  PresentationType = "midterm"
	PresentationTypeFinal    PresentationType = "final"
)

func ValidPresentationTypes() []PresentationType {
	return []PresentationType{
		PresentationTypeProposal,
		PresentationTypeMidterm,
		PresentationTypeFinal,
	}
}

func (rt PresentationType) IsValid() bool {
	for _, validType := range ValidPresentationTypes() {
		if rt == validType {
			return true
		}
	}
	return false
}
