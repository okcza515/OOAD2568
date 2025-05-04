// MEP-1003 Student Recruitment

remark: Initiate the data(Faculty, Department, Instructors) by using common (\ModEd\common\cli\README.md)

go run common/cli/CommonDataCLI.go --database="data/ModEd.bin" --path="data/common/DepartmentList.csv"

go run common/cli/CommonDataCLI.go --database="data/ModEd.bin" --path="data/common/FacultyList.csv"

go run common/cli/CommonDataCLI.go --database="data/ModEd.bin" --path="data/common/InstructorsList.csv"
