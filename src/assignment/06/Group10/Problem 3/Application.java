//65070501079 Pitchayuth Jampong

public class Application {
    public static void main(String[] args) {
        Editor editor = new Editor();
        EventListener logListener = new LogOpenListener("log.txt");
        EventListener emailListener = new EmailNotificationListener("faanbodyslim@gmail.com");

        // Subscribe listeners to events
        editor.events.subscribe("open", logListener);
        editor.events.subscribe("save", emailListener);

        try {
            editor.openFile("test.txt");
            editor.saveFile();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
