import javax.swing.*;
import java.awt.event.ActionEvent;

public class DeleteButton extends JButton {

    private Editor mediator;

    public DeleteButton(String text, Editor mediator) {
        super(text);
        this.mediator = mediator;
    }

    @Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.deleteNote();
    }
}