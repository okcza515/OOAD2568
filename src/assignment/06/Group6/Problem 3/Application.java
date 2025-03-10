public class Application {
    public static void main(String[] args) {
        Editor editor = new Editor();
        
        // Subscribe log and email listeners
        editor.subscribe("open", new LogOpenListener("log.txt"));
        editor.subscribe("save", new EmailNotificationListener("admin@example.com"));
        editor.subscribe("save", new EmailNotificationListener("support@example.com")); // Additional recipient

        try {
            editor.openFile("test.txt");
            editor.saveFile();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
