
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class DeleteButton extends JButton{

    private EditorMediator _mediator;
	
	public DeleteButton(EditorMediator mediator) {
        super("Del");
        this._mediator = mediator;
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        _mediator.deleteNote();
    }
}
