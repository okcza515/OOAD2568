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


### process flow
1. Registration Form
- The student fills out a registration form to apply for admission using the Register Applicant option in the User menu.

2. Score Criteria Check
- The system checks whether the applicant's entered academic scores meet the predefined criteria.
- If the scores meet the requirement → status is set to pending (awaiting interview).
- If not → status is set to rejected.

3. Schedule Interview
- The admin schedules interviews for applicants with a pending status.
- This includes assigning an instructor and setting up the interview time and details.
- Accessible via the Schedule Interview option in the Admin menu.

4. Evaluate Interview
- The assigned instructor evaluates each applicant by entering interview scores into the system.
- Done through the Evaluate an Applicant option in the Instructor menu.

5. Interview Criteria Check
- The system compares the interview score with the required threshold.
- If the score meets the criteria → status is updated to accepted.

6. User Confirmation
- The user confirms their intent to enroll using the View Application Status option in the User menu.

7. Admin Confirmation
- The admin reviews applicants with an accepted status and confirms them.
- Once confirmed, the applicant officially becomes a student.
- This is done through the Transfer Confirmed Applicants to Students option in the Admin menu.