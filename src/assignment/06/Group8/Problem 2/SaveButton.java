import javax.swing.JButton;

public class SaveButton extends JButton {
    private EditorMediator mediator;

    public SaveButton(EditorMediator mediator) {
        super("Save");
        this.mediator = mediator;
        addActionListener(e -> mediator.saveNote());
    }
}
