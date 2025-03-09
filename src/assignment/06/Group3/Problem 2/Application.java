
public class Application {

	public static void main(String[] args) {
		Editor editor = new Editor();

		mediator.registerComponent(new Title());
		mediator.registerComponent(new TextBox());
		mediator.registerComponent(new AddButton());
		mediator.registerComponent(new DeleteButton());
		mediator.registerComponent(new SaveButton());
		mediator.registerComponent(new List(new DefaultListModel()));
		mediator.registerComponent(new Filter());

		editor.createGUI();
	}

}

// 65070501001 Kantapong Vongapnich