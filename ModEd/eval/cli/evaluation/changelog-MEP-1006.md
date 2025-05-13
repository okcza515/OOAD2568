# MEP-1006: การประเมินผลการใช้ Design Patterns
# Consult: MEP-1010
📅 13 พฤษภาคม 2568

## 🎨 Design Patterns ที่ใช้

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

## Migration



