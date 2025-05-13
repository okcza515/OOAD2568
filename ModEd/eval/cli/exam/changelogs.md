# MEP-1007: Examination Evaluation  
📅 13 May 2025
## Consult MEP-1012 (กลุ่ม 3)

## ✅ Added
- เพิ่ม Model (`ShortAnswer`,`QuestionAnswer`,`MultipleChoiceAnswer`)
- เพิ่ม CLI 
- เพิ่ม Handler (`Exam`,`Question`,`Submission`)
- แก้ไขและเพิ่ม `Controller` (`QuestionController`,`SubmissionController`)
- เพิ่ม Migration ให้กับ MigrateManager ของ core และ implement ลงใน ExamModuleHandler

## ♻️ Refactor
- นำ `core model` และ `core controller` มาใช้เพื่อให้โครงสร้างโค้ดเป็นมาตรฐานเดียวกัน
- ลบ Controller ที่ไม่ได้ใช้ (`ResultController`)

## 🎨 Design Patterns
- เพิ่ม Command ใช้กับ CLI (exam), ExamCommand, LoadExamCommand
- เพิ่ม State ใช้กับ ExamMenuState, QuestionMenuStateHandler, SubmissionMenuStateHandler
- เพิ่ม Strategy ใช้กับ Grading Submission (`Submission Controller`), IGradingStrategy
