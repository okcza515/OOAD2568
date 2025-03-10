
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class SaveButton extends JButton{
    protected Mediator mediator;

	public SaveButton(Mediator mediator) {
        super("Save");
        this.mediator = mediator;
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.saveChanges();
    }
}
