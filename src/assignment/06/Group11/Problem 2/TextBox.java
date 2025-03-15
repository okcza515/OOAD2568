
import javax.swing.JTextArea;
import java.awt.event.KeyEvent;

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
