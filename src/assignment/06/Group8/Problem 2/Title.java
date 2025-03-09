
import javax.swing.JTextField;
import java.awt.event.KeyEvent;

public class Title extends JTextField{
    private EditorMediator mediator;

    public Title(EditorMediator mediator) {
        this.mediator = mediator;
    }
	
    @Override
	protected void processComponentKeyEvent(KeyEvent keyEvent) {
        mediator.markNote();
    }
}
