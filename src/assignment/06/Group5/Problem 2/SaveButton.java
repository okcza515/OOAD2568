import javax.swing.JButton;


public class SaveButton extends JButton {
    private final Editor editor;  // Mark as final to prevent the warning

    // Constructor: Accepts an instance of Editor
    public SaveButton(Editor editor) {  
        this.editor = editor;  // Save reference
        this.setText("Save");

        // ✅ Use editor inside a method to ensure it's recognized
        this.addActionListener(e -> triggerSave());
    }

    // Explicit method using editor
    private void triggerSave() {
        if (editor != null) {
            editor.saveChanges();  // Call saveChanges() on the instance
        }
    }
}
