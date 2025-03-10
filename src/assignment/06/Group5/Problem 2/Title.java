
import javax.swing.JTextField;
import java.awt.event.KeyEvent;

class Title extends JTextField {
    private Mediator mediator;

    public Title(Mediator mediator) {
        super(); // Call the JTextField constructor
        this.mediator = mediator;
    }

    @Override
    protected void processComponentKeyEvent(KeyEvent keyEvent) {
        mediator.markNote();
    }
}