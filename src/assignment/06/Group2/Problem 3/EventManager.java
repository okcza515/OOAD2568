import java.util.ArrayList;
import java.util.List;
import java.io.File;

public class EventManager {
    private List<EventListener> listeners;

    public EventManager(){
        this.listeners = new ArrayList<>();
    }

    public void subscribe(EventListener listener){
        listeners.add(listener);
    }

    public void unsubscribe(EventListener listener){
        listeners.remove(listener);
    }

    public void notify(String eventType, File file){
        for(EventListener listener : listeners){
            listener.update(eventType, file);
        }
    }

}
//Korawit Sritotum 65070503402