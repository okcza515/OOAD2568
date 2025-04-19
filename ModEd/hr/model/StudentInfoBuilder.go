package model

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

func (b *StudentInfoBuilder) WithYear(year int) *StudentInfoBuilder {
	b.info.Year = year
	return b
}

func (b *StudentInfoBuilder) Build() *StudentInfo {
	return b.info
}
