
public class Application {

	public static void main(String[] args) {
		Editor editor = new Editor();
        editor.events.subscribe("open", new LogOpenListener("just3rd.txt"));
        editor.events.subscribe("save", new EmailNotificationListener("just3rd@gmail.com"));
		try {
            editor.openFile("test.txt");
            editor.saveFile();
        } catch (Exception e) {
            e.printStackTrace();
        }
	}

}
