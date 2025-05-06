# Core Validation

## How to use

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

## Example
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
