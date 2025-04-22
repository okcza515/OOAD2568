package model

import (
	commonModel "ModEd/common/model"
	"fmt"
	"time"
)

type InstructorInfoBuilder struct {
	instructorInfo *InstructorInfo
	err            error
}

func NewInstructorInfoBuilder() *InstructorInfoBuilder {
	return &InstructorInfoBuilder{
		instructorInfo: &InstructorInfo{},
	}
}

func (b *InstructorInfoBuilder) WithFirstName(firstName string) *InstructorInfoBuilder {
	b.instructorInfo.FirstName = firstName
	return b
}

func (b *InstructorInfoBuilder) WithLastName(lastName string) *InstructorInfoBuilder {
	b.instructorInfo.LastName = lastName
	return b
}

func (b *InstructorInfoBuilder) WithEmail(email string) *InstructorInfoBuilder {
	b.instructorInfo.Email = email
	return b
}

func (b *InstructorInfoBuilder) WithStartDate(s string) *InstructorInfoBuilder {
	if b.err != nil {
		return b
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		b.err = fmt.Errorf("invalid start date %q: %w", s, err)
	} else {
		b.instructorInfo.StartDate = &t
	}
	return b
}

func (b *InstructorInfoBuilder) WithDepartment(department string) *InstructorInfoBuilder {
	b.instructorInfo.Department = &department
	return b
}

func (b *InstructorInfoBuilder) WithGender(gender string) *InstructorInfoBuilder {
	b.instructorInfo.Gender = gender
	return b
}

func (b *InstructorInfoBuilder) WithCitizenID(citizenID string) *InstructorInfoBuilder {
	b.instructorInfo.CitizenID = citizenID
	return b
}

func (b *InstructorInfoBuilder) WithPhoneNumber(phoneNumber string) *InstructorInfoBuilder {
	b.instructorInfo.PhoneNumber = phoneNumber
	return b
}

func (b *InstructorInfoBuilder) WithSalary(salary int) *InstructorInfoBuilder {
	b.instructorInfo.Salary = salary
	return b
}

func (b *InstructorInfoBuilder) WithAcademicPosition(academicPosition AcademicPosition) *InstructorInfoBuilder {
	b.instructorInfo.AcademicPosition = academicPosition
	return b
}

func (b *InstructorInfoBuilder) WithDepartmentPosition(departmentPosition DepartmentPosition) *InstructorInfoBuilder {
	b.instructorInfo.DepartmentPosition = departmentPosition
	return b
}

func (b *InstructorInfoBuilder) WithInstructor(instructor commonModel.Instructor) *InstructorInfoBuilder {
	b.instructorInfo.FirstName = instructor.FirstName
	b.instructorInfo.LastName = instructor.LastName
	b.instructorInfo.Email = instructor.Email
	b.instructorInfo.StartDate = instructor.StartDate
	b.instructorInfo.Department = instructor.Department
	return b
}

func (b *InstructorInfoBuilder) Build() *InstructorInfo {
	return b.instructorInfo
}
