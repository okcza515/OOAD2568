##  How to Run

1. Install Go (if not installed): [https://go.dev/dl/](https://go.dev/dl/)
2. Run the application using terminal:
```bash
cd recruit
go run RecruitMain.go
```

##  Key Files

- `RecruitMain.go` - Program entry point
- `cli/InstructorCLI.go` - Command Line Interface for instructors
- `util/*` - General utility functions


## Additional Notes

// MEP-1003 Student Recruitment

remark: Initiate the data(Faculty, Department, Instructors) by using common (\ModEd\common\cli\README.md)

go run common/cli/CommonDataCLI.go common/cli/MenuItemHandlers.go common/cli/CLIFunction.go --database="data/ModEd.bin"  --database="data/ModEd.bin" --path="data/common/InstructorsList.csv"

go run common/cli/CommonDataCLI.go common/cli/MenuItemHandlers.go common/cli/CLIFunction.go --database="data/ModEd.bin"  --database="data/ModEd.bin" --database="data/ModEd.bin" --path="data/common/FacultyList.csv"

go run common/cli/CommonDataCLI.go common/cli/MenuItemHandlers.go common/cli/CLIFunction.go --database="data/ModEd.bin"  --database="data/ModEd.bin" --database="data/ModEd.bin" --database="data/ModEd.bin" --path="data/common/DepartmentList.csv"
