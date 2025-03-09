
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class AddButton extends JButton {
    private EditorMediators mediators;

    public AddButton(EditorMediators mediators) {
        super("Add");
        this.mediators = mediators;
    }

    @Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediators.addNewNote(new Note());
    }
}

// Ratchnon Tarawan 65070503464