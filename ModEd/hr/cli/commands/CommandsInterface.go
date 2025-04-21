package commands

type UpdateCommand struct{}
type ImportCommand struct{}
type ExportCommand struct{}
type ListStudentsCommand struct{}
type AddStudentCommand struct{}
type UpdateStudentStatusCommand struct{}
type ImportStudentsCommand struct{}
type DeleteStudentCommand struct{}
type UpdateStudentCommand struct{}

type AnswerResignationCommand struct{} 
type CancelResignationCommand struct{}  
// type RequestResignationCommand struct{} 
type RequestCommand struct{}

type MigrateStudentsCommand struct{}
type ExportStudentsCommand struct{}

type RequestLeaveStudentCommand struct{} // นักศึกษายื่นคำร้องขอลา
type RequestLeaveInstructorCommand struct{}      // เจ้าหน้าที่ตอบคำร้องขอลา
type UpdateInstructorCommand struct{}

type RequestRaiseCommand struct{} 