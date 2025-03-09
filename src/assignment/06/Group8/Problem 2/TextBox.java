
import javax.swing.JTextArea;
import java.awt.event.KeyEvent;

public class TextBox extends JTextArea{
    private EditorMediator mediator;

    public TextBox(EditorMediator mediator) {
        this.mediator = mediator;
    }
	
	@Override
    protected void processComponentKeyEvent(KeyEvent keyEvent) {
        mediator.markNote();
    }
}
