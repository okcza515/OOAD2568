
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class AddButton extends JButton{
    protected Mediator mediator;
	
	public AddButton(Mediator mediator) {
        super("Add");
        this.mediator = mediator;
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.addNewNote(new Note());
    }

}
