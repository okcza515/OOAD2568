// Sawitt Ngamvilaisiriwong 65070503469

import javax.swing.JTextField;
import java.awt.event.KeyEvent;

public class Title extends JTextField{
	
	protected void processComponentKeyEvent(KeyEvent keyEvent) {
        Editor.markNote();
    }
}