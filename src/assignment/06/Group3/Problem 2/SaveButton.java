import javax.swing.JButton;
import java.awt.event.ActionEvent;

public class SaveButton extends JButton implements Component{
    private Mediator mediator;

 	public SaveButton() {
        super("Save");
    }
	
    @Override
    public void setMediator(Mediator  mediator) {
        this.mediator = mediator;
    }

	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.saveChanges();
    }

    @Override
    public String getName() {
        return "SaveButton";
    }
}

//65070501088 Sopida Keawjongkool