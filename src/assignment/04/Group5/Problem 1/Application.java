//65070501074 Napat Sinjindawong
public class Application {

	//private static HtmlDialog dialog;
	private static Dialog dialog;

	public static void main(String[] args) {
		configure();
		runBusinessLogic();
	}

	static void configure() {
		//dialog = new HtmlDialog();

		//New code after implement Windows dialog
		if (System.getProperty("os.name").equals("Window")) {
			dialog = new WindowsDialog();
		} else {
			dialog = new HtmlDialog();
		}

	}

	static void runBusinessLogic() {
		dialog.renderWindow();
	}

}