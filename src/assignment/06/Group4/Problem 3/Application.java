
public class Application {

	public static void main(String[] args) {
		Editor editor = new Editor();

        editor.events.subscribe("open", new LogOpenListener("log.txt"));
        editor.events.subscribe("save", new EmailNotificationListener("sorrawit@cpe1119.dev"));
        editor.events.subscribe("save", new EmailNotificationListener("chanapat@cpe1119.dev"));

        EmailNotificationListener chaiyapat = new EmailNotificationListener("chaiyapat@cpe1119.dev");
        editor.events.subscribe("save",chaiyapat);
        editor.events.unsubscribe("save", chaiyapat);

        editor.events.subscribe("fake-log", new LogOpenListener("fake-log.txt"));
        editor.events.subscribe("fake-event", new EmailNotificationListener("chanapat2@cpe1119.dev"));

		try {
            editor.openFile("test.txt");
            editor.saveFile();
        } catch (Exception e) {
            e.printStackTrace();
        }
	}

}
