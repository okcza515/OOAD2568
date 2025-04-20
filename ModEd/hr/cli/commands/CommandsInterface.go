package commands

type ListStudentsCommand struct{}
type AddStudentCommand struct{}
type UpdateStudentStatusCommand struct{}
type ImportStudentsCommand struct{}
type DeleteStudentCommand struct{}
type UpdateStudentCommand struct{}

type RequestResignationCommand struct{}       // นักศึกษายื่นคำร้องขอลาออก
type ApproveResignationCommand struct{}       // ผู้ดูแลระบบอนุมัติคำขอลาออก
type RejectResignationCommand struct{}        // ผู้ดูแลระบบปฏิเสธคำขอลาออก
type CancelResignationCommand struct{}       // นักศึกษาเพิกถอนคำร้องขอลาออก

type MigrateStudentsCommand struct{}
type ExportStudentsCommand struct{}

type RequsetLeaveCommand struct{}          // นักศึกษายื่นคำร้องขอลา
