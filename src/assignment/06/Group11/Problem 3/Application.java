import java.util.Arrays;

public class Application {
    public static void main(String[] args) {
        Editor editor = new Editor();
        
        // Setting up listeners
        LogOpenListener logListener = new LogOpenListener("file_ops.log");
        EmailNotificationListener emailListener = new EmailNotificationListener("admin@example.com");
        
        // You can add more email recipients
        emailListener.addEmail("manager@example.com");
        
        // Or create with multiple recipients at once
        // EmailNotificationListener multiEmailListener = new EmailNotificationListener(
        //     Arrays.asList("admin@example.com", "manager@example.com", "dev@example.com")
        // );
        
        // Subscribe listeners to different events
        editor.getEventManager().subscribe("open", logListener);
        editor.getEventManager().subscribe("save", emailListener);
        
        // Using the editor
        try {
            editor.openFile("test.txt");
            editor.saveFile();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}