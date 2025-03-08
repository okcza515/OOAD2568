import javax.swing.JButton;

public class DeleteButton extends JButton {
    private EditorMediator mediator;

    public DeleteButton(EditorMediator mediator) {
        super("Del");
        this.mediator = mediator;
        addActionListener(e -> mediator.deleteNote());
    }
}
