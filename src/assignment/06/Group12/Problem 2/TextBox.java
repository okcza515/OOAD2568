import javax.swing.*;
import java.awt.event.KeyEvent;

public class TextBox extends JTextArea {

    private Editor mediator;

    public TextBox(Editor mediator) {
        super();
        this.mediator = mediator;
    }

    @Override
    protected void processComponentKeyEvent(KeyEvent keyEvent) {
        mediator.markNote();
        super.processComponentKeyEvent(keyEvent);
    }
}