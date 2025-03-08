
import java.io.File;

public class Editor {
    private File file;
    public Publisher publisher;

    public Editor() {

        this.publisher = new Publisher("open", "save");
    }

    public void openFile(String filePath) {
        this.file = new File(filePath);
        publisher.notify("open", file);
    }

    public void saveFile() throws Exception {
        //email = new EmailNotificationListener("admin@example.com");
        if(this.file != null) {
            publisher.notify("save", file);
        } else {
            throw new Exception("Please open a file first.");
        }
    }
}

//65070501018 Natchanon Phattamanuruk
//65070501074 Napat Sinjindawong
