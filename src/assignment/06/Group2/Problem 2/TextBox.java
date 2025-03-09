
import javax.swing.JTextArea;
import java.awt.event.KeyEvent;

public class TextBox extends JTextArea {
    private EditorMediators mediators;

    public TextBox(EditorMediators mediators) {
        this.mediators = mediators;
    }

    @Override
    protected void processComponentKeyEvent(KeyEvent keyEvent) {
        mediators.markNote();
    }
}

// Sikares Nuntipatsakul 65070503439