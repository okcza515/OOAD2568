
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class DeleteButton extends JButton implements Element {
    private Mediator mediator;

    public DeleteButton() {
        super("Delete");
    }

    @Override
    public void setMediator(Mediator mediator) {
        this.mediator = mediator;
    }

	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.deleteNote();
    }

    @Override
    public String getName() {
        return "Delete Button";
    }
}
