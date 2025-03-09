
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class SaveButton extends JButton{

    private EditorMediators mediators;

	public SaveButton(EditorMediators mediators) {
        super("Save");
        this.mediators = mediators;
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediators.saveChanges();
    }
}

// Sikares Nuntipatsakul 65070503439