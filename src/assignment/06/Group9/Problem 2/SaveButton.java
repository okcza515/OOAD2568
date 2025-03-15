// Sawitt Ngamvilaisiriwong 65070503469

import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class SaveButton extends JButton{

	public SaveButton() {
        super("Save");
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        Editor.saveChanges();
    }
}