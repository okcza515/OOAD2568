
public class Application {

	public static void main(String[] args) {
		Editor editor = new Editor();
		try {
            editor.openFile("test.txt");
            editor.saveFile();
        } catch (Exception e) {
            e.printStackTrace();
        }
	}

}
