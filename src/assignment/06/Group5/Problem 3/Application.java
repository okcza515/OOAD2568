
public class Application {

	public static void main(String[] args) {
		Editor editor = new Editor();

		try {
            editor.openFile("test2.txt");
            editor.saveFile();
        } catch (Exception e) {
            e.printStackTrace();
        }
	}

}

//65070501074 Napat Sinjindawong