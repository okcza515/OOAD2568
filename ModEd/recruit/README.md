## MEP-1003 Student Recruitment

### Overview

This module provides a command-line interface and controller-based structure to manage the student recruitment process. It supports CRUD operations and interactive flows for three main roles: Applicant, Instructor, and Admin. The core logic for each role is implemented in the /controller directory, which acts as the bridge between the CLI and the data models.

- Applicant: Applicants can register for programs and view their application status. The system handles user input validation, storage, and basic interaction feedback.

- Instructor: Instructors participate in the evaluation process. They can view applicant lists, schedule interviews, and submit evaluation scores through a guided CLI workflow.

- Admin: Admins oversee and manage the entire recruitment pipeline. They are responsible for maintaining program structure, including class, course, and curriculum data, as well as accessing application reports and ensuring data integrity.

Each role uses a dedicated CLI entry point located in the /cli directory, and interacts with the shared data through modular controller logic.

### Applicant Interface

When choosing <span style="color:rgb(21, 168, 218)">[1]</span>, the user will enter the `User Menu`. This menu has two features available to applicants.

First, the applicant must register by selecting <span style="color:rgb(21, 168, 218)">[1]</span> `Register Applicant`. After selecting this option, the system will display the `Applicant Registration menu`, which allows registration either manually or by uploading a CSV/JSON file. Once registered, the `Available Application Rounds` menu will appear. There are four rounds to choose from, and each round requires different information.

After entering the necessary details for the selected round, the `Available Faculties` will be displayed. You can then choose a faculty for that round. Once a faculty is selected, the `Available Departments` will be shown. After selecting a department, the `Available Programs` will appear. When a program is selected, the applicant registration process is complete. The command line will confirm with “Registration completed” and display the Applicant Report ID, which you should remember for use in the next feature.

Returning to the User Menu, select <span style="color:rgb(21, 168, 218)">[2]</span> `View Application Status`. The command line will prompt: “Enter Application Report ID to view the report.” Enter the Report ID you received during registration. The system will then display your recruitment process information, which will update after each stage of the process is completed.

### Admin Interface

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

`remark`: you can search by status <span style="color:yellow">Pending</span>,
<span style="color:green">Accepted</span>,
<span style="color:red">Rejected</span>,
etc

--> Select `2:Schedule Interview` for Schedule Applicant and Instructor to interview. By enter both ID and Interview date and time to create interview schedule.

--> Select `3:Delete Interview` for Delete the interview that in Student recruit system, By enter interview ID.

--> Select `4:Transfer Confirm Applicant to student` for Transfer applicant that confirmed recruitment to student

### Instructor Interface

#### Select `Instructor` role on the start menu:

#### Login with:

- Instructor ID : `Use ID in ModEd.bin`

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

### How to Run

1. Install Go (if not installed): [https://go.dev/dl/](https://go.dev/dl/)
2. Run the application using terminal:

```bash
cd recruit
go run RecruitMain.go
```

### Additional Notes

> ⚠️ Before using the Recruit module, make sure to initialize the data for Faculty, Department, and Instructors using the CLI tools provided in the common module (\ModEd\common\cli\README.md).

```go
go run common/cli/CommonDataCLI.go common/cli/MenuItemHandlers.go common/cli/CLIFunction.go --database="data/ModEd.bin" --database="data/ModEd.bin" --path="data/common/InstructorsList.csv"
```

```go
go run common/cli/CommonDataCLI.go common/cli/MenuItemHandlers.go common/cli/CLIFunction.go --database="data/ModEd.bin" --database="data/ModEd.bin" --database="data/ModEd.bin" --path="data/common/FacultyList.csv"
```

```go
go run common/cli/CommonDataCLI.go common/cli/MenuItemHandlers.go common/cli/CLIFunction.go --database="data/ModEd.bin" --database="data/ModEd.bin" --database="data/ModEd.bin" --database="data/ModEd.bin" --path="data/common/DepartmentList.csv"
```
