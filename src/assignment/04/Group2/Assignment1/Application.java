public class Application {

	private static DialogFactory dialog;

	public static void main(String[] args) {
		configure();
		runBusinessLogic();
	}

	static void configure() {
		// New code after implement Windows dialog
		if (System.getProperty("os.name").equals("Mac OS X")) {
			dialog = new WindowDialogFactory();
		} else {
			dialog = new HtmlDialogFactory();
		}
	}

	static void runBusinessLogic() {
		Dialog dialog_ = dialog.CreateDialog();
		dialog_.renderWindow();
	}

}
//Korawit Sritotum 65070503402
