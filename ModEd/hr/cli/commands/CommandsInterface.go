package commands

type ListStudentsCommand struct{}
type AddStudentCommand struct{}
type UpdateStudentStatusCommand struct{}
type ImportStudentsCommand struct{}
type DeleteStudentCommand struct{}
type UpdateStudentCommand struct{}

type RequestResignationCommand struct{}       // นักศึกษายื่นคำร้องขอลาออก
type AnswerResignationCommand struct{}     // เจ้าหน้าที่ตอบคำร้องขอลาออก
type CancelResignationCommand struct{}       // นักศึกษาเพิกถอนคำร้องขอลาออก

type MigrateStudentsCommand struct{}
type ExportStudentsCommand struct{}
