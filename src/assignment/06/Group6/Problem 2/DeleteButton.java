
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class DeleteButton extends JButton{
	protected Mediator mediator;

	public DeleteButton(Mediator mediator) {
        super("Del");
        this.mediator = mediator;
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.deleteNote();
    }
}
