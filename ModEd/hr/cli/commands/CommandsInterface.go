package commands

type ListStudentsCommand struct{}
type AddStudentCommand struct{}
type UpdateStudentStatusCommand struct{}
type ImportStudentsCommand struct{}
type DeleteStudentCommand struct{}
type UpdateStudentCommand struct{}

type AnswerResignationCommand struct{}  // เจ้าหน้าที่ตอบคำร้องขอลาออก
type CancelResignationCommand struct{}  // นักศึกษาเพิกถอนคำร้องขอลาออก
type RequestResignationCommand struct{} // นักศึกษายื่นคำร้องขอลาออก
type ApproveResignationCommand struct{} // ผู้ดูแลระบบอนุมัติคำขอลาออก
type RejectResignationCommand struct{}  // ผู้ดูแลระบบปฏิเสธคำขอลาออก

type MigrateStudentsCommand struct{}
type ExportStudentsCommand struct{}

type RequsetLeaveCommand struct{} // นักศึกษายื่นคำร้องขอลา
type UpdateInstructorCommand struct{}
