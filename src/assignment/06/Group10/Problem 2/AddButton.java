// 65070501049 Roodfan Maimahad

import java.awt.event.ActionEvent;
import javax.swing.JButton;

public class AddButton extends JButton implements Element {

    private Mediator mediator;
	public AddButton() {
        super("Add");
    }

    @Override
    public void setMediator(Mediator mediator) {
        this.mediator = mediator;
    }

    @Override
    public String getName() {
        return "Add Button";
    }
	
	@Override
    protected void fireActionPerformed(ActionEvent actionEvent) {
        mediator.addNewNote(new Note());
    }

}