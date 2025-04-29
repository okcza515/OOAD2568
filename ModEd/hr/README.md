** Design Pattern
- CLI 
  - Command Design Pattern
- Controller
  - Remove Unused Facade
- Model
  - Factory
    - Used to create different types of leave requests (Instructor vs. Student) based on a role parameter, returning a common interface (`RequestLeaveProductInterface`) and centralizing creation logic.
- Util
  - Chain of Responsibility
    - Chaining multiple different validation rules. (Required, Length and Regex)
    - Create Multiple independent chains for different flags within same command

** SOLIDs
- **Single Responsibility Principle (SRP):**
  - Packages (`model`, `controller`, `cli`, `util`) have distinct responsibilities.
  - Controllers (`StudentHRController`, `ResignationInstructorHRController`, etc.) handle specific domains.
  - Models (`RequestResignationInstructor`) represent single data concepts.
- **Open/Closed Principle (OCP):**
  - The use of separate controllers allows adding new HR features by adding new controllers without modifying existing ones.
  - Command design pattern allows us to add new command easily.
- **Interface Segregation Principle (ISP):**
  - Specific controllers and models ensure clients don't depend on interfaces/structs with methods/fields they don't use.