
import java.io.File;

public class LogOpenListener implements EventListener {
	 private File log;

    public LogOpenListener(String fileName) {
        this.log = new File(fileName);
    }

    @Override
    public void update(String event, File file) {
        System.out.println(log + " got " + event + " in " + file.getName());
    }
}
