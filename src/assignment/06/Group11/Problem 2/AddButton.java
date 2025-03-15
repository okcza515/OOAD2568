
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class AddButton extends JButton implements Element {

    private Mediator mediator;
	public AddButton() {
        super("Add");
    }

    @Override
    public void setMediator(Mediator mediator) {
        this.mediator = mediator;
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.addNewNote(new Note());
    }

    @Override
    public String getName() {
        return "Add Button";
    }

}
