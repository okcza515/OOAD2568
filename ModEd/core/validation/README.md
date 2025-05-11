# Core Validation

## Field Validator

### Supported Validator

These are currently implementation of supported validator. The explanation for each function is within the function name.

- `IsNumberValid`
- `ParseNumber`
- `IsUintValid`
- `ParseUint`
- `IsStringNotEmpty`
- `IsEmailValid`
- `IsPhoneNumberValid`
- `IsDateTimeValid`
- `IsStudentID`
- `IsValueAllowed`

### Example

```go
package main

import (
    "fmt"
    "ModEd/core/validation" // Assuming this is the correct import path
)

func main() {
    // Get an instance of the validator
    validator := validation.NewValidator()

    // Example 1: Validate an email
    email1 := "test@example.com"
    isValidEmail1 := validator.IsEmailValid(email1)
    fmt.Printf("Is '%s' a valid email? %t\n", email1, isValidEmail1) 
    // Output: Is 'test@example.com' a valid email? true

    email2 := "invalid-email"
    isValidEmail2 := validator.IsEmailValid(email2)
    fmt.Printf("Is '%s' a valid email? %t\n", email2, isValidEmail2) 
    // Output: Is 'invalid-email' a valid email? false

    // Example 2: Validate if a value is allowed from a list
    allowedGenders := []string{"Male", "Female", "Other"}
    gender1 := "Male"
    isGenderAllowed1 := validator.IsValueAllowed(gender1, allowedGenders)
    fmt.Printf("Is gender '%s' allowed? %t\n", gender1, isGenderAllowed1) 
    // Output: Is gender 'Male' allowed? true

    gender2 := "Unknown"
    isGenderAllowed2 := validator.IsValueAllowed(gender2, allowedGenders)
    fmt.Printf("Is gender '%s' allowed? %t\n", gender2, isGenderAllowed2) 
    // Output: Is gender 'Unknown' allowed? false

    // Example 3: Validate a student ID
    studentID1 := "12345678901"
    isValidStudentID1 := validator.IsStudentID(studentID1)
    fmt.Printf("Is student ID '%s' valid? %t\n", studentID1, isValidStudentID1) 
    // Output: Is student ID '12345678901' valid? true

    studentID2 := "12345"
    isValidStudentID2 := validator.IsStudentID(studentID2)
    fmt.Printf("Is student ID '%s' valid? %t\n", studentID2, isValidStudentID2) 
    // Output: Is student ID '12345' valid? false
}
```


## Model Validator

### How to use

1. **Define your struct** with `gorm` and/or `validation` tags. The values for these tags should correspond to the `ModelValidatorEnum` constants (e.g., "email", "not null", "studentId").
    * `gorm:"not null"` can be used for basic non-empty checks.
    * `validation:"<rule>"` can be used for specific rules like `validation:"email"` or `validation:"studentId,phone"`.
1. **Create a `Validate()` method** on your struct (optional, but a common pattern).
2. Inside this method (or wherever you need to perform validation):
    a.  Get an instance of the `ModelValidator`:
        ```go
        modelValidator := validation.NewModelValidator()
        ```
    b.  Call the `ModelValidate` method, passing a pointer to your struct instance:
        ```go
        err := modelValidator.ModelValidate(yourStructInstance)
        ```
3. **Check the returned `error`**. If it's `nil`, the model is valid according to the defined tags. Otherwise, the error will describe the first validation failure encountered.

### Example
- Inside model:
    ```go
    package main

    import (
        "fmt"
        "time" // Import time package for time.Time

        "your_module_path/core/validation" // Replace with your actual module path
    )

    // Sample struct with validation tags
    type User struct {
        ID        uint      `gorm:"primaryKey"`
        Name      string    `gorm:"not null"`
        Email     string    `validation:"email"`
        StudentID string    `validation:"studentId"`
        Phone     string    `validation:"phone"`
        CreatedAt time.Time `validation:"datetime"` // Example for datetime, ensure format matches
        Optional  string    // No validation tags
    }

    func (u *User) Validate() error {
        // Get the singleton instance of ModelValidator
        modelValidator := validation.NewModelValidator()

        // Validate the struct using ModelValidator
        if err := modelValidator.ModelValidate(u); err != nil {
            return err
        }

        return nil
    }
    ```

- Inside where you want to validate:
    ```go
    // --- Existing Code ---
    if err := User.Validate(); err != nil {
        fmt.Println("Validation error:", err)
        return err
    }
    // --- Existing Code ---

    ```

## Fluent Validator

### Overview

The Fluent Validator provides a simple API to:

- Define a field with a prompt message.
- Chain required validations and specific rule validations (e.g., email, phone, or student code).
- Automatically prompt the user for input until all rules pass.

### Example Usage

```go
    func main() {
        // Define an input getter function that prompts the user for input on the console.
        inputGetter := func(prompt string) string {
            var input string
            fmt.Printf("%s: ", prompt)
            fmt.Scanln(&input)
            return input
        }

        // Create a new ValidationChain with the input getter.
        chain := validation.NewValidationChain(inputGetter)

        // Example: Validate an email address.
        email := chain.
            Field(validation.FieldConfig{
                Name:   "email",
                Prompt: "Enter your email",
            }).
            Required().
            IsEmail().
            GetInput()

        fmt.Printf("Validated Email: %s\n", email)

        // Example: Validate a student code (11 digits).
        studentCode := chain.
            Field(validation.FieldConfig{
                Name:   "student code",
                Prompt: "Enter your student code",
            }).
            Required().
            IsStudentCode().
            GetInput()

        fmt.Printf("Validated Student Code: %s\n", studentCode)

        // Example: Validate a phone number (10 digits, starting with 0).
        phoneNumber := chain.
            Field(validation.FieldConfig{
                Name:   "phone number",
                Prompt: "Enter your phone number",
            }).
            Required().
            IsPhoneNumber().
            GetInput()

        fmt.Printf("Validated Phone Number: %s\n", phoneNumber)
    }
```
