
public class Application {

	public static void main(String[] args) {
		Editor editor = new Editor();

        // Registering listeners
        LogOpenListener logListener1 = new LogOpenListener("log1.txt");
        EmailNotificationListener emailListener1 = new EmailNotificationListener("admin@example.com");
        EmailNotificationListener emailListener2 = new EmailNotificationListener("user@kmutt.com");

        editor.publisher.subscribe("open", logListener1);
        editor.publisher.subscribe("save", emailListener1);
        editor.publisher.subscribe("save", emailListener2);

        try {
            System.out.println("------------Test Subscribed------------\n");
            editor.openFile("test2.txt");
            editor.saveFile();
        } catch (Exception e) {
            e.printStackTrace();
        }

        editor.publisher.subscribe("open", logListener1);
        editor.publisher.unsubscribe("save", emailListener1);

        try {
            System.out.println("------------Test Unsubscribed------------\n");
            editor.openFile("test3.txt");
            editor.saveFile();
        } catch (Exception e) {
            e.printStackTrace();
        }
	}

}

//65070501018 Natchanon Phattamanuruk
//65070501074 Napat Sinjindawong