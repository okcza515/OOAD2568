
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class DeleteButton extends JButton{
	
    private EditorMediator mediator;

	public DeleteButton(EditorMediator mediator) {
        super("Del");
        this.mediator = mediator;
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.deleteNote();
    }
}
