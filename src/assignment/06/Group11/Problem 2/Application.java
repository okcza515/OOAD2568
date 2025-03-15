
import javax.swing.*;
public class Application {

	public static void main(String[] args) {
		Mediator editor = new Editor();

		editor.createElement(new Title());
		editor.createElement(new TextBox());
		editor.createElement(new AddButton());
		editor.createElement(new DeleteButton());
		editor.createElement(new SaveButton());
		editor.createElement(new List(new DefaultListModel()));
		editor.createElement(new Filter());
		editor.createGUI();
	}

}
