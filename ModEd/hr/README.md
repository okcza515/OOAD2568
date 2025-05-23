# ModED MEP-1004 : HR Module

## Changelogs 13/05/2025

### Structure

- #### CLI

  - Change all command cli to menu based using core module.
  - Use strategy, state, singleton.
  - CLI version is still there to be compared to.
  - Using new field validator.
  - All handler now uses Strategy design pattern where controller function is injected into handler, allowing different behaviors to be injected dynamically.
  - All menu is working.
    - CRUD StudentInfo, Instructor
    - Request Raise|Leave|Resign
    - Review Raise|Leave|Resign
    - Pull data from common
    - Import data from file
    - Export data to file

- #### Controller

  - Centralized all controller access and initialization using HRControllerManager using Facade
  - Add Export Method
    - Implement Serializer in CSVMapper and JSONMapper and fix typo in DataMapper regarding .json length.

- #### Model

  - Add model validation rule.

### Issues

- Lastest issues is common change controller name last minute. So all method that are using common is not working.
- Incorrect convention on common module.
  - Using Map string instead of object.

## Changelogs 06/05/2025

### Structure

- #### CLI

  - All command that has been implemented is working.
  - Refactor Command CLI to be able to support dynamically adding command at runtime.

    ```go
      invoker.RegisterCommand("export", &commands.ExportStudentsCommand{})
      invoker.RegisterCommand("add-student", &commands.AddStudentCommand{})
      invoker.RegisterCommand("request-student-leave", &commands.RequestStudentLeaveCommand{})
    ```

    - When executing, the invoker tries to determine the command by matching from the longest to a single word. This support multi-word command names like ("request-student-resign") without switch statements.
    - Decoupling and Single Responsibility:
      - Each command (founds in commands folder) implements the `Command` interface with a single `Execute` method. Each command handles only a specific domain.
    - Command layer is slim. It only passes on validated parameter to the controller.
      - For example:
  
      ```go
      controller := controller.CreateResignationInstructorHRController(tx)
      return controller.SubmitResignationInstructor(*id, *reason)
      ```

    - Open/Close Priciple:
      - Adding new commands only involves create new command structs that implement the `Command` interface without modifying the invoker's code.
    - Interface Segregation:
      - Individual commands only expose the specific methods required by the CLI which is `Execute`, reducing dependencies between modules.
  - Reducing duplicate requests code:
    - Most of the code for different type of request is mostly the same(i.e. flag parsing, validator).
    - Create a `RequestHelpers` for handling `LeaveRequest` and `ResignationRequest` for both student and instructor.
  - Reducing duplciate review requests code:

    - All of the code for different type of review request is the same. Just with different controller.
    - Controller logic is passed in as a callback, keeping review commands flexible and decoupled

    ```go
      func handleReviewCommand(
        args []string,
        tx *gorm.DB,
        commandName string,
        idDescription string,
        controllerFunc ControllerReviewFunc,
        entityType string,
      ) error {
        fs := flag.NewFlagSet(commandName, flag.ExitOnError)
        requestID := fs.String("id", "", idDescription)
        action := fs.String("action", "", "Action to perform (approve or reject)")
        reason := fs.String("reason", "", "Reason if the request is rejected")
        fs.Parse(args)

        validator := util.NewValidationChain(fs)
        validator.Field("id").Required()
        validator.Field("action").Required().AllowedValues([]string{"approve", "reject"})
        if strings.ToLower(*action) == "reject" {
          validator.Field("reason").Required()
        }
        if err := validator.Validate(); err != nil {
          fs.Usage()
          return fmt.Errorf("validation error: %v", err)
        }

        if err := controllerFunc(tx, *requestID, strings.ToLower(*action), *reason); err != nil {
          return fmt.Errorf("failed to review %s request: %v", entityType, err)
        }

        fmt.Printf("%s request '%s' %sed successfully!\n", entityType, *requestID, strings.ToLower(*action))
        return nil
      }
    ```
  - Reducing Duplicate Import/Export Code
    - All of the code for different types of import/export commands is the same, differing only in the controller logic and file type.
    - Controller logic is passed in as a callback, keeping import/export commands flexible and decoupled.

- #### Controller

  - Redo all controller method
  - Add Generic Review Helper with Strategy Pattern

    - `fetcher` and `saver` represent families of algorithms (different ways to fetch or save data).
    - Each specific fetch or save logic is encapsulated within a function (`fetcher` or `saver` type).
    - The `ReviewRequest` function can work with any concrete `fetcher` or `saver` function that matches the required signature. The caller provides the specific strategy (the concrete function) to use.
    - `ReviewRequest` don't need to know specific implementation details. It also decoupled from the implementation details of how data is fetch or saved. It only depends on function signatures (the strategy interface).
    - `Reviewable` interface defines a contract `ApplyStatus` that any request type must fufill to be processbed by `ReviewRequest`.
    - The `ReviewRequest` function defining the algorithm and accepting strategies:

    ```go
        type fetcher func(id uint) (Reviewable, error)
        type saver func(Reviewable) error

        // ReviewRequest does the common parsing / fetching / status logic.
        func ReviewRequest(
            requestID, action, reason string,
            getByID fetcher, // Strategy for fetching
            save saver,      // Strategy for saving
        ) error {
            id, err := strconv.ParseUint(requestID, 10, 32)
            if err != nil {
                return fmt.Errorf("invalid request ID: %v", err)
            }
            // Use the fetcher strategy
            req, err := getByID(uint(id))
            if err != nil {
                return fmt.Errorf("failed to fetch request: %v", err)
            }
            // Use the Reviewable interface method
            if err := req.ApplyStatus(action, reason); err != nil {
                return err
            }
            // Use the saver strategy
            if err := save(req); err != nil {
                return fmt.Errorf("failed to save request: %v", err)
            }
            return nil
        }
    ```

    - The `Reviewable` interface defining the contract:

    ```go
    // Reviewable is any request that can be approved or rejected.
    type Reviewable interface {
        ApplyStatus(action, reason string) error
    }
    ```

    - A caller `ResignationStudentHRController` providing concrete strategies:

    ```go
    func (c *ResignationStudentHRController) ReviewStudentResignRequest(
        tx *gorm.DB,
        requestID, action, reason string,
    ) error {
            // Calling ReviewRequest with specific strategies
            return ReviewRequest(
                requestID,
                action,
                reason,
                // Concrete fetcher strategy: uses the controller's getByID method
                func(id uint) (Reviewable, error) {
                    return c.getByID(id) // c.getByID returns *model.RequestResignationStudent which implements Reviewable
                },
                // Concrete saver strategy: uses the transaction's Save method
                func(r Reviewable) error {
                    return tx.Save(r).Error
                },
            )
    }
    ```
  - Refactor `ImportInstructors` Method
    - Improved the `ImportInstructors` method in the `InstructorHRController` to enhance readability, maintainability, and reduce duplication.

