public class Application {

	private static Dialog dialog;

	public static void main(String[] args) {
		configure();
		runBusinessLogic();
	}

	static void configure() {
		dialog = new HtmlDialog();

		String osName = System.getProperty("os.name").toLowerCase();

		if (osName.contains("mac") || osName.contains("win")) {
			dialog = new WindowsDialog();
		} else {
			dialog = new HtmlDialog();
		}
	}

	static void runBusinessLogic() {
		dialog.renderWindow();
	}

}
