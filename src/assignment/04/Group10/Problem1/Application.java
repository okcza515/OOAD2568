
public class Application {

	private static DialogFactory dialog;

	public static void main(String[] args) {
		configure();
		runBusinessLogic();
	}

	static void configure() {
		dialog = new HtmlDialog();
		// New code after implement Windows dialog
		if (System.getProperty("os.name").equals("Mac OS X")) {
			dialog = new WindowsDialog();
		} else {
			dialog = new HtmlDialog();
		}
	}

	static void runBusinessLogic() {
		dialog.renderWindow();
	}

}
// 65070501049 Roodfan Maimahad
