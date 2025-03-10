import javax.swing.JTextField;
import java.awt.event.KeyEvent;

public class Title extends JTextField{
private EditorMediator _mediator;

    public Title(EditorMediator mediator){
        this._mediator = mediator;
    }

protected void processComponentKeyEvent(KeyEvent keyEvent) {
        _mediator.markNote();
    }
}

//65070501069 Kanitsorn Darunaitorn