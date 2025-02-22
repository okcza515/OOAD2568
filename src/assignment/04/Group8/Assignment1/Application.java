
public class Application {

	private static Dialog dialog;

	public static void main(String[] args) {
		configure();
		runBusinessLogic();
	}

	static void configure() {
		//dialog = new HtmlDialog();
		// New code after implement Windows dialog
		if (System.getProperty("os.name").equals("Windows")) {
			dialog = new WindowsDialog();
		} else {
			dialog = new HtmlDialog();
		}
	}

	static void runBusinessLogic() {
		dialog.renderWindow();
	}

}
