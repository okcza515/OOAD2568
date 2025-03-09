
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class DeleteButton extends JButton {
    private EditorMediators mediators;

    public DeleteButton(EditorMediators mediators) {
        super("Del");
        this.mediators = mediators;
    }

    @Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediators.deleteNote();
    }
}
// Ratchnon Tarawan 65070503464
