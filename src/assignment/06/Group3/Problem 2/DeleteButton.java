
import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class DeleteButton extends JButton implements Component {

    private Mediator mediator
	public DeleteButton() {
        super("Del");
    }

    @Override
    public void setMediator(Mediator mediator) {
        this.mediator = mediator;
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        Editor.deleteNote();
    }

    @Override
    public String getName() {
        return "DeleteButton";
    }
}

// 65070501039 Pongpon Butseemart