
import javax.swing.JTextField;
import java.awt.event.KeyEvent;

public class Title extends JTextField {
    private EditorMediators mediators;

    public Title(EditorMediators mediators) {
        this.mediators = mediators;
    }

    @Override

    protected void processComponentKeyEvent(KeyEvent keyEvent) {
        mediators.markNote();
    }
}

// Sikares Nuntipatsakul 65070503439