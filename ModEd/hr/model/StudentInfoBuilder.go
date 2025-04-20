package model

import (
	commonModel "ModEd/common/model"
	"time"
)

type StudentInfoBuilder struct {
	info *StudentInfo
}

func NewStudentInfoBuilder() *StudentInfoBuilder {
	return &StudentInfoBuilder{
		info: &StudentInfo{},
	}
}

func (b *StudentInfoBuilder) WithStudentCode(code string) *StudentInfoBuilder {
	b.info.StudentCode = code
	return b
}

func (b *StudentInfoBuilder) WithFirstName(firstName string) *StudentInfoBuilder {
	b.info.FirstName = firstName
	return b
}

func (b *StudentInfoBuilder) WithLastName(lastName string) *StudentInfoBuilder {
	b.info.LastName = lastName
	return b
}

func (b *StudentInfoBuilder) WithEmail(email string) *StudentInfoBuilder {
	b.info.Email = email
	return b
}

func (b *StudentInfoBuilder) WithStartDate(startDate string) *StudentInfoBuilder {
	parsedDate, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		panic("Invalid start date format: " + startDate)
	}
	b.info.StartDate = parsedDate
	return b
}

func (b *StudentInfoBuilder) WithBirthDate(birthDate string) *StudentInfoBuilder {
	parsedDate, err := time.Parse("2006-01-02", birthDate)
	if err != nil {
		panic("Invalid birth date format: " + birthDate)
	}
	b.info.BirthDate = parsedDate
	return b
}

func (b *StudentInfoBuilder) WithProgram(program *commonModel.ProgramType) *StudentInfoBuilder {
	b.info.Program = *program
	return b
}

func (b *StudentInfoBuilder) WithStatus(status *commonModel.StudentStatus) *StudentInfoBuilder {
	b.info.Status = status
	return b
}

func (b *StudentInfoBuilder) WithGender(gender string) *StudentInfoBuilder {
	b.info.Gender = gender
	return b
}

func (b *StudentInfoBuilder) WithCitizenID(cid string) *StudentInfoBuilder {
	b.info.CitizenID = cid
	return b
}

func (b *StudentInfoBuilder) WithPhoneNumber(phone string) *StudentInfoBuilder {
	b.info.PhoneNumber = phone
	return b
}

// func (b *StudentInfoBuilder) WithAdvisor(advisor commonModel.Instructor) *StudentInfoBuilder {
// 	b.info.Advisor = advisor
// 	return b
// }

// func (b *StudentInfoBuilder) WithDepartment(department commonModel.Department) *StudentInfoBuilder {
// 	b.info.Department = department
// 	return b
// }

func (b *StudentInfoBuilder) WithStudent(student commonModel.Student) *StudentInfoBuilder {
	b.info.StudentCode = student.StudentCode
	b.info.FirstName = student.FirstName
	b.info.LastName = student.LastName
	b.info.Email = student.Email
	b.info.StartDate = student.StartDate
	b.info.BirthDate = student.BirthDate
	b.info.Program = student.Program
	b.info.Status = student.Status
	return b
}

func (b *StudentInfoBuilder) Build() *StudentInfo {
	return b.info
}
