
import javax.swing.JTextArea;
import java.awt.event.KeyEvent;

class TextBox extends JTextArea {
    private Mediator mediator;

    public TextBox(Mediator mediator) {
        this.mediator = mediator;
    }

    @Override
    protected void processComponentKeyEvent(KeyEvent keyEvent) {
        mediator.markNote();
    }
}