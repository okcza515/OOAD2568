
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class AddButton extends JButton{
	
    private EditorMediator mediator;
    
	public AddButton(EditorMediator mediator) {
        super("Add");
        this.mediator = mediator;
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.addNewNote(new Note());
    }
}
