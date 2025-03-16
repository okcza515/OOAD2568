public class Application {

    public static void main(String[] args) {
        Editor editor = new Editor();

        // Subscribe the logging listener for open events
        editor.subscribe(new LogOpenListener("log.txt"));

        // Subscribe an email listener for save events
        editor.subscribe(new EmailNotificationListener("admin@example.com"));

        try {
            editor.openFile("test.txt");
            editor.saveFile();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

}