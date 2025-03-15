// Sawitt Ngamvilaisiriwong 65070503469

import javax.swing.JTextArea;
import java.awt.event.KeyEvent;

public class TextBox extends JTextArea{
	
	@Override
    protected void processComponentKeyEvent(KeyEvent keyEvent) {
        Editor.markNote();
    }
}