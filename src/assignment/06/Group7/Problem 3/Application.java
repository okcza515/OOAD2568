
public class Application {

	public static void main(String[] args) {
		Editor editor = new Editor();
        editor.events.subscribe("open", new LogOpenListener("/path/log/file.txt"));
        editor.events.subscribe("save", new EmailNotificationListener("puwadech.into@kmutt.ac.th"));


		try {
            editor.openFile("test.txt");
            editor.saveFile();
        } catch (Exception e) {
            e.printStackTrace();
        }
	}

}
