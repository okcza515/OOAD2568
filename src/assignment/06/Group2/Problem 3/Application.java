
public class Application {

	public static void main(String[] args) {
		Editor editor = new Editor();

        LogOpenListener logOpenListener = new LogOpenListener("log.txt");
        EmailNotificationListener emailNotificationListener = new EmailNotificationListener("gingerx@mail.com");

        editor.events.subscribe("open", logOpenListener);
        editor.events.subscribe("save", emailNotificationListener);
		try {
            editor.openFile("test.txt");
            editor.saveFile();
        } catch (Exception e) {
            e.printStackTrace();
        }
	}

}
//Korawit Sritotum 65070503402