- #### Model

  - Combine all request into single abstract factory
    - Centralize creation of different types of requests (resignation, leave, raise) in a single factory interface.
    - Define an `AbstractFactory` interface with methods for creating different types of request:
      - `CreateLeave(id, leaveType, reason, dateStr string)`
      - `CreateResignation(id, reason string)`
      - `CreateRaise(id, reason string, targetSalary int)`
    - Two concrete factories implement this interface:
      - `StudentFactory` for student related requests
      - `InstructorFactory` for instructor related requests
    - A client calls the `GetFactory(role string)` function to obtain the proper factory. After that, the client can call the corresponding creation method to get a concrete instance.
    - Encapsulate instantiation logic for each concrete request (`RequestResignationStudent`,`RequestLeaveInstructor`).
    - Single Responsibility: Each factory only handles creating request objects, while the controller stays focused on business logic and user interactions.
    - Decoupling: The controller uses the factory interface, makes it independent of specific construction and types.
    - Adding new request types only involves extending the factory, rather than modifying client code that uses it. Aligns with **_Open/Closed Principle_**.
  - Using composition to reduce model declaration duplication for all `Request` interface as all model only differ with `StudentCode` or `InstructorCode`

    - Multiple types of request models (resignation, leave, raise requests) share common fields (like status, reason, timestamp).
    - Seperate repeated fields to deadicated structs

    ```go
      type BaseStandardRequest struct {
        gorm.Model
        Reason string `gorm:"type:text"`
        Status string `gorm:"default:Pending"`
      }

      // SetStatus implements RequestStatus.
      func (b *BaseStandardRequest) SetStatus(status string) {
        b.Status = status
      }

      // SetReason implements RequestStatus.
      func (b *BaseStandardRequest) SetReason(reason string) {
        b.Reason = reason
      }

      func (b *BaseStandardRequest) ApplyStatus(action Action, reason string) error {
        return ApplyStatus(b, action, reason)
      }
    ```

    - Using common handler to support both `BaseStandardRequest` and `BaseLeaveRequest`
    ```go
      // commonActionHandlers maps actions to their handler functions for any RequestStatus.
      var commonActionHandlers = map[Action]func(RequestStatus, string){
        ActionApprove: func(r RequestStatus, _ string) {
          r.SetStatus("approve")
        },
        ActionReject: func(r RequestStatus, reason string) {
          r.SetStatus("reject")
          r.SetReason(reason)
        },
      }

      // ApplyStatus updates the status on any RequestStatus using the common map.
      func ApplyStatus(r RequestStatus, action Action, reason string) error {
        if handler, ok := commonActionHandlers[action]; ok {
          handler(r, reason)
          return nil
        }
        return fmt.Errorf("invalid action: %v", action)
      }
    ``` 

    - When new types of requests are added that share common fields, they can embed the base struct, keeping instantiation logic and field consistent.

- #### Util

  - Chain of Responsibility
    - Chaining multiple different validation rules (Required, Length and Regex).
    - Create Multiple independent chains for different flags within same command .
    - **_Updated_** : Refactor for ease of use.
    - Change from `Method Chaining` to `Fluent API per Field`.
    - Method Chaining Version:

      ```go
        err := util.NewValidationChain(fs).
          Required("id").
          Required("fname").
          Required("lname").
          Required("email").
          Required("gender").
          Required("citizenID").
          Required("phoneNumber").
          Length("id", 11).
          Regex("id", `^[0-9]{11}$`).
          Regex("email", `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).
          Validate()
        ```

    - Fluent API per Field Version:

      ```go
        validator := util.NewValidationChain(fs)
        validator.Field("id").Required().IsStudentID()
        validator.Field("fname").Required()
        validator.Field("lname").Required()
        validator.Field("email").Required().IsEmail()
        validator.Field("gender").Required().AllowedValues([]string{"Male", "Female", "Other"})
        validator.Field("citizenID").Required().Length(13)
        validator.Field("phone").Required()
        err := validator.Validate()
      ```

    - API become easier to use. Developers work on one field at a time and chain the rules directly.
    - It is simpler to add or remove validations for individual fields without interfering with other validations.
    - Combine frequent uses validator for example: `IsStudentID()`, `IsEmail()`, `IsDate()`.

### Issues

- Common CLI import instructor does not work.

### Questions

- 
- Is it correct to create a adapter for menu to execute the command using flags?