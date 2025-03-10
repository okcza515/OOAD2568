// 65070501019 Natlada Simasathien

import java.awt.event.ActionEvent;
import javax.swing.JButton;

public class SaveButton extends JButton implements Element{
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
        return "Save Button";
    }
}
