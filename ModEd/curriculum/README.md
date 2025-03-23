## Curriculum Core Interfaces

### Overview

This document provides guidelines on how to use the CRUD interfaces for class, course, and curriculum in the `/controller` directory. These interfaces are essential for managing the educational content and structure within the system.

### Class Interface
The Class (คาบเรียน) interface allows you to create, read, update, and delete class records. Class = Course + other necessary information (e.g., instructors and schedule).

#### Data Model

| Field        | Type                | Requirement   | Relation Type |
|--------------|---------------------|------------|---------------|
| ClassId      | `uint`              |            |               |
| CourseId     | `uint`              | required   |               |
| Course       | `Course`            | required   | [1:N]         |
| Section      | `int`               | required   |               |
| Schedule     | `time.Time`         | required   |               |
| Instructor   | `[]model.Instructor`|            | [N:N]         |
| StudentList  | `[]model.Student`   |            | [N:N]         |
| CreatedAt    | `time.Time`         |auto-generated|               |
| UpdatedAt    | `time.Time`         |auto-generated|               |
| DeletedAt    | `gorm.DeletedAt`    |auto-generated|               |

#### Methods

- `createClass(Class)`: Creates a new class with the provided data.
- `getClass(id)`: Retrieves the class with the specified ID.
- `updateClass(id, Class)`: Updates the class with the specified ID using the provided data.
- `deleteClass(id)`: Deletes the class with the specified ID.

### Course Interface

The Course (วิชาเรียน) interface provides methods to manage courses.

#### Data Model
| Field        | Type                | Requirement   | Relation Type |
|--------------|---------------------|---------------|---------------|
| CourseId     | `uint`              |               |               |
| Name         | `string`            | required      |               |
| Description  | `string`            | required      |               |
| CurriculumId | `uint`              | required      |               |
| Curriculum   | `Curriculum`        | required      | [1:N]         |
| Optional     | `boolean`           | required      |               |
| Prerequisite | `[]Course`          | required      | [N:N]         |
| ClassList    | `[]Course`          | auto-generated| [N:N]         |
| CourseStatus | `CourseStatus`      | required      |               |
| CreatedAt    | `datetime`          | auto-generated|               |
| UpdatedAt    | `datetime`          | auto-generated|               |
| DeletedAt    | `gorm.Delete`       | auto-generated|               |


#### Methods

- `createCourse(data)`: Creates a new course with the provided data.
- `getCourse(id)`: Retrieves the course with the specified ID.
- `updateCourse(id, data)`: Updates the course with the specified ID using the provided data.
- `deleteCourse(id)`: Deletes the course with the specified ID.

### Curriculum Interface - Work In-Progress 

The Curriculum (หลักสูตรการเรียน) interface is used to handle curriculum-related operations.

#### Data Model

*The relations in `common` module seems to be incorrect, so we are waiting for the fix.*

#### Methods

- `createCurriculum(data)`: Creates a new curriculum with the provided data.
- `getCurriculum(id)`: Retrieves the curriculum with the specified ID.
- `updateCurriculum(id, data)`: Updates the curriculum with the specified ID using the provided data.
- `deleteCurriculum(id)`: Deletes the curriculum with the specified ID.

### Implementation

To use these interfaces, import the necessary modules from the `/controller` directory and call the appropriate methods with the required parameters.

1. Import interface

    ```go
    //import database connector
    import "ModEd/curriculum/utils"
    //and
    import classController "ModEd/curriculum/controller/class"
    // or
    import courseController "ModEd/curriculum/controller/course"
    //or
    import curriculumController "ModEd/curriculum/controller/curriculum"
    ```

2. Initialize Database

    ```go
    db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "path-to/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}
    ```
3. Access Data (example for Class data)

    ```go
    classId, err := classController.CreateClass(&newClass)
	if err != nil {
		t.Fatalf("Failed to create class: %v", err)
	}

	retrievedClass, err := classController.GetClass(classId)
	if err != nil {
		t.Fatalf("Failed to get class: %v", err)
	}

	deletedClass, err := classController.DeleteClass(retrievedClass.ID)
	if err != nil {
		t.Fatalf("Failed to delete class: %v", err)
	}
    ```

For full implementation visit `/testing`:
- [Class](testing/class_interface_test.go)
- [Course](testing/course_interface_test.go)
- [Curriculum](testing/curriculum_interface_test.go)

### Miscellanous
#### Command to dump database schema as sql file

```sh
sqlite3 test.db
.output database_schema.sql
.schema
.output stdout
```
