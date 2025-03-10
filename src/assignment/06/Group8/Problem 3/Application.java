
public class Application {
    public static void main(String[] args) {
        Editor editor = new Editor();

        // Register listeners
        editor.events.subscribe("open", new LogOpenListener("log.txt"));
        editor.events.subscribe("save", new EmailNotificationListener("admin@example.com"));

        try {
            editor.openFile("test.txt");
            editor.saveFile();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
