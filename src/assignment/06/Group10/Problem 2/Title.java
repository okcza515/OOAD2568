// 65070501019 Natlada Simasathien

import java.awt.event.KeyEvent;
import javax.swing.JTextField;

public class Title extends JTextField implements Element {

    private Mediator mediator;

    @Override
    public void setMediator(Mediator mediator) {
        this.mediator = mediator;
    }

    @Override
	protected void processComponentKeyEvent(KeyEvent keyEvent) {
        mediator.markNote();
    }

    @Override
    public String getName() {
        return "Title";
    }
}