import java.io.File;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Editor {
    private File file;
    private Map<String, List<EventListener>> listeners = new HashMap<>();

    public Editor() {
        listeners.put("open", new ArrayList<>());
        listeners.put("save", new ArrayList<>());
    }

    public void subscribe(String eventType, EventListener listener) {
        if (listeners.containsKey(eventType)) {
            listeners.get(eventType).add(listener);
        }
    }

    public void unsubscribe(String eventType, EventListener listener) {
        if (listeners.containsKey(eventType)) {
            listeners.get(eventType).remove(listener);
        }
    }

    private void notifyListeners(String eventType) {
        if (file != null && listeners.containsKey(eventType)) {
            for (EventListener listener : listeners.get(eventType)) {
                listener.update(eventType, file);
            }
        }
    }

    public void openFile(String filePath) {
        this.file = new File(filePath);
        notifyListeners("open");
    }

    public void saveFile() throws Exception {
        if (this.file != null) {
            notifyListeners("save");
        } else {
            throw new Exception("Please open a file first.");
        }
    }
}
