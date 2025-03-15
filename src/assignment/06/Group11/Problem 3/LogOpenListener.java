// LogOpenListener.java
import java.io.File;

public class LogOpenListener implements EventListener {
    private String filename;

    public LogOpenListener(String filename) {
        this.filename = filename;
    }

    @Override
    public void update(String eventType, File file) {
        System.out.println("Save to log " + filename + ": Someone has performed " + eventType
                + " operation with the following file: " + file.getName());
    }
}