import java.io.File;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class EventPublisher {
    private final Map<String, List<EventListener>> listeners = new HashMap<>();

    public EventPublisher(List<String> events){
        for(String event : events){
            this.listeners.put(event, new ArrayList<>());
        }
    }

    public void subscribe(String eventType, EventListener listener){
        if(listeners.containsKey(eventType)){
            listeners.get(eventType).add(listener);
        }
    }

    public void unsubscribe(String eventType, EventListener listener){
        if(listeners.containsKey(eventType)){
            listeners.get(eventType).remove(listener);
        }
    }
    public void notifySubscribe(String eventType, File file){
        if(file != null && listeners.containsKey(eventType)){
            for(EventListener listener : listeners.get(eventType)){
                listener.update(eventType, file);
            }
        }
    }
}
