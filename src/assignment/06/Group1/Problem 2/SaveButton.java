
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class SaveButton extends JButton{

    private EditorMediator mediator;
    
	public SaveButton(EditorMediator mediator) {
        super("Save");
        this.mediator = mediator;
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.saveChanges();
    }
}
