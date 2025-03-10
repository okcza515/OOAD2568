
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class AddButton extends JButton{
    private EditorMediator _mediator;

	public AddButton(EditorMediator mediator) {
        super("Add");
        this._mediator = mediator;
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        _mediator.addNewNote(new Note());
    }

}
