import java.io.File;
import java.util.List;

public class Editor {
  private File file;
  public EventPublisher events;

  public Editor() {
    this.events = new EventPublisher(List.of("open", "save"));
  }

  public void openFile(String filePath) {
    this.file = new File(filePath);
    events.notifySubscribe("open", file);
  }

  public void saveFile() throws Exception {
    if (this.file != null) {
      events.notifySubscribe("save", file);
    } else {
      throw new Exception("Please open a file first.");
    }
  }
}
