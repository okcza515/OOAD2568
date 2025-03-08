import javax.swing.JButton;

public class AddButton extends JButton {
    private EditorMediator mediator;

    public AddButton(EditorMediator mediator) {
        super("Add");
        this.mediator = mediator;
        addActionListener(e -> mediator.addNote());
    }
}
