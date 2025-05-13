# MEP-1006: การประเมินผลการใช้ Design Patterns
# Consult: MEP-1010
📅 13 พฤษภาคม 2568

## 🎨 Design Patterns ที่ใช้

### Adapter Pattern
- **การใช้งาน**: `QuizAdapter` interface, `QuizAdapter`

### Command Pattern
- **การใช้งาน**: `Command` interface, `CommandExecutor`, `EvaluationCommand`, `LoadExamCommand`

### Strategy Pattern
- **การใช้งาน**: `FuncStrategy`, `ChangeMenuHandlerStrategy`

### State Pattern
- **การใช้งาน**: `MenuState`, `QuizMenuStateHandler`, `ExamMenuState`, `MenuStateManager`, `SubmissionMenuStateHandler`, `QuestionMenuState`

### Factory Method Pattern
- **การใช้งาน**: `NewExamMenuState`, `NewMenuStateManager`, `NewQuestionMenuStateHandler`, `NewSubmissionMenuStateHandler`

### Composite Pattern
- **การใช้งาน**: โครงสร้างเมนูในส่วนของ handlers ซึ่งรายการเมนูสามารถเป็นคำสั่งหรือเมนูย่อยได้

### Refactor
- ทำการ refactor code โดยนำ core model และ core controller เพื่อให้ code อยู่ในมาตฐานเดียวกัน

### Migration
- ทำการ Migrate module กับทาง core และ implement Migration ใน EvalModuleHandler

## 📐 หลักการ SOLID ที่ใช้

### Single Responsibility Principle (SRP)
- **การใช้งาน**: การแยกคลาสตาม functionality เช่น `ProgressHandler`, `QuizHandler`, `EvaluationHandler` แต่ละคลาสรับผิดชอบหน้าที่เฉพาะของตัวเอง

### Open/Closed Principle (OCP)
- **การใช้งาน**: ระบบ handler ที่สามารถเพิ่มคำสั่งใหม่ได้โดยไม่ต้องแก้ไขโค้ดเดิม ผ่านการใช้ Strategy Pattern



