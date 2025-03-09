
import javax.swing.JTextField;
import java.awt.event.KeyEvent;

public class Title extends JTextField{
    private EditorMediator mediators;

    public Title(EditorMediator mediators) {
        this.mediators = mediators;
    }
    
    @Override
	
	protected void processComponentKeyEvent(KeyEvent keyEvent) {
        mediators.markNote();
    }
}

// Sikares Nuntipatsakul 65070503439