import javax.swing.*;
import java.awt.event.KeyEvent;

public class Title extends JTextField {

    private Editor mediator;

    public Title(Editor mediator) {
        super();
        this.mediator = mediator;
    }

    @Override
    protected void processComponentKeyEvent(KeyEvent keyEvent) {
        mediator.markNote();
        super.processComponentKeyEvent(keyEvent);
    }
}