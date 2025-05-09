import java.io.File;

// 65070501011 Chayapol Wongpuwarak
public class LogOpenListener implements EventListener {
    private final String logFileName;

    public LogOpenListener(String fileName) {
        this.logFileName = fileName;
    }

    @Override
    public void update(String eventType, File file) {
        System.out.println("Save to log " + logFileName + ": Someone performed " + eventType 
                + " operation on the file: " + file.getName());
    }
}
