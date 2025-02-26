
public class Application {

	private static IDialog dialog;
	

	public static void main(String[] args) {
		configure();
		runBusinessLogic();
	}

	static void configure() {
		dialog = new HtmlDialog();
		
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
