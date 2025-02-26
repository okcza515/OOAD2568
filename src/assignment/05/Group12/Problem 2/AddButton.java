
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class AddButton extends JButton{
	
	public AddButton() {
        super("Add");
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        Editor.addNewNote(new Note());
    }

}
