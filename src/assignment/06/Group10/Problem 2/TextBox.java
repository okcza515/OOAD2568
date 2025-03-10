// 65070501019 Natlada Simasathien

import java.awt.event.KeyEvent;
import javax.swing.JTextArea;

public class TextBox extends JTextArea implements Element{
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
        return "Text Box";
    }
}