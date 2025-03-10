
import javax.swing.JTextArea;
import java.awt.event.KeyEvent;

public class TextBox extends JTextArea{
    protected Mediator mediator;

    public TextBox(Mediator mediator) {
        this.mediator = mediator;
    }
	
	@Override
    protected void processComponentKeyEvent(KeyEvent keyEvent) {
        mediator.markNote();
    }
}
