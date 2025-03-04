package model

type StudentInterface interface {
}

func CreateStudent(st ProgramType) StudentInterface {
	if st == REGULAR {
		student := RegularStudent{}
		student.Program = REGULAR
		return student
	} else if st == INTERNATIONAL {
		student := InternationalStudent{}
		student.Program = INTERNATIONAL
		return student
	} else {
		return Student{}
	}
}
