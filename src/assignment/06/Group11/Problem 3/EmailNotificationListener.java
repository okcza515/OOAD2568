import java.io.File;
import java.util.List;
import java.util.ArrayList;

public class EmailNotificationListener implements EventListener {
    private List<String> emails;

    public EmailNotificationListener(String email) {
        this.emails = new ArrayList<>();
        this.emails.add(email);
    }

    public EmailNotificationListener(List<String> emails) {
        this.emails = new ArrayList<>(emails);
    }

    public void addEmail(String email) {
        if (!emails.contains(email)) {
            emails.add(email);
        }
    }

    public void removeEmail(String email) {
        emails.remove(email);
    }

    @Override
    public void update(String eventType, File file) {
        for (String email : emails) {
            System.out.println("Email to " + email + ": Someone has performed " + eventType
                    + " operation with the following file: " + file.getName());
        }
    }
}