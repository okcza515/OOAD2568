//MEP-1008
package model

type ProjectEvaluationTypeEnum string

const (
	ASSIGNMENT   ProjectEvaluationTypeEnum = "Assignment"
	PROPOSAL     ProjectEvaluationTypeEnum = "Proposal"
	REPORT       ProjectEvaluationTypeEnum = "Report"
	PRESENTATION ProjectEvaluationTypeEnum = "Presentation"
)

func IsValidAssignmentType(assignmentType string) bool {
	switch ProjectEvaluationTypeEnum(assignmentType) {
	case ASSIGNMENT, PROPOSAL, REPORT, PRESENTATION:
		return true
	default:
		return false
	}
}
