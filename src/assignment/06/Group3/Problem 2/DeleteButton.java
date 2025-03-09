
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class DeleteButton extends JButton{
	
	public DeleteButton() {
        super("Del");
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        Editor.deleteNote();
    }
}
