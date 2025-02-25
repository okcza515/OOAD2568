
public class Application {

	private static Dialog dialog;

	public static void main(String[] args) {
		configure();
		runBusinessLogic();
	}

	static void configure() {
		//New code after implement Windows dialog
		if (System.getProperty("os.name").equals("Mac OS X")
				|| System.getProperty("os.name").startsWith("Windows")
		) {
			dialog = new WindowsDialog();
		} else {
			dialog = new HtmlDialog();
		}
	}

	static void runBusinessLogic() {
		dialog.renderWindow();
	}

}
// 65070501051