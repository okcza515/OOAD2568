
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class SaveButton extends JButton{

    private EditorMediator _mediator;

	public SaveButton(EditorMediator mediator) {
        super("Save");
        this._mediator = mediator;
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        _mediator.saveChanges();
    }
}
