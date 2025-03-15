import java.io.File;

// Common interface for all event listeners
public interface EventListener {
    void update(String eventType, File file);
}