
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class SaveButton extends JButton implements Element {
    private Mediator mediator;

    @Override
    public void setMediator(Mediator mediator) {
        this.mediator = mediator;
    }

	public SaveButton() {
        super("Save");
    }

	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.saveChanges();
    }

    @Override
    public String getName() {
        return "Save Button";
    }
}
