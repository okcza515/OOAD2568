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

## User guide

### Admin role
#### Select `Admin` role on the start menu:

#### Login with: 
- Username: `admin`
- password: `1234`

#### Then, you will see the `Admin menu`:

    Admin Menu:
    1. View Application Report
    2. Schedule Interview
    3. Delete Interview
    4. Transfer Confirmed Applicants to Students
 --> Select `1:Application Report` for view Applicant infomations, the functions will show up.

    1. view All Applicatoin Reports
    2. View Interview Reports
    3. View Application Reports by Status:

`remark`: you can search by status  <span style="color:yellow">Pending</span>, 
<span style="color:green">Accepted</span>, 
<span style="color:red">Rejected</span>, 
etc


--> Select `2:Schedule Interview` for Schedule Applicant and Instructor to interview. By enter both ID and Interview date and time to create interview schedule.

--> Select `3:Delete Interview` for Delete the interview that in Student recruit system, By enter interview ID.

--> Select `4:Transfer Confirm Applicant to student` for Transfer applicant that confirmed recruitment to student 

### User role

### Instructor role
#### Select `Instructor` role on the start menu:

#### Login with: 
- Instructor ID :  `Use ID in ModEd.bin`
#### Main Menu Options
```bash
1. View All Interview Details
2. View Pending Interview Details
3. View Evaluated Interview Details
4. Evaluate Applicant
3. Menu Details
```
`Option 1`: View All Interview Details
Shows all applicants assigned to this instructor

Includes applicant ID, application round, and evaluation scores (if any)

Useful to get an overview of the full workload

`Option 2`: View Pending Interview Details
Displays only those applicants that have not yet been evaluated

Helps plan upcoming interviews

`Option 3`: View Evaluated Interview Details
Shows applicants that have already been evaluated

Useful for review or reporting

`Option 4`: Evaluate Applicant
1. Select an applicant ID from your list

2. You will be prompted with evaluation criteria (loaded from JSON file)

3. Enter scores or comments as required

4. Confirm to submit your evaluation

5. Evaluation is saved to memory or file



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