
import java.io.File;

public class EmailNotificationListener implements EventListener {

	private String email;

    public EmailNotificationListener(String email) {
        this.email = email;
    }

    @Override
    public void update(String event, File file) {
        System.out.println("Email:" + email + "has" + event + "in the" + file.getName());
    }
}
