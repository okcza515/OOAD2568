import java.io.File;

public class LogOpenListener implements EventListener {
    private String logFileName;

    public LogOpenListener(String logFileName) {
        this.logFileName = logFileName;
    }

    @Override
    public void update(String eventType, File file) {
        System.out.println("Save to log " + logFileName + ": Someone performed " + eventType +
                " operation with file: " + file.getName());
    }
}
