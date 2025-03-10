
import javax.swing.JButton;
import java.awt.event.ActionEvent;

class AddButton extends JButton {
    private Mediator mediator;

    public AddButton(Mediator mediator) {
        super("Add");
        this.mediator = mediator;
    }

    @Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.addNote();
    }
}