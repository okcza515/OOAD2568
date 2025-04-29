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
    import (
		classController "ModEd/curriculum/controller/class"
		//or
		courseController "ModEd/curriculum/controller/course"
		//or
		curriculumController "ModEd/curriculum/controller/curriculum"
	) 
    ```

2. Access Data (example for Class data)

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
3. (Advance Query) Using `preload` with functions
	The `GetCourse` method supports preloading related entities. For example, if a course has a `Prerequisite` relationship, you can preload it as follows:

	```go
	// Retrieve a course with its prerequisites preloaded
	course, err := courseController.GetCourse(courseId, "Prerequisite")
	if err != nil {
    	fmt.Printf("Failed to get course: %v\n", err)
    return
	}
	```
	After retrieving the course, you can directly access the `Prerequisite` field without making additional queries. This approach is useful for reducing the number of database queries when working with related entities.

	Example of how the data will look like:
	```
	{
		"course_id": 1,
		"name": "Introduction to Programming",
		"description": "This course covers the basics of programming using Python.",
		"curriculum_id": 101,
		"optional": false,
		"course_status": 1,
		"created_at": "2025-01-01T10:00:00Z",
		"updated_at": "2025-01-15T10:00:00Z",
		"class_list": [
			{
			"class_id": 1,
			"section": 1,
			"schedule": "2025-02-01T09:00:00Z",
			"instructor": "John Doe"
			}
		],
		"prerequisite": [
			{
			"course_id": 2,
			"name": "Basic Mathematics",
			"description": "This course covers fundamental mathematical concepts.",
			"curriculum_id": 101,
			"optional": false,
			"course_status": 1,
			"created_at": "2024-01-01T10:00:00Z",
			"updated_at": "2024-01-15T10:00:00Z"
			}
		]
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
