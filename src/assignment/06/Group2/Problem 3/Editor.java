
import java.io.File;

public class Editor {
	private File file;
	private LogOpenListener log;
	private EmailNotificationListener email;
	public EventManager events;
	
	public Editor(){
		this.events = new EventManager("open", "save");
	}
	
	public void openFile(String filePath) {
		this.file = new File(filePath);
		log = new LogOpenListener(file.getName());
		log.update("open", file);
	}
	
	public void saveFile() throws Exception{
		email = new EmailNotificationListener("admin@example.com");
		if(this.file != null) {
			email.update("save", file);
		}else {
			throw new Exception("Please open a file first.");
		}
	}
}
//Korawit Sritotum 65070503402
