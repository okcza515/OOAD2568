import java.io.File;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Publisher {
    private Map<String, List<Listener>> listeners = new HashMap<>();

    public Publisher(String... events) {
        for (String event : events) {
            listeners.put(event, new ArrayList<>());
        }
    }

    public void subscribe(String eventType, Listener listener) {
        listeners.get(eventType).add(listener);
    }

    public void unsubscribe(String eventType, Listener listener) {
        listeners.get(eventType).remove(listener);
    }

    public void notify(String eventType, File file) {
        for (Listener listener : listeners.get(eventType)) {
            listener.update(eventType, file);
        }
    }
}

//65070501018 Natchanon Phattamanuruk
//65070501074 Napat Sinjindawong