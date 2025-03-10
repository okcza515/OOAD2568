
import javax.swing.JButton;
import java.awt.event.ActionEvent;

class DeleteButton extends JButton {
    private Mediator mediator;

    public DeleteButton(Mediator mediator) {
        super("Del");
        this.mediator = mediator;
    }

    @Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.deleteNote();
    }
}