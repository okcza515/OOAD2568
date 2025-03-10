import javax.swing.JTextArea;
import java.awt.event.KeyEvent;

public class TextBox extends JTextArea{
    private EditorMediator _mediator;

    public TextBox(EditorMediator mediator){
        this._mediator = mediator;
    }
	@Override
    protected void processComponentKeyEvent(KeyEvent keyEvent) {
        _mediator.markNote();
    }
}

//65070501081 Phakawat Rattanasopa