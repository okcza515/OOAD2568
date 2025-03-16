import javax.swing.*;
import java.awt.event.ActionEvent;

public class SaveButton extends JButton {

    private Editor mediator;

    public SaveButton(String text, Editor mediator) {
        super(text);
        this.mediator = mediator;
    }

    @Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.saveChanges();
    }
}