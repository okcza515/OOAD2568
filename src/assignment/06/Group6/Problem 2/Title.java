
import javax.swing.JTextField;
import java.awt.event.KeyEvent;

public class Title extends JTextField{
    protected Mediator mediator;
	
	protected void processComponentKeyEvent(KeyEvent keyEvent) {
        mediator.markNote();
    }
}
