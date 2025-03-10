import java.io.File;

// 65070501011 Chayapol Wongpuwarak
public class EmailNotificationListener implements EventListener {

	private final String email;

	public EmailNotificationListener(String email) {
		this.email = email;
	}

	@Override
	public void update(String eventType, File file) {
		System.out.println("Email to " + email + ": Someone has performed " + eventType
				+ " operation with the following file: " + file.getName());
	}
}
