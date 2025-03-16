import javax.swing.*;
import java.awt.event.ActionEvent;

public class AddButton extends JButton {

    private Editor mediator;

    public AddButton(String text, Editor mediator) {
        super(text);
        this.mediator = mediator;
    }

    @Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.addNewNote(new Note());
    }
}