import java.io.File;
import java.util.ArrayList;
import java.util.List;

public class Editor {
    private File file;
    private List<EventListener> listeners = new ArrayList<>();

    public Editor() {
    }

    public void subscribe(EventListener listener) {
        listeners.add(listener);
    }

    public void unsubscribe(EventListener listener) {
        listeners.remove(listener);
    }

    private void notify(String eventType, File file) {
        for (EventListener listener : listeners) {
            listener.update(eventType, file);
        }
    }

    public void openFile(String filePath) {
        this.file = new File(filePath);
        notify("open", file);
    }

    public void saveFile() throws Exception {
        if (this.file != null) {
            notify("save", file);
        } else {
            throw new Exception("Please open a file first.");
        }
    }
}