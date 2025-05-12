package model

type StudentInterface interface {
	Validate() error
}

func NewStudentByProgram(st ProgramType) StudentInterface {
	if st == REGULAR {
		student := RegularStudent{}
		student.Program = REGULAR
		return &student
	} else if st == INTERNATIONAL {
		student := InternationalStudent{}
		student.Program = INTERNATIONAL
		return &student
	} else {
		student := Student{}
		student.Program = st
		return &student
	}
}
